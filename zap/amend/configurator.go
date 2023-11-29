package amend

import (
	"encoding/xml"
	"io"

	"github.com/hasty/alchemy/matter"
)

func (r *renderer) writeConfigurator(dp xmlDecoder, e xmlEncoder, el xml.StartElement) (err error) {

	r.bitmaps = make(map[*matter.Bitmap]bool)
	r.enums = make(map[*matter.Enum]bool)
	r.clusters = make(map[*matter.Cluster]bool)
	r.structs = make(map[*matter.Struct]bool)

	var cluster *matter.Cluster
	var clusterIDs []string
	for _, m := range r.models {
		switch v := m.(type) {
		case *matter.Cluster:
			if cluster == nil {
				cluster = v
				for _, bm := range v.Bitmaps {
					r.bitmaps[bm] = false
				}
				for _, en := range v.Enums {
					r.enums[en] = false
				}
				for _, s := range v.Structs {
					r.structs[s] = false
				}
			}
			clusterIDs = append(clusterIDs, v.ID.HexString())
			r.clusters[v] = false
		}
	}

	configuratorAttributes := map[string]string{
		//"domain": matter.DomainNames[r.doc.Domain],
	}

	var configuratorTokens []xml.Token
	configuratorTokens, err = Extract(dp, el)
	if err != nil {
		return
	}

	ts := &tokenSet{tokens: configuratorTokens}

	var hasCommentCharDataPending bool
	var lastIgnoredCharData xml.CharData

	for {
		var tok xml.Token
		tok, err = ts.Token()
		if tok == nil || err == io.EOF {
			err = io.EOF
			return
		} else if err != nil {
			return
		}
		var lastSection = matter.SectionUnknown
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "bitmap":
				name := getAttributeValue(t.Attr, "name")
				if name == "Feature" {
					err = r.writeFeatures(ts, e, t, cluster, clusterIDs)
					if err != nil {
						return
					}
					break
				}
				if lastSection != matter.SectionDataTypeBitmap {
					err = r.flushUnusedConfiguratorValues(e, lastSection, configuratorAttributes, clusterIDs)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionDataTypeBitmap
				err = r.amendBitmap(ts, e, t, cluster, clusterIDs)
			case "enum":
				if lastSection != matter.SectionDataTypeEnum {
					err = r.flushUnusedConfiguratorValues(e, lastSection, configuratorAttributes, clusterIDs)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionDataTypeEnum
				err = r.amendEnum(ts, e, t, cluster, clusterIDs)

			case "struct":
				if lastSection != matter.SectionDataTypeStruct {
					err = r.flushUnusedConfiguratorValues(e, lastSection, configuratorAttributes, clusterIDs)
					if err != nil {
						return
					}
				}

				lastSection = matter.SectionDataTypeStruct
				err = r.amendStruct(ts, e, t, cluster, clusterIDs)

			case "cluster":
				if lastSection != matter.SectionCluster {
					err = r.flushUnusedConfiguratorValues(e, lastSection, configuratorAttributes, clusterIDs)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionCluster
				err = r.amendCluster(ts, e, t)
			case "clusterExtension":
				err = writeThrough(ts, e, t)
			default:
				if v, ok := configuratorAttributes[t.Name.Local]; ok {
					t.Attr = setAttributeValue(t.Attr, "name", v)
					delete(configuratorAttributes, t.Name.Local)
				}
				err = e.EncodeToken(tok)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "cluster":
			case "enum", "struct", "bitmap":
			case "configurator":
				err = r.flushBitmaps(e, clusterIDs)
				if err != nil {
					return
				}
				err = r.flushEnums(e, clusterIDs)
				if err != nil {
					return
				}
				err = r.flushStructs(e, clusterIDs)
				if err != nil {
					return
				}
				return e.EncodeToken(tok)
			default:
				err = e.EncodeToken(tok)
			}
		case xml.Comment:
			if lastIgnoredCharData != nil {
				err = e.EncodeToken(lastIgnoredCharData)
				if err != nil {
					return
				}
				lastIgnoredCharData = nil
			}
			err = e.EncodeToken(t)
			hasCommentCharDataPending = true
		case xml.CharData:
			if hasCommentCharDataPending {
				err = e.EncodeToken(t)
				hasCommentCharDataPending = false
				lastIgnoredCharData = nil
			} else {
				lastIgnoredCharData = t
			}

		default:
			err = e.EncodeToken(tok)
			lastIgnoredCharData = nil
		}
		if err != nil {
			return
		}
	}
}

func (r *renderer) flushUnusedConfiguratorValues(e xmlEncoder, lastSection matter.Section, configuratorValues map[string]string, clusterIDs []string) (err error) {
	switch lastSection {
	case matter.SectionUnknown:
		for k, v := range configuratorValues {
			err = e.EncodeToken(xml.StartElement{Name: xml.Name{Local: k}, Attr: []xml.Attr{{Name: xml.Name{Local: "name"}, Value: v}}})
			if err != nil {
				return
			}
			err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: k}})
			if err != nil {
				return
			}
			delete(configuratorValues, k)
		}
	case matter.SectionDataTypeBitmap:
		err = r.flushBitmaps(e, clusterIDs)
	case matter.SectionDataTypeEnum:
		err = r.flushEnums(e, clusterIDs)
	case matter.SectionDataTypeStruct:
		err = r.flushStructs(e, clusterIDs)
	}
	return
}

func (r *renderer) flushBitmaps(e xmlEncoder, clusterIDs []string) (err error) {
	for bm, handled := range r.bitmaps {
		if handled {
			continue
		}
		err = r.writeBitmap(e, xml.StartElement{Name: xml.Name{Local: "bitmap"}}, bm, clusterIDs, true)
		r.bitmaps[bm] = true
		if err != nil {
			return
		}
	}
	return
}

func (r *renderer) flushEnums(e xmlEncoder, clusterIDs []string) (err error) {
	for en := range r.enums {
		err = r.writeEnum(e, xml.StartElement{Name: xml.Name{Local: "enum"}}, en, clusterIDs, true)
		delete(r.enums, en)
		if err != nil {
			return
		}
	}
	return
}

func (r *renderer) flushStructs(e xmlEncoder, clusterIDs []string) (err error) {
	for s, handled := range r.structs {
		if handled {
			continue
		}
		err = r.writeStruct(e, xml.StartElement{Name: xml.Name{Local: "struct"}}, s, clusterIDs, true)
		r.structs[s] = true
		if err != nil {
			return
		}
	}
	return
}

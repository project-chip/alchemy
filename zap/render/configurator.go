package render

import (
	"encoding/xml"
	"io"

	"github.com/hasty/matterfmt/matter"
)

func (r *renderer) writeConfigurator(d xmlDecoder, e xmlEncoder) (err error) {
	bitmaps := make(map[*matter.Bitmap]struct{})
	enums := make(map[*matter.Enum]struct{})
	clusters := make(map[*matter.Cluster]struct{})
	structs := make(map[*matter.Struct]struct{})
	var cluster *matter.Cluster
	var clusterIDs []string
	for _, m := range r.models {
		switch v := m.(type) {
		case *matter.Cluster:
			if cluster == nil {
				cluster = v
				for _, bm := range v.Bitmaps {
					bitmaps[bm] = struct{}{}
				}
				for _, en := range v.Enums {
					enums[en] = struct{}{}
				}
				for _, s := range v.Structs {
					structs[s] = struct{}{}
				}
			}
			clusterIDs = append(clusterIDs, v.ID.HexString())
			clusters[v] = struct{}{}
		}
	}

	configuratorAttributes := map[string]string{
		"domain": matter.DomainNames[r.doc.Domain],
	}

	err = e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "configurator"}})
	if err != nil {
		return
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
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
					err = r.writeFeatures(d, e, t, cluster, clusterIDs)
					if err != nil {
						return
					}
					break
				}
				if lastSection != matter.SectionDataTypeBitmap {
					err = r.flushUnusedConfiguratorValues(e, lastSection, configuratorAttributes, clusterIDs, bitmaps, enums, structs)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionDataTypeBitmap
				err = r.amendBitmap(d, e, t, cluster, clusterIDs, bitmaps)
			case "enum":
				if lastSection != matter.SectionDataTypeBitmap {
					err = r.flushUnusedConfiguratorValues(e, lastSection, configuratorAttributes, clusterIDs, bitmaps, enums, structs)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionDataTypeEnum
				err = r.amendEnum(d, e, t, cluster, clusterIDs, enums)

			case "struct":
				if lastSection != matter.SectionDataTypeStruct {
					err = r.flushUnusedConfiguratorValues(e, lastSection, configuratorAttributes, clusterIDs, bitmaps, enums, structs)
					if err != nil {
						return
					}
				}

				lastSection = matter.SectionDataTypeStruct
				err = r.amendStruct(d, e, t, cluster, clusterIDs, structs)

			case "cluster":
				lastSection = matter.SectionCluster
				err = r.amendCluster(d, e, t, clusters)
			default:
				if v, ok := configuratorAttributes[t.Name.Local]; ok {
					t.Attr = setAttributeValue(t.Attr, "name", v)
					delete(configuratorAttributes, t.Name.Local)
				}
				err = e.EncodeToken(tok)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "cluster", "enum", "struct", "bitmap":
			case "configurator":
				err = r.flushBitmaps(e, bitmaps, clusterIDs)
				if err != nil {
					return
				}
				err = r.flushEnums(e, enums, clusterIDs)
				if err != nil {
					return
				}
				err = r.flushStructs(e, structs, clusterIDs)
				if err != nil {
					return
				}
				return e.EncodeToken(tok)
			default:
				err = e.EncodeToken(tok)
			}
		case xml.CharData:
		default:
			err = e.EncodeToken(tok)
		}
		if err != nil {
			return
		}
	}
}

func (r *renderer) flushUnusedConfiguratorValues(e xmlEncoder, lastSection matter.Section, configuratorValues map[string]string, clusterIDs []string, bitmaps map[*matter.Bitmap]struct{}, enums map[*matter.Enum]struct{}, structs map[*matter.Struct]struct{}) (err error) {
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
		err = r.flushBitmaps(e, bitmaps, clusterIDs)
	case matter.SectionDataTypeEnum:
		err = r.flushEnums(e, enums, clusterIDs)
	case matter.SectionDataTypeStruct:
		err = r.flushStructs(e, structs, clusterIDs)
	}
	return
}

func (r *renderer) flushBitmaps(e xmlEncoder, bitmaps map[*matter.Bitmap]struct{}, clusterIDs []string) (err error) {
	for bm := range bitmaps {
		err = r.writeBitmap(e, xml.StartElement{Name: xml.Name{Local: "bitmap"}}, bm, clusterIDs)
		delete(bitmaps, bm)
		if err != nil {
			return
		}
	}
	return
}

func (r *renderer) flushEnums(e xmlEncoder, enums map[*matter.Enum]struct{}, clusterIDs []string) (err error) {
	for en := range enums {
		err = r.writeEnum(e, xml.StartElement{Name: xml.Name{Local: "enum"}}, en, clusterIDs)
		delete(enums, en)
		if err != nil {
			return
		}
	}
	return
}

func (r *renderer) flushStructs(e xmlEncoder, structs map[*matter.Struct]struct{}, clusterIDs []string) (err error) {
	for s := range structs {
		err = r.writeStruct(e, xml.StartElement{Name: xml.Name{Local: "struct"}}, s, clusterIDs)
		delete(structs, s)
		if err != nil {
			return
		}
	}
	return
}

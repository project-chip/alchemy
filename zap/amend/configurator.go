package amend

import (
	"encoding/xml"
	"io"
	"slices"
	"strings"

	"github.com/hasty/alchemy/conformance"
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
				r.addTypes(v.Attributes)
				for _, cmd := range v.Commands {
					r.addTypes(cmd.Fields)
				}
				for _, e := range v.Events {
					r.addTypes(e.Fields)
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

	needFeatures := len(cluster.Features) > 0

	var lastSection = matter.SectionUnknown
	for {
		var tok xml.Token
		tok, err = ts.Token()
		if tok == nil || err == io.EOF {
			err = io.EOF
			return
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "bitmap":
				name := getAttributeValue(t.Attr, "name")
				if name == "Feature" {
					err = r.amendFeatures(ts, e, t, cluster, clusterIDs)
					if err != nil {
						return
					}
					needFeatures = false
					break
				}
				if lastSection != matter.SectionDataTypeBitmap {
					err = r.flushUnusedConfiguratorValues(e, lastSection, configuratorAttributes)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionDataTypeBitmap
				err = r.amendBitmap(ts, e, t, cluster)
			case "enum":
				if lastSection != matter.SectionDataTypeEnum {
					err = r.flushUnusedConfiguratorValues(e, lastSection, configuratorAttributes)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionDataTypeEnum
				err = r.amendEnum(ts, e, t, cluster)

			case "struct":
				if lastSection != matter.SectionDataTypeStruct {
					err = r.flushUnusedConfiguratorValues(e, lastSection, configuratorAttributes)
					if err != nil {
						return
					}
				}

				lastSection = matter.SectionDataTypeStruct
				err = r.amendStruct(ts, e, t, cluster)

			case "cluster":
				if lastSection != matter.SectionCluster {
					err = r.flushUnusedConfiguratorValues(e, lastSection, configuratorAttributes)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionCluster
				err = r.amendCluster(ts, e, t)
			case "clusterExtension":
				err = writeThrough(ts, e, t)
			case "tag":
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
				err = r.flushBitmaps(e)
				if err != nil {
					return
				}
				err = r.flushEnums(e)
				if err != nil {
					return
				}
				err = r.flushStructs(e)
				if err != nil {
					return
				}
				if needFeatures {
					r.writeFeatures(ts, e, xml.StartElement{Name: xml.Name{Local: "bitmap"}}, cluster, clusterIDs)
				}
				return e.EncodeToken(tok)
			default:
				err = e.EncodeToken(tok)
			}
		case xml.Comment:
			err = newLine(e)
			if err != nil {
				return
			}
			err = e.EncodeToken(t)
			hasCommentCharDataPending = true
		case xml.CharData:
			if hasCommentCharDataPending {
				err = e.EncodeToken(t)
				hasCommentCharDataPending = false
			}
		default:
			err = e.EncodeToken(tok)
		}
		if err != nil {
			return
		}
	}
}

func (r *renderer) addTypes(fs matter.FieldSet) {
	for _, f := range fs {
		if f.Type == nil {
			continue
		}
		if conformance.IsZigbee(fs, f.Conformance) {
			continue
		}
		r.addType(f.Type)
	}
}

func (r *renderer) addType(dt *matter.DataType) {
	if dt == nil {
		return
	}

	if dt.IsArray() {
		r.addType(dt.EntryType)
		return
	}
	switch model := dt.Model.(type) {
	case *matter.Bitmap:
		r.bitmaps[model] = false
	case *matter.Enum:
		r.enums[model] = false
	case *matter.Struct:
		r.structs[model] = false
		r.addTypes(model.Fields)
	}
}

func (r *renderer) flushUnusedConfiguratorValues(e xmlEncoder, lastSection matter.Section, configuratorValues map[string]string) (err error) {
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
		err = newLine(e)
	case matter.SectionDataTypeBitmap:
		err = r.flushBitmaps(e)
	case matter.SectionDataTypeEnum:
		err = r.flushEnums(e)
	case matter.SectionDataTypeStruct:
		err = r.flushStructs(e)
	}
	return
}

func (r *renderer) flushBitmaps(e xmlEncoder) (err error) {
	bms := make([]*matter.Bitmap, 0, len(r.bitmaps))
	for bm, handled := range r.bitmaps {
		if handled {
			continue
		}
		bms = append(bms, bm)
	}
	slices.SortFunc(bms, func(a, b *matter.Bitmap) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, bm := range bms {
		err = r.writeBitmap(e, xml.StartElement{Name: xml.Name{Local: "bitmap"}}, bm, true)
		r.bitmaps[bm] = true
		if err != nil {
			return
		}
	}
	return
}

func (r *renderer) flushEnums(e xmlEncoder) (err error) {
	ens := make([]*matter.Enum, 0, len(r.enums))
	for en, handled := range r.enums {
		if handled {
			continue
		}
		ens = append(ens, en)
	}
	slices.SortFunc(ens, func(a, b *matter.Enum) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, en := range ens {
		err = r.writeEnum(e, xml.StartElement{Name: xml.Name{Local: "enum"}}, en, true)
		delete(r.enums, en)
		if err != nil {
			return
		}
	}
	return
}

func (r *renderer) flushStructs(e xmlEncoder) (err error) {
	structs := make([]*matter.Struct, 0, len(r.structs))
	for s, handled := range r.structs {
		if handled {
			continue
		}
		structs = append(structs, s)
	}
	slices.SortFunc(structs, func(a, b *matter.Struct) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, s := range structs {
		err = r.writeStruct(e, xml.StartElement{Name: xml.Name{Local: "struct"}}, s, r.getClusterCodes(s), true)
		r.structs[s] = true
		if err != nil {
			return
		}
	}
	return
}

package amend

import (
	"encoding/xml"
	"io"
	"slices"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

func (r *renderer) writeConfigurator(dp parse.XmlDecoder, e xmlEncoder, el xml.StartElement, configurator *zap.Configurator) (err error) {

	var exampleCluster *matter.Cluster
	for c := range configurator.Clusters {
		if c != nil {
			exampleCluster = c
			break
		}
	}

	configuratorAttributes := map[string]string{
		//"domain": matter.DomainNames[r.doc.Domain],
	}

	var configuratorTokens *parse.XmlTokenSet
	configuratorTokens, err = parse.XmlExtract(dp, el)
	if err != nil {
		return
	}

	var hasCommentCharDataPending bool

	needFeatures := exampleCluster.Features != nil && len(exampleCluster.Features.Bits) > 0

	var lastSection = matter.SectionUnknown
	for {
		var tok xml.Token
		tok, err = configuratorTokens.Token()
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
					err = r.amendFeatures(configuratorTokens, e, t, exampleCluster, configurator.ClusterIDs)
					if err != nil {
						return
					}
					needFeatures = false
					break
				}
				if lastSection != matter.SectionDataTypeBitmap {
					err = r.flushUnusedConfiguratorValues(e, configurator, lastSection, configuratorAttributes)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionDataTypeBitmap
				err = r.amendBitmap(configuratorTokens, e, t, exampleCluster)
			case "enum":
				if lastSection != matter.SectionDataTypeEnum {
					err = r.flushUnusedConfiguratorValues(e, configurator, lastSection, configuratorAttributes)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionDataTypeEnum
				err = r.amendEnum(configuratorTokens, e, t, exampleCluster)

			case "struct":
				if lastSection != matter.SectionDataTypeStruct {
					err = r.flushUnusedConfiguratorValues(e, configurator, lastSection, configuratorAttributes)
					if err != nil {
						return
					}
				}

				lastSection = matter.SectionDataTypeStruct
				err = r.amendStruct(configuratorTokens, e, t, exampleCluster)

			case "cluster":
				if lastSection != matter.SectionCluster {
					err = r.flushUnusedConfiguratorValues(e, configurator, lastSection, configuratorAttributes)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionCluster
				err = r.amendCluster(configuratorTokens, e, t)
			case "clusterExtension":
				err = configuratorTokens.WriteElement(e, t)
			case "tag":
				err = configuratorTokens.WriteElement(e, t)
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
					r.writeFeatures(configuratorTokens, e, xml.StartElement{Name: xml.Name{Local: "bitmap"}}, exampleCluster, configurator.ClusterIDs)
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

func (r *renderer) flushUnusedConfiguratorValues(e xmlEncoder, configurator *zap.Configurator, lastSection matter.Section, configuratorValues map[string]string) (err error) {
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
	bms := make([]*matter.Bitmap, 0, len(r.configurator.Bitmaps))
	for bm, handled := range r.configurator.Bitmaps {
		if len(handled) == 0 {
			continue
		}
		bms = append(bms, bm)
	}
	slices.SortFunc(bms, func(a, b *matter.Bitmap) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, bm := range bms {
		err = r.writeBitmap(e, xml.StartElement{Name: xml.Name{Local: "bitmap"}}, bm, true)
		r.configurator.Bitmaps[bm] = nil
		if err != nil {
			return
		}
	}
	return
}

func (r *renderer) flushEnums(e xmlEncoder) (err error) {
	ens := make([]*matter.Enum, 0, len(r.configurator.Enums))
	for en, handled := range r.configurator.Enums {
		if len(handled) == 0 {
			continue
		}
		ens = append(ens, en)
	}
	slices.SortFunc(ens, func(a, b *matter.Enum) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, en := range ens {
		err = r.writeEnum(e, xml.StartElement{Name: xml.Name{Local: "enum"}}, en, true)
		delete(r.configurator.Enums, en)
		if err != nil {
			return
		}
	}
	return
}

func (r *renderer) flushStructs(e xmlEncoder) (err error) {
	structs := make([]*matter.Struct, 0, len(r.configurator.Structs))
	for s, handled := range r.configurator.Structs {
		if len(handled) == 0 {
			continue
		}
		structs = append(structs, s)
	}
	slices.SortFunc(structs, func(a, b *matter.Struct) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, s := range structs {
		err = r.writeStruct(e, xml.StartElement{Name: xml.Name{Local: "struct"}}, s, r.getClusterCodes(s), true)
		r.configurator.Structs[s] = nil
		if err != nil {
			return
		}
	}
	return
}

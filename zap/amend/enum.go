package amend

import (
	"encoding/xml"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

func (r *renderer) amendEnum(d xmlDecoder, e xmlEncoder, el xml.StartElement, cluster *matter.Cluster, clusterIDs []string) (err error) {
	name := getAttributeValue(el.Attr, "name")

	var matchingEnum *matter.Enum
	var skip bool
	for en, handled := range r.enums {
		if en.Name == name || strings.TrimSuffix(en.Name, "Enum") == name {
			matchingEnum = en
			skip = handled
			r.enums[en] = true
			break
		}
	}

	if matchingEnum == nil || skip {
		Ignore(d, "enum")
		return nil
	}

	var valFormat string
	el.Attr, valFormat = r.setEnumAttributes(el.Attr, matchingEnum)
	err = e.EncodeToken(el)
	if err != nil {
		return
	}

	remainingClusterIDs := make([]string, len(clusterIDs))
	copy(remainingClusterIDs, clusterIDs)

	var valueIndex int

	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = io.EOF
			return
		} else if err != nil {
			return
		}

		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "description":
				writeThrough(d, e, t)
			case "cluster":
				code := getAttributeValue(t.Attr, "code")
				id := matter.ParseID(code)
				if id.Valid() {
					ids := id.HexString()
					remainingClusterIDs = slices.DeleteFunc(remainingClusterIDs, func(s string) bool {
						return ids == s
					})
				}
				writeThrough(d, e, t)
			case "item":
				if len(remainingClusterIDs) > 0 {
					err = r.renderClusterCodes(e, remainingClusterIDs)
					if err != nil {
						return
					}
				}
				for {
					if valueIndex >= len(matchingEnum.Values) {
						Ignore(d, "item")
						break
					} else {
						v := matchingEnum.Values[valueIndex]
						valueIndex++
						if conformance.IsZigbee(v.Conformance) {
							continue
						}
						t.Attr = r.setEnumValueAttributes(v, t.Attr, valFormat)
						writeThrough(d, e, t)
						break
					}
				}

			default:

			}
		case xml.EndElement:
			switch t.Name.Local {
			case "enum":
				if len(remainingClusterIDs) > 0 {
					err = r.renderClusterCodes(e, remainingClusterIDs)
					if err != nil {
						return
					}
				}
				for valueIndex < len(matchingEnum.Values) {
					v := matchingEnum.Values[valueIndex]
					valueIndex++
					if conformance.IsZigbee(v.Conformance) {
						continue
					}
					elName := xml.Name{Local: "item"}
					xfs := xml.StartElement{Name: elName}
					xfs.Attr = r.setEnumValueAttributes(v, xfs.Attr, valFormat)
					err = e.EncodeToken(xfs)
					if err != nil {
						return
					}
					xfe := xml.EndElement{Name: elName}
					err = e.EncodeToken(xfe)
					if err != nil {
						return
					}
				}
				err = e.EncodeToken(t)
				return
			default:
				err = e.EncodeToken(tok)

			}
		case xml.CharData:
		default:
			err = e.EncodeToken(t)
		}
		if err != nil {
			return
		}
	}
}

func (r *renderer) writeEnum(e xmlEncoder, el xml.StartElement, en *matter.Enum, clusterIDs []string, provisional bool) (err error) {
	xfb := el.Copy()

	var valFormat string
	xfb.Attr, valFormat = r.setEnumAttributes(xfb.Attr, en)

	err = e.EncodeToken(xfb)
	if err != nil {
		return
	}
	err = r.renderClusterCodes(e, clusterIDs)
	if err != nil {
		return
	}

	for _, v := range en.Values {
		if conformance.IsZigbee(v.Conformance) {
			continue
		}

		elName := xml.Name{Local: "item"}
		xfs := xml.StartElement{Name: elName}

		xfs.Attr = r.setEnumValueAttributes(v, xfs.Attr, valFormat)
		err = e.EncodeToken(xfs)
		if err != nil {
			return
		}
		xfe := xml.EndElement{Name: elName}
		err = e.EncodeToken(xfe)
		if err != nil {
			return
		}

	}
	return e.EncodeToken(xml.EndElement{Name: xfb.Name})
}

func (*renderer) setEnumValueAttributes(v *matter.EnumValue, xfs []xml.Attr, valFormat string) []xml.Attr {
	val := v.Value
	valNum, er := parse.HexOrDec(val)
	if er == nil {
		val = fmt.Sprintf(valFormat, valNum)
	}

	name := zap.CleanName(v.Name)
	xfs = setAttributeValue(xfs, "name", name)
	xfs = setAttributeValue(xfs, "value", val)
	return xfs
}

func (*renderer) setEnumAttributes(xfb []xml.Attr, en *matter.Enum) ([]xml.Attr, string) {

	var valFormat string
	switch en.Type.BaseType {
	case matter.BaseDataTypeEnum16:
		valFormat = "0x%04X"
	default:
		valFormat = "0x%02X"
	}

	xfb = setAttributeValue(xfb, "name", en.Name)
	xfb = setAttributeValue(xfb, "type", zap.ConvertDataTypeNameToZap(en.Type.Name))
	return xfb, valFormat
}

package amend

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"slices"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/zap"
)

func (r *renderer) amendEnum(d xmlDecoder, e xmlEncoder, el xml.StartElement, cluster *matter.Cluster) (err error) {
	name := getAttributeValue(el.Attr, "name")

	slog.Info("checking enums", "name", name, "en", len(r.configurator.Enums), "c", len(cluster.Enums))
	var matchingEnum *matter.Enum
	var skip bool
	for en, handled := range r.configurator.Enums {
		slog.Info("checking enum", "name", name, "en", en.Name)
		if en.Name == name || strings.TrimSuffix(en.Name, "Enum") == name {
			matchingEnum = en
			skip = handled
			r.configurator.Enums[en] = true
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

	remainingClusterIDs := r.getClusterCodes(matchingEnum)

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
				err = writeThrough(d, e, t)
			case "cluster":
				code := getAttributeValue(t.Attr, "code")
				id := matter.ParseNumber(code)
				if id.Valid() {
					ids := id.HexString()
					remainingClusterIDs = slices.DeleteFunc(remainingClusterIDs, func(s string) bool {
						return ids == s
					})
				}
				err = writeThrough(d, e, t)
			case "item":
				if len(remainingClusterIDs) > 0 {
					err = r.renderClusterCodes(e, remainingClusterIDs)
					if err != nil {
						return
					}
					remainingClusterIDs = nil
				}
				for {
					if valueIndex >= len(matchingEnum.Values) {
						Ignore(d, "item")
						break
					} else {
						v := matchingEnum.Values[valueIndex]
						valueIndex++
						if conformance.IsZigbee(matchingEnum.Values, v.Conformance) {
							continue
						}
						t.Attr = r.setEnumValueAttributes(v, t.Attr, valFormat)
						err = writeThrough(d, e, t)
						if err != nil {
							return
						}
						break
					}
				}

			default:
				slog.Warn("unexpected element in enum", "name", t.Name.Local)
				err = Ignore(d, t.Name.Local)
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
					if conformance.IsZigbee(matchingEnum.Values, v.Conformance) {
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
				if err != nil {
					return
				}
				err = newLine(e)
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

func (r *renderer) writeEnum(e xmlEncoder, el xml.StartElement, en *matter.Enum, provisional bool) (err error) {
	xfb := el.Copy()

	var valFormat string
	xfb.Attr, valFormat = r.setEnumAttributes(xfb.Attr, en)

	err = e.EncodeToken(xfb)
	if err != nil {
		return
	}
	err = r.renderClusterCodes(e, r.getClusterCodes(en))
	if err != nil {
		return
	}

	for _, v := range en.Values {
		if conformance.IsZigbee(en.Values, v.Conformance) {
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
	err = e.EncodeToken(xml.EndElement{Name: xfb.Name})
	if err != nil {
		return
	}
	return newLine(e)
}

func (*renderer) setEnumValueAttributes(v *matter.EnumValue, xfs []xml.Attr, valFormat string) []xml.Attr {
	name := zap.CleanName(v.Name)
	xfs = setAttributeValue(xfs, "name", name)
	if v.Value.Valid() {
		xfs = setAttributeValue(xfs, "value", fmt.Sprintf(valFormat, v.Value.Value()))
	}
	return xfs
}

func (*renderer) setEnumAttributes(xfb []xml.Attr, en *matter.Enum) ([]xml.Attr, string) {

	var valFormat string
	switch en.Type.BaseType {
	case types.BaseDataTypeEnum16:
		valFormat = "0x%04X"
	default:
		valFormat = "0x%02X"
	}

	xfb = setAttributeValue(xfb, "name", en.Name)
	xfb = setAttributeValue(xfb, "type", zap.ConvertDataTypeNameToZap(en.Type.Name))
	return xfb, valFormat
}

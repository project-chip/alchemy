package amend

import (
	"encoding/xml"
	"io"
	"slices"
	"strings"

	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
)

func (r *renderer) amendStruct(d xmlDecoder, e xmlEncoder, el xml.StartElement, cluster *matter.Cluster, clusterIDs []string) (err error) {
	name := getAttributeValue(el.Attr, "name")

	var skip bool
	var matchingStruct *matter.Struct
	for s, handled := range r.structs {
		if s.Name == name || strings.TrimSuffix(s.Name, "Struct") == name {
			matchingStruct = s
			skip = handled
			r.structs[s] = true
			break
		}
	}

	if matchingStruct == nil || skip {
		Ignore(d, "struct")
		return
	}

	if r.errata.SeparateStructs != nil {
		if _, ok := r.errata.SeparateStructs[name]; ok {
			for _, clusterID := range clusterIDs {
				err = r.writeStruct(e, el, matchingStruct, []string{clusterID}, false)
				if err != nil {
					return
				}
			}
			return
		}
	}

	el.Attr = r.setStructAttributes(el.Attr, matchingStruct, false)
	err = e.EncodeToken(el)
	if err != nil {
		return
	}

	remainingClusterIDs := make([]string, len(clusterIDs))
	copy(remainingClusterIDs, clusterIDs)

	var fieldIndex int

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
					if fieldIndex >= len(matchingStruct.Fields) {
						Ignore(d, "item")
						break
					} else {
						f := matchingStruct.Fields[fieldIndex]
						fieldIndex++
						if conformance.IsZigbee(f.Conformance) {
							continue
						}
						t.Attr = setAttributeValue(t.Attr, "fieldId", f.ID.IntString())
						t.Attr = r.setFieldAttributes(f, t.Attr, matchingStruct.Fields)
						writeThrough(d, e, t)
						break
					}
				}

			default:

			}
		case xml.EndElement:
			switch t.Name.Local {
			case "struct":
				if len(remainingClusterIDs) > 0 {
					err = r.renderClusterCodes(e, remainingClusterIDs)
					if err != nil {
						return
					}
				}
				for fieldIndex < len(matchingStruct.Fields) {
					f := matchingStruct.Fields[fieldIndex]
					fieldIndex++
					elName := xml.Name{Local: "item"}
					xfs := xml.StartElement{Name: elName}
					xfs.Attr = setAttributeValue(xfs.Attr, "fieldId", f.ID.IntString())
					xfs.Attr = r.setFieldAttributes(f, xfs.Attr, matchingStruct.Fields)
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

func (r *renderer) writeStruct(e xmlEncoder, el xml.StartElement, s *matter.Struct, clusterIDs []string, provisional bool) (err error) {
	xfb := el.Copy()
	xfb.Name = xml.Name{Local: "struct"}
	xfb.Attr = r.setStructAttributes(xfb.Attr, s, provisional)
	err = e.EncodeToken(xfb)
	if err != nil {
		return
	}

	err = r.renderClusterCodes(e, clusterIDs)
	if err != nil {
		return
	}

	for _, v := range s.Fields {
		if conformance.IsZigbee(v.Conformance) {
			continue
		}

		elName := xml.Name{Local: "item"}
		xfs := xml.StartElement{Name: elName}
		xfs.Attr = setAttributeValue(xfs.Attr, "fieldId", v.ID.IntString())
		xfs.Attr = setAttributeValue(xfs.Attr, "name", v.Name)
		xfs.Attr = writeDataType(s.Fields, v, xfs.Attr)
		xfs.Attr = r.renderConstraint(s.Fields, v, xfs.Attr)
		if v.Quality.Has(matter.QualityNullable) {
			xfs.Attr = setAttributeValue(xfs.Attr, "isNullable", "true")
		} else {
			xfs.Attr = removeAttribute(xfs.Attr, "isNullable")
		}
		if !conformance.IsMandatory(v.Conformance) {
			xfs.Attr = setAttributeValue(xfs.Attr, "optional", "true")
		} else {
			xfs.Attr = removeAttribute(xfs.Attr, "optional")
		}
		if v.Access.FabricSensitive {
			xfs.Attr = setAttributeValue(xfs.Attr, "isFabricSensitive", "true")
		} else {
			xfs.Attr = removeAttribute(xfs.Attr, "isFabricSensitive")
		}
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

func (*renderer) setStructAttributes(xfb []xml.Attr, s *matter.Struct, provisional bool) []xml.Attr {
	xfb = setAttributeValue(xfb, "name", s.Name)
	if provisional {
		xfb = setAttributeValue(xfb, "apiMaturity", "provisional")
	}
	if s.FabricScoped {
		xfb = setAttributeValue(xfb, "isFabricScoped", "true")
	} else {
		xfb = removeAttribute(xfb, "isFabricScoped")
	}
	return xfb
}
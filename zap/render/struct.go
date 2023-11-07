package render

import (
	"encoding/xml"

	"github.com/hasty/alchemy/matter"
)

func (r *renderer) amendStruct(d xmlDecoder, e xmlEncoder, el xml.StartElement, cluster *matter.Cluster, clusterIDs []string, structs map[*matter.Struct]struct{}) (err error) {
	name := getAttributeValue(el.Attr, "name")

	var matchingStruct *matter.Struct
	for s := range structs {
		if s.Name == name {
			matchingStruct = s
			delete(structs, s)
			break
		}
	}

	if matchingStruct == nil {
		return writeThrough(d, e, el)
	}

	Ignore(d, "struct")

	return r.writeStruct(e, el, matchingStruct, clusterIDs)
}

func (r *renderer) writeStruct(e xmlEncoder, el xml.StartElement, s *matter.Struct, clusterIDs []string) (err error) {
	xfb := el.Copy()
	xfb.Name = xml.Name{Local: "struct"}
	xfb.Attr = setAttributeValue(xfb.Attr, "name", s.Name)
	err = e.EncodeToken(xfb)
	if err != nil {
		return
	}

	err = r.renderClusterCodes(e, clusterIDs)
	if err != nil {
		return
	}

	for _, v := range s.Fields {
		if v.Conformance == "Zigbee" {
			continue
		}

		elName := xml.Name{Local: "item"}
		xfs := xml.StartElement{Name: elName}
		xfs.Attr = setAttributeValue(xfs.Attr, "name", v.Name)
		xfs.Attr = writeDataType(v.Type, xfs.Attr)
		if v.Quality.Has(matter.QualityNullable) {
			xfs.Attr = setAttributeValue(xfs.Attr, "isNullable", "true")
		}
		if v.Conformance != "M" {
			xfs.Attr = setAttributeValue(xfs.Attr, "optional", "true")
		}
		xfs.Attr = r.renderConstraint(xfs.Attr, v.Constraint)
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

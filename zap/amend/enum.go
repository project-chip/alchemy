package amend

import (
	"encoding/xml"
	"fmt"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (r *renderer) amendEnum(d xmlDecoder, e xmlEncoder, el xml.StartElement, cluster *matter.Cluster, clusterIDs []string, enums map[*matter.Enum]struct{}) (err error) {
	name := getAttributeValue(el.Attr, "name")

	var matchingEnum *matter.Enum
	for en := range enums {
		if en.Name == name {
			matchingEnum = en
			delete(enums, en)
			break
		}
	}
	Ignore(d, "enum")

	if matchingEnum == nil {
		return nil
	}

	return r.writeEnum(e, el, matchingEnum, clusterIDs, false)
}

func (r *renderer) writeEnum(e xmlEncoder, el xml.StartElement, en *matter.Enum, clusterIDs []string, provisional bool) (err error) {
	xfb := el.Copy()
	xfb.Attr = setAttributeValue(xfb.Attr, "name", en.Name)
	if en.Type != "" {
		xfb.Attr = setAttributeValue(xfb.Attr, "type", en.Type)
	} else {
		xfb.Attr = setAttributeValue(xfb.Attr, "type", "enum8")
	}
	if provisional {
		xfb.Attr = setAttributeValue(xfb.Attr, "apiMaturity", "provisional")
	}

	err = e.EncodeToken(xfb)
	if err != nil {
		return
	}
	err = r.renderClusterCodes(e, clusterIDs)
	if err != nil {
		return
	}

	for _, v := range en.Values {
		if v.Conformance == "Zigbee" {
			continue
		}

		val := v.Value
		valNum, er := parse.HexOrDec(val)
		if er == nil {
			val = fmt.Sprintf("0x%02X", valNum)
		}

		elName := xml.Name{Local: "item"}
		xfs := xml.StartElement{Name: elName}
		xfs.Attr = setAttributeValue(xfs.Attr, "name", v.Name)
		xfs.Attr = setAttributeValue(xfs.Attr, "value", val)
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

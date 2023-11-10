package amend

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
	"github.com/hasty/alchemy/zap/render"
)

func writeDataType(dt *matter.DataType, attr []xml.Attr) []xml.Attr {
	if dt == nil {
		panic("missing data type")
		return attr
	}
	dts := zap.ConvertDataTypeToZap(dt.Name)
	if dt.IsArray {
		attr = setAttributeValue(attr, "type", "ARRAY")
		attr = setAttributeValue(attr, "entryType", dts)
	} else {
		attr = setAttributeValue(attr, "type", dts)
		attr = removeAttribute(attr, "entryType")
	}
	return attr
}

func writeCommandDataType(dt *matter.DataType, attr []xml.Attr) []xml.Attr {
	if dt == nil {
		return attr
	}
	dts := zap.ConvertDataTypeToZap(dt.Name)
	if dt.IsArray {
		attr = setAttributeValue(attr, "array", "true")
		attr = setAttributeValue(attr, "type", dts)
	} else {
		attr = setAttributeValue(attr, "type", dts)
		attr = removeAttribute(attr, "array")
	}
	return attr
}

func (r *renderer) writeAttribute(cluster *matter.Cluster, e xmlEncoder, el xml.StartElement, a *matter.Field, clusterPrefix string) (err error) {

	el.Name = xml.Name{Local: "attribute"}
	el.Attr = setAttributeValue(el.Attr, "code", a.ID.HexString())
	el.Attr = setAttributeValue(el.Attr, "side", "server")
	el.Attr = writeDataType(a.Type, el.Attr)
	define := render.GetDefine(a.Name, clusterPrefix, r.errata)
	el.Attr = setAttributeValue(el.Attr, "define", define)
	if a.Quality.Has(matter.QualityNullable) {
		el.Attr = setAttributeValue(el.Attr, "isNullable", "true")
	} else {
		el.Attr = removeAttribute(el.Attr, "isNullable")
	}
	if a.Quality.Has(matter.QualityReportable) {
		el.Attr = setAttributeValue(el.Attr, "reportable", "true")
	} else {
		el.Attr = removeAttribute(el.Attr, "reportable")
	}
	if a.Default != "" {
		switch a.Default {
		case "null":
			switch a.Type.Name {
			case "uint8":
				el.Attr = setAttributeValue(el.Attr, "default", "0xFF")
			case "uint16":
				el.Attr = setAttributeValue(el.Attr, "default", "0xFFFF")
			case "uint32":
				el.Attr = setAttributeValue(el.Attr, "default", "0xFFFFFFFF")
			case "uint64":
				el.Attr = setAttributeValue(el.Attr, "default", "0xFFFFFFFFFFFFFFFF")
			default:
				el.Attr = removeAttribute(el.Attr, "default")
			}
		default:
			def, e := parse.HexOrDec(a.Default)
			if e == nil {
				el.Attr = setAttributeValue(el.Attr, "default", strconv.Itoa(int(def)))
			} else {
				el.Attr = removeAttribute(el.Attr, "default")
			}
		}
	} else {
		el.Attr = removeAttribute(el.Attr, "default")
	}

	el.Attr = r.renderConstraint(cluster.Attributes, el.Attr, a.Type, a.Constraint)

	if a.Conformance != "M" {
		el.Attr = setAttributeValue(el.Attr, "optional", "true")
	} else {
		el.Attr = setAttributeValue(el.Attr, "optional", "false")
	}

	if a.Quality.Has(matter.QualityFixed) || (a.Access.Read == matter.PrivilegeUnknown || a.Access.Read == matter.PrivilegeView) && a.Access.Write == matter.PrivilegeUnknown || r.errata.SuppressAttributePermissions {
		if a.Access.Write != matter.PrivilegeUnknown {
			el.Attr = setAttributeValue(el.Attr, "writable", "true")
		} else {
			el.Attr = setAttributeValue(el.Attr, "writable", "false")
		}
		err = e.EncodeToken(el)
		if err != nil {
			return
		}
		err = e.EncodeToken(xml.CharData(a.Name))
		if err != nil {
			return
		}
	} else {
		if a.Access.Write != matter.PrivilegeUnknown {
			el.Attr = setAttributeValue(el.Attr, "writable", "true")
		} else {
			el.Attr = setAttributeValue(el.Attr, "writable", "false")
		}
		err = e.EncodeToken(el)
		if err != nil {
			return
		}

		if a.Access.Read != matter.PrivilegeUnknown {
			err = r.renderAccess(e, "read", a.Access.Read)
			if err != nil {
				return
			}
		}
		if a.Access.Write != matter.PrivilegeUnknown {
			err = r.renderAccess(e, "write", a.Access.Write)
			if err != nil {
				return
			}
		}
		elName := xml.Name{Local: "description"}
		xfs := xml.StartElement{Name: elName}
		err = e.EncodeToken(xfs)
		if err != nil {
			return
		}
		err = e.EncodeToken(xml.CharData(a.Name))
		if err != nil {
			return
		}
		xfe := xml.EndElement{Name: elName}
		err = e.EncodeToken(xfe)
		if err != nil {
			return
		}
	}

	err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "attribute"}})
	if err != nil {
		return
	}
	return
}

func (*renderer) renderConstraint(fs matter.FieldSet, attr []xml.Attr, t *matter.DataType, c matter.Constraint) []xml.Attr {
	if t == nil || t.IsArray {
		return attr
	}
	if c == nil {
		return attr
	}

	max, min := c.MinMax(&matter.ConstraintContext{Fields: fs})
	if t.IsString() {
		if max.Defined() {
			attr = setAttributeValue(attr, "length", fmt.Sprintf("%d", max.Int64))
		}
		if min.Defined() {
			attr = setAttributeValue(attr, "minLength", fmt.Sprintf("%d", min.Int64))
		}
	} else {
		if min.Defined() {
			attr = setAttributeValue(attr, "min", fmt.Sprintf("%d", min.Int64))
		}
		if max.Defined() {
			attr = setAttributeValue(attr, "max", fmt.Sprintf("%d", max.Int64))
		}
	}
	return attr
}

func (r *renderer) amendAttribute(cluster *matter.Cluster, ts *tokenSet, e xmlEncoder, el xml.StartElement, attributes map[*matter.Field]struct{}, clusterPrefix string) (err error) {
	code := getAttributeValue(el.Attr, "code")

	attributeID := matter.ParseID(code)
	if !attributeID.Valid() {
		//err = fmt.Errorf("invalid attribute code: %s", code)
		return writeThrough(ts, e, el)
	}
	var field *matter.Field
	for a := range attributes {
		if a.ID.Equals(attributeID) && a.Conformance != "Zigbee" {
			field = a
			delete(attributes, a)
			break
		}
	}

	if field == nil {
		return writeThrough(ts, e, el)
	}

	Ignore(ts, "attribute")

	if field.Conformance == "D" {
		return nil
	}

	return r.writeAttribute(cluster, e, el, field, clusterPrefix)
}

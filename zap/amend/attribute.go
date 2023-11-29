package amend

import (
	"encoding/xml"
	"io"

	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
	"github.com/hasty/alchemy/zap/render"
)

func writeAttributeDataType(fs matter.FieldSet, f *matter.Field, attr []xml.Attr) []xml.Attr {
	dts := zap.FieldToZapDataType(fs, f)
	if f.Type.IsArray() {
		attr = setAttributeValue(attr, "type", "ARRAY")
		attr = setAttributeValue(attr, "entryType", dts)
	} else {
		attr = setAttributeValue(attr, "type", dts)
		attr = removeAttribute(attr, "entryType")
	}
	return attr
}

func writeDataType(fs matter.FieldSet, f *matter.Field, attr []xml.Attr) []xml.Attr {
	if f.Type == nil {
		return attr
	}
	dts := zap.FieldToZapDataType(fs, f)
	if f.Type.IsArray() {
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
	el.Attr = r.setAttributeAttributes(el.Attr, a, cluster, clusterPrefix)

	if a.Quality.Has(matter.QualityFixed) || ((a.Access.Read == matter.PrivilegeUnknown || a.Access.Read == matter.PrivilegeView) && (a.Access.Write == matter.PrivilegeUnknown || a.Access.Write == matter.PrivilegeOperate)) || r.errata.SuppressAttributePermissions {

		err = e.EncodeToken(el)
		if err != nil {
			return
		}
		err = e.EncodeToken(xml.CharData(a.Name))
		if err != nil {
			return
		}
	} else {

		err = e.EncodeToken(el)
		if err != nil {
			return
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

		if a.Access.Read != matter.PrivilegeUnknown && a.Access.Read != matter.PrivilegeView {
			err = r.renderAccess(e, "read", a.Access.Read)
			if err != nil {
				return
			}
		}
		if a.Access.Write != matter.PrivilegeUnknown && a.Access.Write != matter.PrivilegeOperate {
			err = r.renderAccess(e, "write", a.Access.Write)
			if err != nil {
				return
			}
		}

	}

	err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "attribute"}})
	if err != nil {
		return
	}
	return
}

func (r *renderer) setAttributeAttributes(el []xml.Attr, a *matter.Field, cluster *matter.Cluster, clusterPrefix string) []xml.Attr {
	// We heard you like attributes in your attributes, so...
	el = setAttributeValue(el, "code", a.ID.HexString())
	el = setAttributeValue(el, "side", "server")
	el = writeAttributeDataType(cluster.Attributes, a, el)
	define := render.GetDefine(a.Name, clusterPrefix, r.errata)
	el = setAttributeValue(el, "define", define)
	if a.Quality.Has(matter.QualityNullable) {
		el = setAttributeValue(el, "isNullable", "true")
	} else {
		el = removeAttribute(el, "isNullable")
	}
	if a.Quality.Has(matter.QualityReportable) {
		el = setAttributeValue(el, "reportable", "true")
	} else {
		el = removeAttribute(el, "reportable")
	}
	if a.Default != "" {
		defaultValue := zap.GetDefaultValue(&matter.ConstraintContext{Field: a, Fields: cluster.Attributes})
		if defaultValue.Defined() {
			el = setAttributeValue(el, "default", defaultValue.ZapString(a.Type))
		} else {
			el = removeAttribute(el, "default")
		}
	} else {
		el = removeAttribute(el, "default")
	}

	el = r.renderConstraint(cluster.Attributes, a, el)
	if a.Quality.Has(matter.QualityFixed) || ((a.Access.Read == matter.PrivilegeUnknown || a.Access.Read == matter.PrivilegeView) && (a.Access.Write == matter.PrivilegeUnknown || a.Access.Write == matter.PrivilegeOperate)) || r.errata.SuppressAttributePermissions {
		if a.Access.Write != matter.PrivilegeUnknown {
			el = setAttributeValue(el, "writable", "true")
		} else {
			el = setAttributeValue(el, "writable", "false")
		}
		if !conformance.IsMandatory(a.Conformance) {
			el = setAttributeValue(el, "optional", "true")
		} else {
			el = setAttributeValue(el, "optional", "false")
		}
	} else {
		if a.Access.Write != matter.PrivilegeUnknown {
			el = setAttributeValue(el, "writable", "true")
		} else {
			el = setAttributeValue(el, "writable", "false")
		}
		if !conformance.IsMandatory(a.Conformance) {
			el = setAttributeValue(el, "optional", "true")
		} else {
			el = setAttributeValue(el, "optional", "false")
		}
	}
	return el
}

func (r *renderer) renderConstraint(fs matter.FieldSet, f *matter.Field, attr []xml.Attr) []xml.Attr {
	from, to := zap.GetMinMax(&matter.ConstraintContext{Field: f, Fields: fs})

	if !from.Defined() {
		attr = removeAttribute(attr, "min")
		attr = removeAttribute(attr, "minLength")
	}
	if !to.Defined() {
		attr = removeAttribute(attr, "max")
		attr = removeAttribute(attr, "length")
	}

	if f.Type != nil && (f.Type.IsString() || f.Type.IsArray()) {
		if to.Defined() {
			attr = setAttributeValue(attr, "length", to.ZapString(f.Type))
		}
		if from.Defined() {
			attr = setAttributeValue(attr, "minLength", from.ZapString(f.Type))
		}
		attr = removeAttribute(attr, "min")
		attr = removeAttribute(attr, "max")
	} else {
		if from.Defined() {
			attr = setAttributeValue(attr, "min", from.ZapString(f.Type))
		}
		if to.Defined() {
			attr = setAttributeValue(attr, "max", to.ZapString(f.Type))
		}
		attr = removeAttribute(attr, "minLength")
		attr = removeAttribute(attr, "length")
	}
	return attr
}

func (r *renderer) amendAttribute(cluster *matter.Cluster, ts *tokenSet, e xmlEncoder, el xml.StartElement, attributes map[*matter.Field]struct{}, clusterPrefix string) (err error) {
	code := getAttributeValue(el.Attr, "code")

	attributeID := matter.ParseID(code)
	if !attributeID.Valid() {
		//err = fmt.Errorf("invalid attribute code: %s", code)
		return nil
	}
	var field *matter.Field
	for a := range attributes {
		if a.ID.Equals(attributeID) && !conformance.IsZigbee(a.Conformance) {
			field = a
			delete(attributes, a)
			break
		}
	}

	if field == nil || conformance.IsDeprecated(field.Conformance) {
		Ignore(ts, "attribute")
		return nil
	}

	el.Attr = r.setAttributeAttributes(el.Attr, field, cluster, clusterPrefix)
	err = e.EncodeToken(el)
	if err != nil {
		return
	}

	needsReadAccess := field.Access.Read != matter.PrivilegeUnknown && field.Access.Read != matter.PrivilegeView
	needsWriteAccess := field.Access.Write != matter.PrivilegeUnknown && field.Access.Write != matter.PrivilegeOperate
	needsAccess := needsReadAccess || needsWriteAccess
	wroteDescription := false

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
			case "description":
				if needsAccess {
					writeThrough(ts, e, t)
				} else {
					Ignore(ts, "description")
				}
				wroteDescription = true
			case "access":
				if needsAccess && !wroteDescription {
					elName := xml.Name{Local: "description"}
					xfs := xml.StartElement{Name: elName}
					err = e.EncodeToken(xfs)
					if err != nil {
						return
					}
					err = e.EncodeToken(xml.CharData(field.Name))
					if err != nil {
						return
					}
					xfe := xml.EndElement{Name: elName}
					err = e.EncodeToken(xfe)
					if err != nil {
						return
					}
				}
				op := getAttributeValue(t.Attr, "op")
				switch op {
				case "read":
					if needsReadAccess {
						t.Attr = r.setAccessAttributes(t.Attr, op, field.Access.Read)
						needsReadAccess = false
						err = e.EncodeToken(t)
					} else {
						Ignore(ts, "access")
					}
				case "write":
					if needsWriteAccess {
						t.Attr = r.setAccessAttributes(t.Attr, op, field.Access.Write)
						needsWriteAccess = false
						err = e.EncodeToken(t)
					} else {
						Ignore(ts, "access")
					}
				}
			default:
				Ignore(ts, t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "attribute":
				if needsAccess {
					if needsReadAccess {
						err = r.renderAccess(e, "read", field.Access.Read)
						if err != nil {
							return
						}
					}
					if needsWriteAccess {
						err = r.renderAccess(e, "write", field.Access.Write)
						if err != nil {
							return
						}
					}
				} else if !wroteDescription {
					err = e.EncodeToken(xml.CharData(field.Name))
					if err != nil {
						return
					}
				}
				err = e.EncodeToken(tok)
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

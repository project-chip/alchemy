package parse

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/constraint"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/zap"
)

func readField(d *xml.Decoder, e xml.StartElement, entityType types.EntityType, name string) (field *matter.Field, err error) {
	field = matter.NewField()
	field.Access = matter.DefaultAccess(entityType)
	err = readFieldAttributes(e, field, name)
	if err != nil {
		return
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case name:
				return
			default:
				err = fmt.Errorf("unexpected %s end element: %s", name, t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected %s level type: %T", name, t)
		}
		if err != nil {
			return
		}
	}
}

func readFieldAttributes(e xml.StartElement, field *matter.Field, name string) error {
	var max, min, length, minLength string
	var fieldType, entryType string
	var isArray bool
	var optional, timed, writable string
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "id", "fieldId", "code": // Pick a lane, jeez
			field.ID = matter.ParseNumber(a.Value)
		case "name":
			field.Name = a.Value
		case "isFabricSensitive":
			if a.Value == "true" {
				field.Access.FabricSensitivity = matter.FabricSensitivitySensitive
			} else {
				field.Access.FabricSensitivity = matter.FabricSensitivityInsensitive
			}
		case "isNullable":
			if a.Value == "true" {
				field.Quality |= matter.QualityNullable
			}
		case "type":
			fieldType = a.Value
		case "max":
			max = a.Value
		case "min":
			min = a.Value
		case "entryType":
			entryType = a.Value
		case "array":
			isArray = (a.Value == "true")
		case "default", "defaut": // Ugh
			field.Default = a.Value
		case "length", "lenght": // Sigh
			length = a.Value
		case "minLength":
			length = a.Value
		case "optional":
			optional = a.Value
		case "writable":
			writable = a.Value
		case "reportable":
			if a.Value == "true" {
				field.Quality |= matter.QualityReportable
			}
		case "mustUseTimedWrite":
			timed = a.Value
		case "apiMaturity":
		case "side":
		case "define":
		case "introducedIn":
		default:
			return fmt.Errorf("unexpected %s attribute: %s", name, a.Name.Local)
		}
	}
	if optional != "true" {
		field.Conformance = conformance.Set{&conformance.Mandatory{}}
	} else {
		field.Conformance = conformance.Set{&conformance.Optional{}}
	}
	if timed == "true" {
		field.Access.Timing = matter.TimingTimed
	} else {
		field.Access.Timing = matter.TimingUntimed
	}
	if writable == "true" {
		// We don't know at this point what the write privilege is, so assume very high
		field.Access.Write = matter.PrivilegeAdminister
	}

	fieldType = zap.ConvertZapToDataTypeName(fieldType)
	entryType = zap.ConvertZapToDataTypeName(entryType)
	if isArray {
		if len(entryType) > 0 {
			field.Type = types.NewDataType(entryType, true)
		} else {
			field.Type = types.NewDataType(fieldType, true)
		}
	} else if fieldType == "ARRAY" || fieldType == "array" {
		field.Type = types.NewDataType(entryType, true)
	} else {
		field.Type = types.NewDataType(fieldType, false)
	}
	var cons string
	if len(max) > 0 && len(min) > 0 {
		cons = fmt.Sprintf("%s to %s", min, max)
	} else if len(max) > 0 {
		cons = fmt.Sprintf("max %s", max)
	} else if len(min) > 0 {
		cons = fmt.Sprintf("min %s", min)
	} else if len(length) > 0 {
		if len(minLength) > 0 {
			cons = fmt.Sprintf("%s to %s", minLength, length)
		} else {
			cons = fmt.Sprintf("max %s", length)
		}
	}
	if cons != "" {
		var err error
		field.Constraint, err = constraint.ParseString(cons)
		if err != nil {
			field.Constraint = &constraint.GenericConstraint{Value: cons}
		}
	}
	return nil
}

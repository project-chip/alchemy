package parse

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/hasty/alchemy/constraint"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func readField(d *xml.Decoder, e xml.StartElement, name string) (field *matter.Field, err error) {
	field = &matter.Field{}
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
		case xml.CharData:
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
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "id", "fieldId", "code": // Pick a lane, jeez
			field.ID = matter.ParseID(a.Value)
		case "name":
			field.Name = a.Value
		case "isFabricSensitive":
			field.Access.FabricSensitive = a.Value == "true"
		case "isNullable":
			field.Quality |= matter.QualityNullable
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
		case "default":
			field.Default = a.Value
		case "length", "lenght": // Sigh
			length = a.Value
		case "minLength":
			length = a.Value
		case "optional":
			if a.Value != "true" {
				field.Conformance = "M"
			}
		case "writable": // Ugh

		default:
			return fmt.Errorf("unexpected %s attribute: %s", name, a.Name.Local)
		}
	}
	fieldType = zap.ConvertZapToDataType(fieldType)
	entryType = zap.ConvertZapToDataType(entryType)
	if isArray {
		if len(entryType) > 0 {
			field.Type = matter.NewDataType(entryType, true)
		} else {
			field.Type = matter.NewDataType(fieldType, true)
		}
	} else if fieldType == "ARRAY" {
		field.Type = matter.NewDataType(entryType, true)
	} else {
		field.Type = matter.NewDataType(fieldType, false)
	}
	if len(max) > 0 && len(min) > 0 {
		field.Constraint = constraint.ParseConstraint(fmt.Sprintf("%s to %s", min, max))
	} else if len(max) > 0 {
		field.Constraint = constraint.ParseConstraint(fmt.Sprintf("max %s", max))
	} else if len(min) > 0 {
		field.Constraint = constraint.ParseConstraint(fmt.Sprintf("min %s", min))
	} else if len(length) > 0 {
		if len(minLength) > 0 {
			field.Constraint = constraint.ParseConstraint(fmt.Sprintf("%s to %s", minLength, length))
		} else {
			field.Constraint = constraint.ParseConstraint(fmt.Sprintf("max %s", length))
		}
	}
	return nil
}

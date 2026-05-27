package parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func readField(path string, d *xml.Decoder, e xml.StartElement, entityType types.EntityType, name string, parent types.Entity) (field *matter.Field, err error) {
	field = matter.NewField(nil, parent, entityType)
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
		case xml.StartElement:
			switch t.Name.Local {
			case "quality":
				field.Quality, err = parseQuality(d, t)
			default:
				if isConformanceElement(t) {
					var cs conformance.Conformance
					cs, err = parseConformance(d, t)
					if err == nil {
						field.Conformance = append(field.Conformance, cs)
					}
				} else {
					err = fmt.Errorf("unexpected %s level element: %s", name, t.Name.Local)
				}
			}
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

func readFieldAttributes(e xml.StartElement, field *matter.Field, name string) (err error) {
	var max, min, length, minLength string
	var fieldType, entryType string
	var fabricSensitive bool
	var rank types.DataTypeRank
	var optional, timed, writable string
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "id", "fieldId", "code": // Pick a lane, jeez
			field.ID = matter.ParseNumber(a.Value)
		case "name":
			field.Name = a.Value
		case "isFabricSensitive":
			fabricSensitive, err = strconv.ParseBool(a.Value)
			if err != nil {
				return err
			}
			if fabricSensitive {
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
			var isArray bool
			isArray, err = strconv.ParseBool(a.Value)
			if err != nil {
				return err
			}
			if isArray {
				rank = types.DataTypeRankList
			}
		case "default", "defaut": // Ugh
			field.Fallback = constraint.ParseLimit(a.Value)
		case "length", "lenght": // Sigh
			length = a.Value
		case "minLength":
			minLength = a.Value
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
		case "mustUseAtomicWrite":
			if a.Value == "true" {
				field.Quality |= matter.QualityAtomicWrite
			}
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
		field.Access.Write = matter.PrivilegeOperate
	}

	fieldBaseType := zap.ToBaseDataType(fieldType)
	entryBaseType := zap.ToBaseDataType(entryType)
	if fieldBaseType == types.BaseDataTypeList {
		switch entryBaseType {
		case types.BaseDataTypeCustom:
			field.Type = types.NewCustomDataType(entryType, types.DataTypeRankList)
		case types.BaseDataTypeUnknown:
			field.Type = types.NewNamedDataType(fieldType, fieldBaseType, types.DataTypeRankList)
		default:
			field.Type = types.NewNamedDataType(entryType, entryBaseType, types.DataTypeRankList)
		}

	} else {
		switch fieldBaseType {
		case types.BaseDataTypeCustom:
			field.Type = types.NewCustomDataType(fieldType, rank)
		default:
			field.Type = types.NewNamedDataType(fieldType, fieldBaseType, rank)
		}
	}
	var cons string
	if len(max) > 0 && len(min) > 0 {
		cons = fmt.Sprintf("%s to %s", min, max)
	} else if len(max) > 0 {
		cons = fmt.Sprintf("max %s", max)
	} else if len(min) > 0 {
		cons = fmt.Sprintf("min %s", min)
	} else if len(length) > 0 && len(minLength) > 0 {
		cons = fmt.Sprintf("%s to %s", minLength, length)
	} else if len(length) > 0 {
		cons = fmt.Sprintf("max %s", length)
	} else if len(minLength) > 0 {
		cons = fmt.Sprintf("min %s", minLength)
	}
	if cons != "" {
		field.Constraint = constraint.ParseString(cons)
	}
	return nil
}

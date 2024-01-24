package testplan

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/constraint"
	"github.com/hasty/alchemy/matter/types"
	"github.com/iancoleman/strcase"
)

func renderAttributes(cluster *matter.Cluster, b *strings.Builder) {
	if len(cluster.Attributes) == 0 {
		return
	}
	b.WriteString("==== Attributes\n\n")
	names := make([]string, 0, len(cluster.Attributes))
	var longest int
	for _, a := range cluster.Attributes {
		name := fmt.Sprintf("A_%s", strcase.ToScreamingSnake(a.Name))
		if len(name) > longest {
			longest = len(name)
		}
		names = append(names, name)
	}
	for i, name := range names {
		b.WriteString(":")
		b.WriteString(fmt.Sprintf("%-*s", longest, name))
		b.WriteString(": ")
		b.WriteString(cluster.Attributes[i].Name)
		b.WriteRune('\n')
	}
	b.WriteRune('\n')
	for i, name := range names {
		b.WriteString(fmt.Sprintf(":PICS_S%-*s : {PICS_S}.A%04x({%s})\n", longest, name, i, name))
	}
	b.WriteString("\n\n|===\n")
	b.WriteString("| *Variable* | *Description* | *Mandatory/Optional* | *Notes/Additional Constraints*\n")
	for i, a := range cluster.Attributes {
		name := names[i]
		b.WriteString(fmt.Sprintf("| {PICS_S%s} | {devimp} the _{%s}_ attribute?| ", name, name))
		renderConformance(b, cluster, cluster.Features, a.Conformance, "{PICS_SF_%s}")
		b.WriteString(" |\n")
	}
	b.WriteString("|===\n")

}

var attributesTestHeader = `
// ################# TEST CASE TEMPLATE: START #################
==== [TC-{picsCode}-2.1] Attributes with Server as DUT
===== Category
Functional conformance

===== Purpose
This test case verifies the primary functionality of the {clustername} cluster server.

===== PICS

* {PICS_S}

===== Required Devices
:reqDevices: reqDevices_C_TH_and_S_DUT
include::../common/required_devices.adoc[]

===== Device Topology

TH and DUT are on the same fabric.

===== Test Setup

{comDutTH}.

`

func renderAttributesTest(cluster *matter.Cluster, b *strings.Builder) {
	b.WriteString(attributesTestHeader)

	b.WriteString("===== Test Procedure\n")
	b.WriteString("[cols=\",,,,\"]\n")
	b.WriteString("|===\n")
	b.WriteString("| **#** | *Ref* | *PICS* | *Test Step* | *Expected Outcome* \n")
	b.WriteString("| 1 | | | {comDutTH}. |\n")

	for i, a := range cluster.Attributes {
		name := strcase.ToScreamingSnake(a.Name)

		var reply string
		var valrange string

		var dt = a.Type
		if dt != nil {

			if a.Type.IsArray() {
				dt = dt.EntryType
				reply = fmt.Sprintf(" - {DUTreply} a list of %s entries", typeString(cluster, dt))
				if a.Constraint != nil {
					switch c := a.Constraint.(type) {
					case *constraint.AllConstraint:
					case *constraint.ExactConstraint:
					case *constraint.RangeConstraint:
						valrange = fmt.Sprintf("\n - Verify that the list has between %s and %s entries", limitString(cluster, c.Minimum), limitString(cluster, c.Maximum))
					case *constraint.MinConstraint:
						valrange = fmt.Sprintf("\n - Verify that the list has %s or more entries", limitString(cluster, c.Minimum))
					case *constraint.MaxConstraint:
						valrange = fmt.Sprintf("\n - Verify that the list has no more than %s entries", limitString(cluster, c.Maximum))
					}
				}
			} else {
				reply = " {DUTreply} "
				if a.Quality.Has(matter.QualityNullable) {
					reply += "either null or "
				}
				tn := typeString(cluster, dt)
				switch tn[0] {
				case 'a', 'e', 'i', 'o', 'u': // Not perfect, but not importing a dictionary for this
					reply += "an "
				default:
					reply += "a "
				}
				reply += fmt.Sprintf("%s value.", tn)
				if a.Constraint != nil {
					switch c := a.Constraint.(type) {
					case *constraint.AllConstraint:
					case *constraint.ExactConstraint:
					case *constraint.RangeConstraint:
						valrange = fmt.Sprintf(" {valrange} %s", c.AsciiDocString(a.Type))
					}
				}
			}

		}

		b.WriteString(fmt.Sprintf("| %d  | {REF_SA_%s} | {PICS_SA_%s} | {THread} _{A_%s}_ attribute. |", i+2, name, name, name))
		if len(reply) > 0 {
			b.WriteString(reply)
		}
		if len(valrange) > 0 {
			b.WriteString(valrange)
		}
		b.WriteString("\n")
	}
	b.WriteString("|===\n")
	b.WriteString("===== Notes/Testing Considerations\n\n\n")
	b.WriteString("// ################# TEST CASE TEMPLATE: END #################\n")
}

func typeString(cluster *matter.Cluster, dt *types.DataType) string {
	switch dt.BaseType {
	case types.BaseDataTypeCustom:
		entity := cluster.Reference(dt.Name)
		switch entity := entity.(type) {
		case *matter.Enum:
			return entity.Type.Name
		case *matter.Bitmap:
			return entity.Type.Name
		}
		return dt.Name
	case types.BaseDataTypeVoltage, types.BaseDataTypePower, types.BaseDataTypeEnergy, types.BaseDataTypeAmperage:
		return "int64"
	default:
		return dt.Name
	}
}

func limitString(cluster *matter.Cluster, limit constraint.ConstraintLimit) string {
	switch limit := limit.(type) {
	case *constraint.BooleanLimit:
		return strconv.FormatBool(limit.Value)
	case *constraint.ExpLimit:
		return fmt.Sprintf("%d^%d^", limit.Value, limit.Exp)
	case *constraint.HexLimit:
		return fmt.Sprintf("0x%X", limit.Value)
	case *constraint.IntLimit:
		return strconv.FormatInt(limit.Value, 10)
	case *constraint.TemperatureLimit:
		return fmt.Sprintf("%s°C", limit.Value.String())
	case *constraint.ReferenceLimit:
		ref := cluster.Reference(limit.Value)
		if ref == nil {
			return fmt.Sprintf("ERR: unknown reference %s", limit.Value)
		}
		switch ref := ref.(type) {
		case *matter.Field:
			return fmt.Sprintf("{A_%s}", strcase.ToScreamingSnake(ref.Name))
		default:
			return fmt.Sprintf("ERR: unknown reference type %T (%s)", ref, limit.Value)
		}
	default:
		return "unknown limit"
	}
}

/*
**#** |*Ref*                                     |*Expected Outcome*
|1     |                                          |
|2     | {REF_PWR_SA_POWER_MODE}                  | {DUTreply} a enum8 value. {valrange} 0 and 1.
|3     | {REF_PWR_SA_NUMBER_OF_MEASUREMENT_TYPES} | {DUTreply} a uint8 value. {valrange} 1 and 255.
|4     | {REF_PWR_SA_ACCURACY}                    | - {DUTreply} a list of MeasurementAccuracyStruct entries
                                                    - Verify that the list has one or more entries
|5     | {REF_PWR_SA_RANGES}                      | {DUTreply} a list of MeasurementRangeStruct entries
                                                    - Verify that the list has less than {A_NUMBER_OF_MEASUREMENT_TYPES} entries
|6     | {REF_PWR_SA_VOLTAGE}                     | {DUTreply} either null or an int64 value. {valrange} -2^62^ to 2^62^.
|7     | {REF_PWR_SA_ACTIVE_CURRENT}              | {DUTreply} either null or an int64 value. {valrange} -2^62^ to 2^62^.
|8     | {REF_PWR_SA_REACTIVE_CURRENT}            | {DUTreply} either null or an int64 value. {valrange} -2^62^ to 2^62^.
|9     | {REF_PWR_SA_APPARENT_CURRENT}            | {DUTreply} either null or an int64 value. {valrange} 0 to 2^62^.
|10    | {REF_PWR_SA_ACTIVE_POWER}                | {DUTreply} either null or an int64 value. {valrange} -2^62^ to 2^62^.
|11    | {REF_PWR_SA_REACTIVE_POWER}              | {DUTreply} either null or an int64 value. {valrange} -2^62^ to 2^62^.
|12    | {REF_PWR_SA_APPARENT_POWER}              | {DUTreply} either null or an int64 value. {valrange} -2^62^ to 2^62^.
|13    | {REF_PWR_SA_RMS_VOLTAGE}                 | {DUTreply} either null or an int64 value. {valrange} -2^62^ to 2^62^.
|14    | {REF_PWR_SA_RMS_CURRENT}                 | {DUTreply} either null or an int64 value. {valrange} -2^62^ to 2^62^.
|15    | {REF_PWR_SA_RMS_POWER}                   | {DUTreply} either null or an int64 value. {valrange} -2^62^ to 2^62^.
|16    | {REF_PWR_SA_FREQUENCY}                   | {DUTreply} either null or an int64 value. {valrange} 0 to 1000000.
|17    | {REF_PWR_SA_HARMONIC_CURRENTS}           | {DUTreply} a list of HarmonicMeasurementStruct entries
|18    | {REF_PWR_SA_HARMONIC_PHASES}             | {DUTreply} a list of HarmonicMeasurementStruct entries
|19    | {REF_PWR_SA_POWER_FACTOR}                | {DUTreply} either null or an int64 value. {valrange} -10000 to 10000.
|20    | {REF_PWR_SA_NEUTRAL_CURRENT}             | {DUTreply} either null or an int64 value. {valrange} -2^62^ to 2^62^.
|===
*/
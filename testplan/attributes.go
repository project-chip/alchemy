package testplan

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func renderAttributes(doc *spec.Doc, cluster *matter.Cluster, b *strings.Builder) {
	if len(cluster.Attributes) == 0 {
		return
	}
	b.WriteString("==== Attributes\n\n\n")
	names := make([]string, 0, len(cluster.Attributes))
	var longest int
	for _, a := range cluster.Attributes {
		name := entityIdentifier(a)
		if len(name) > longest {
			longest = len(name)
		}
		names = append(names, name)
	}
	for i, name := range names {
		b.WriteString(":")
		b.WriteString(fmt.Sprintf("%-*s", longest, name))
		b.WriteString(" : ")
		b.WriteString(cluster.Attributes[i].Name)
		b.WriteRune('\n')
	}
	b.WriteString("\n\n")
	for i, name := range names {
		b.WriteString(fmt.Sprintf(":PICS_S%-*s : {PICS_S}.A%04x({%s})\n", longest, name, i, name))
	}
	b.WriteString("\n\n|===\n")
	b.WriteString("| *Variable* | *Description* | *Mandatory/Optional* | *Notes/Additional Constraints*\n")
	for i, a := range cluster.Attributes {
		name := names[i]
		b.WriteString(fmt.Sprintf("| {PICS_S%s} | {devimp} the _{%s}_ attribute?| ", name, name))
		if len(a.Conformance) > 0 {
			b.WriteString("{PICS_S}: ")
			renderPicsConformance(b, doc, cluster, a.Conformance)
		}
		b.WriteString(" |\n")
	}
	b.WriteString("|===\n\n")

}

var attributesTestHeader = `
// ################# TEST CASE TEMPLATE: START #################
==== [TC-{picsCode}-2.1] Attributes with Server as DUT
===== Category
Functional conformance

===== Purpose
This test case verifies the non-global attributes of the {clustername} cluster server.

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
	b.WriteString("[cols=\"5%,5%,10%,40%,40%\"]\n")
	b.WriteString("|===\n")
	b.WriteString("| **#** | *Ref* | *PICS* | *Test Step* | *Expected Outcome* \n")
	b.WriteString("| 1 | | | {comDutTH}. |\n")

	for i, a := range cluster.Attributes {
		name := entityIdentifier(a)

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
				firstLetter, _ := utf8.DecodeRuneInString(tn)
				switch unicode.ToLower(firstLetter) {
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
						valrange = fmt.Sprintf(" {valrange} %s", c.ASCIIDocString(a.Type))
					}
				}
			}

		}

		b.WriteString(fmt.Sprintf("| %d  | {REF_%s_S%s} | {PICS_S%s} | {THread} _{%s}_ attribute. |", i+2, cluster.PICS, name, name, name))
		if len(reply) > 0 {
			b.WriteString(reply)
		}
		if len(valrange) > 0 {
			b.WriteString(valrange)
		}
		b.WriteString("\n")
	}
	b.WriteString("|===\n\n")
	b.WriteString("===== Notes/Testing Considerations\n\n\n")
	b.WriteString("// ################# TEST CASE TEMPLATE: END #################\n")
}

func typeString(cluster *matter.Cluster, dt *types.DataType) string {
	switch dt.BaseType {
	case types.BaseDataTypeCustom:
		entity, ok := cluster.Identifier(dt.Name)
		if !ok {
			return dt.Name
		}
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

func limitString(cluster *matter.Cluster, limit constraint.Limit) string {
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
		ref, ok := cluster.Identifier(limit.Value)
		if !ok {
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

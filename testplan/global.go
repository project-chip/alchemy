package testplan

import (
	"fmt"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/iancoleman/strcase"
)

var globalHeader = `
== Test Cases
// ################# TEST CASE TEMPLATE: START #################
'''
=== Generic Test Cases
'''
// ################# GLOBAL ATTRIBUTES TEST CASE: START #################
==== [TC-{picsCode}-1.1] Global Attributes with {DUT_Server}

===== Category
Functional conformance.

===== Purpose
This test case verifies the behavior of the global attributes of the cluster server.

===== PICS

* {PICS_S}

===== Required Devices
:reqDevices: reqDevices_C_TH_and_S_DUT
include::../common/required_devices.adoc[]

`

func renderGlobalAttributesTestCase(cluster *matter.Cluster, b *strings.Builder) (err error) {
	b.WriteString(globalHeader)

	b.WriteString("===== Test Procedure\n\n")
	b.WriteString("\n[cols=\",,,,\"]\n")
	b.WriteString("|===\n")
	b.WriteString("| **#** | *Ref* | *PICS* | *Test Step* | *Expected Outcome* \n")
	b.WriteString("| 1 | | | {comDutTH}. |\n")
	step := 1
	if len(cluster.Revisions) > 0 {
		revision := cluster.Revisions[len(cluster.Revisions)-1]
		step++
		b.WriteString(fmt.Sprintf("| %d | {REF_CLUSTERREVISION} | | {THread} _ClusterRevision_ attribute. | {DUTreply} the _ClusterRevision_ attribute and has the value %s.\n", step, revision.Number))
	}
	if cluster.Features != nil && len(cluster.Features.Bits) > 0 {
		step++
		b.WriteString(fmt.Sprintf("| %d | {REF_FEATUREMAP} | | {THread} _FeatureMap_ attribute. | {DUTreply} the _FeatureMap_ attribute and have the following bits set:\n", step))
		for _, bit := range cluster.Features.Bits {
			var from, to uint64

			from, to, err = bit.Bits()
			if err != nil {
				return
			}
			for i := from; i <= to; i++ {
				b.WriteString(fmt.Sprintf("- bit %d: {shallBeOneIff} {PICS_SF_%s}\n", i, bit.Code))
			}
		}
		b.WriteString("{remainingBitsZero}\n")
	}
	step++
	b.WriteString(fmt.Sprintf("| %d | {REF_ATTRIBUTELIST} |  | {THread} _AttributeList_ attribute.  | {DUTreply} the _AttributeList_ attribute and have the list of supported attributes\n", step))
	if len(cluster.Attributes) > 0 {

		var mandatory, optional, feature []*matter.Field
		expressions := make(map[*matter.Field]conformance.Expression)
		optionality := make(map[*matter.Field]bool)
		for _, a := range cluster.Attributes {
			for _, c := range a.Conformance {
				switch c := c.(type) {
				case *conformance.Mandatory:
					optionality[a] = false
					if c.Expression == nil {
						mandatory = append(mandatory, a)
						break
					}
					feature = append(feature, a)
					expressions[a] = c.Expression

				case *conformance.Optional:
					optionality[a] = true
					if c.Expression == nil {
						optional = append(optional, a)
						continue
					}
					feature = append(feature, a)
					expressions[a] = c.Expression
				default:
					continue
				}
				break
			}
		}
		if len(mandatory) > 0 {
			b.WriteString("{mandatoryEntries} +\n")
			for i, a := range mandatory {
				if i > 0 {
					b.WriteString(", ")
				}
				b.WriteString(a.ID.HexString())
			}
			b.WriteString("\n{attsGlobalNumbers}.\n")
		}
		for _, a := range optional {
			b.WriteString("{optionalEntries} +\n")
			b.WriteString(fmt.Sprintf("- %s: {shallIncludeIff} {PICS_SA_%s}\n", a.ID.HexString(), strcase.ToScreamingSnake(a.Name)))
		}
		if len(feature) > 0 {
			b.WriteString("{featureEntries} +\n")
			for _, a := range feature {
				b.WriteString(fmt.Sprintf("- %s: {shallIncludeIf} ", a.ID.HexString()))
				exp := expressions[a]
				if optionality[a] {
					b.WriteString("(")
					renderExpression(b, cluster, exp, "{PICS_SF_%s}")
					b.WriteString(fmt.Sprintf(" & {PICS_SA_%s})", strcase.ToScreamingSnake(a.Name)))
				} else {
					renderExpression(b, cluster, exp, "{PICS_SF_%s}")
				}
				b.WriteString(" and {shallNotInclude}.\n")
			}
		}
	}
	b.WriteString(`
|5^*^  | {REF_EVENTLIST}            |       | {THread} _EventList_ attribute.            | {DUTreply} the _EventList_ attribute and have the list of supported events:
{noEntryStdRgn} +
|6     | {REF_ACCEPTEDCOMMANDLIST}  |       | {THread} _AcceptedCommandList_ attribute.  | {DUTreply} the _AcceptedCommandList_ attribute and have the list of Accepted Command:
{noEntryStdRgn} +
|7     | {REF_GENERATEDCOMMANDLIST} |       | {THread} _GeneratedCommandList_ attribute. | {DUTreply} the _GeneratedCommandList_ attribute and have the list of Generated Command:
{noEntryStdRgn} +
|===

===== Notes/Testing Considerations
^*^ Step 5 is currently not supported and SHALL be skipped.
// ################# GLOBAL ATTRIBUTES TEST CASE: END #################

// ################# TEST CASE TEMPLATE: END #################
`)
	return
}





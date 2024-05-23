package testplan

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/spec"
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

func renderGlobalAttributesTestCase(doc *spec.Doc, cluster *matter.Cluster, b *strings.Builder) (err error) {
	b.WriteString(globalHeader)

	b.WriteString("===== Test Procedure\n")
	b.WriteString("[cols=\"5%,5%,10%,40%,40%\"]\n")
	b.WriteString("|===\n")
	b.WriteString("| **#** | *Ref* | *PICS* | *Test Step* | *Expected Outcome* \n")
	b.WriteString("| 1 | | | {comDutTH}. |\n")
	if len(cluster.Revisions) > 0 {
		revision := cluster.Revisions[len(cluster.Revisions)-1]
		b.WriteString(fmt.Sprintf("| 2 | {REF_CLUSTERREVISION} | | {THread} _ClusterRevision_ attribute. | {DUTreply} the _ClusterRevision_ attribute and has the value %s.\n", revision.Number))
	}
	if cluster.Features != nil && len(cluster.Features.Bits) > 0 {
		b.WriteString("| 3 | {REF_FEATUREMAP} | | {THread} _FeatureMap_ attribute. | {DUTreply} the _FeatureMap_ attribute and have the following bits set: \n")
		for _, bit := range cluster.Features.Bits {
			f := bit.(*matter.Feature)
			var from, to uint64

			from, to, err = bit.Bits()
			if err != nil {
				return
			}
			for i := from; i <= to; i++ {
				b.WriteString(fmt.Sprintf("- bit %d: {shallBeOneIff} {PICS_SF_%s}\n", i, f.Code))
			}
		}
		b.WriteString("+\n{remainingBitsZero}\n")
	}
	writeAttributeListAttribute(b, doc, cluster)
	writeEventListAttribute(b, doc, cluster)
	b.WriteString(`
| 6     | {REF_ACCEPTEDCOMMANDLIST}  |       | {THread} _AcceptedCommandList_ attribute.  | {DUTreply} the _AcceptedCommandList_ attribute and have the list of Accepted Command:
{noEntryStdRgn} +
| 7     | {REF_GENERATEDCOMMANDLIST} |       | {THread} _GeneratedCommandList_ attribute. | {DUTreply} the _GeneratedCommandList_ attribute and have the list of Generated Command:
{noEntryStdRgn} +
|===

===== Notes/Testing Considerations
^*^ Step 5 is currently not supported and SHALL be skipped.
// ################# GLOBAL ATTRIBUTES TEST CASE: END #################

// ################# TEST CASE TEMPLATE: END #################
`)
	return
}

func writeAttributeListAttribute(b *strings.Builder, doc *spec.Doc, cluster *matter.Cluster) {
	b.WriteString("| 4 | {REF_ATTRIBUTELIST} |  | {THread} _AttributeList_ attribute.  | {DUTreply} the _AttributeList_ attribute and have the list of supported attributes\n")
	if len(cluster.Attributes) == 0 {
		b.WriteString("{noEntryStdRgn} +\n")
		return
	}

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
		b.WriteString("{mandatoryEntries} +\n - ")
		for i, a := range mandatory {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(a.ID.HexString())
		}
		b.WriteString(" +\n{attsGlobalNumbers}.\n")
	}
	for _, a := range optional {
		b.WriteString("+ \n{optionalEntries} +\n ")
		b.WriteString(fmt.Sprintf("- %s: {shallIncludeIff} {PICS_S%s}\n", a.ID.HexString(), entityIdentifier(a)))
	}
	if len(feature) > 0 {
		b.WriteString(" + \n{featureEntries} +\n")
		for _, a := range feature {
			b.WriteString(fmt.Sprintf("- %s: {shallIncludeIf} ", a.ID.HexString()))
			var conf conformance.Set
			exp := expressions[a]
			if optionality[a] {
				conf = append(conf, &conformance.Mandatory{
					Expression: &conformance.LogicalExpression{
						Operand: "&", Left: exp,
						Right: []conformance.Expression{&conformance.IdentifierExpression{
							ID: a.Name,
						}},
					},
				})
			} else {
				conf = append(conf, &conformance.Mandatory{Expression: exp})
			}
			renderPicsConformance(b, doc, cluster, conf)
			b.WriteString(" and {shallNotInclude}.\n")
		}
	}

}

func writeEventListAttribute(b *strings.Builder, doc *spec.Doc, cluster *matter.Cluster) {
	b.WriteString("| 5^*^ | {REF_EVENTLIST} | | {THread} _EventList_ attribute. | {DUTreply} the _EventList_ attribute and have the list of supported events:\n")
	if len(cluster.Events) == 0 {
		b.WriteString("{noEntryStdRgn} +\n")
		return
	}

	var mandatory, optional, feature []*matter.Event
	expressions := make(map[*matter.Event]conformance.Expression)
	optionality := make(map[*matter.Event]bool)
	for _, event := range cluster.Events {
		for _, c := range event.Conformance {
			switch c := c.(type) {
			case *conformance.Mandatory:
				optionality[event] = false
				if c.Expression == nil {
					mandatory = append(mandatory, event)
					break
				}
				feature = append(feature, event)
				expressions[event] = c.Expression

			case *conformance.Optional:
				optionality[event] = true
				if c.Expression == nil {
					optional = append(optional, event)
					continue
				}
				feature = append(feature, event)
				expressions[event] = c.Expression

			default:
				slog.Warn("Unable to determine conformance for event", slog.String("clusterName", cluster.Name), slog.String("eventName", event.Name), slog.Any("conformance", c))
				continue
			}
			break
		}
	}
	if len(mandatory) > 0 {
		b.WriteString("{mandatoryEntries} +\n")
		for i, event := range mandatory {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(event.ID.HexString())
		}
		b.WriteString("+\n")
	}
	for _, event := range optional {
		b.WriteString("{optionalEntries} +\n")
		b.WriteString(fmt.Sprintf("- %s: {shallIncludeIff} {PICS_S_%s} +\n", event.ID.HexString(), entityIdentifier(event)))
	}
	if len(feature) > 0 {
		b.WriteString("{featureEntries} +\n")
		for _, event := range feature {
			b.WriteString(fmt.Sprintf("- %s: {shallIncludeIf} ", event.ID.HexString()))
			var conf conformance.Set
			conf = append(conf, event.Conformance...)
			if optionality[event] {
				conf = append(conf, &conformance.Mandatory{Expression: &conformance.IdentifierExpression{ID: event.Name}})
			}
			renderPicsConformance(b, doc, cluster, conf)
			b.WriteString(" and {shallNotInclude}. +\n")
		}
	}

}

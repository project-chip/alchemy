package testplan

import (
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
)

type clusterUnderTest struct {
	cluster           *matter.Cluster
	features          []*matter.Feature
	attributes        []*matter.Field
	commandsAccepted  []*matter.Command
	commandsGenerated []*matter.Command
	events            []*matter.Event
}

func renderClusterTestPlan(doc *spec.Doc, cluster *matter.Cluster) (output string, err error) {
	cut := filterCluster(doc, cluster)
	var out strings.Builder
	err = renderHeader(cluster, &out)
	if err != nil {
		return
	}
	err = renderServer(doc, cut, &out)
	if err != nil {
		return
	}
	out.WriteString(testCases)
	err = renderGlobalAttributesTestCase(doc, cut, &out)
	if err != nil {
		return
	}
	renderAttributesTest(cluster, &out)
	output = out.String()
	return
}

func filterCluster(doc *spec.Doc, cluster *matter.Cluster) *clusterUnderTest {
	cut := &clusterUnderTest{
		cluster:    cluster,
		attributes: make([]*matter.Field, 0, len(cluster.Attributes)),
	}

	if cluster.Features != nil {
		for _, bit := range cluster.Features.Bits {
			if !checkConformance(bit.Conformance(), cluster.Features) {
				continue
			}

			f := bit.(*matter.Feature)
			cut.features = append(cut.features, f)
		}
	}
	for _, attribute := range cluster.Attributes {
		if !checkConformance(attribute.Conformance, cluster.Attributes) {
			continue
		}

		cut.attributes = append(cut.attributes, attribute)
	}

	for _, command := range cluster.Commands {
		if !checkConformance(command.Conformance, cluster.Features) {
			continue
		}
		switch command.Direction {
		case matter.InterfaceServer:
			cut.commandsAccepted = append(cut.commandsAccepted, command)
		}
	}

	for _, command := range cluster.Commands {
		if !checkConformance(command.Conformance, cluster.Features) {
			continue
		}
		switch command.Direction {
		case matter.InterfaceClient:
			for _, c := range cut.commandsAccepted {
				if c.Response == command.Name {
					cut.commandsGenerated = append(cut.commandsGenerated, command)
					break
				}
			}
		}
	}

	for _, event := range cluster.Events {
		if !checkConformance(event.Conformance, cluster.Attributes) {
			continue
		}
		cut.events = append(cut.events, event)
	}
	return cut
}

func checkConformance(c conformance.Set, store conformance.IdentifierStore) bool {
	return !(conformance.IsZigbee(store, c) || conformance.IsDisallowed(c) || conformance.IsDeprecated(c))
}

var testCases = `== Test Case List

|===
| *TC UUID*         | *Test Case Name*
| TC-{picsCode}-1.1 | Global Attributes with {DUT_Server}
| TC-{picsCode}-2.1 | Attributes with Server as DUT
| TC-{picsCode}-2.2 | Primary Functionality with Server as DUT
|===


`

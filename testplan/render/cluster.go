package render

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
)

type clusterUnderTest struct {
	cluster           *matter.Cluster
	features          []*feature
	attributes        []*matter.Field
	commandsAccepted  []*matter.Command
	commandsGenerated []*matter.Command
	events            []*matter.Event
}

func filterCluster(doc *spec.Doc, cluster *matter.Cluster) (cut *clusterUnderTest, err error) {
	cut = &clusterUnderTest{
		cluster:    cluster,
		attributes: make([]*matter.Field, 0, len(cluster.Attributes)),
	}

	if cluster.Features != nil {
		for _, bit := range cluster.Features.Bits {
			if !checkConformance(bit.Conformance(), cluster.Features) {
				continue
			}

			feat := bit.(*matter.Feature)
			f := &feature{Code: feat.Code, Summary: feat.Summary(), Conformance: feat.Conformance()}
			f.From, f.To, err = feat.Bits()
			if err != nil {
				return
			}
			if f.From <= f.To {
				for i := f.From; i <= f.To; i++ {
					f.Bits = append(f.Bits, i)
				}
			}
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
				if c.Response != nil && c.Response.Name == command.Name {
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
	return
}

func checkConformance(c conformance.Set, store conformance.IdentifierStore) bool {
	return !(conformance.IsZigbee(store, c) || conformance.IsDisallowed(c) || conformance.IsDeprecated(c))
}

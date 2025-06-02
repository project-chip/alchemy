package testplan

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
)

type Plan struct {
	Cluster           *matter.Cluster
	Features          []*Feature
	Attributes        []*matter.Field
	CommandsAccepted  []*matter.Command
	CommandsGenerated []*matter.Command
	Events            []*matter.Event
}

func NewPlan(doc *spec.Doc, cluster *matter.Cluster) (p *Plan, err error) {
	p = &Plan{
		Cluster:    cluster,
		Attributes: make([]*matter.Field, 0, len(cluster.Attributes)),
	}

	if cluster.Features != nil {
		for _, bit := range cluster.Features.Bits {
			if !checkConformance(bit.Conformance()) {
				continue
			}

			feat := bit.(*matter.Feature)
			f := &Feature{Code: feat.Code, Summary: feat.Summary(), Conformance: feat.Conformance()}
			f.From, f.To, err = feat.Bits()
			if err != nil {
				return
			}
			if f.From <= f.To {
				for i := f.From; i <= f.To; i++ {
					f.Bits = append(f.Bits, i)
				}
			}
			p.Features = append(p.Features, f)
		}
	}
	for _, attribute := range cluster.Attributes {
		if !checkConformance(attribute.Conformance) {
			continue
		}

		p.Attributes = append(p.Attributes, attribute)
	}

	for _, command := range cluster.Commands {
		if !checkConformance(command.Conformance) {
			continue
		}
		switch command.Direction {
		case matter.InterfaceServer:
			p.CommandsAccepted = append(p.CommandsAccepted, command)
		}
	}

	for _, command := range cluster.Commands {
		if !checkConformance(command.Conformance) {
			continue
		}
		switch command.Direction {
		case matter.InterfaceClient:
			for _, c := range p.CommandsAccepted {
				if c.Response != nil && c.Response.Name == command.Name {
					p.CommandsGenerated = append(p.CommandsGenerated, command)
					break
				}
			}
		}
	}

	for _, event := range cluster.Events {
		if !checkConformance(event.Conformance) {
			continue
		}
		p.Events = append(p.Events, event)
	}
	return
}

func checkConformance(c conformance.Set) bool {
	return !(conformance.IsZigbee(c) || conformance.IsDisallowed(c) || conformance.IsDeprecated(c))
}

func (p *Plan) ToTests() (tests []*Test) {
	t := &Test{}
	tests = append(tests, t)
	return
}

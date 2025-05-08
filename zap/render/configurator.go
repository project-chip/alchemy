package render

import (
	"bytes"
	"log/slog"
	"regexp"
	"slices"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

type configuratorRenderer struct {
	configurator *zap.Configurator
	elementMap   map[*etree.Element]types.Entity
	generator    *TemplateGenerator
}

func newConfiguratorRenderer(generator *TemplateGenerator, configurator *zap.Configurator) *configuratorRenderer {
	return &configuratorRenderer{
		generator:    generator,
		configurator: configurator,
		elementMap:   make(map[*etree.Element]types.Entity),
	}
}

func (cr *configuratorRenderer) render(x *etree.Document, exampleCluster *matter.Cluster) (out string, err error) {

	err = cr.patchComments(cr.configurator, x)
	if err != nil {
		return
	}

	configuratorElement := x.SelectElement("configurator")
	if configuratorElement == nil {
		configuratorElement = x.CreateElement("configurator")
	}
	configuratorElement.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	configuratorElement.CreateAttr("xsi:noNamespaceSchemaLocation", "../../zcl.xsd")

	domainElement := configuratorElement.SelectElement("domain")
	if domainElement == nil {
		domainElement = etree.NewElement("domain")
		xml.AppendElement(configuratorElement, domainElement)
		domainElement.CreateAttr("name", cr.configurator.Domain)
	}

	if exampleCluster != nil {
		err = cr.generateFeatures(configuratorElement, exampleCluster.Features)
		if err != nil {
			return
		}
	}

	err = cr.generateBitmaps(cr.configurator.Bitmaps, configuratorElement)
	if err != nil {
		return
	}

	err = cr.generateEnums(cr.configurator.Enums, configuratorElement)
	if err != nil {
		return
	}

	err = cr.generateStructs(cr.configurator.Structs, configuratorElement)
	if err != nil {
		return
	}

	err = cr.renderClusters(configuratorElement)
	if err != nil {
		return
	}
	if cr.generator.options.SpecOrder && !cr.configurator.Global {
		err = cr.reorderConfigurator(configuratorElement)
		if err != nil {
			return
		}
	}
	return xmlToString(x)
}

func xmlToString(x *etree.Document) (string, error) {
	x.Indent(2)
	var b bytes.Buffer
	_, err := x.WriteTo(&b)
	if err != nil {
		return "", err
	}
	s := b.String()
	s = postProcessTemplate(s)
	return s, nil
}

var tagClosePattern = regexp.MustCompile(`(?m)/(?P<Tag>bitmap|cluster|command|enum|event|struct)>\n(\s+)<`)

func postProcessTemplate(s string) string {
	// etree removes extraneous whitespace between tags, so this restores it for commonly separated tags in ZAP templates
	s = tagClosePattern.ReplaceAllString(s, "/$1>\n\n$2<")
	return s
}

func amendExistingClusterCodes(parent *etree.Element, entity types.Entity, clusterIDs []*matter.Number) (amendedCodes []*matter.Number, remainingCodes []*matter.Number) {
	clusterCodeElements := parent.SelectElements("cluster")
	remainingCodes = make([]*matter.Number, len(clusterIDs))
	copy(remainingCodes, clusterIDs)
	for _, cce := range clusterCodeElements {
		ca := cce.SelectAttr("code")
		if ca == nil {
			slog.Debug("missing cluster code", "val", entity)
			parent.RemoveChild(cce)
			continue
		}
		id := matter.ParseNumber(ca.Value)

		if !id.Valid() || !matter.ContainsNumber(clusterIDs, id) {
			parent.RemoveChild(cce)
			continue
		}
		amendedCodes = append(amendedCodes, id)
		remainingCodes = slices.DeleteFunc(remainingCodes, func(s *matter.Number) bool { return s.Equals(id) })
	}
	return
}

func flushClusterCodes(parent *etree.Element, clusterIDs []*matter.Number) {
	for _, clusterID := range clusterIDs {
		if !clusterID.Valid() {
			continue
		}
		ce := etree.NewElement("cluster")
		patchNumberAttribute(ce, clusterID, "code")
		xml.AppendElement(parent, ce, "cluster")
	}
}

package render

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

type NamespacePatcher struct {
	sdkRoot string
	spec    *spec.Specification
}

func NewNamespacePatcher(sdkRoot string, spec *spec.Specification) *NamespacePatcher {
	dtp := &NamespacePatcher{sdkRoot: sdkRoot, spec: spec}
	return dtp
}

func (p NamespacePatcher) Name() string {
	return "Patching namespaces"
}

func (p NamespacePatcher) Process(cxt context.Context, inputs []*pipeline.Data[[]*matter.Namespace]) (outputs []*pipeline.Data[[]byte], err error) {

	namespaceXMLPath := filepath.Join(p.sdkRoot, "/src/app/zap-templates/zcl/data-model/chip/semantic-tag-namespace-enums.xml")

	var namespaceXML []byte
	namespaceXML, err = os.ReadFile(namespaceXMLPath)
	if err != nil {
		return
	}

	doc := etree.NewDocument()
	err = doc.ReadFromBytes(namespaceXML)
	if err != nil {
		return
	}

	configurator := doc.SelectElement("configurator")
	if configurator == nil {
		err = fmt.Errorf("missing configurator element in %s", namespaceXMLPath)
		return
	}

	namespacesByName := make(map[string]*matter.Namespace)
	for _, input := range inputs {
		for _, dt := range input.Content {
			namespacesByName[matterNamespaceName(dt)] = dt
		}
	}

	enumElements := configurator.SelectElements("enum")
	for _, enumElement := range enumElements {
		name := enumElement.SelectAttrValue("name", "")
		ns, ok := namespacesByName[name]
		if !ok {
			continue
		}
		enumElement.Child = nil
		for _, val := range ns.SemanticTags {
			ve := enumElement.CreateElement("item")
			ve.CreateAttr("value", val.ID.ShortHexString())
			ve.CreateAttr("name", val.Name)
		}
		delete(namespacesByName, name)
	}

	for name, ns := range namespacesByName {
		nse := etree.NewElement("enum")
		nse.CreateAttr("name", name)
		nse.CreateAttr("type", "enum8")
		for _, val := range ns.SemanticTags {
			ve := nse.CreateElement("item")
			ve.CreateAttr("value", val.ID.ShortHexString())
			ve.CreateAttr("name", val.Name)
		}
		xml.InsertElementByAttribute(configurator, nse, "name")
	}

	var out string
	doc.Indent(2)
	out, err = doc.WriteToString()
	if err != nil {
		return
	}
	out = postProcessTemplate(out)
	outputs = append(outputs, pipeline.NewData(namespaceXMLPath, []byte(out)))

	return
}

func matterNamespaceName(ns *matter.Namespace) string {
	name := strings.TrimSpace(text.TrimCaseInsensitivePrefix(text.TrimCaseInsensitiveSuffix(ns.Name, " Namespace"), "Common "))
	name = matter.Case(name)
	if name == "Area" {
		// Backwards compatibility for one namespace
		name += "Type"
	}
	return name + "Tag"
}

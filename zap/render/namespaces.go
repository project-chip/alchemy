package render

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/internal/vcs"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/zap"
)

// namespaceIDEnumName names the enum that lists every standard namespace ID.
const namespaceIDEnumName = "SemanticTagNamespaceId"

type NamespacePatcher struct {
	sdkRoot     string
	spec        *spec.Specification
	specVersion string
}

func NewNamespacePatcher(sdkRoot string, specification *spec.Specification) *NamespacePatcher {
	p := &NamespacePatcher{sdkRoot: sdkRoot, spec: specification}
	if specification != nil && specification.Root != "" {
		var err error
		p.specVersion, err = vcs.GitDescribe(specification.Root)
		if err != nil {
			slog.Error("Unable to determine spec git tag", slog.Any("error", err))
		}
	}
	return p
}

func (p NamespacePatcher) Name() string {
	return "Patching namespaces"
}

// Process regenerates the whole namespace file from every namespace in the specification, so the
// namespace ID enum is complete even when only some namespace documents were passed in.
func (p NamespacePatcher) Process(cxt context.Context, inputs []*pipeline.Data[*asciidoc.Document]) (outputs []*pipeline.Data[[]byte], err error) {

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

	namespaces := slices.Clone(p.spec.Namespaces)
	slices.SortStableFunc(namespaces, func(a, b *matter.Namespace) int {
		return a.ID.Compare(b.ID)
	})

	enumNames := make(map[string]*matter.Namespace, len(namespaces))
	for _, ns := range namespaces {
		name := matterNamespaceName(p.spec, ns)
		if other, ok := enumNames[name]; ok {
			err = fmt.Errorf("namespaces %q and %q both map to enum %s; add an sdk type-names override to the errata of one of them", other.Name, ns.Name, name)
			return
		}
		enumNames[name] = ns
	}

	for _, name := range removeEnums(configurator, enumNames) {
		slog.Warn("removing enum that no spec namespace generates", slog.String("path", namespaceXMLPath), slog.String("enum", name))
	}

	// Each namespace ID is named after its tag enum, so kPosition pairs with PositionTag.
	idEnum := configurator.CreateElement("enum")
	idEnum.CreateAttr("name", namespaceIDEnumName)
	idEnum.CreateAttr("type", "enum8")
	for _, ns := range namespaces {
		ve := idEnum.CreateElement("item")
		ve.CreateAttr("value", ns.ID.ShortHexString())
		ve.CreateAttr("name", strings.TrimSuffix(matterNamespaceName(p.spec, ns), "Tag"))
	}

	var docs []*asciidoc.Document
	for _, ns := range namespaces {
		enumName := matterNamespaceName(p.spec, ns)
		nsErrata := namespaceErrata(p.spec, ns)
		nse := configurator.CreateElement("enum")
		nse.CreateAttr("name", enumName)
		nse.CreateAttr("type", "enum8")
		for _, val := range ns.SemanticTags {
			ve := nse.CreateElement("item")
			ve.CreateAttr("value", val.ID.ShortHexString())
			ve.CreateAttr("name", semanticTagName(nsErrata, enumName, val))
		}
		if ref, ok := p.spec.DocRefs[ns]; ok && !slices.Contains(docs, ref) {
			docs = append(docs, ref)
		}
	}

	cr := &configuratorRenderer{
		generator: &TemplateGenerator{specVersion: p.specVersion},
	}
	err = cr.patchComments(&zap.Configurator{Docs: docs, OutPath: namespaceXMLPath}, doc)
	if err != nil {
		return
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

// removeEnums removes every enum from the configurator, and returns the names of the ones that
// neither a spec namespace nor the namespace ID enum generates.
func removeEnums(configurator *etree.Element, generated map[string]*matter.Namespace) (unknown []string) {
	for _, enumElement := range configurator.SelectElements("enum") {
		name := enumElement.SelectAttrValue("name", "")
		if _, ok := generated[name]; !ok && name != namespaceIDEnumName {
			unknown = append(unknown, name)
		}
		configurator.RemoveChild(enumElement)
	}
	return
}

func namespaceBaseName(ns *matter.Namespace) string {
	return strings.TrimSpace(text.TrimCaseInsensitiveSuffix(ns.Name, " Namespace"))
}

// matterNamespaceName names a namespace's tag enum. "Common " is dropped from the name unless
// another namespace already uses the shorter name, as "Closure" does for "Common Closure".
func matterNamespaceName(specification *spec.Specification, ns *matter.Namespace) string {
	baseName := namespaceBaseName(ns)
	shortName := strings.TrimSpace(text.TrimCaseInsensitivePrefix(baseName, "Common "))
	if shortName != baseName {
		for _, other := range specification.Namespaces {
			if other != ns && strings.EqualFold(namespaceBaseName(other), shortName) {
				shortName = baseName
				break
			}
		}
	}
	name := matter.Case(shortName)
	if name == "Area" {
		// Backwards compatibility for one namespace
		name += "Type"
	}
	name = name + "Tag"
	if override, ok := namespaceErrata(specification, ns).SDK.TypeNames[name]; ok {
		return override
	}
	return name
}

// namespaceErrata returns the errata of the document that defines the namespace.
func namespaceErrata(specification *spec.Specification, ns *matter.Namespace) *errata.Errata {
	if doc, ok := specification.DocRefs[ns]; ok {
		return specification.Errata.Get(doc.Path.Relative)
	}
	return errata.DefaultErrata
}

// semanticTagName names a tag in the tag enum enumName. An errata field of that enum whose name is
// the tag's spec name replaces it with its override-name, verbatim.
func semanticTagName(nsErrata *errata.Errata, enumName string, tag *matter.SemanticTag) string {
	if nsErrata.SDK.Types == nil {
		return tag.Name
	}
	override, ok := nsErrata.SDK.Types.Enums[enumName]
	if !ok || override == nil {
		return tag.Name
	}
	for _, field := range override.Fields {
		if field.Name == tag.Name && field.OverrideName != "" {
			return field.OverrideName
		}
	}
	return tag.Name
}

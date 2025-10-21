package dm

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter"
)

func getNamespacePath(dmRoot string, path asciidoc.Path, namespaceName string) string {
	p := path.Base()
	file := strings.TrimSuffix(p, path.Ext())
	if len(namespaceName) > 0 {
		file += "-" + namespaceName
	}
	return filepath.Join(dmRoot, fmt.Sprintf("/namespaces/%s.xml", file))
}

func renderNamespace(doc *asciidoc.Document, namespace *matter.Namespace) (output string, err error) {

	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(getLicense())
	nse := x.CreateElement("namespace")
	nse.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	nse.CreateAttr("xsi:schemaLocation", "types types.xsd namespace namespace.xsd")
	if namespace.ID != nil && namespace.ID.Valid() {
		nse.CreateAttr("id", namespace.ID.HexString())
	}
	nse.CreateAttr("name", namespace.Name)

	if len(namespace.SemanticTags) > 0 {
		tagse := nse.CreateElement("tags")
		for _, tag := range namespace.SemanticTags {
			tage := tagse.CreateElement("tag")
			if tag.ID != nil && tag.ID.Valid() {
				tage.CreateAttr("id", tag.ID.HexString())
			}
			tage.CreateAttr("name", tag.Name)
			if tag.Description != "" {
				tage.CreateElement("description").SetText(tag.Description)
			}
		}
	}

	x.Indent(2)

	var b bytes.Buffer
	_, err = x.WriteTo(&b)
	output = b.String()

	return
}

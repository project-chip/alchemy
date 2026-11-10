package spec

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

// referenceFinder resolves cross-reference IDs from a fixed map.
type referenceFinder struct {
	entityFinderCommon
	entities map[string]types.Entity
}

func (rf *referenceFinder) findEntityByReference(reference string, label string, source log.Source) types.Entity {
	return rf.entities[reference]
}

func (rf *referenceFinder) findEntityByIdentifier(identifier string, source log.Source) types.Entity {
	return nil
}

func (rf *referenceFinder) suggestIdentifiers(identifier string, suggestions map[types.Entity]int) {
}

func TestTagFieldNamespaceFromNamespaceReference(t *testing.T) {
	root := t.TempDir()
	for name, contents := range map[string]string{
		"main.adoc": "include::namespaces/Namespace-Common-Landmark.adoc[]\n\n" +
			"[[ref_LandmarkTagField]]\n== LandmarkTag Field\n\n" +
			"This field SHALL be the ID of a landmark semantic tag, located within the <<ref_CommonLandmarkNamespace>>.\n",
		"namespaces/Namespace-Common-Landmark.adoc": "[[ref_CommonLandmarkNamespace]]\n= Common Landmark Semantic Tag Namespace\n\n" +
			"|===\n| ID | Namespace\n| 0x11 | Common Landmark Namespace\n|===\n\n" +
			"|===\n| ID | Name\n| 0x00 | AirConditioner\n|===\n",
	} {
		path := filepath.Join(root, name)
		if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(contents), 0600); err != nil {
			t.Fatal(err)
		}
	}
	document, err := readFile(filepath.Join(root, "main.adoc"), root)
	if err != nil {
		t.Fatal(err)
	}
	library := NewLibrary(document, config.Library{}, nil, NewDocCache(root))
	if err := library.Preparse(root, nil); err != nil {
		t.Fatal(err)
	}
	library.indexCrossReferences()
	if err := library.indexAnchors(); err != nil {
		t.Fatal(err)
	}

	if !library.isNamespaceAnchor("ref_CommonLandmarkNamespace") {
		t.Error("ref_CommonLandmarkNamespace is not recognized as a namespace anchor")
	}
	if library.isNamespaceAnchor("ref_LandmarkTagField") {
		t.Error("ref_LandmarkTagField is recognized as a namespace anchor")
	}

	anchors := library.anchors["ref_CommonLandmarkNamespace"]
	if len(anchors) != 1 {
		t.Fatalf("found %d ref_CommonLandmarkNamespace anchors, want 1", len(anchors))
	}
	namespaceSection, ok := anchors[0].Element.(*asciidoc.Section)
	if !ok {
		t.Fatalf("ref_CommonLandmarkNamespace anchors a %T, want a section", anchors[0].Element)
	}
	ns, err := library.toNamespace(library, anchors[0].Document, namespaceSection)
	if err != nil || ns == nil {
		t.Fatalf("toNamespace returned %v, %v", ns, err)
	}
	if entities, _ := library.EntitiesForElement(namespaceSection); !slices.Contains(entities, types.Entity(ns)) {
		t.Error("the namespace section does not resolve to the namespace entity")
	}

	fieldSection := library.anchors["ref_LandmarkTagField"][0].Element.(*asciidoc.Section)
	field := &matter.Field{Name: "LandmarkTag", Type: types.NewDataType(types.BaseDataTypeTag, types.DataTypeRankScalar)}
	if err := library.findTagNamespace(library, document, fieldSection, field, map[string]*matter.Field{}); err != nil {
		t.Fatal(err)
	}
	var reference *asciidoc.CrossReference
	parse.Traverse(document, library, fieldSection, library.Children(fieldSection), func(doc *asciidoc.Document, ref *asciidoc.CrossReference, parent asciidoc.ParentElement, index int) parse.SearchShould {
		reference = ref
		return parse.SearchShouldStop
	})
	if field.Type.Source != reference {
		t.Fatalf("findTagNamespace set the type source to %v, want the namespace reference", field.Type.Source)
	}

	builder := &Builder{Spec: &Specification{}}
	builder.getTagNamespace(library, field, &referenceFinder{entities: map[string]types.Entity{"ref_CommonLandmarkNamespace": ns}})
	if field.Type.Entity != ns {
		t.Errorf("getTagNamespace resolved %v, want the Common Landmark namespace", field.Type.Entity)
	}
}

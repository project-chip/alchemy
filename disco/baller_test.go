package disco

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func TestDiscoBallTopLevelSection_AddXrefstyle(t *testing.T) {
	b := &Baller{
		options: DiscoOptions{AddXrefstyle: true},
	}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	workspaceRoot := filepath.Dir(wd)

	rootDoc, err := spec.ReadFile("testdata/root.adoc", workspaceRoot)
	if err != nil {
		t.Fatalf("failed to read root doc: %v", err)
	}
	includeDoc, err := spec.ReadFile("testdata/include.adoc", workspaceRoot)
	if err != nil {
		t.Fatalf("failed to read include doc: %v", err)
	}

	lib := &spec.Library{Root: rootDoc}

	// Test Root Doc
	{
		top := parse.FindFirst[*asciidoc.Section](rootDoc, asciidoc.RawReader, rootDoc)
		if top == nil {
			t.Fatalf("no top level section found in root doc")
		}

		dc := &discoContext{
			Context: context.Background(),
			doc:     rootDoc,
			library: lib,
		}

		err = b.discoBallTopLevelSection(dc, top, matter.DocTypeCluster)
		if err != nil {
			t.Fatalf("unexpected error on root doc: %v", err)
		}

		found := false
		for _, el := range rootDoc.Elements {
			if ae, ok := el.(*asciidoc.AttributeEntry); ok && ae.Name == "xrefstyle" {
				found = true
				break
			}
		}
		if !found {
			t.Error("expected xrefstyle to be added to root document")
		}
	}

	// Test Include Doc
	{
		top := parse.FindFirst[*asciidoc.Section](includeDoc, asciidoc.RawReader, includeDoc)
		if top == nil {
			t.Fatalf("no top level section found in include doc")
		}

		dc := &discoContext{
			Context: context.Background(),
			doc:     includeDoc,
			library: lib,
		}

		err = b.discoBallTopLevelSection(dc, top, matter.DocTypeCluster)
		if err != nil {
			t.Fatalf("unexpected error on include doc: %v", err)
		}

		found := false
		for _, el := range includeDoc.Elements {
			if ae, ok := el.(*asciidoc.AttributeEntry); ok && ae.Name == "xrefstyle" {
				found = true
				break
			}
		}
		if found {
			t.Error("expected xrefstyle to NOT be added to include document")
		}
	}
}

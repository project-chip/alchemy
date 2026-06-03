package errata

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/project-chip/alchemy/config"
)

func TestLoadErrataWithOverlay(t *testing.T) {
	tmpDir := t.TempDir()

	mainErrataContent := `
errata:
  app_clusters/SomeCluster.adoc:
    sdk:
      cluster-name: "OriginalName"
      type-names:
        OldType: "NewType"
`
	overlayErrataContent := `
errata:
  app_clusters/SomeCluster.adoc:
    sdk:
      cluster-name: "OverlayName"
      type-names:
        OverlayOld: "OverlayNew"
        OldType: "OverriddenType"
`

	mainPath := filepath.Join(tmpDir, "main_errata.yaml")
	overlayPath := filepath.Join(tmpDir, "overlay_errata.yaml")

	if err := os.WriteFile(mainPath, []byte(mainErrataContent), 0644); err != nil {
		t.Fatalf("failed to write main errata: %v", err)
	}

	if err := os.WriteFile(overlayPath, []byte(overlayErrataContent), 0644); err != nil {
		t.Fatalf("failed to write overlay errata: %v", err)
	}

	// Create a fake config with root pointing to tmpDir
	cfg := &config.Config{}

	col, err := LoadErrata(cfg, mainPath, overlayPath)
	if err != nil {
		t.Fatalf("LoadErrata failed: %v", err)
	}

	e := col.Get("app_clusters/SomeCluster.adoc")
	if e == nil {
		t.Fatalf("expected errata for app_clusters/SomeCluster.adoc, got nil")
	}

	// Verify overriden fields with overlay higher priority
	if e.SDK.ClusterName != "OverlayName" {
		t.Errorf("expected ClusterName to be 'OverlayName', got '%s'", e.SDK.ClusterName)
	}

	if val, ok := e.SDK.TypeNames["OldType"]; !ok || val != "OverriddenType" {
		t.Errorf("expected OldType to be 'OverriddenType', got '%s'", val)
	}

	if val, ok := e.SDK.TypeNames["OverlayOld"]; !ok || val != "OverlayNew" {
		t.Errorf("expected OverlayOld to be 'OverlayNew', got '%s'", val)
	}
}

func TestLoadErrataSDKTypesMerge(t *testing.T) {
	tmpDir := t.TempDir()

	mainErrataContent := `
errata:
  app_clusters/SomeCluster.adoc:
    sdk:
      types:
        attributes:
          SomeAttribute:
            override-name: "MainOverride"
            conformance: "O"
            access: "read"
            fields:
              - name: "FieldA"
                type: "int"
        commands:
          SomeCommand:
            conformance: "M"
      extra-types:
        structs:
          SomeExtraStruct:
            name: "MainExtraStruct"
`
	overlayErrataContent := `
errata:
  app_clusters/SomeCluster.adoc:
    sdk:
      types:
        attributes:
          SomeAttribute:
            override-name: "OverlayOverride"
            conformance: "M"
            fields:
              - name: "FieldA"
                type: "string"
              - name: "FieldB"
                type: "bool"
          AnotherAttribute:
            conformance: "D"
        commands:
          SomeCommand:
            conformance: "O"
            response: "SomeResponse"
      extra-types:
        structs:
          SomeExtraStruct:
            name: "OverlayExtraStruct"
            override-name: "OverlayExtraOverride"
`

	mainPath := filepath.Join(tmpDir, "main_errata.yaml")
	overlayPath := filepath.Join(tmpDir, "overlay_errata.yaml")

	if err := os.WriteFile(mainPath, []byte(mainErrataContent), 0644); err != nil {
		t.Fatalf("failed to write main errata: %v", err)
	}

	if err := os.WriteFile(overlayPath, []byte(overlayErrataContent), 0644); err != nil {
		t.Fatalf("failed to write overlay errata: %v", err)
	}

	cfg := &config.Config{}

	col, err := LoadErrata(cfg, mainPath, overlayPath)
	if err != nil {
		t.Fatalf("LoadErrata failed: %v", err)
	}

	e := col.Get("app_clusters/SomeCluster.adoc")
	if e == nil {
		t.Fatalf("expected errata, got nil")
	}

	// 1. Verify types.attributes["SomeAttribute"] merge
	someAttr, ok := e.SDK.Types.Attributes["SomeAttribute"]
	if !ok {
		t.Fatalf("expected SomeAttribute to exist in attributes")
	}
	if someAttr.OverrideName != "OverlayOverride" {
		t.Errorf("expected override-name to be 'OverlayOverride', got '%s'", someAttr.OverrideName)
	}
	if someAttr.Conformance != "M" {
		t.Errorf("expected conformance to be 'M', got '%s'", someAttr.Conformance)
	}
	if someAttr.Access != "read" {
		t.Errorf("expected access to be retained as 'read', got '%s'", someAttr.Access)
	}
	if len(someAttr.Fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(someAttr.Fields))
	}
	if someAttr.Fields[0].Name != "FieldA" || someAttr.Fields[0].Type != "string" {
		t.Errorf("expected field 0 to be FieldA/string, got %s/%s", someAttr.Fields[0].Name, someAttr.Fields[0].Type)
	}
	if someAttr.Fields[1].Name != "FieldB" || someAttr.Fields[1].Type != "bool" {
		t.Errorf("expected field 1 to be FieldB/bool, got %s/%s", someAttr.Fields[1].Name, someAttr.Fields[1].Type)
	}

	// 2. Verify types.attributes["AnotherAttribute"] creation
	anotherAttr, ok := e.SDK.Types.Attributes["AnotherAttribute"]
	if !ok {
		t.Fatalf("expected AnotherAttribute to exist")
	}
	if anotherAttr.Conformance != "D" {
		t.Errorf("expected conformance to be 'D', got '%s'", anotherAttr.Conformance)
	}

	// 3. Verify types.commands["SomeCommand"] merge
	someCmd, ok := e.SDK.Types.Commands["SomeCommand"]
	if !ok {
		t.Fatalf("expected SomeCommand to exist")
	}
	if someCmd.Conformance != "O" {
		t.Errorf("expected conformance to be 'O', got '%s'", someCmd.Conformance)
	}
	if someCmd.Response != "SomeResponse" {
		t.Errorf("expected response to be 'SomeResponse', got '%s'", someCmd.Response)
	}

	// 4. Verify extra-types.structs["SomeExtraStruct"] merge
	extraStruct, ok := e.SDK.ExtraTypes.Structs["SomeExtraStruct"]
	if !ok {
		t.Fatalf("expected SomeExtraStruct to exist")
	}
	if extraStruct.Name != "OverlayExtraStruct" {
		t.Errorf("expected name to be 'OverlayExtraStruct', got '%s'", extraStruct.Name)
	}
	if extraStruct.OverrideName != "OverlayExtraOverride" {
		t.Errorf("expected override-name to be 'OverlayExtraOverride', got '%s'", extraStruct.OverrideName)
	}
}

package provisional_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/provisional"
)

func TestCompareProvisional_FileBased_Cases(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError})))

	cxt := context.Background()
	baseRoot := "testdata/base"
	headRoot := "testdata/head"

	processingOptions := pipeline.ProcessingOptions{}

	violations, err := provisional.Pipeline(cxt, baseRoot, headRoot, nil, processingOptions, nil)
	if err != nil {
		t.Fatalf("Pipeline failed: %v", err)
	}

	// We read test_cluster.adoc which contains 4 cases:
	//
	// Case 1: ProvCommand1 and Field1
	// - ProvCommand1 is an existing provisional command (matched by ID 0x00 to base).
	// - Field1 is a NEW non-provisional field added to it, inside if-def.
	// - Field1 triggers NonProvisional violation, but it is MASKED by provisional parent ProvCommand1.
	// - Result: PASSES. Verifies ancestor resolving logic.
	//
	// Case 2: ProvCommand2
	// - Marked provisional (P) in table.
	// - NOT inside ifdef::in-progress[].
	// - Result: FAILS (NotIfDefd). Provisional entities must be in ifdef.
	//
	// Case 3: ProvCommand3
	// - Marked provisional (P) in table.
	// - Inside ifdef::in-progress[].
	// - Result: PASSES. It is properly ifdef'd.
	//
	// Case 4: NonProvCommand
	// - Marked mandatory (M) (not provisional) in table.
	// - Inside ifdef::in-progress[].
	// - Result: FAILS (NonProvisional). Non-provisional entities should not be in in-progress ifdef.

	expectedNotIfDefd := 0
	expectedNonProvisional := 0

	for _, vs := range violations {
		for _, v := range vs {
			if v.Type.Has(spec.ViolationTypeNotIfDefd) {
				expectedNotIfDefd++
			}
			if v.Type.Has(spec.ViolationTypeNonProvisional) {
				expectedNonProvisional++
			}
		}
	}

	// We expect 1 NotIfDefd: for ProvCommand2 (no ifdef).
	if expectedNotIfDefd != 1 {
		t.Errorf("Expected 1 NotIfDefd violations, got %d", expectedNotIfDefd)
	}

	// We expect 1 NonProvisional: for NonProvCommand (Case 4: not provisional in ifdef).
	if expectedNonProvisional != 1 {
		t.Errorf("Expected 1 NonProvisional violation, got %d", expectedNonProvisional)
	}
}

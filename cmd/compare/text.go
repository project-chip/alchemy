package compare

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/project-chip/alchemy/compare"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func writeText(w io.Writer, diffs []*compare.ClusterDifferences) {
	for _, cd := range diffs {
		writeClusterDifference(w, cd)
	}
}

func writeClusterDifference(w io.Writer, cd *compare.ClusterDifferences) {
	writeEntityDiff(w, 0, &cd.IdentifiedDiff, types.EntityTypeCluster)
	writeEntityDiffs(w, 1, cd.Attributes, types.EntityTypeAttribute)
	writeEntityDiffs(w, 1, cd.Features, types.EntityTypeFeature)
	writeEntityDiffs(w, 1, cd.Bitmaps, types.EntityTypeBitmap)
	writeEntityDiffs(w, 1, cd.Enums, types.EntityTypeEnum)
	writeEntityDiffs(w, 1, cd.Structs, types.EntityTypeStruct)
	writeEntityDiffs(w, 1, cd.Events, types.EntityTypeEvent)
	writeEntityDiffs(w, 1, cd.Commands, types.EntityTypeCommand)
}

func writeEntityDiffs(w io.Writer, indent int, diffs []compare.Diff, entityType types.EntityType) {
	prefix := fmt.Sprintf("%*s", indent, "\t")
	var missing []*compare.MissingDiff
	var identified []*compare.IdentifiedDiff
	for _, d := range diffs {
		switch d := d.(type) {
		case *compare.IdentifiedDiff:
			identified = append(identified, d)
		case *compare.MissingDiff:
			missing = append(missing, d)
		default:
			fmt.Fprintf(w, "\tunrecognized %s diff: %T:\n", entityType, d)
		}
	}
	if len(missing) > 0 {
		for _, m := range missing {
			writeMissingDiff(w, m, prefix)
		}
		fmt.Fprintln(w)
	}
	if len(identified) > 0 {
		slices.SortFunc(identified, func(a *compare.IdentifiedDiff, b *compare.IdentifiedDiff) int {
			return strings.Compare(a.Name, b.Name)
		})
		for _, m := range identified {
			writeIdentifiedDiffs(indent, m, w)
		}
		fmt.Fprintln(w)
	}
}

func writeEntityDiff(w io.Writer, indent int, d compare.Diff, entityType types.EntityType) {
	prefix := strings.Repeat("\t", indent)
	switch d := d.(type) {
	case *compare.IdentifiedDiff:
		writeIdentifiedDiffs(indent, d, w)
	case *compare.MissingDiff:
		writeMissingDiff(w, d, prefix)
	default:
		fmt.Fprintf(w, "\tunrecognized %s diff: %T:\n", entityType, d)
	}
}

func writeIdentifiedDiffs(indent int, id *compare.IdentifiedDiff, w io.Writer) {
	prefix := strings.Repeat("\t", indent)
	fmt.Fprintf(w, "%s%s %s:\n", prefix, id.Name, id.Entity)
	var identified []*compare.IdentifiedDiff
	var others []compare.Diff
	for _, idd := range id.Diffs {
		switch idd := idd.(type) {
		case *compare.IdentifiedDiff:
			identified = append(identified, idd)
		default:
			others = append(others, idd)
		}
	}
	for _, idd := range others {
		switch idd := idd.(type) {
		case *compare.MissingDiff:
			writeMissingDiff(w, idd, prefix)
		case *compare.StringDiff:
			writeStringDiff(w, idd, id.Entity, id.Name, prefix)
		case *compare.PropertyDiff[matter.Privilege]:
			writePrivilegeDiff(w, idd, prefix, id.Entity, id.Name)
		case *compare.PropertyDiff[matter.FabricSensitivity]:
			writeSensitivityDiff(w, idd, prefix, id.Entity, id.Name)
		case *compare.PropertyDiff[matter.FabricScoping]:
			writeScopingDiff(w, idd, prefix, id.Entity, id.Name)
		case *compare.PropertyDiff[matter.Timing]:
			writeTimingDiff(w, idd, prefix, id.Entity, id.Name)
		case *compare.ConformanceDiff:
			writeConformanceDiff(w, idd, id.Entity, id.Name, prefix)
		case *compare.BoolDiff:
			writeBoolDiff(w, idd, id.Entity, id.Name, prefix)
		default:
			fmt.Fprintf(w, "%sunrecognized identified diff: %T:\n", prefix, idd)
		}
	}
	if len(identified) > 0 {
		slices.SortFunc(identified, func(a *compare.IdentifiedDiff, b *compare.IdentifiedDiff) int {
			return strings.Compare(a.Name, b.Name)
		})
		for _, idd := range identified {
			writeIdentifiedDiffs(indent+1, idd, w)
		}
	}
	if indent < 2 {
		fmt.Fprintf(w, "\n")
	}

}

func writePrivilegeDiff(w io.Writer, ad *compare.PropertyDiff[matter.Privilege], prefix string, entityType types.EntityType, name string) {
	fmt.Fprintf(w, "%s%s %s has %s access %s, but is %s in the spec\n", prefix, name, entityType.String(), ad.Property, ad.ZAP, ad.Spec)
}

func writeScopingDiff(w io.Writer, ad *compare.PropertyDiff[matter.FabricScoping], prefix string, entityType types.EntityType, name string) {
	fmt.Fprintf(w, "%s%s %s has fabric scoping %s, but is %s in the spec\n", prefix, name, entityType.String(), ad.ZAP, ad.Spec)
}

func writeSensitivityDiff(w io.Writer, ad *compare.PropertyDiff[matter.FabricSensitivity], prefix string, entityType types.EntityType, name string) {
	fmt.Fprintf(w, "%s%s %s has fabric sensitivity %s, but is %s in the spec\n", prefix, name, entityType.String(), ad.ZAP, ad.Spec)
}

func writeTimingDiff(w io.Writer, ad *compare.PropertyDiff[matter.Timing], prefix string, entityType types.EntityType, name string) {
	fmt.Fprintf(w, "%s%s %s has timing %s, but is %s in the spec\n", prefix, name, entityType.String(), ad.ZAP, ad.Spec)
}

func writeMissingDiff(w io.Writer, md *compare.MissingDiff, prefix string) {
	switch md.Property {
	case compare.DiffPropertyUnknown:
		switch md.Source {
		case compare.SourceSpec:
			fmt.Fprintf(w, "%s%s %s is missing in the spec\n", prefix, md.Name, md.Entity)
		case compare.SourceZAP:
			fmt.Fprintf(w, "%s%s %s is missing in the ZAP template\n", prefix, md.Name, md.Entity)
		}
	default:
		switch md.Source {
		case compare.SourceSpec:
			fmt.Fprintf(w, "%s%s %s is missing %s in the spec\n", prefix, md.Name, md.Entity, md.Property.String())
		case compare.SourceZAP:
			fmt.Fprintf(w, "%s%s %s is missing %s in the ZAP template\n", prefix, md.Name, md.Entity, md.Property.String())
		}

	}
}

func writeStringDiff(w io.Writer, sd *compare.StringDiff, entityType types.EntityType, name string, prefix string) {
	zapValue := sd.ZAP
	specValue := sd.Spec
	switch sd.Property {
	case compare.DiffPropertyDefault, compare.DiffPropertyMax, compare.DiffPropertyLength, compare.DiffPropertyMin, compare.DiffPropertyMinLength:
		if zapValue == "" {
			zapValue = "not set"
		}
		if specValue == "" {
			specValue = "not set"
		}
	default:
		zapValue = fmt.Sprintf("\"%s\"", zapValue)
		specValue = fmt.Sprintf("\"%s\"", specValue)
	}
	fmt.Fprintf(w, "%s%s %s %s is %s, but should be %s\n", prefix, name, entityType, sd.Property.String(), zapValue, specValue)
}

func writeBoolDiff(w io.Writer, sd *compare.BoolDiff, entityType types.EntityType, name string, prefix string) {
	fmt.Fprintf(w, "%s%s %s %s is %v, but should be %v\n", prefix, name, entityType, sd.Property.String(), sd.ZAP, sd.Spec)
}

func writeConformanceDiff(w io.Writer, sd *compare.ConformanceDiff, entityType types.EntityType, name string, prefix string) {
	if sd.ZAP == conformance.StateMandatory {
		if len(sd.SpecConfornance) > 0 {
			mc, ok := sd.SpecConfornance[0].(*conformance.Mandatory)
			if ok {
				if mc.Expression != nil {
					fmt.Fprintf(w, "%s%s %s %s is marked as mandatory, but is only mandatory when %s\n", prefix, name, entityType, sd.Property.String(), mc.Expression.Description())
					return
				}
			}
		}
	}
	fmt.Fprintf(w, "%s%s %s %s is marked as %s, but should be %s\n", prefix, name, entityType, sd.Property.String(), sd.ZAP, sd.Spec)
}

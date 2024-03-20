package compare

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/hasty/alchemy/compare"
	"github.com/hasty/alchemy/matter/types"
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
		slices.SortFunc[[]*compare.IdentifiedDiff](identified, func(a *compare.IdentifiedDiff, b *compare.IdentifiedDiff) int {
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
	if len(identified) > 0 {
		slices.SortFunc[[]*compare.IdentifiedDiff](identified, func(a *compare.IdentifiedDiff, b *compare.IdentifiedDiff) int {
			return strings.Compare(a.Name, b.Name)
		})
		for _, idd := range identified {
			writeIdentifiedDiffs(indent+1, idd, w)
		}
	}
	for _, idd := range others {
		switch idd := idd.(type) {
		case *compare.MissingDiff:
			writeMissingDiff(w, idd, prefix)
		case *compare.StringDiff:
			writeStringDiff(w, idd, id.Entity, id.Name, prefix)
		case *compare.AccessDiff:
			writeAccessDiff(w, idd, prefix, id.Entity, id.Name)
		case *compare.ConformanceDiff:
			writeConformanceDiff(w, idd, id.Entity, id.Name, prefix)
		case *compare.BoolDiff:
			writeBoolDiff(w, idd, id.Entity, id.Name, prefix)
		default:
			fmt.Fprintf(w, "%suncreognized identified diff: %T:\n", prefix, idd)
		}
	}
	if indent < 2 {
		fmt.Fprintf(w, "\n")
	}

}

func writeAccessDiff(w io.Writer, ad *compare.AccessDiff, prefix string, entityType types.EntityType, name string) {
	if ad.Spec.Read != ad.ZAP.Read {
		fmt.Fprintf(w, "%s%s %s has read access %s, but is %s in the spec\n", prefix, name, entityType.String(), ad.ZAP.Read, ad.Spec.Read)
	}
	if ad.Spec.Write != ad.ZAP.Write {
		fmt.Fprintf(w, "%s%s %s has write access %s, but is %s in the spec\n", prefix, name, entityType.String(), ad.ZAP.Write, ad.Spec.Write)
	}
	if ad.Spec.OptionalWrite != ad.ZAP.OptionalWrite {
		fmt.Fprintf(w, "%s%s %s has optional write %v, but is %v in the spec\n", prefix, name, entityType.String(), ad.ZAP.OptionalWrite, ad.Spec.OptionalWrite)
	}
	if ad.Spec.Invoke != ad.ZAP.Invoke {
		fmt.Fprintf(w, "%s%s %s has invoke access %s, but is %s in the spec\n", prefix, name, entityType.String(), ad.ZAP.Invoke, ad.Spec.Invoke)
	}
	if ad.Spec.FabricScoping != ad.ZAP.FabricScoping {
		fmt.Fprintf(w, "%s%s %s has fabric scoping %s, but is %s in the spec\n", prefix, name, entityType.String(), ad.ZAP.FabricScoping, ad.Spec.FabricScoping)
	}
	if ad.Spec.FabricSensitivity != ad.ZAP.FabricSensitivity {
		fmt.Fprintf(w, "%s%s %s has fabric sensitivity %s, but is %s in the spec\n", prefix, name, entityType.String(), ad.ZAP.FabricSensitivity, ad.Spec.FabricSensitivity)
	}
	if ad.Spec.Timing != ad.ZAP.Timing {
		fmt.Fprintf(w, "%s%s %s has timing %s, but is %s in the spec\n", prefix, name, entityType.String(), ad.ZAP.Timing, ad.Spec.Timing)
	}
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
	fmt.Fprintf(w, "%s%s %s %s is marked as %s, but should be %s\n", prefix, name, entityType, sd.Property.String(), sd.ZAP, sd.Spec)
}

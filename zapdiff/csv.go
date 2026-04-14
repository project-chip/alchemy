package zapdiff

import (
	"encoding/csv"
	"io"
	"log/slog"
	"sort"
)

// WriteMismatchesToCSV writes the given mismatches to the writer in CSV format.
func WriteMismatchesToCSV(w io.Writer, mm []XmlMismatch, l XmlMismatchLevel) (err error) {
	csvWriter := csv.NewWriter(w)
	defer csvWriter.Flush()

	// Write header
	header := []string{"Level", "Type", "File", "Element Xpath", "Details"}
	if err = csvWriter.Write(header); err != nil {
		slog.Error("failed to write CSV header", "error", err)
		return
	}

	sort.Slice(mm, func(i, j int) bool {
		// Level (Descending), Path, Type, EntityUniqueIdentifier, Details
		if mm[i].Level() != mm[j].Level() {
			return mm[i].Level() > mm[j].Level()
		}
		if mm[i].Path != mm[j].Path {
			return mm[i].Path < mm[j].Path
		}
		if mm[i].Type != mm[j].Type {
			return mm[i].Type.String() < mm[j].Type.String()
		}
		if mm[i].EntityUniqueIdentifier != mm[j].EntityUniqueIdentifier {
			return mm[i].EntityUniqueIdentifier < mm[j].EntityUniqueIdentifier
		}
		return mm[i].Details < mm[j].Details
	})

	// Write mismatches
	for _, m := range mm {
		if m.Level() >= l {
			row := []string{
				m.Level().String(),
				m.Type.String(),
				m.Path,
				m.EntityUniqueIdentifier,
				m.Details,
			}
			if err = csvWriter.Write(row); err != nil {
				slog.Error("Warning: failed to write row to CSV", "err", err)
				return
			}
		}
	}

	return
}

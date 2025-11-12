package cli

import (
	"encoding/csv"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/zapdiff"
)

type ZAPDiff struct {
	SdkRoot1      string `default:"connectedhomeip" help:"the first clone of project-chip/connectedhomeip" group:"SDK Commands:"`
	SdkRoot2      string `default:"connectedhomeip" help:"the second clone of project-chip/connectedhomeip" group:"SDK Commands:"`
	Out           string `default:"." help:"path to output mismatch.csv file" group:"SDK Commands:"`
	MismatchLevel int    `default:"3" help:"The minimum mismatch level to report (1-3)" group:"SDK Commands:"`
}

func (z *ZAPDiff) Run(cc *Context) (err error) {
	var mismatchPrintLevel zapdiff.XmlMismatchLevel
	if z.MismatchLevel < 1 || z.MismatchLevel > 3 {
		slog.Warn("invalid mismatch level. must be between 1 and 3.", "level", z.MismatchLevel)
		mismatchPrintLevel = zapdiff.MismatchLevel3 // Default
	} else {
		mismatchPrintLevel = zapdiff.XmlMismatchLevel(z.MismatchLevel - 1) // Convert 1-3 to 0-2
	}

	err = sdk.CheckAlchemyVersion(z.SdkRoot1)
	if err != nil {
		return
	}

	err = sdk.CheckAlchemyVersion(z.SdkRoot2)
	if err != nil {
		return
	}

	p1 := filepath.Join(z.SdkRoot1, "src", "app", "zap-templates", "zcl", "data-model", "chip")
	ff1, err := listXMLFiles(p1)
	if err != nil {
		slog.Error("error listing files", "dir", p1, "error", err)
		return err
	}

	p2 := filepath.Join(z.SdkRoot2, "src", "app", "zap-templates", "zcl", "data-model", "chip")
	ff2, err := listXMLFiles(p2)
	if err != nil {
		slog.Error("error listing files", "dir", p2, "error", err)
		return err
	}

	mm := zapdiff.Pipeline(ff1, ff2, "sdk-1", "sdk-2")

	csvOutputPath := filepath.Join(z.Out, "mismatches.csv")
	err = writeMismatchesToCSV(csvOutputPath, mm, mismatchPrintLevel)
	if err != nil {
		slog.Error("Failed to write CSV output", "error", err)
	}

	return
}

func listXMLFiles(p string) (paths []string, err error) {
	var entries []os.DirEntry
	entries, err = os.ReadDir(p)
	if err != nil {
		return
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".xml") {
			paths = append(paths, filepath.Join(p, e.Name()))
		}
	}

	return
}

func writeMismatchesToCSV(p string, mm []zapdiff.XmlMismatch, l zapdiff.XmlMismatchLevel) (err error) {
	f, err := os.Create(p)
	if err != nil {
		slog.Error("failed to create file", "path", p, "error", err)
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	// Write header
	header := []string{"Level", "Type", "File", "Element Xpath", "Details"}
	if err = w.Write(header); err != nil {
		slog.Error("failed to write CSV header", "error", err)
		return
	}

	sort.Slice(mm, func(i, j int) bool {
		// Level (Descending), Path, Type, ElementID, Details
		if mm[i].Level() != mm[j].Level() {
			return mm[i].Level() > mm[j].Level()
		}
		if mm[i].Path != mm[j].Path {
			return mm[i].Path < mm[j].Path
		}
		if mm[i].Type != mm[j].Type {
			return mm[i].Type.String() < mm[j].Type.String()
		}
		if mm[i].ElementID != mm[j].ElementID {
			return mm[i].ElementID < mm[j].ElementID
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
				m.ElementID,
				m.Details,
			}
			if err = w.Write(row); err != nil {
				slog.Error("Warning: failed to write row to CSV", "err", err)
				return
			}
		}
	}

	slog.Info("Successfully wrote mismatches to CSV", "dir", p)
	return
}

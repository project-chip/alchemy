package cli

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/project-chip/alchemy/zapdiff"
)

type ZAPDiff struct {
	XmlRoot1      string `help:"root of first set of ZAP XMLs" group:"SDK Commands:" required:"true"`
	XmlRoot2      string `help:"root of second set of ZAP XMLs" group:"SDK Commands:" required:"true"`
	Label1        string `default:"ZapXML-1" help:"label for first set of ZAP XMLs" group:"SDK Commands:"`
	Label2        string `default:"ZapXML-2" help:"label for second set of ZAP XMLs" group:"SDK Commands:"`
	Out           string `default:"." help:"path to output mismatch files" group:"SDK Commands:"`
	Format        string `default:"both" help:"output format: csv, html, or both" group:"SDK Commands:"`
	MismatchLevel int    `default:"3" help:"the minimum mismatch level to report (1-3)" group:"SDK Commands:"`
}

func (z *ZAPDiff) Run(cc *Context) (err error) {

	var mismatchPrintLevel zapdiff.XmlMismatchLevel
	if z.MismatchLevel < 1 || z.MismatchLevel > 3 {
		slog.Warn("invalid mismatch level. must be between 1 and 3.", "level", z.MismatchLevel)
		mismatchPrintLevel = zapdiff.MismatchLevel3 // Default
	} else {
		mismatchPrintLevel = zapdiff.XmlMismatchLevel(z.MismatchLevel - 1) // Convert 1-3 to 0-2
	}

	ff1, err := listXMLFiles(z.XmlRoot1)
	if err != nil {
		slog.Error("error listing files", "dir", z.XmlRoot1, "error", err)
		return err
	}

	ff2, err := listXMLFiles(z.XmlRoot2)
	if err != nil {
		slog.Error("error listing files", "dir", z.XmlRoot2, "error", err)
		return err
	}

	mm := zapdiff.Pipeline(ff1, ff2, z.Label1, z.Label2)

	generateCSV := z.Format == "csv" || z.Format == "both" || z.Format == ""
	generateHTML := z.Format == "html" || z.Format == "both" || z.Format == ""

	if generateCSV {
		csvOutputPath := filepath.Join(z.Out, "mismatches.csv")
		f, err := os.Create(csvOutputPath)
		if err != nil {
			slog.Error("failed to create CSV file", "path", csvOutputPath, "error", err)
			return err
		}
		defer f.Close()
		err = zapdiff.WriteMismatchesToCSV(f, mm, mismatchPrintLevel)
		if err != nil {
			slog.Error("Failed to write CSV output", "error", err)
		} else {
			slog.Info("Successfully wrote mismatches to CSV", "dir", csvOutputPath)
		}
	}

	if generateHTML {
		htmlOutputPath := filepath.Join(z.Out, "mismatches.html")
		f, err := os.Create(htmlOutputPath)
		if err != nil {
			slog.Error("failed to create HTML file", "path", htmlOutputPath, "error", err)
			return err
		}
		defer f.Close()
		err = zapdiff.WriteMismatchesToHTML(f, mm, mismatchPrintLevel, z.XmlRoot1, z.XmlRoot2)
		if err != nil {
			slog.Error("Failed to write HTML output", "error", err)
		} else {
			slog.Info("Successfully wrote mismatches to HTML", "dir", htmlOutputPath)
		}
	}

	return
}

func listXMLFiles(p string) (paths []string, err error) {
	fi, err := os.Stat(p)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return []string{p}, nil
	}
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


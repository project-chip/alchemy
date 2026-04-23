package zapdiff

import (
	"log/slog"
	"path/filepath"

	"github.com/beevik/etree"
)

type filePair struct {
	p1 string
	p2 string
}

type elementPair struct {
	e1 *etree.Element
	e2 *etree.Element
}

func Pipeline(ff1, ff2 []string, n1, n2 string) (mm []XmlMismatch) {
	mm = make([]XmlMismatch, 0)

	// Filter manual files
	f1 := excludeNonAlchemyFiles(ff1)
	f2 := excludeNonAlchemyFiles(ff2)
	ff := getFilePairs(f1, f2)

	mm = append(mm, fileListDiff(f1, f2, n1, n2)...)

	for _, f := range ff {
		baseName := filepath.Base(f.p1)

		d1 := etree.NewDocument()
		d2 := etree.NewDocument()

		err := d1.ReadFromFile(f.p1)
		if err != nil {
			slog.Warn("Failed to parse", "file", f.p1, "error", err)
			continue
		}
		err = d2.ReadFromFile(f.p2)
		if err != nil {
			slog.Warn("Failed to parse", "file", f.p2, "error", err)
			continue
		}

		r1 := d1.Root()
		r2 := d2.Root()

		if r1 == nil {
			slog.Warn("File has no root element", "file", baseName, "clone", n1)
			continue
		}
		if r2 == nil {
			slog.Warn("File has no root element", "file", baseName, "clone", n2)
			continue
		}

		emm := checkMismatches(elementPair{e1: r1, e2: r2}, baseName, n1, n2)
		mm = append(mm, emm...)
	}
	return
}

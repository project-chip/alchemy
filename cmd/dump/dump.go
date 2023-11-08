package dump

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/parse"
)

func Dump(cxt context.Context, filepaths []string, options Options) error {
	files, err := files.Paths(filepaths)
	if err != nil {
		return err
	}
	for i, f := range files {
		if len(files) > 0 {
			fmt.Fprintf(os.Stderr, "Dumping %s (%d of %d)...\n", f, (i + 1), len(files))
		}
		doc, err := ascii.Open(f, options.AsciiSettings...)
		docType, err := doc.DocType()
		if err != nil {
			return err
		}
		if options.Ascii {
			for _, top := range parse.Skim[*ascii.Section](doc.Elements) {
				ascii.AssignSectionTypes(docType, top)
			}
			dumpElements(doc, doc.Elements, 0)
		} else if options.Json {
			models, err := doc.ToModel()
			if err != nil {
				return err
			}
			encoder := json.NewEncoder(os.Stdout)
			//encoder.SetIndent("", "\t")
			return encoder.Encode(models)
		} else {
			dumpElements(doc, doc.Base.Elements, 0)
		}
	}
	return nil
}

func dumpTOC(tocs []*types.ToCSection, indent int) {
	for _, toc := range tocs {
		fmt.Print(strings.Repeat("\t", indent))
		fmt.Printf("{toc %d} %s.%s\n", toc.Level, toc.Number, toc.Title)
		if len(toc.Children) > 0 {
			dumpTOC(toc.Children, indent+1)
		}
	}

}

func dumpLocation(l *types.Location) {
	if l != nil {
		fmt.Printf("%s %s}", l.Scheme, l.Path.(string))
	} else {
		fmt.Printf("missing location")
	}
}

func snippet(str string) string {
	v := []rune(str)
	if 42 < len(v) {
		str = string(v[:20]) + "…" + string(v[len(v)-20:])
	}
	return strings.ReplaceAll(str, "\n", "")
}

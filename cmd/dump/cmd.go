package dump

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type Command struct {
	spec.ParserOptions `embed:""`
	pipeline.ProcessingOptions
	common.ASCIIDocAttributes `embed:""`

	Paths []string `arg:""`
	Ascii bool     `name:"ascii" help:"dump asciidoc object model"`
	Json  bool     `name:"json" help:"dump json object model"`
}

func (d *Command) Run(cc *cli.Context) (err error) {
	files, err := paths.Expand(d.Paths)
	if err != nil {
		return fmt.Errorf("error building paths: %w", err)
	}
	if d.Json {
		encoder := json.NewEncoder(os.Stdout)
		s, _, err := spec.Parse(cc, d.ParserOptions, d.ProcessingOptions, nil, d.ASCIIDocAttributes.ToList())
		if err != nil {
			return err
		}

		for i, f := range files {

			path, err := spec.NewSpecPath(f, s.Root)
			if err != nil {
				return fmt.Errorf("error resolving path doc %s: %w", f, err)
			}
			doc, ok := s.Docs[path.Relative]
			if !ok {
				if len(files) > 0 {
					fmt.Fprintf(os.Stderr, "Skipping %s (%d of %d)...\n", f, (i + 1), len(files))
				}
				continue
			}
			if len(files) > 0 {
				fmt.Fprintf(os.Stderr, "Dumping %s (%d of %d)...\n", f, (i + 1), len(files))
			}
			entities := s.EntitiesForDocument(doc)
			encoder.Encode(entities)

		}

		/*doc, _, err := spec.ParseFile(path, d.Root, d.ASCIIDocAttributes.ToList()...)
		if err != nil {
			return fmt.Errorf("error opening doc %s: %w", f, err)
		}
		entities, err := doc.Entities()
		if err != nil {
			return fmt.Errorf("error parsing entities %s: %w", f, err)
		}
		globalObjects, err := doc.GlobalObjects()
		if err != nil {
			return fmt.Errorf("error parsing global objects %s: %w", f, err)
		}
		entities = append(entities, globalObjects...)*/
		//encoder.SetIndent("", "\t")

	} else {
		for i, f := range files {
			if len(files) > 0 {
				fmt.Fprintf(os.Stderr, "Dumping %s (%d of %d)...\n", f, (i + 1), len(files))
			}
			if d.Ascii {
				doc, err := spec.ReadFile(f, ".")
				if err != nil {
					return fmt.Errorf("error opening doc %s: %w", f, err)
				}

				sic := newDumpInfoCache(asciidoc.RawReader)

				spec.AssignDocType(sic, sic, asciidoc.RawReader, doc)
				for top := range parse.Skim[*asciidoc.Section](asciidoc.RawReader, doc, doc.Children()) {
					err := spec.AssignSectionTypes(sic, sic, sic, doc, top)
					if err != nil {
						return err
					}
				}
				dumpElements(sic, doc, doc.Children(), 0)
			} else {
				doc, err := spec.ReadFile(f, ".")
				if err != nil {
					return fmt.Errorf("error opening doc %s: %w", f, err)
				}
				dumpElements(newDumpInfoCache(asciidoc.RawReader), doc, doc.Children(), 0)
			}
		}
	}

	return nil
}

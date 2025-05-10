package dump

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/matter/spec"
)

type Command struct {
	spec.ParserOptions        `embed:""`
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
	for i, f := range files {
		if len(files) > 0 {
			fmt.Fprintf(os.Stderr, "Dumping %s (%d of %d)...\n", f, (i + 1), len(files))
		}
		if d.Ascii {
			doc, err := spec.ReadFile(f, ".")
			if err != nil {
				return fmt.Errorf("error opening doc %s: %w", f, err)
			}

			for top := range parse.Skim[*spec.Section](doc.Elements()) {
				err := spec.AssignSectionTypes(doc, top)
				if err != nil {
					return err
				}
			}
			dumpElements(doc, doc.Elements(), 0)
		} else if d.Json {
			err = errata.LoadErrataConfig(d.Root)
			if err != nil {
				return
			}
			path, err := asciidoc.NewPath(f, d.Root)
			if err != nil {
				return fmt.Errorf("error resolving doc path %s: %w", f, err)
			}

			doc, err := spec.Parse(path, d.Root, d.ASCIIDocAttributes.ToList()...)
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
			entities = append(entities, globalObjects...)
			encoder := json.NewEncoder(os.Stdout)
			//encoder.SetIndent("", "\t")
			return encoder.Encode(entities)
		} else if d.Inline {
			err = errata.LoadErrataConfig(d.Root)
			if err != nil {
				return
			}
			path, err := asciidoc.NewPath(f, d.Root)
			if err != nil {
				return fmt.Errorf("error resolving doc path %s: %w", f, err)
			}
			doc, err := spec.Parse(path, d.Root, d.ASCIIDocAttributes.ToList()...)
			if err != nil {
				return fmt.Errorf("error parsing %s: %w", f, err)
			}
			dumpElements(doc, doc.Elements(), 0)
		} else {
			doc, err := spec.ReadFile(f, ".")
			if err != nil {
				return fmt.Errorf("error opening doc %s: %w", f, err)
			}
			dumpElements(doc, doc.Base.Elements(), 0)
		}
	}
	return nil
}

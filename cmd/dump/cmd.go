package dump

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "dump filename_pattern",
	Short: "dump the parse tree of Matter documents specified by filename_pattern",
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		asciiSettings := common.ASCIIDocAttributes(cmd)
		asciiOut, _ := cmd.Flags().GetBool("ascii")
		jsonOut, _ := cmd.Flags().GetBool("json")
		unifiedOut, _ := cmd.Flags().GetBool("inline")
		specRoot, _ := cmd.Flags().GetString("specRoot")

		files, err := paths.Expand(args)
		if err != nil {
			return fmt.Errorf("error building paths: %w", err)
		}
		for i, f := range files {
			if len(files) > 0 {
				fmt.Fprintf(os.Stderr, "Dumping %s (%d of %d)...\n", f, (i + 1), len(files))
			}
			if asciiOut {
				doc, err := spec.ReadFile(f, ".")
				if err != nil {
					return fmt.Errorf("error opening doc %s: %w", f, err)
				}

				for _, top := range parse.Skim[*spec.Section](doc.Elements()) {
					err := spec.AssignSectionTypes(doc, top)
					if err != nil {
						return err
					}
				}
				dumpElements(doc, doc.Elements(), 0)
			} else if jsonOut {
				err = errata.LoadErrataConfig(specRoot)
				if err != nil {
					return
				}
				path, err := asciidoc.NewPath(f, specRoot)
				if err != nil {
					return fmt.Errorf("error resolving doc path %s: %w", f, err)
				}

				doc, err := spec.ParseFile(path, specRoot, asciiSettings...)
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
			} else if unifiedOut {
				err = errata.LoadErrataConfig(specRoot)
				if err != nil {
					return
				}
				path, err := asciidoc.NewPath(f, specRoot)
				if err != nil {
					return fmt.Errorf("error resolving doc path %s: %w", f, err)
				}
				doc, err := spec.InlineParse(path, specRoot, asciiSettings...)
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
	},
}

func init() {
	Command.Flags().Bool("ascii", false, "dump asciidoc object model")
	Command.Flags().Bool("json", false, "dump json object model")
	Command.Flags().Bool("inline", false, "use inline parser")
	Command.Flags().String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
}

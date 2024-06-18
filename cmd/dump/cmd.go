package dump

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter/spec"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "dump",
	Short: "dump the parse tree of Matter documents",
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		asciiSettings := common.ASCIIDocAttributes(cmd)
		asciiOut, _ := cmd.Flags().GetBool("ascii")
		jsonOut, _ := cmd.Flags().GetBool("json")

		files, err := files.Paths(args)
		if err != nil {
			return fmt.Errorf("error building paths: %w", err)
		}
		for i, f := range files {
			if len(files) > 0 {
				fmt.Fprintf(os.Stderr, "Dumping %s (%d of %d)...\n", f, (i + 1), len(files))
			}
			if asciiOut {
				doc, err := spec.ReadFile(f)
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
				asciiSettings = append(asciiSettings, spec.GithubSettings()...)
				doc, err := spec.ParseFile(f, asciiSettings...)
				if err != nil {
					return fmt.Errorf("error opening doc %s: %w", f, err)
				}
				entities, err := doc.Entities()
				if err != nil {
					return fmt.Errorf("error parsing entities %s: %w", f, err)
				}
				encoder := json.NewEncoder(os.Stdout)
				//encoder.SetIndent("", "\t")
				return encoder.Encode(entities)
			} else {
				doc, err := spec.ReadFile(f)
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
}

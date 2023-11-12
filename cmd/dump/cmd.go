package dump

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/parse"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "dump",
	Short: "dump the parse tree of Matter documents",
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		asciiSettings := common.AsciiDocAttributes(cmd)
		asciiOut, _ := cmd.Flags().GetBool("ascii")
		jsonOut, _ := cmd.Flags().GetBool("json")

		files, err := files.Paths(args)
		if err != nil {
			return err
		}
		for i, f := range files {
			if len(files) > 0 {
				fmt.Fprintf(os.Stderr, "Dumping %s (%d of %d)...\n", f, (i + 1), len(files))
			}
			doc, err := ascii.Open(f, asciiSettings...)
			docType, err := doc.DocType()
			if err != nil {
				return err
			}
			if asciiOut {
				for _, top := range parse.Skim[*ascii.Section](doc.Elements) {
					ascii.AssignSectionTypes(docType, top)
				}
				dumpElements(doc, doc.Elements, 0)
			} else if jsonOut {
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
	},
}

func init() {
	Command.Flags().Bool("ascii", false, "dump asciidoc object model")
	Command.Flags().Bool("json", false, "dump json object model")
}

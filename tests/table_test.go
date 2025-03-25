package tests

import (
	"testing"

	"github.com/sanity-io/litter"
)

func init() {
	litter.Config.HomePackage = "github.com/project-chip/alchemy"
}

var tableTests = parseTests{
	{"table simple", "table_simple.adoc", tableSimple, nil},
	{"table comment", "table_comment.adoc", tableComment, nil},
	{"table new line", "table_new_line.adoc", tableNewLine, nil},
	{"table blank cells", "table_blank_cells.adoc", tableBlankCells, nil},
	{"table indented cell", "table_indented_cell.adoc", tableIndentedCell, nil},
	{"table intermediate empty line", "table_intermediate_empty_line.adoc", tableIntermediateEmptyLine, nil},
	{"table end empty line", "table_end_empty_line.adoc", tableEndEmptyLine, nil},
	{"table special character", "table_special_character.adoc", tableSpecialCharacter, nil},
}

func TestBlockTables(t *testing.T) {
	tableTests.run(t)
}

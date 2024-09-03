//go:build generate

package main

import "strings"

func parserPatch(parser string) string {
	parser = strings.ReplaceAll(parser, "globalStore storeDict", "globalStore storeDict\n\tdelimitedBlockState delimitedBlockState\n\ttableColumnsAttribute *asciidoc.TableColumnsAttribute")
	parser = strings.ReplaceAll(parser, "globalStore: make(storeDict),", "globalStore: make(storeDict),\n\t\t\tdelimitedBlockState: make(delimitedBlockState),")
	return parser
}

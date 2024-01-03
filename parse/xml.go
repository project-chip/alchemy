package parse

import (
	"github.com/beevik/etree"
)

func FormatXML(x string) (string, error) {
	doc := etree.NewDocument()
	err := doc.ReadFromString(x)
	if err != nil {
		return "", err
	}
	indent := etree.NewIndentSettings()
	indent.Spaces = 2
	indent.PreserveLeafWhitespace = true
	doc.IndentWithSettings(indent)
	return doc.WriteToString()
}

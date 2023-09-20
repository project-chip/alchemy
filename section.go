package main

import "github.com/bytesparadise/libasciidoc/pkg/types"

type section struct {
	name string

	base *types.Section

	secType docSectionType

	elements []interface{}
}

type docSectionType uint8

const (
	docSectionTypePreface docSectionType = 0
	docSectionTypeAttributes
	docSectionTypeFeatures
	docSectionTypeDataTypes
	docSectionTypeCommands
)

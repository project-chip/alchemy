//go:build generate

package main

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/dave/dst/dstutil"
)

func parserPatch(file *dst.File) (err error) {
	// This adds a couple extra fields to the "current" struct
	var currentStruct *dst.StructType
	var newParser *dst.FuncDecl
	var parseRuleRefExpr *dst.FuncDecl
	dstutil.Apply(file, nil, func(c *dstutil.Cursor) bool {
		ok := true
		switch node := c.Node().(type) {
		case *dst.TypeSpec:
			switch node.Name.Name {
			case "current":
				currentStruct, ok = node.Type.(*dst.StructType)
			}
		case *dst.FuncDecl:
			switch node.Name.Name {
			case "newParser":
				newParser = node
			case "parseRuleRefExpr":
				parseRuleRefExpr = node
			}
		case nil:
		default:
		}
		return ok
	})
	if currentStruct == nil {
		err = fmt.Errorf("unable to find 'current' struct in parser")
		return
	}
	if newParser == nil {
		err = fmt.Errorf("unable to find 'newParser' struct in parser")
		return
	}
	patchCurrent(currentStruct)
	patchNewParser(newParser)
	if parseRuleRefExpr != nil {
		patchParseRuleRefExpr(parseRuleRefExpr)
	}
	return
}

func patchCurrent(currentStruct *dst.StructType) {
	// We add a couple extra fields to the "current" struct to keep track of nested delimited blocks, and cols attributes on tables
	delimitedBlockStateField := &dst.Field{
		Names: []*dst.Ident{dst.NewIdent("delimitedBlockState")},
		Type:  dst.NewIdent("delimitedBlockState"),
	}
	delimitedBlockStateField.Decorations().End.Append("// Alchemy patch: we keep track of levels of delimited blocks here")

	currentStruct.Fields.List = append(currentStruct.Fields.List, delimitedBlockStateField)

	tableColumnsAttributeField := &dst.Field{
		Names: []*dst.Ident{dst.NewIdent("tableColumnsAttribute")},
		Type:  dst.NewIdent("*asciidoc.TableColumnsAttribute"),
	}
	tableColumnsAttributeField.Decorations().End.Append("// Alchemy patch: we keep track of explicit table columns here to aid in allocating columns")
	currentStruct.Fields.List = append(currentStruct.Fields.List, tableColumnsAttributeField)
}

func patchNewParser(newParser *dst.FuncDecl) {
	// We need to initialize the delimitedBlockState map that we added to the "current" struct
	var pAssign *dst.AssignStmt
	for _, stmt := range newParser.Body.List {
		switch stmt := stmt.(type) {
		case *dst.AssignStmt:
			id, ok := stmt.Lhs[0].(*dst.Ident)
			if !ok {
				continue
			}
			switch id.Name {
			case "p":
				pAssign = stmt
			}

		}
	}
	if pAssign == nil {
		return
	}
	ue, ok := pAssign.Rhs[0].(*dst.UnaryExpr)
	if !ok {
		return
	}
	cl, ok := ue.X.(*dst.CompositeLit)
	if !ok {
		return
	}
	// Find the initialization of the "cur" field, and add a line to make the delimitedBlockState map
	for _, x := range cl.Elts {
		switch x := x.(type) {
		case *dst.KeyValueExpr:
			switch key := x.Key.(type) {
			case *dst.Ident:
				if key.Name == "cur" {
					val, ok := x.Value.(*dst.CompositeLit)
					if !ok {
						return
					}
					kv := &dst.KeyValueExpr{
						Key: dst.NewIdent("delimitedBlockState"),
						Value: &dst.CallExpr{
							Fun: dst.NewIdent("make"),
							Args: []dst.Expr{
								dst.NewIdent("delimitedBlockState"),
							},
						},
					}
					kv.Decs.After = dst.NewLine
					val.Elts = append(val.Elts, kv)
				}
			}
		}
	}
}

func patchParseRuleRefExpr(parseRuleRefExpr *dst.FuncDecl) {
	for _, stmt := range parseRuleRefExpr.Body.List {
		switch stmt := stmt.(type) {
		case *dst.IfStmt:
			switch cond := stmt.Cond.(type) {
			case *dst.SelectorExpr:
				if cond.Sel.Name == "debug" {
					// There's a debugging line here that we broke by removing the name cache for performance reasons
					stmt.Body.List = nil
				}
			}
		}
	}
}

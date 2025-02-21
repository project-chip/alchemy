package generate

import (
	"bytes"
	"fmt"
	"go/token"
	"slices"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
)

func optimizeParser(source string, customPatcher ParserPatcher) (out string, err error) {
	// There are several aspects of the pigeon-generated parser that we don't use but have performance impacts, so we patch the source to remove/improve them
	var file *dst.File
	file, err = decorator.Parse(source)
	if err != nil {
		return
	}
	var currentStruct, parserStruct, charClassMatcherStruct *dst.StructType
	var newParser, parseSeqExpr *dst.FuncDecl
	dstutil.Apply(file, nil, func(c *dstutil.Cursor) bool {
		ok := true
		switch node := c.Node().(type) {
		case *dst.TypeSpec:
			switch node.Name.Name {
			case "current":
				currentStruct, ok = node.Type.(*dst.StructType)
			case "parser":
				parserStruct, ok = node.Type.(*dst.StructType)
			case "charClassMatcher":
				charClassMatcherStruct, ok = node.Type.(*dst.StructType)
			}
		case *dst.FuncDecl:
			switch node.Name.Name {
			case "newParser":
				newParser = node
			case "parseSeqExpr":
				parseSeqExpr = node
			}
		case nil:
		}
		return ok
	})
	if currentStruct == nil {
		err = fmt.Errorf("unable to find 'current' struct in parser")
		return
	}
	if parserStruct == nil {
		err = fmt.Errorf("unable to find 'parser' struct in parser")
		return
	}
	if parseSeqExpr == nil {
		err = fmt.Errorf("unable to find 'parseSeqExpr' func in parser")
		return
	}
	if newParser == nil {
		err = fmt.Errorf("unable to find 'newParser' func in parser")
		return
	}
	if charClassMatcherStruct == nil {
		err = fmt.Errorf("unable to find 'charClassMatcher' struct in parser")
		return
	}
	patchCurrent(currentStruct)
	patchParserStruct(parserStruct)
	patchParseSeqExpr(parseSeqExpr)
	patchCharClassMatcherStruct(charClassMatcherStruct)
	patchNewParser(newParser)

	if customPatcher != nil {
		err = customPatcher(file)
		if err != nil {
			return
		}
	}
	buf := bytes.NewBuffer(make([]byte, 0, len(source)))
	err = decorator.Fprint(buf, file)
	if err != nil {
		return
	}
	out = buf.String()
	return
}

func patchCurrent(currentStruct *dst.StructType) {
	// We add a reference to the overall parser in the "current" struct; this allows us
	// access to the file path and cursor position within the grammar
	parserField := &dst.Field{
		Names: []*dst.Ident{dst.NewIdent("parser")},
		Type:  dst.NewIdent("*parser"),
	}
	parserField.Decorations().Before = dst.EmptyLine
	parserField.Decorations().End.Append("// Alchemy patch: we keep a reference to the parent parser here, so inline code can access it")
	currentStruct.Fields.List = append(currentStruct.Fields.List, parserField)
}

func patchParserStruct(parserStruct *dst.StructType) {
	// We add a position field to the parser so elements can state their location in the file
	offsetField := &dst.Field{
		Names: []*dst.Ident{dst.NewIdent("offset")},
		Type:  dst.NewIdent("position"),
	}
	offsetField.Decorations().Before = dst.EmptyLine
	offsetField.Decorations().End.Append("// Alchemy patch: we add an offset field to track element positions in the doc")
	parserStruct.Fields.List = append(parserStruct.Fields.List, offsetField)
}

func patchParseSeqExpr(parseSeqExpr *dst.FuncDecl) {
	// The parseSeqExpr function allocates a vals slice whether it uses it or not; it's faster to do this lazily
	valsIndex := -1
	for i, stmt := range parseSeqExpr.Body.List {
		switch stmt := stmt.(type) {
		case *dst.AssignStmt:
			id, ok := stmt.Lhs[0].(*dst.Ident)
			if !ok {
				continue
			}
			switch id.Name {
			case "vals":
				valsIndex = i
			}
		}
	}
	if valsIndex == -1 {
		return
	}
	valsDecl := &dst.DeclStmt{
		Decl: &dst.GenDecl{
			Tok: token.VAR,
			Specs: []dst.Spec{
				&dst.ValueSpec{
					Names: []*dst.Ident{dst.NewIdent("vals")},
					Type:  &dst.ArrayType{Elt: dst.NewIdent("any")},
				},
			},
		},
	}
	valsDecl.Decorations().End.Append("// Alchemy patch: we lazily allocate this array, as it's infrequently populated")
	valsDecl.Decorations().After = dst.EmptyLine
	parseSeqExpr.Body.List[valsIndex] = valsDecl
}

func patchCharClassMatcherStruct(charClassMatcherStruct *dst.StructType) {
	// There's a -optimize-basic-latin option for pigeon that turns out to be slower, but the parser allocates an array for it whether
	// it's used or not; if it's there, we comment it out
	basicLatinCharsIndex := -1
	for i, f := range charClassMatcherStruct.Fields.List {
		if f.Names[0].Name == "basicLatinChars" {
			basicLatinCharsIndex = i
			break
		}
	}
	if basicLatinCharsIndex < 0 {
		return
	}
	charClassMatcherStruct.Fields.List = slices.Delete(charClassMatcherStruct.Fields.List, basicLatinCharsIndex, basicLatinCharsIndex+1)
	charClassMatcherStruct.Fields.List[basicLatinCharsIndex].Decorations().Start.Prepend("//basicLatinChars [128]bool")
	charClassMatcherStruct.Fields.List[basicLatinCharsIndex].Decorations().Start.Prepend("// Alchemy patch: we don't use this optimization, so don't allocate the array")
}

func patchNewParser(newParser *dst.FuncDecl) {
	// We added a new field to the "current" struct to reference the parent parser, so we need to add
	// a line to set that field after the parser is created. It's added right after the setOptions call
	setOptionsIndex := -1
	for i, stmt := range newParser.Body.List {

		switch stmt := stmt.(type) {
		case *dst.ExprStmt:
			switch x := stmt.X.(type) {
			case *dst.CallExpr:
				switch f := x.Fun.(type) {
				case *dst.SelectorExpr:
					if f.Sel.Name == "setOptions" {
						setOptionsIndex = i
						break
					}
				}
			}

		}
	}
	if setOptionsIndex >= 0 {
		assign := &dst.AssignStmt{
			Lhs: []dst.Expr{
				&dst.SelectorExpr{
					X: &dst.SelectorExpr{
						X:   dst.NewIdent("p"),
						Sel: dst.NewIdent("cur"),
					},
					Sel: dst.NewIdent("parser"),
				},
			},
			Tok: token.ASSIGN,
			Rhs: []dst.Expr{
				dst.NewIdent("p"),
			},
		}
		assign.Decorations().Before = dst.None
		assign.Decorations().End.Append(" // Alchemy patch: We copy the parser pointer to the current parse state so it is available to inline code")
		newParser.Body.List[setOptionsIndex].Decorations().After = dst.NewLine
		newParser.Body.List = slices.Insert[[]dst.Stmt, dst.Stmt](newParser.Body.List, setOptionsIndex+1, assign)
	}
}

package main

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/lithammer/dedent"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/sanity-io/litter"
)

var testPattern = regexp.MustCompile(`(?m)test '(?P<Name>[^']+)' do[^<]+<<~'EOS'(?P<Ascii>(.|\n)*?)EOS`)

type adTestGroup struct {
	Name      string
	ArrayName string
	Tests     []*adTest
}

type adTest struct {
	Name            string
	QuotedName      string
	TestName        string
	Asciidoc        string
	AsciidocPath    string
	DocObject       string
	ParsedDocObject string
}

func init() {
	litter.Config.DisablePointerReplacement = true
	litter.Config.FieldExclusions = regexp.MustCompile("^Parent$")
}

func main() {
	text, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	testName := filepath.Base(os.Args[1])
	testName = strings.TrimSuffix(testName, filepath.Ext(testName))

	var tests []*adTest
	matches := testPattern.FindAllStringSubmatch(string(text), -1)
	testNames := make(map[string]struct{})
	for _, match := range matches {
		fmt.Printf("test: %s\n", match[1])
		tn := match[1]
		_, exists := testNames[tn]
		if exists {
			continue
		}
		tests = append(tests, &adTest{Name: tn, QuotedName: strconv.Quote(tn), Asciidoc: match[2]})
		testNames[tn] = struct{}{}
	}

	for _, t := range tests {
		a := dedent.Dedent(t.Asciidoc)
		doc, err := parse.Reader("", strings.NewReader(a))
		if err != nil {
			slog.Error("failed parsing", "test", t.Name, "err", err)
			continue
		}
		name := strings.ReplaceAll(t.Name, ",", "")
		name = strings.ReplaceAll(name, "/", " ")
		name = strings.ReplaceAll(name, "\"", " ")

		docPath, err := asciidoc.NewPath("test.adoc", ".")
		if err != nil {
			slog.Error("failed creating path", "test", t.Name, "err", err)
			continue
		}
		parsedDoc, err := parse.Inline(spec.NewPreparseContext(docPath, "."), "test.adoc", strings.NewReader(a))
		if err != nil {
			slog.Error("failed creating path", "test", t.Name, "err", err)
			continue
		}

		t.TestName = strcase.ToLowerCamel(name)
		t.DocObject = cleanObjectDump(litter.Sdump(doc))
		t.ParsedDocObject = cleanObjectDump(litter.Sdump(parsedDoc))
		t.AsciidocPath = path.Join("asciidoctor/", testName+"_"+strcase.ToSnakeWithIgnore(name, ",")+".adoc")
		renderedDocPath := path.Join("tests/asciidoctor/", testName+"_"+strcase.ToSnakeWithIgnore(name, ",")+".adoc")
		err = os.WriteFile(renderedDocPath, []byte(a), os.ModeAppend|0644)
		if err != nil {
			panic(err)
		}
	}

	testName = strings.TrimSuffix(testName, "_test")

	t := adTestGroup{
		Name:      strcase.ToCamel(testName),
		ArrayName: strcase.ToLowerCamel(testName),
		Tests:     tests,
	}

	tt := template.New("Asciidoctor Test Template")
	tt, err = tt.Parse(testTemplate)
	if err != nil {
		panic(err)
	}

	tf, err := os.Create(path.Join("tests/", testName+"_test.go"))

	if err != nil {
		panic(err)
	}
	err = tt.Execute(tf, t)
	if err != nil {
		panic(err)
	}
	tf.Close()
}

func cleanObjectDump(s string) string {
	return strings.ReplaceAll(s, "[github.com/project-chip/alchemy/", "[")
}

var testTemplate = `package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func Test{{.Name}}(t *testing.T) {
	{{.ArrayName}}Tests.run(t)
}

var {{.ArrayName}}Tests = parseTests{
	{{range .Tests}}
	{ {{.QuotedName}}, "{{.AsciidocPath}}", {{.TestName}} },
	{{end}}
}

{{range .Tests}}
var {{.TestName}} = {{.GoObject}}
{{end}}

`

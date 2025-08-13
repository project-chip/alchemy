package main

import (
	"bytes"
	"fmt"
	"go/format"
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
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/sanity-io/litter"
)

var testPattern = regexp.MustCompile(`(?m)test '(?P<Name>[^']+)' do[^<]+<<~'EOS'(?P<Ascii>(.|\n)*?)EOS`)

type adTestGroup struct {
	Name      string
	ArrayName string
	Tests     []*adTest
}

type adTest struct {
	Name         string
	QuotedName   string
	TestName     string
	Asciidoc     string
	AsciidocPath string
	DocObject    string
}

func init() {
	litter.Config.DisablePointerReplacement = true
	litter.Config.FieldExclusions = regexp.MustCompile("^Parent$")
}

func main() {
	testPaths, err := paths.Expand(os.Args[1:])
	if err != nil {
		panic(err)
	}
	for _, testPath := range testPaths {
		text, err := os.ReadFile(testPath)
		if err != nil {
			panic(err)
		}

		testName := filepath.Base(testPath)
		testName = strings.TrimSuffix(testName, filepath.Ext(testName))

		var tests []*adTest
		matches := testPattern.FindAllStringSubmatch(string(text), -1)
		testNames := make(map[string]struct{})
		for _, match := range matches {
			tn := match[1]
			_, exists := testNames[tn]
			if exists {
				continue
			}
			tests = append(tests, &adTest{Name: tn, QuotedName: strconv.Quote(tn), Asciidoc: match[2]})
			testNames[tn] = struct{}{}
		}

		var workingTests []*adTest

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

			fmt.Printf("parsing test: %s\n", t.Name)

			t.TestName = strcase.ToLowerCamel(testName + " " + name)
			t.DocObject = cleanObjectDump(litter.Sdump(doc))
			fileName := testName + "_" + cleanFileName(strcase.ToSnakeWithIgnore(name, ",")) + ".adoc"
			t.AsciidocPath = path.Join("asciidoctor/", fileName)
			renderedDocPath := path.Join("tests/asciidoctor/", fileName)
			err = os.WriteFile(renderedDocPath, []byte(a), os.ModeAppend|0644)
			if err != nil {
				panic(err)
			}
			workingTests = append(workingTests, t)
		}

		if len(workingTests) == 0 {
			continue
		}

		testName = strings.TrimSuffix(testName, "_test")

		t := adTestGroup{
			Name:      strcase.ToCamel(testName),
			ArrayName: strcase.ToLowerCamel(testName),
			Tests:     workingTests,
		}

		tt := template.New("Asciidoctor Test Template")
		tt, err = tt.Parse(testTemplate)
		if err != nil {
			panic(err)
		}

		out := new(bytes.Buffer)

		err = tt.Execute(out, t)
		if err != nil {
			panic(err)
		}

		tf, err := os.Create(path.Join("tests/", testName+"_test.go"))

		if err != nil {
			panic(err)
		}

		var formatted []byte
		formatted, err = format.Source(out.Bytes())
		if err != nil {
			panic(err)
		}

		tf.Write(formatted)
		tf.Close()

	}
}

var pointerCommentPattern = regexp.MustCompile(`(?m)\/\/ p[0-9]+$`)

var invalidFilenameChars = strings.NewReplacer("\\", "", "/", "", ":", "", "*", "", "?", "", "\"", "", "<", "", ">", "", "|", "")

func cleanFileName(path string) string {
	return invalidFilenameChars.Replace(path)
}

func cleanObjectDump(s string) string {
	return pointerCommentPattern.ReplaceAllString(strings.ReplaceAll(s, "[github.com/project-chip/alchemy/", "["), "")
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
	{{"{"}}{{.QuotedName}}, "{{.AsciidocPath}}", {{.TestName}} },
	{{end}}
}

{{range .Tests}}
var {{.TestName}} = {{.DocObject}}
{{end}}

`

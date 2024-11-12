package parse

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/goccy/go-yaml"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type TestYamlParser struct {
	rootPath string
}

func NewTestYamlParser(rootPath string) (TestYamlParser, error) {
	if !filepath.IsAbs(rootPath) {
		var err error
		rootPath, err = filepath.Abs(rootPath)
		if err != nil {
			return TestYamlParser{}, err
		}
	}
	return TestYamlParser{rootPath: rootPath}, nil
}

func (p TestYamlParser) Name() string {
	return "Parsing YAML tests"
}

func (p TestYamlParser) Process(cxt context.Context, input *pipeline.Data[struct{}], index int32, total int32) (outputs []*pipeline.Data[*Test], extras []*pipeline.Data[struct{}], err error) {
	var r []byte
	r, err = os.ReadFile(input.Path)
	if err != nil {
		return
	}
	t := Test{Path: input.Path}
	cm := yaml.CommentMap{}
	err = yaml.UnmarshalWithOptions(r, &t, yaml.DisallowUnknownField(), yaml.CommentToMap(cm), yaml.UseOrderedMap())
	if err != nil {
		slog.WarnContext(cxt, "Failed parsing test YAML", slog.String("path", input.Path), slog.Any("error", err))
		err = nil
		return
	}
	commentPattern := regexp.MustCompile(`\$\.tests\[([0-9]+)\]\.label`)
	for key, c := range cm {
		comments := commentPattern.FindStringSubmatch(key)
		if len(comments) == 0 {
			continue
		}
		index, err := strconv.Atoi(comments[1])
		if err != nil {
			continue
		}
		if len(t.Tests) <= index {
			continue
		}
		step := t.Tests[index]
		for _, comment := range c {
			step.Comments = append(step.Comments, comment.Texts...)
		}
	}
	outputs = append(outputs, pipeline.NewData(input.Path, &t))
	return
}

func stringUnmarshaller(s *string, b []byte) error {
	return nil
}

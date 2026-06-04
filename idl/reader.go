package idl

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	"github.com/project-chip/alchemy/internal/pipeline"
)

type Reader struct {
}

func NewReader() (Reader, error) {
	return Reader{}, nil
}

func (p Reader) Name() string {
	return "Parsing ZAP files"
}

func (p Reader) Process(cxt context.Context, input *pipeline.Data[struct{}], index int32, total int32) (outputs []*pipeline.Data[*File], extras []*pipeline.Data[struct{}], err error) {
	slog.Info("reading zap path", "path", input.Path)
	var contents *os.File
	contents, err = os.Open(input.Path)
	if err != nil {
		return
	}
	defer contents.Close()

	decoder := json.NewDecoder(contents)

	decoder.DisallowUnknownFields()

	var zf File

	err = decoder.Decode(&zf)
	if err != nil {
		slog.Error("Error reading ZAP path", slog.Any("error", err), slog.String("source", input.Path))
		return
	}

	outputs = append(outputs, &pipeline.Data[*File]{Path: input.Path, Content: &zf})
	return
}

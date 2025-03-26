package files

import (
	"context"
	"os"

	"github.com/project-chip/alchemy/internal/pipeline"
)

type Reader struct {
	name string
}

func (sp *Reader) Name() string {
	return sp.name
}

func (sp *Reader) Process(cxt context.Context, input *pipeline.Data[struct{}], index int32, total int32) (outputs []*pipeline.Data[[]byte], extras []*pipeline.Data[struct{}], err error) {
	var b []byte
	b, err = os.ReadFile(input.Path)
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData(input.Path, b))
	return
}

func NewReader(name string) *Reader {
	return &Reader{name: name}
}

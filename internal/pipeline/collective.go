package pipeline

import (
	"context"
	"fmt"
	"os"

	"github.com/puzpuzpuz/xsync/v3"
)

type CollectiveProcessor[I, O any] interface {
	Processor
	Process(cxt context.Context, inputs []*Data[I]) (outputs []*Data[O], err error)
}

func processCollective[I, O any](cxt context.Context, processor CollectiveProcessor[I, O], input *xsync.MapOf[string, *Data[I]]) (output *xsync.MapOf[string, *Data[O]], err error) {
	name := processor.Name()
	if len(name) > 0 {
		cyan.Fprintf(os.Stderr, "%s...\n", name)
	}

	total := int32(input.Size())
	inputs := make([]*Data[I], 0, total)
	input.Range(func(key string, value *Data[I]) bool {
		inputs = append(inputs, value)
		return true
	})
	var outputs []*Data[O]
	outputs, err = processor.Process(cxt, inputs)
	if err != nil {
		return
	}
	output = xsync.NewMapOfPresized[string, *Data[O]](len(outputs))
	for _, o := range outputs {
		_, loaded := output.LoadAndStore(o.Path, o)
		if loaded {
			err = fmt.Errorf("duplicate path in output: %s", o.Path)
			return
		}
	}
	return
}

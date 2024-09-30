package pipeline

import (
	"context"
	"fmt"
	"os"
)

type CollectiveProcessor[I, O any] interface {
	Processor
	Process(cxt context.Context, inputs []*Data[I]) (outputs []*Data[O], err error)
}

func processCollective[I, O any](cxt context.Context, processor CollectiveProcessor[I, O], input Map[string, *Data[I]]) (output Map[string, *Data[O]], err error) {
	name := processor.Name()
	if len(name) > 0 {
		cyan.Fprintf(os.Stderr, "%s...\n", name)
	}

	total := int32(input.Size())
	output = NewMapPresized[string, *Data[O]](input.Size())
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
	for _, o := range outputs {
		_, loaded := output.LoadAndStore(o.Path, o)
		if loaded {
			err = fmt.Errorf("duplicate path in output: %s", o.Path)
			return
		}
	}
	return
}

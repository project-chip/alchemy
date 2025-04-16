package pipeline

import (
	"context"
	"fmt"
)

type CollectiveProcessor[I, O any] interface {
	Processor
	Process(cxt context.Context, inputs []*Data[I]) (outputs []*Data[O], err error)
}

type CollectiveProcess[I, O any] func(cxt context.Context, inputs []*Data[I]) (outputs []*Data[O], err error)

func Collective[I, O any](cxt context.Context, options ProcessingOptions, processor CollectiveProcessor[I, O], input Map[string, *Data[I]]) (output Map[string, *Data[O]], err error) {
	inputs := DataMapToSlice(input)
	output = NewMapPresized[string, *Data[O]](input.Size())
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

type anonymousCollectiveProcessor[I, O any] struct {
	name    string
	process CollectiveProcess[I, O]
}

func (acp *anonymousCollectiveProcessor[I, O]) Name() string {
	return acp.name
}

func (acp *anonymousCollectiveProcessor[I, O]) Process(cxt context.Context, inputs []*Data[I]) (outputs []*Data[O], err error) {
	return acp.process(cxt, inputs)
}

func CollectiveFunc[I, O any](name string, process CollectiveProcess[I, O]) CollectiveProcessor[I, O] {
	return &anonymousCollectiveProcessor[I, O]{process: process}
}

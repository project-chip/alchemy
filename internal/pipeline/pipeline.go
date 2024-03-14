package pipeline

import (
	"context"
	"fmt"
	"reflect"

	"github.com/puzpuzpuz/xsync/v3"
)

type Targeter func(cxt context.Context) ([]string, error)

func Start[T any](cxt context.Context, targeter Targeter) (*xsync.MapOf[string, *Data[T]], error) {
	paths, err := targeter(cxt)
	if err != nil {
		return nil, err
	}
	output := xsync.NewMapOf[string, *Data[T]]()
	for _, p := range paths {
		_, loaded := output.LoadAndStore(p, &Data[T]{Path: p})
		if loaded {
			return nil, fmt.Errorf("duplicate path in target: %s", p)
		}
	}
	return output, nil
}

func Process[I, O any](cxt context.Context, options Options, processor Processor, input *xsync.MapOf[string, *Data[I]]) (output *xsync.MapOf[string, *Data[O]], err error) {
	switch processor.Type() {
	case ProcessorTypeCollective:
		proc, ok := processor.(CollectiveProcessor[I, O])
		if !ok {
			proc = processor.(CollectiveProcessor[I, O])
			err = fmt.Errorf("processor \"%s\" claimed to be collective, but does not implement CollectiveProcessor interface", processor.Name())
			return
		}
		return processCollective[I, O](cxt, proc, input)
	case ProcessorTypeIndividual:
		proc, ok := processor.(IndividualProcessor[I, O])
		if !ok {
			proc = processor.(IndividualProcessor[I, O])
			err = fmt.Errorf("processor \"%s\" claimed to be individual, but does not implement IndividualProcessor interface", processor.Name())
			fooType := reflect.TypeOf(processor)
			for i := 0; i < fooType.NumMethod(); i++ {
				method := fooType.Method(i)
				fmt.Printf("method: %v\n", method)
			}
			return
		}
		if options.Serial {
			return ProcessSerialFunc[I, O](cxt, options, input, proc.Name(), proc.Process)
		}
		return ProcessParallelFunc[I, O](cxt, options, input, proc.Name(), proc.Process)
	}
	return
}

func ProcessSerialFunc[I, O any](cxt context.Context, options Options, input *xsync.MapOf[string, *Data[I]], name string, f IndividualProcess[I, O]) (output *xsync.MapOf[string, *Data[O]], err error) {
	total := int32(input.Size())
	queue := make(chan *Data[I], total)
	inputs := DataMapToSlice[I](input)
	SortData[I](inputs)
	for _, i := range inputs {
		select {
		case queue <- i:
		default:
			err = fmt.Errorf("queue full")
			return

		}
	}
	return processSerial[I, O](cxt, name, f, queue, total)
}

func ProcessParallelFunc[I, O any](cxt context.Context, options Options, input *xsync.MapOf[string, *Data[I]], name string, f IndividualProcess[I, O]) (output *xsync.MapOf[string, *Data[O]], err error) {
	total := int32(input.Size())
	queue := make(chan *Data[I], total)
	input.Range(func(key string, value *Data[I]) bool {
		select {
		case queue <- value:
			return true
		default:
			err = fmt.Errorf("queue full")
			return false

		}
	})
	return processParallel[I, O](cxt, name, f, queue, total)
}

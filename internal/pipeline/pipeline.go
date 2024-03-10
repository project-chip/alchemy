package pipeline

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"
	"unicode/utf8"

	"github.com/puzpuzpuz/xsync/v3"
	"github.com/schollz/progressbar/v3"
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

func Process[I, O any](cxt context.Context, options Options, processor Processor[I, O], input *xsync.MapOf[string, *Data[I]]) (output *xsync.MapOf[string, *Data[O]], err error) {
	switch processor.Type() {
	case ProcessorTypeSerial:
		return processAll[I, O](cxt, processor, input)
	case ProcessorTypeParallel:
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
		if err != nil {
			return
		}
		if options.Serial {
			return processSerial[I, O](cxt, processor, queue, total)
		}
		return processParallel[I, O](cxt, processor, queue, total)
	}
	return
}

func processAll[I, O any](cxt context.Context, processor Processor[I, O], input *xsync.MapOf[string, *Data[I]]) (output *xsync.MapOf[string, *Data[O]], err error) {
	fmt.Fprintf(os.Stderr, "%s...\n", processor.Name())
	total := int32(input.Size())
	inputs := make([]*Data[I], 0, total)
	input.Range(func(key string, value *Data[I]) bool {
		inputs = append(inputs, value)
		return true
	})
	var outputs []*Data[O]
	outputs, err = processor.ProcessAll(cxt, inputs)
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

func processSerial[I, O any](cxt context.Context, processor Processor[I, O], queue chan *Data[I], total int32) (output *xsync.MapOf[string, *Data[O]], err error) {
	var counter int32
	output = xsync.NewMapOfPresized[string, *Data[O]](int(total))
	fmt.Fprintf(os.Stderr, "%s...\n", processor.Name())
	for {
		var input *Data[I]
		select {
		case input = <-queue:
		default:
		}
		if input == nil {
			return
		}
		var outputs []*Data[O]
		var extras []*Data[I]
		outputs, extras, err = processor.Process(cxt, input, counter, total)
		if err != nil {
			return nil, err
		}
		counter++
		for _, e := range extras {
			select {
			case queue <- e:
				total++
			default:
				err = fmt.Errorf("queue full")
				return
			}
		}
		for _, o := range outputs {
			_, loaded := output.LoadAndStore(o.Path, o)
			if loaded {
				err = fmt.Errorf("duplicate path in output: %s", o.Path)
				return
			}
		}
	}
}

func processParallel[I, O any](cxt context.Context, processor Processor[I, O], queue chan *Data[I], total int32) (output *xsync.MapOf[string, *Data[O]], err error) {

	output = xsync.NewMapOfPresized[string, *Data[O]](int(total))
	fmt.Fprintf(os.Stderr, "%s...\n", processor.Name())
	bar := progressbar.Default(int64(total))
	var complete int32
	wg := newWorkGroup()
	for {
		var input *Data[I]
		select {
		case input = <-queue:
		default:
		}
		if input == nil {
			break
		}
		wg.run(cxt, func() error {
			done := atomic.AddInt32(&complete, 1)
			outputs, extras, err := processor.Process(cxt, input, done, total)
			if err != nil {
				return err
			}
			bar.Describe(progressFileName(input.Path))
			bar.Add(1)
			for _, e := range extras {
				select {
				case queue <- e:
					newTotal := atomic.AddInt32(&total, 1)
					bar.ChangeMax(int(newTotal))
				default:
					return fmt.Errorf("queue full")
				}
			}
			for _, o := range outputs {
				_, loaded := output.LoadAndStore(o.Path, o)
				if loaded {
					return fmt.Errorf("duplicate path in output: %s", o.Path)
				}
			}
			return nil
		})
	}
	err = wg.Wait()
	return
}

const fileNameSize = 30

func progressFileName(file string) string {
	file = filepath.Base(file)
	length := utf8.RuneCountInString(file)
	if length > fileNameSize {
		file = string([]rune(file)[length-fileNameSize:])
	}
	return fmt.Sprintf("%-*s", fileNameSize, file)
}

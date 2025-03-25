package pipeline

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func processSerial[I, O any](cxt context.Context, name string, processor IndividualProcess[I, O], queue chan *Data[I], total int32) (output Map[string, *Data[O]], err error) {
	var counter int32
	processed := make(map[string]bool, total)
	output = NewMapPresized[string, *Data[O]](int(total))
	cyan.Fprintf(os.Stderr, "%s...\n", name)

	for {
		var input *Data[I]
		var done bool
		select {
		case input = <-queue:
			done = processed[input.Path]
		default:
		}
		if input == nil {
			return
		}
		if done {
			continue
		}
		var outputs []*Data[O]
		var extras []*Data[I]
		outputs, extras, err = processor(cxt, input, counter, total)
		if err != nil {
			return nil, err
		}
		counter++
		gray.Fprintf(os.Stderr, "%s...\n", input.Path)
		processed[input.Path] = true
		for _, e := range extras {
			_, ok := processed[e.Path]
			if ok {
				continue
			}
			select {
			case queue <- e:
				total++
				processed[e.Path] = false
			default:
				err = fmt.Errorf("queue full")
				return
			}
		}
		for _, o := range outputs {
			_, loaded := output.LoadOrStore(o.Path, o)
			if loaded {
				slog.WarnContext(cxt, "duplicate path in output", slog.String("path", o.Path))
				return
			}
		}
	}
}

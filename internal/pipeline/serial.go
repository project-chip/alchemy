package pipeline

import (
	"context"
	"iter"
	"log/slog"
	"os"
)

type queue[V any] struct {
	initial    iter.Seq[V]
	additional []V
}

func (q *queue[V]) Next() iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range q.initial {
			if !yield(v) {
				return
			}
		}
		for {
			if len(q.additional) == 0 {
				return
			}
			v := q.additional[0]
			q.additional = q.additional[1:]
			if !yield(v) {
				return
			}
		}
	}
}

func processSerial[I, O any](cxt context.Context, name string, processor IndividualProcess[I, O], values iter.Seq[*Data[I]], total int32) (output Map[string, *Data[O]], err error) {
	var counter int32
	processed := make(map[string]bool, total)
	output = NewMapPresized[string, *Data[O]](int(total))
	cyan.Fprintf(os.Stderr, "%s...\n", name)

	queue := &queue[*Data[I]]{initial: values}

	for input := range queue.Next() {
		done := processed[input.Path]
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
			queue.additional = append(queue.additional, e)
		}
		for _, o := range outputs {
			_, loaded := output.LoadOrStore(o.Path, o)
			if loaded {
				slog.WarnContext(cxt, "duplicate path in output", slog.String("path", o.Path))
				return
			}
		}
	}
	return
}

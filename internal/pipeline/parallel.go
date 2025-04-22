package pipeline

import (
	"context"
	"fmt"
	"iter"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"sync/atomic"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
)

var cyan = color.New(color.FgCyan).Add(color.Bold)
var gray = color.New(color.FgBlack).Add(color.Bold)

func Parallel[I, O any](cxt context.Context, options ProcessingOptions, processor IndividualProcessor[I, O], input Map[string, *Data[I]]) (output Map[string, *Data[O]], err error) {
	total := int32(input.Size())

	var values iter.Seq[*Data[I]]
	if options.Serial {
		inputs := DataMapToSlice(input)
		SortData(inputs)
		values = slices.Values(inputs)
	} else {
		values = dataMapValues(input)
	}

	if options.Serial {
		return processSerial(cxt, processor.Name(), processor.Process, values, total)
	}
	return processParallel(cxt, processor.Name(), processor.Process, values, total, !options.NoProgress)
}

func processParallel[I, O any](cxt context.Context, name string, processor IndividualProcess[I, O], values iter.Seq[*Data[I]], total int32, showProgress bool) (output Map[string, *Data[O]], err error) {

	processed := NewConcurrentMapPresized[string, bool](int(total))
	output = NewConcurrentMapPresized[string, *Data[O]](int(total))
	cyan.Fprintf(os.Stderr, "%s...\n", name)
	var bar *progressbar.ProgressBar
	if showProgress {
		bar = progressbar.Default(int64(total))
	}
	var complete int32
	queue := make(chan *Data[I], total)
	for value := range values {
		select {
		case queue <- value:
			processed.LoadOrStore(value.Path, false)
		default:
			err = fmt.Errorf("queue full")
		}
		if err != nil {
			return
		}
	}

	for {
		wg := newWorkGroup()
		for {
			var input *Data[I]
			var done bool
			select {
			case input = <-queue:
				done, _ = processed.LoadOrStore(input.Path, false)
			default:
			}
			if input == nil {
				break
			}
			if done {
				if bar != nil {
					_ = bar.Add(1)
				}
				continue
			}
			wg.run(cxt, func() error {
				done := atomic.AddInt32(&complete, 1)
				completed, _ := processed.Load(input.Path)
				if completed {
					slog.WarnContext(cxt, "skipping already processed input", slog.String("path", input.Path))
					return nil
				}
				outputs, extras, err := processor(cxt, input, done, total)
				if err != nil {
					return err
				}
				if bar != nil {
					bar.Describe(progressFileName(input.Path))
					_ = bar.Add(1)
				}
				processed.Store(input.Path, true)
				for _, e := range extras {
					_, loaded := processed.LoadOrStore(e.Path, false)
					if loaded {
						slog.DebugContext(cxt, "skipping already queued input", slog.String("path", input.Path))
						continue
					}
					select {
					case queue <- e:
						newTotal := atomic.AddInt32(&total, 1)
						if bar != nil {
							bar.ChangeMax(int(newTotal))
						}
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
		if err != nil {
			return
		}
		if len(queue) == 0 {
			break
		}
	}
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

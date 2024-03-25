package pipeline

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync/atomic"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
)

var cyan = color.New(color.FgCyan).Add(color.Bold)

func processParallel[I, O any](cxt context.Context, name string, processor IndividualProcess[I, O], queue chan *Data[I], total int32, showProgress bool) (output Map[string, *Data[O]], err error) {

	processed := NewConcurrentMapPresized[string, bool](int(total))
	output = NewConcurrentMapPresized[string, *Data[O]](int(total))
	cyan.Fprintf(os.Stderr, "%s...\n", name)
	var bar *progressbar.ProgressBar
	if showProgress {
		bar = progressbar.Default(int64(total))
	}
	var complete int32

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
				bar.Add(1)
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
				bar.Add(1)
			}
			processed.Store(input.Path, true)
			for _, e := range extras {
				_, loaded := processed.LoadOrStore(e.Path, false)
				if loaded {
					slog.WarnContext(cxt, "skipping already queued input", slog.String("path", input.Path))
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

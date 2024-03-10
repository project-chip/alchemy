package pipeline

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"sync"

	"golang.org/x/sync/semaphore"
)

type ErrorGroup []error

type workGroup struct {
	wait sync.WaitGroup
	s    *semaphore.Weighted

	lock   sync.Mutex
	errors ErrorGroup
}

func newWorkGroup() *workGroup {
	return &workGroup{
		s: semaphore.NewWeighted(int64(runtime.GOMAXPROCS(0))),
	}
}

func (wg *workGroup) run(cxt context.Context, f func() error) {
	wg.wait.Add(1)

	err := wg.s.Acquire(cxt, 1)
	if err != nil {
		wg.lock.Lock()
		wg.errors = append(wg.errors, fmt.Errorf("semaphore error: %w", err))
		wg.lock.Unlock()
		wg.wait.Done()
		return
	}

	go func() {
		err := f()
		if err != nil {
			wg.lock.Lock()
			wg.errors = append(wg.errors, err)
			wg.lock.Unlock()
		}
		wg.wait.Done()
		wg.s.Release(1)
	}()
}

func (wg *workGroup) Wait() error {
	wg.wait.Wait()
	if len(wg.errors) == 0 {
		return nil
	}
	return wg.errors
}

func (eg ErrorGroup) Error() string {
	var s strings.Builder
	for _, e := range eg {
		s.WriteString(fmt.Sprintf("error: %s\n", e.Error()))
	}
	return s.String()
}

package pipeline

import "sync"

type Once[T any] struct {
	once sync.Once
	err  error
	t    T
}

func (o *Once[T]) Do(f func() (T, error)) (T, error) {
	o.once.Do(func() {
		o.t, o.err = f()
	})
	return o.t, o.err
}

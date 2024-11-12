package pipeline

import (
	"context"
	"fmt"
)

type Paths Map[string, *Data[struct{}]]

type Targeter func(cxt context.Context) ([]string, error)

func Start(cxt context.Context, targeter Targeter) (Paths, error) {
	paths, err := targeter(cxt)
	if err != nil {
		return nil, err
	}
	output := NewMapPresized[string, *Data[struct{}]](len(paths))
	for _, p := range paths {
		_, loaded := output.LoadAndStore(p, &Data[struct{}]{Path: p})
		if loaded {
			return nil, fmt.Errorf("duplicate path in target: %s", p)
		}
	}
	return output, nil
}

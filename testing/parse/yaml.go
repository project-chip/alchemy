package parse

import (
	"fmt"
	"log/slog"
)

type unmarshaller interface {
	UnmarshalMap(c map[string]any) error
}

func extractValue[T any](val map[string]any, key string, defaultValue ...T) (value T) {
	var v any
	var ok bool
	if v, ok = val[key]; !ok {
		if len(defaultValue) > 0 {
			value = defaultValue[0]
		}
		return
	}
	switch v := v.(type) {
	case T:
		value = v
		delete(val, key)
		return
	case map[string]any:
		if u, ok := any(&value).(unmarshaller); ok {
			err := u.UnmarshalMap(v)
			if err == nil {
				delete(val, key)
				return
			}
		}
	}
	slog.Info("unable to cast YAML value", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", value)), slog.String("got", fmt.Sprintf("%T", v)))
	return
}

func extractObject[T any](val map[string]any, key string) (value *T) {
	var v any
	var ok bool
	if v, ok = val[key]; !ok {
		return
	}
	switch v := v.(type) {
	case T:
		value = &v
		delete(val, key)
		return
	case map[string]any:
		value = new(T)
		if u, ok := any(value).(unmarshaller); ok {
			err := u.UnmarshalMap(v)
			if err == nil {
				delete(val, key)
				return
			}
		}
	}
	slog.Info("unable to cast YAML value", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", value)), slog.String("got", fmt.Sprintf("%T", v)))
	return
}

func extractArray[T any](val map[string]any, key string) (value []T) {
	var v any
	var ok bool
	if v, ok = val[key]; !ok {
		return
	}
	switch v := v.(type) {
	case []T:
		value = v
		delete(val, key)
		return
	case []any:
		for _, v := range v {
			if v, ok := v.(T); ok {
				value = append(value, v)
			} else {
				slog.Info("unable to cast YAML array value element", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", value)), slog.String("got", fmt.Sprintf("%T", v)))

			}
		}
		delete(val, key)
		return
	}
	slog.Info("unable to cast YAML array value", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", value)), slog.String("got", fmt.Sprintf("%T", v)))
	return
}

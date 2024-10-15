package parse

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/goccy/go-yaml"
	"github.com/project-chip/alchemy/internal/log"
)

type unmarshaller interface {
	UnmarshalMap(c yaml.MapSlice) error
}

func deleteMapItem(val yaml.MapSlice, key string) yaml.MapSlice {
	return slices.DeleteFunc(val, func(mi yaml.MapItem) bool {
		if mi, ok := mi.Key.(string); ok {
			return mi == key
		}
		return false
	})
}

func extractValue[T any](val yaml.MapSlice, key string, defaultValue ...T) (value T, out yaml.MapSlice) {
	out = val
	if key == "defaultValue" {
		slog.Info("getting default value")
	}
	v, ok := ValueFromMapSlice(val, key)
	if !ok {
		if len(defaultValue) > 0 {
			value = defaultValue[0]
		}
		return
	}
	if key == "defaultValue" {
		slog.Info("got default value", log.Type("type", v))
	}

	switch v := v.(type) {
	case T:
		value = v
		out = deleteMapItem(val, key)
		return
	case yaml.MapSlice:
		if u, ok := any(&value).(unmarshaller); ok {
			err := u.UnmarshalMap(v)
			if err == nil {
				out = deleteMapItem(val, key)
				return
			}
		}
	}
	slog.Info("unable to cast YAML value", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", value)), slog.String("got", fmt.Sprintf("%T", v)))
	return
}

func ValueFromMapSlice(val yaml.MapSlice, key string) (v any, ok bool) {
	for _, si := range val {
		switch sik := si.Key.(type) {
		case string:
			if sik == key {
				v = si.Value
				ok = true
				return
			}
		}
	}
	return
}

func extractObject[T any](val yaml.MapSlice, key string) (value *T, out yaml.MapSlice) {
	out = val
	v, ok := ValueFromMapSlice(val, key)
	if !ok {
		return
	}
	switch v := v.(type) {
	case T:
		value = &v
		out = deleteMapItem(val, key)
		return
	case yaml.MapSlice:
		value = new(T)
		if u, ok := any(value).(unmarshaller); ok {
			err := u.UnmarshalMap(v)
			if err == nil {
				out = deleteMapItem(val, key)
				return
			}
		}
	}
	slog.Info("unable to cast YAML value", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", value)), slog.String("got", fmt.Sprintf("%T", v)))
	return
}

func extractArray[T any](val yaml.MapSlice, key string) (value []T, out yaml.MapSlice) {
	out = val
	v, ok := ValueFromMapSlice(val, key)
	if !ok {
		return
	}
	switch v := v.(type) {
	case T:
		value = append(value, v)
		out = deleteMapItem(val, key)
		return
	case []T:
		value = v
		out = deleteMapItem(val, key)
		return
	case []any:
		for _, v := range v {
			el, ok := v.(T)
			if ok {
				value = append(value, el)
			} else {
				slog.Info("unable to cast YAML array value element", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", el)), slog.String("got", fmt.Sprintf("%T", v)))

			}
		}
		out = deleteMapItem(val, key)
		return
	}
	slog.Info("unable to cast YAML array value", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", value)), slog.String("got", fmt.Sprintf("%T", v)))
	return
}

func extractArrayAny(val yaml.MapSlice, key string) (value []any, out yaml.MapSlice) {
	out = val
	v, ok := ValueFromMapSlice(val, key)
	if !ok {
		return
	}
	switch v := v.(type) {
	case []any:
		value = v
		out = deleteMapItem(val, key)
		return
	}
	slog.Info("unable to cast YAML any array value", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", value)), slog.String("got", fmt.Sprintf("%T", v)))
	return
}

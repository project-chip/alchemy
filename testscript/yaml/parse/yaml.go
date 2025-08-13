package parse

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/goccy/go-yaml"
)

type mapSliceUnmarshaller interface {
	UnmarshalMapSlice(c yaml.MapSlice) error
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
	value, out, _ = tryExtractValue(val, key, defaultValue...)
	return
}

func tryExtractValue[T any](val yaml.MapSlice, key string, defaultValue ...T) (value T, out yaml.MapSlice, ok bool) {
	out = val
	var v any
	v, ok = ValueFromMapSlice(val, key)
	if !ok {
		if len(defaultValue) > 0 {
			value = defaultValue[0]
			ok = true
		}
		return
	}
	switch v := v.(type) {
	case T:
		value = v
		out = deleteMapItem(val, key)
		return
	case yaml.MapSlice:
		var u mapSliceUnmarshaller
		if u, ok = any(&value).(mapSliceUnmarshaller); ok {
			err := u.UnmarshalMapSlice(v)
			if err == nil {
				out = deleteMapItem(val, key)
				return
			}
			ok = false
		}
	case nil:
		out = deleteMapItem(val, key)
		return
	default:
		ok = false
	}
	if !ok {
		slog.Warn("unable to cast YAML value", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", value)), slog.String("got", fmt.Sprintf("%T", v)))
	}
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
		if u, ok := any(value).(mapSliceUnmarshaller); ok {
			err := u.UnmarshalMapSlice(v)
			if err == nil {
				out = deleteMapItem(val, key)
				return
			}
		}
	}
	slog.Warn("unable to cast YAML value", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", value)), slog.String("got", fmt.Sprintf("%T", v)))
	return
}

func extractArray[T any](val yaml.MapSlice, key string) (value []T, out yaml.MapSlice) {
	value, out, _ = tryExtractArray[T](val, key)
	return
}

func tryExtractArray[T any](val yaml.MapSlice, key string) (value []T, out yaml.MapSlice, ok bool) {
	out = val

	var v any
	v, ok = ValueFromMapSlice(val, key)
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
			var el T
			el, ok = v.(T)
			if ok {
				value = append(value, el)
			} else {
				slog.Warn("unable to cast YAML array value element", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", el)), slog.String("got", fmt.Sprintf("%T", v)))
				return
			}
		}
		out = deleteMapItem(val, key)
		return
	default:
		ok = false
		slog.Warn("unable to cast YAML array value", slog.String("key", key), slog.String("expected", fmt.Sprintf("%T", value)), slog.String("got", fmt.Sprintf("%T", v)))
		return
	}
}

func extractArrayAny(val yaml.MapSlice, key string) (value []any, out yaml.MapSlice) {
	value, out, _ = tryExtractArrayAny(val, key)
	return
}

func tryExtractArrayAny(val yaml.MapSlice, key string) (value []any, out yaml.MapSlice, ok bool) {
	out = val
	var v any
	v, ok = ValueFromMapSlice(val, key)
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
	ok = false
	return
}

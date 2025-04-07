package internal

import (
	"fmt"
	"log/slog"

	"github.com/goccy/go-yaml"
)

type MapMarshaler interface {
	ToMap() (map[string]any, error)
	FromMap(map[string]any) error
}

type YamlMap[V MapMarshaler] struct {
	OrderedMap[string, V]
}

func NewYamlMap[V MapMarshaler]() *YamlMap[V] {
	return &YamlMap[V]{
		OrderedMap: OrderedMap[string, V]{
			items: make(map[string]V),
		},
	}
}

func (om *YamlMap[V]) MarshalYAML() ([]byte, error) {
	ms := yaml.MapSlice{}
	for _, key := range om.keys {
		ms = append(ms, yaml.MapItem{Key: key, Value: om.items[key]})
	}
	return yaml.Marshal(ms)
}

func (t *YamlMap[V]) UnmarshalYAML(unmarshal func(any) error) (err error) {
	var ms yaml.MapSlice
	slog.Info("Unmarshalling ordered map yaml")
	if err = unmarshal(&ms); err != nil {
		err = fmt.Errorf("error unmarshaling map: %w", err)
		return err
	}
	for _, item := range ms {
		key, ok := item.Key.(string)
		if !ok {
			err = fmt.Errorf("unexpected key type: %T", item.Key)
			return
		}

		var value V
		switch val := item.Value.(type) {
		case map[string]any:
			err = value.FromMap(val)
			if err != nil {
				return
			}
		default:
			err = fmt.Errorf("unexpected value type for key %v: %T", key, item.Value)
			return
		}
		t.Set(key, value)
	}
	return
}

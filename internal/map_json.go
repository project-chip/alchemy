package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type JSONMap struct {
	OrderedMap[string, any]
}

func NewJSONMap() *JSONMap {
	return &JSONMap{
		OrderedMap: OrderedMap[string, any]{
			items: make(map[string]any),
		},
	}
}

func (om *JSONMap) MarshalJSON() ([]byte, error) {
	if om == nil || len(om.keys) == 0 {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteRune('{')
	for i, key := range om.keys {
		s := fmt.Sprintf("%v", key)
		v := om.items[key]
		var b []byte
		var err error
		switch v := v.(type) {
		case []any:
			if len(v) == 0 {
				b = []byte("[]")
			} else {
				b, err = json.Marshal(&v)
			}
		default:
			b, err = json.Marshal(&v)
		}

		if err != nil {
			return nil, err
		}
		if i > 0 {
			buf.WriteRune(',')
		}
		buf.WriteString(fmt.Sprintf("%q:", s))
		buf.Write(b)
	}
	buf.WriteRune('}')
	return buf.Bytes(), nil
}

func (o *JSONMap) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		o.items = make(map[string]any)
		o.keys = nil
		return nil
	}
	dec := json.NewDecoder(bytes.NewReader(b))
	var tok json.Token
	tok, err = dec.Token()
	if err != nil {
		err = fmt.Errorf("error reading first token at %d: %w", dec.InputOffset(), err)
		return
	}
	if tok, ok := tok.(json.Delim); !ok || tok != '{' {
		err = fmt.Errorf("expected JSON object")
		return
	}
	o.items = make(map[string]any)
	o.keys = nil

	for dec.More() {
		tok, err = dec.Token()
		if err != nil {
			err = fmt.Errorf("error reading token at %d: %w", dec.InputOffset(), err)
			return err
		}

		var key string
		switch tok := tok.(type) {
		case string:
			if err = json.Unmarshal([]byte(strconv.Quote(tok)), &key); err != nil {
				err = fmt.Errorf("error unmarshalling key token at %d: %w", dec.InputOffset(), err)
				return err
			}
		default:
			err = fmt.Errorf("unexpected token in JSON at %d: %v", dec.InputOffset(), tok)
			return
		}

		var value any
		value, err = decodeValue(dec)
		if err != nil {
			return
		}

		o.Set(key, value)
	}
	tok, err = dec.Token()
	if err != nil {
		err = fmt.Errorf("error decoding closing object token at %d: %w", dec.InputOffset(), err)
		return
	}
	if tok, ok := tok.(json.Delim); !ok || tok != '}' {
		err = fmt.Errorf("unexpected token before end of JSON object at %d: %v", dec.InputOffset(), tok)
		return
	}
	if dec.More() {
		err = fmt.Errorf("unexpected data after JSON object at %d", dec.InputOffset())
		return
	}
	return
}

func decodeValue(dec *json.Decoder) (value any, err error) {

	var raw json.RawMessage
	if err = dec.Decode(&raw); err != nil {
		err = fmt.Errorf("error decoding raw value at %d: %w", dec.InputOffset(), err)
		return
	}

	raw = bytes.TrimSpace(raw)

	if len(raw) > 0 && (raw[0] == '{' || raw[0] == '[') {
		switch raw[0] {
		case '{':
			child := NewJSONMap()
			if err = child.UnmarshalJSON(raw); err != nil {
				return
			}
			value = child
		case '[':
			var array []any
			array, err = decodeJSONArray(raw)
			if err != nil {
				return
			}
			value = array
		default:
			err = fmt.Errorf("unexpected value type at %d: %c", dec.InputOffset(), raw[0])
			return
		}
	} else {
		var decoded any
		if err = json.Unmarshal(raw, &decoded); err != nil {
			err = fmt.Errorf("error decoding value at %d: %w", dec.InputOffset(), err)
			return
		}
		value = decoded
	}
	return
}

func decodeJSONArray(b json.RawMessage) (array []any, err error) {
	dec := json.NewDecoder(bytes.NewReader(b))
	var tok json.Token
	tok, err = dec.Token()
	if err != nil {
		err = fmt.Errorf("error decoding opening array token at %d: %w", dec.InputOffset(), err)
		return
	}
	if tok, ok := tok.(json.Delim); !ok || tok != '[' {
		err = fmt.Errorf("expected JSON array at %d", dec.InputOffset())
		return
	}
	for dec.More() {

		var value any
		value, err = decodeValue(dec)
		if err != nil {
			return
		}

		array = append(array, value)
	}

	tok, err = dec.Token()
	if err != nil {
		err = fmt.Errorf("error decoding closing array token at %d: %w", dec.InputOffset(), err)
		return
	}
	if tok, ok := tok.(json.Delim); !ok || tok != ']' {
		err = fmt.Errorf("unexpected token before end of JSON array at %d: %v", dec.InputOffset(), tok)
		return
	}
	if dec.More() {
		err = fmt.Errorf("unexpected data after JSON object")
		return
	}
	return
}

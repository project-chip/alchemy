package conformance

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Choice struct {
	Set   string      `json:"set,omitempty"`
	Limit ChoiceLimit `json:"limit,omitempty"`
}

func (c *Choice) String() string {
	if c.Limit != nil {
		return c.Limit.String(c.Set)
	}
	return fmt.Sprintf("set: %s", c.Set)
}

func (c *Choice) AsciiDocString() string {
	if c.Limit != nil {
		return fmt.Sprintf("%s%s", c.Set, c.Limit.AsciiDocString())
	}
	return c.Set
}

func (c *Choice) Equal(oc *Choice) bool {
	if c == nil {
		return oc == nil
	} else if oc == nil {
		return false
	}
	if c.Set != oc.Set {
		return false
	}
	if !c.Limit.Equal(oc.Limit) {
		return false
	}
	return true
}

func (c *Choice) Clone() *Choice {
	return &Choice{Set: c.Set, Limit: c.Limit.Clone()}
}

type ChoiceLimit interface {
	String(set string) string
	AsciiDocString() string

	Equal(cl ChoiceLimit) bool
	Clone() ChoiceLimit
}

type ChoiceExactLimit struct {
	Limit int `json:"limit"`
}

func (c *ChoiceExactLimit) String(set string) string {
	return fmt.Sprintf("with exactly %d of set %s", c.Limit, set)
}

func (c *ChoiceExactLimit) AsciiDocString() string {
	if c.Limit <= 1 {
		return ""
	}
	return strconv.Itoa(c.Limit)
}

func (c *ChoiceExactLimit) Equal(cl ChoiceLimit) bool {
	if c == nil {
		return cl == nil
	} else if cl == nil {
		return false
	}
	ocl, ok := cl.(*ChoiceExactLimit)
	if !ok {
		return false
	}
	if c.Limit != ocl.Limit {
		return false
	}
	return true
}

func (c *ChoiceExactLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "exact",
		"limit": c.Limit,
	}
	return json.Marshal(js)
}

func (c *ChoiceExactLimit) Clone() ChoiceLimit {
	return &ChoiceExactLimit{Limit: c.Limit}
}

type ChoiceMinLimit struct {
	Min int `json:"min"`
}

func (c *ChoiceMinLimit) String(set string) string {
	return fmt.Sprintf("with at least %d of set %s", c.Min, set)
}

func (c *ChoiceMinLimit) AsciiDocString() string {
	if c.Min > 1 {
		return fmt.Sprintf("%d+", c.Min)
	}
	return "+"
}

func (c *ChoiceMinLimit) Equal(cl ChoiceLimit) bool {
	if c == nil {
		return cl == nil
	} else if cl == nil {
		return false
	}
	ocl, ok := cl.(*ChoiceMinLimit)
	if !ok {
		return false
	}
	if c.Min != ocl.Min {
		return false
	}
	return true
}

func (c *ChoiceMinLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "min",
		"min":  c.Min,
	}
	return json.Marshal(js)
}

func (c *ChoiceMinLimit) Clone() ChoiceLimit {
	return &ChoiceMinLimit{Min: c.Min}
}

type ChoiceMaxLimit struct {
	Max int `json:"max"`
}

func (c *ChoiceMaxLimit) String(set string) string {
	return fmt.Sprintf("with at most %d of set %s", c.Max, set)
}

func (c *ChoiceMaxLimit) AsciiDocString() string {
	if c.Max > 1 {
		return fmt.Sprintf("%d-", c.Max)
	}
	return "-"
}

func (c *ChoiceMaxLimit) Equal(cl ChoiceLimit) bool {
	if c == nil {
		return cl == nil
	} else if cl == nil {
		return false
	}
	ocl, ok := cl.(*ChoiceMaxLimit)
	if !ok {
		return false
	}
	if c.Max != ocl.Max {
		return false
	}
	return true
}

func (c *ChoiceMaxLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "max",
		"max":  c.Max,
	}
	return json.Marshal(js)
}

func (c *ChoiceMaxLimit) Clone() ChoiceLimit {
	return &ChoiceMaxLimit{Max: c.Max}
}

type ChoiceRangeLimit struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

func (c *ChoiceRangeLimit) String(set string) string {
	return fmt.Sprintf("with between %d and %d of set %s", c.Min, c.Max, set)
}

func (c *ChoiceRangeLimit) AsciiDocString() string {
	return fmt.Sprintf("%d-%d", c.Min, c.Max)
}

func (c *ChoiceRangeLimit) Equal(cl ChoiceLimit) bool {
	if c == nil {
		return cl == nil
	} else if cl == nil {
		return false
	}
	ocl, ok := cl.(*ChoiceRangeLimit)
	if !ok {
		return false
	}
	if c.Min != ocl.Min || c.Max != ocl.Max {
		return false
	}
	return true
}

func (c *ChoiceRangeLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "range",
		"min":  c.Min,
		"max":  c.Max,
	}
	return json.Marshal(js)
}

func (c *ChoiceRangeLimit) Clone() ChoiceLimit {
	return &ChoiceRangeLimit{Min: c.Min, Max: c.Max}
}

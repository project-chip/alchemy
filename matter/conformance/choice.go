package conformance

import (
	"fmt"
)

type Choice struct {
	Set   string
	Limit ChoiceLimit
}

func (c *Choice) String() string {
	if c.Limit != nil {
		return c.Limit.String(c.Set)
	}
	return fmt.Sprintf("set: %s", c.Set)
}

type ChoiceLimit interface {
	String(set string) string
}

type ChoiceExactLimit struct {
	Limit int
}

func (c *ChoiceExactLimit) String(set string) string {
	return fmt.Sprintf("with exactly %d of set %s", c.Limit, set)
}

type ChoiceMinLimit struct {
	Min int
}

func (c *ChoiceMinLimit) String(set string) string {
	return fmt.Sprintf("with at least %d of set %s", c.Min, set)
}

type ChoiceMaxLimit struct {
	Max int
}

func (c *ChoiceMaxLimit) String(set string) string {
	return fmt.Sprintf("with at most %d of set %s", c.Max, set)
}

type ChoiceRangeLimit struct {
	Min int
	Max int
}

func (c *ChoiceRangeLimit) String(set string) string {
	return fmt.Sprintf("with between %d and %d of set %s", c.Min, c.Max, set)
}

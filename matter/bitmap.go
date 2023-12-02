package matter

import (
	"fmt"
	"regexp"

	"github.com/hasty/alchemy/parse"
)

type Bitmap struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Type        *DataType `json:"type,omitempty"`
	Bits        BitSet    `json:"bits,omitempty"`
}

func (c *Bitmap) ModelType() Entity {
	return EntityBitmap
}

type Bit struct {
	Bit         string      `json:"bit,omitempty"`
	Name        string      `json:"name,omitempty"`
	Summary     string      `json:"summary,omitempty"`
	Conformance Conformance `json:"conformance,omitempty"`
}

var bitRangePattern = regexp.MustCompile(`^(?P<From>[0-9]+)\.{2,}(?P<To>[0-9]+)$`)

func (bv *Bit) Mask() (uint64, error) {
	val, err := parse.HexOrDec(bv.Bit)
	if err == nil {
		return 1 << (val), nil
	}
	matches := bitRangePattern.FindStringSubmatch(bv.Bit)
	if len(matches) > 2 {
		from, err := parse.HexOrDec(matches[1])
		if err != nil {
			return 0, err
		}
		to, err := parse.HexOrDec(matches[2])
		if err != nil {
			return 0, err
		}
		if from > to {
			return 0, fmt.Errorf("incorrect order of bit mask range: %d..%d", from, to)
		}
		var val uint64
		for i := from; i <= to; i++ {
			val |= (1 << i)
		}
		return val, err
	}
	return 0, fmt.Errorf("invalid bit mask range: %s", bv.Bit)
}

func (bv *Bit) GetConformance() Conformance {
	return bv.Conformance
}

type BitSet []*Bit

func (bs BitSet) ConformanceReference(name string) HasConformance {
	for _, b := range bs {
		if b.Name == name {
			return b
		}
	}
	return nil
}

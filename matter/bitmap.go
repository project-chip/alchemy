package matter

import (
	"fmt"
	"regexp"

	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/parse"
)

type Bitmap struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Type        *DataType `json:"type,omitempty"`
	Bits        BitSet    `json:"bits,omitempty"`
}

func (c *Bitmap) Entity() Entity {
	return EntityBitmap
}

func (c *Bitmap) Size() int {
	if c.Type == nil {
		return 8
	}
	switch c.Type.BaseType {
	case BaseDataTypeMap64:
		return 64
	case BaseDataTypeMap32:
		return 32
	case BaseDataTypeMap16:
		return 16
	default:
		return 8
	}
}

type Bit struct {
	Bit         string                  `json:"bit,omitempty"`
	Name        string                  `json:"name,omitempty"`
	Summary     string                  `json:"summary,omitempty"`
	Conformance conformance.Conformance `json:"conformance,omitempty"`
}

func (c *Bit) Entity() Entity {
	return EntityBitmapValue
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

func (bv *Bit) GetConformance() conformance.Conformance {
	return bv.Conformance
}

type BitSet []*Bit

func (bs BitSet) ConformanceReference(name string) conformance.HasConformance {
	for _, b := range bs {
		if b.Name == name {
			return b
		}
	}
	return nil
}

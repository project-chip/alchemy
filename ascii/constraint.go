package ascii

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hasty/alchemy/matter"
)

type constraintPatternMatch uint8

const (
	constraintPatternMatchUnknown constraintPatternMatch = iota
	constraintPatternMatchHexFrom
	constraintPatternMatchHexTo
	constraintPatternMatchIntFrom
	constraintPatternMatchIntTo
	constraintPatternMatchTemperatureFrom
	constraintPatternMatchTemperatureTo
	constraintPatternMatchPercentFrom
	constraintPatternMatchPercentTo
	constraintPatternMatchHexLimit
	constraintPatternMatchIntLimit
	constraintPatternMatchMinMax
	constraintPatternMatchItemConstraint
)

var rangePattern = regexp.MustCompile(`^(?:(?P<IntFrom>\-?[0-9]+)|(?P<HexFrom>0[xX][0-9A-Fa-f]+)|(?:(?P<TempFrom>\-?[0-9]+(?:\.[0-9]+)?)°C)|(?:(?P<PercentFrom>\-?[0-9]+(?:\.[0-9]+)?)%))\s+to\s+(?:(?P<IntTo>\-?[0-9]+)|(?P<HexTo>0[xX][0-9A-Fa-f]+)|(?:(?P<TempTo>\-?[0-9]+(?:\.[0-9]+)?)°C)|(?:(?P<PercentTo>\-?[0-9]+(?:\.[0-9]+)?)%))$`)
var rangePatternMatchMap map[int]constraintPatternMatch

var limitPattern = regexp.MustCompile(`(?i)^(?P<MinMax>min|max)[\x{25A0}\x{00A0}\s]+(?:(?P<IntLimit>-?[0-9]+)|(?P<HexLimit>0[xX][0-9A-Fa-f]+))(?: chars)?(?:[\x{25A0}\x{00A0}\s]*\[(?P<ItemConstraint>[^\]]+)\])?$`)
var limitPatternMatchMap map[int]constraintPatternMatch

func init() {
	rangePatternMatchMap = make(map[int]constraintPatternMatch)
	for i, name := range rangePattern.SubexpNames() {
		switch name {
		case "IntFrom":
			rangePatternMatchMap[i] = constraintPatternMatchIntFrom
		case "HexFrom":
			rangePatternMatchMap[i] = constraintPatternMatchHexFrom
		case "TempFrom":
			rangePatternMatchMap[i] = constraintPatternMatchTemperatureFrom
		case "PercentFrom":
			rangePatternMatchMap[i] = constraintPatternMatchPercentFrom
		case "IntTo":
			rangePatternMatchMap[i] = constraintPatternMatchIntTo
		case "HexTo":
			rangePatternMatchMap[i] = constraintPatternMatchHexTo
		case "TempTo":
			rangePatternMatchMap[i] = constraintPatternMatchTemperatureTo
		case "PercentTo":
			rangePatternMatchMap[i] = constraintPatternMatchPercentTo
		}
	}
	limitPatternMatchMap = make(map[int]constraintPatternMatch)
	for i, name := range limitPattern.SubexpNames() {
		switch name {
		case "MinMax":
			limitPatternMatchMap[i] = constraintPatternMatchMinMax
		case "HexLimit":
			limitPatternMatchMap[i] = constraintPatternMatchHexLimit
		case "IntLimit":
			limitPatternMatchMap[i] = constraintPatternMatchIntLimit
		case "ItemConstraint":
			limitPatternMatchMap[i] = constraintPatternMatchItemConstraint
		}
	}
}

func ParseConstraint(parentType *matter.DataType, s string) matter.Constraint {
	s = strings.TrimSpace(s)
	switch strings.ToLower(s) {
	case "all", "none", "any", "empty":
		return &matter.AllConstraint{Value: s}
	case "desc":
		return &matter.DescribedConstraint{}
	}

	c := parseNumConstraint(parentType, s)
	if c != nil {
		return c
	}
	matches := rangePattern.FindStringSubmatch(s)
	if matches != nil {
		c = parseRangeValue(parentType, matches)
		if c != nil {
			return c
		}
		c = parseTempRange(parentType, matches)
		if c != nil {
			return c
		}
		c = parsePercentRange(parentType, matches)
		if c != nil {
			return c
		}
	}
	c = parseLimitValue(parentType, s)
	if c != nil {
		return c
	}

	fmt.Printf("returning generic constraint: %s\n", s)
	return &matter.GenericConstraint{Value: s}
}

func parseNumConstraint(parentType *matter.DataType, s string) matter.Constraint {
	intVal, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		switch parentType.Name {
		case "string", "octstr":
			if intVal >= 0 {
				return &matter.MaxLengthConstraint{Length: &matter.IntLimit{Value: intVal}}
			}
		default:
			return &matter.MaxConstraint{Max: &matter.IntLimit{Value: intVal}}
		}
	}
	return nil
}

func parseRangeValue(parentType *matter.DataType, matches []string) matter.Constraint {

	var fromLimit matter.ConstraintLimit
	var toLimit matter.ConstraintLimit
	for i, s := range matches {
		if s == "" {
			continue
		}
		category, ok := rangePatternMatchMap[i]
		if !ok {
			continue
		}
		switch category {
		case constraintPatternMatchIntFrom:
			minInt, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return nil
			}
			fromLimit = &matter.IntLimit{Value: minInt}
		case constraintPatternMatchHexFrom:
			s = strings.TrimPrefix(s, "0x")
			minHex, err := strconv.ParseUint(s, 16, 64)
			if err != nil {
				return nil
			}
			fromLimit = &matter.HexLimit{Value: minHex}
		case constraintPatternMatchIntTo:
			minInt, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return nil
			}
			toLimit = &matter.IntLimit{Value: minInt}
		case constraintPatternMatchHexTo:
			s = strings.TrimPrefix(s, "0x")
			minHex, err := strconv.ParseUint(s, 16, 64)
			if err != nil {
				return nil
			}
			toLimit = &matter.HexLimit{Value: minHex}
		}
	}
	if fromLimit != nil && toLimit != nil {
		if parentType != nil {
			switch parentType.Name {
			case "string", "octstr":
				return &matter.LengthRangeConstraint{
					Min: fromLimit,
					Max: toLimit,
				}
			}

		}
		return &matter.RangeConstraint{
			Min: fromLimit,
			Max: toLimit,
		}

	}
	return nil
}

func parseLimitValue(parentType *matter.DataType, s string) matter.Constraint {
	matches := limitPattern.FindStringSubmatch(s)
	if matches == nil {
		return nil
	}
	var limit matter.ConstraintLimit
	var minMax string
	var itemConstraint string
	for i, s := range matches {
		if s == "" {
			continue
		}
		category, ok := limitPatternMatchMap[i]
		if !ok {
			continue
		}
		switch category {
		case constraintPatternMatchMinMax:
			minMax = s
		case constraintPatternMatchIntLimit:
			intLimit, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return nil
			}
			limit = &matter.IntLimit{Value: intLimit}
		case constraintPatternMatchHexLimit:
			s = strings.TrimPrefix(s, "0x")
			hexLimit, err := strconv.ParseUint(s, 16, 64)
			if err != nil {
				return nil
			}
			limit = &matter.HexLimit{Value: hexLimit}
		case constraintPatternMatchItemConstraint:
			itemConstraint = s
		}
	}
	if limit == nil {
		return nil
	}
	var minMaxConstraint matter.Constraint
	if minMax == "min" {
		if parentType.IsArray {
			return &matter.MinLengthConstraint{Length: limit}
		}
		switch parentType.Name {
		case "string", "octstr":
			return &matter.MinLengthConstraint{Length: limit}
		default:
			minMaxConstraint = &matter.MinConstraint{Min: limit}
		}
	} else {
		if parentType.IsArray {
			return &matter.MaxLengthConstraint{Length: limit}
		}
		switch parentType.Name {
		case "string", "octstr":
			return &matter.MaxLengthConstraint{Length: limit}
		default:
			minMaxConstraint = &matter.MaxConstraint{Max: limit}
		}
	}
	if len(itemConstraint) > 0 {
		ic := ParseConstraint(parentType, itemConstraint)
		if ic == nil {
			return nil
		}
		return &matter.ListConstraint{
			Constraint:      minMaxConstraint,
			EntryConstraint: ic,
		}
	}
	return minMaxConstraint

}

func parseTempRange(parentType *matter.DataType, matches []string) matter.Constraint {

	var tempFrom *matter.TemperatureLimit
	var tempTo *matter.TemperatureLimit
	for i, s := range matches {
		if s == "" {
			continue
		}
		category, ok := rangePatternMatchMap[i]
		if !ok {
			continue
		}
		switch category {
		case constraintPatternMatchTemperatureFrom:
			minTemp, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return nil
			}
			minTemp *= 100
			tempFrom = &matter.TemperatureLimit{Value: uint64(minTemp)}
		case constraintPatternMatchTemperatureTo:
			maxTemp, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return nil
			}
			maxTemp *= 100
			tempTo = &matter.TemperatureLimit{Value: uint64(maxTemp)}
		}
	}
	if tempFrom != nil && tempTo != nil {
		return &matter.RangeConstraint{
			Min: tempFrom,
			Max: tempTo,
		}

	}
	return nil
}

func parsePercentRange(parentType *matter.DataType, matches []string) matter.Constraint {

	var hundredths bool
	switch parentType.Name {
	case "percent":
	case "percent100ths":
		hundredths = true
	default:
		return nil
	}
	var from *matter.PercentLimit
	var to *matter.PercentLimit
	for i, s := range matches {
		if s == "" {
			continue
		}
		category, ok := rangePatternMatchMap[i]
		if !ok {
			continue
		}
		switch category {
		case constraintPatternMatchTemperatureFrom:
			min, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return nil
			}
			from = &matter.PercentLimit{Hundredths: hundredths}
			if hundredths {
				from.Value = uint64(min * 100)
			} else {
				from.Value = uint64(min)

			}
		case constraintPatternMatchTemperatureTo:
			max, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return nil
			}
			to = &matter.PercentLimit{Hundredths: hundredths}
			if hundredths {
				to.Value = uint64(max * 100)
			} else {
				to.Value = uint64(max)

			}
		}
	}
	if from != nil && to != nil {
		return &matter.RangeConstraint{
			Min: from,
			Max: to,
		}

	}
	return nil
}

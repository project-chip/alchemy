package generate

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/internal/log"
)

func checkConstraints(ts *testStep, indent string, sb *strings.Builder) error {
	if ts.Response.Constraints == nil && ts.Response.Value == nil {
		return nil
	}
	val := "val"
	if ts.Response.SaveAs != "" {
		val = ts.Response.SaveAs
	}
	if ts.Response.Constraints != nil {
		if ts.Response.Constraints.MinValue != nil {
			switch minValue := ts.Response.Constraints.MinValue.(type) {
			case uint64:
				assertComparison(indent, "greater_equal", val, strconv.FormatUint(minValue, 10), sb)
			case string:
				assertComparison(indent, "greater_equal", val, minValue, sb)
			default:
				slog.Info("unknown type in response constraints", log.Type("minValue", minValue), slog.Any("val", minValue))

			}
		}
		if ts.Response.Constraints.MaxValue != nil {
			switch maxValue := ts.Response.Constraints.MaxValue.(type) {
			case uint64:
				assertComparison(indent, "less_equal", val, strconv.FormatUint(maxValue, 10), sb)
			case string:
				assertComparison(indent, "less_equal", val, maxValue, sb)

			default:
				slog.Info("unknown type in response constraints", log.Type("maxValue", maxValue), slog.Any("val", maxValue))

			}
		}
		if ts.Response.Constraints.AnyOf != nil {
			switch anyOf := ts.Response.Constraints.AnyOf.(type) {
			case []any:
				sb.WriteString(indent)
				sb.WriteString("asserts.assert_in(")
				sb.WriteString(val)
				sb.WriteString(", [")
				var vals []string
				for _, v := range anyOf {
					vals = append(vals, fmt.Sprintf("%v", v))
				}
				sb.WriteString(strings.Join(vals, ", "))
				sb.WriteString("])\n")
			}

		}
	}
	if ts.Response.Value != nil {
		assertComparison(indent, "equal", val, fmt.Sprintf("%v", ts.Response.Value), sb)
	}
	return nil
}

func assertComparison(indent string, method string, variable string, comparison string, sb *strings.Builder) {
	sb.WriteString(indent)
	sb.WriteString("asserts.assert_")
	sb.WriteString(method)
	sb.WriteString("(")
	sb.WriteString(variable)
	sb.WriteString(", ")
	sb.WriteString(comparison)
	sb.WriteString(")\n")
}

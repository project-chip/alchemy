package generate

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/internal/log"
)

func readAttribute(ts *testStep, indent string, sb *strings.Builder) error {

	sb.WriteString(indent)
	if ts.Response.SaveAs != "" {
		sb.WriteString(ts.Response.SaveAs)
		sb.WriteString(" = ")
	} else if ts.Response.Constraints != nil || ts.Response.Value != nil {
		sb.WriteString("val = ")
	}
	args := []string{"endpoint=endpoint", "cluster=cluster"}
	args = append(args, fmt.Sprintf("attribute=cluster.Attributes.%s", ts.Attribute))
	if ts.FabricFiltered {
		args = append(args, "fabricFiltered=True")
	}
	if ts.Response.Error != "" {
		args = append(args, fmt.Sprintf("error=Status."+strcase.ToCamel(ts.Response.Error)))
		sb.WriteString("await self.read_single_attribute_expect_error(")
	} else {
		sb.WriteString("await self.read_single_attribute_check_success(")
	}
	sb.WriteString(strings.Join(args, ", "))
	sb.WriteString(")\n")
	checkConstraints(ts, indent, sb)
	return nil
}

func writeAttribute(ts *testStep, indent string, sb *strings.Builder) error {

	sb.WriteString(indent)
	if ts.Response.Error != "" {
		sb.WriteString("status = ")
	}
	sb.WriteString("await self.write_single_attribute(attribute_value=cluster.Attributes.")
	sb.WriteString(ts.Attribute)
	sb.WriteString("(")
	if ts.Arguments.Value != nil {
		switch value := ts.Arguments.Value.(type) {
		case uint64:
			sb.WriteString(strconv.FormatUint(value, 10))
		case int64:
			sb.WriteString(strconv.FormatInt(value, 10))
		case string:
			sb.WriteString(value)
		default:
			slog.Info("unknown type in write arguments", log.Type("value", value), slog.Any("val", value))
		}
	}
	sb.WriteString("), endpoint_id=endpoint")
	if ts.Response.Error != "" {
		sb.WriteString(", expect_success=False")
	}
	sb.WriteString(")\n")
	if ts.Response.Error != "" {
		assertComparison(indent, "equal", "status", "Status."+strcase.ToCamel(ts.Response.Error), sb)
	}
	//await self.write_single_attribute(attribute_value=cluster.Attributes.OccupiedHeatingSetpoint(heatSetpoint-1), endpoint_id=endpoint)

	return nil
}

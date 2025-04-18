package python

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/testscript"
)

func minConstraintHelper(action testscript.TestAction) raymond.SafeString {
	switch action := action.(type) {
	case testscript.CheckMinConstraint:
		return raymond.SafeString(pythonLimit(action.Constraint.Minimum, action.Field))
	case testscript.CheckRangeConstraint:
		return raymond.SafeString(pythonLimit(action.Constraint.Minimum, action.Field))
	default:
		slog.Error("Unexpected type on minConstraint", log.Type("type", action))
	}
	return raymond.SafeString("")
}

func maxConstraintHelper(action testscript.TestAction) raymond.SafeString {
	switch action := action.(type) {
	case testscript.CheckMaxConstraint:
		return raymond.SafeString(pythonLimit(action.Constraint.Maximum, action.Field))
	case testscript.CheckRangeConstraint:
		return raymond.SafeString(pythonLimit(action.Constraint.Maximum, action.Field))
	default:
		slog.Error("Unexpected type on maxConstraint", log.Type("type", action))
	}
	return raymond.SafeString("")
}

func pythonLimit(l constraint.Limit, field *matter.Field) string {
	var sb strings.Builder
	buildPythonLimit(l, field, &sb)
	return sb.String()
}

func buildPythonLimit(l constraint.Limit, field *matter.Field, builder *strings.Builder) {
	switch l := l.(type) {
	case *constraint.IdentifierLimit:
		switch entity := l.Entity.(type) {
		case *matter.Field:
			switch entity.EntityType() {
			case types.EntityTypeAttribute:
				builder.WriteString("self.")
			default:
				builder.WriteString("struct.")
			}
			builder.WriteString(entity.Name)
		case nil:
			slog.Warn("Missing entity when evaluating identifier limit", slog.String("id", l.ID), slog.String("fieldName", field.Name), log.Path("source", field))
		default:
			slog.Warn("Unexpected entity type when evaluating identifier limit", log.Type("type", entity), slog.String("id", l.ID), slog.String("fieldName", field.Name), log.Path("source", field))
		}
	case *constraint.ReferenceLimit:
		switch entity := l.Entity.(type) {
		case *matter.Field:
			switch entity.EntityType() {
			case types.EntityTypeAttribute:
				builder.WriteString("self.")
			default:
				builder.WriteString("struct.")
			}
			builder.WriteString(entity.Name)
		case nil:
			slog.Warn("Missing entity when evaluating reference limit", slog.String("reference", l.Reference), slog.String("fieldName", field.Name), log.Path("source", field))
		default:
			slog.Warn("Unexpected entity type when evaluating reference limit", log.Type("type", entity), slog.String("reference", l.Reference), slog.String("fieldName", field.Name), log.Path("source", field))
		}
	case *constraint.TagIdentifierLimit:
		switch entity := l.Entity.(type) {
		case *matter.Field:
			switch entity.EntityType() {
			case types.EntityTypeAttribute:
				builder.WriteString("self.")
			default:
				builder.WriteString("struct.")
			}
			builder.WriteString(entity.Name)
		case nil:
			slog.Warn("Missing entity when evaluating tag identifier limit", slog.String("tag", l.Tag), slog.String("fieldName", field.Name), log.Path("source", field))
		default:
			slog.Warn("Unexpected entity type when evaluating tag identifier limit", log.Type("type", entity), slog.String("tag", l.Tag), slog.String("fieldName", field.Name), log.Path("source", field))
		}
	case *constraint.MathExpressionLimit:
		buildPythonLimit(l.Left, field, builder)
		builder.WriteRune(' ')
		builder.WriteString(l.Operand)
		builder.WriteRune(' ')
		buildPythonLimit(l.Right, field, builder)
	case *constraint.LogicalLimit:
		buildPythonLimit(l.Left, field, builder)
		for _, r := range l.Right {
			builder.WriteRune(' ')
			switch l.Operand {
			case "|":
				builder.WriteString("or")
			case "&":
				builder.WriteString("and")
			}
			builder.WriteRune(' ')
			buildPythonLimit(r, field, builder)
		}
	case *constraint.IntLimit:
		builder.WriteString(l.DataModelString(field.Type))
	case *constraint.HexLimit:
		builder.WriteString(strconv.FormatUint(l.Value, 10))
	case *constraint.TemperatureLimit:
		builder.WriteString(l.DataModelString(field.Type))
	case *constraint.PercentLimit:
		builder.WriteString(l.Value.String())
		return
	case *constraint.ExpLimit:
		builder.WriteString(fmt.Sprintf("%de%d", l.Value, l.Exp))
		return
	case *constraint.ManufacturerLimit:
	case *constraint.CharacterLimit:
		builder.WriteString(l.ByteCount.DataModelString(field.Type))
	default:
		slog.Warn("Unexpected limit type for Python limit", log.Type("type", l), slog.String("fieldName", field.Name), log.Path("source", field))
	}
}

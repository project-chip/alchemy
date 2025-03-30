package python

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/testscript"
)

type conformanceContext struct {
	e types.Entity
}

func (cc *conformanceContext) Entity() types.Entity {
	return cc.e
}

func needsConformanceCheckHelper(action testscript.TestAction, options *raymond.Options) string {
	switch action := action.(type) {
	case *testscript.ReadAttribute:
		if action.Attribute == nil {
			slog.Info("nil attribute")
			return options.Inverse()
		}
		switch action.Attribute.Conformance {
		case nil:
			slog.Info("nil attribute conformance")
			return options.Inverse()
		default:
			slog.Info("non-nil attribute conformance", slog.String("attr", action.Attribute.Name), slog.String("c", action.Attribute.Conformance.ASCIIDocString()))
			if conformance.IsMandatory(action.Attribute.Conformance) {
				slog.Info("non-nil attribute conformance is mandatory", slog.String("c", action.Attribute.Conformance.ASCIIDocString()))
				return options.Inverse()
			}
			return options.Fn()
		}
	default:
		return options.Inverse()
	}
}

func conformanceGuardHelper(action testscript.TestAction) raymond.SafeString {
	var sb strings.Builder
	switch action := action.(type) {
	case *testscript.ReadAttribute:
		slog.Info("conformance guard read attribute", slog.String("attr", action.Attribute.Name))
		err := buildPythonConformance(action.Attribute.Conformance, action.Attribute, &sb)
		if err != nil {
			slog.Error("Error building conformance", slog.Any("error", err))
			return raymond.SafeString("True")
		}
	case *testscript.CheckType:
	default:
		slog.Error("Unexpected action type", log.Type("type", action))
	}
	return raymond.SafeString(sb.String())
}

func buildPythonConformance(c conformance.Conformance, entity types.Entity, builder *strings.Builder) error {
	switch c := c.(type) {
	case *conformance.Deprecated:
		return fmt.Errorf("deprecated conformance cannot be converted to Python")
	case *conformance.Described:
		return fmt.Errorf("described conformance cannot be converted to Python")
	case *conformance.Generic:
		return fmt.Errorf("generic conformance cannot be converted to Python")
	case *conformance.Mandatory:
		return buildPythonMandatoryConformance(c, entity, builder)
	case *conformance.Optional:
		return buildPythonOptionalConformance(c, entity, builder)
	case conformance.Set:
		var count int
		for _, c := range c {
			switch c := c.(type) {
			case *conformance.Provisional:
				continue

			default:
				if count > 0 {
					builder.WriteString(" and ")
				}
				if err := buildPythonConformance(c, entity, builder); err != nil {
					return err
				}
				count++
			}
		}
	default:
		return fmt.Errorf("unexpected conformance converting to Python: %T", c)
	}
	return nil
}

func buildPythonMandatoryConformance(o *conformance.Mandatory, entity types.Entity, builder *strings.Builder) error {
	if o.Expression == nil {
		builder.WriteString("True")
		return nil
	}
	return buildPythonConformanceExpression(o.Expression, entity, builder)
}

func buildPythonOptionalConformance(o *conformance.Optional, entity types.Entity, builder *strings.Builder) error {
	slog.Info("conformance guard optional conformance", log.Type("entity", entity))
	switch entity := entity.(type) {
	case *matter.Field:
		switch entity.EntityType() {
		case types.EntityTypeAttribute:
			builder.WriteString("await self.attribute_guard(endpoint=endpoint, attribute=attributes.")
			builder.WriteString(entity.Name)
			builder.WriteString(")")
		case types.EntityTypeStructField:
			builder.WriteString("struct.")
			builder.WriteString(entity.Name)
			builder.WriteString(" is not None")
		default:
			return fmt.Errorf("unexpected field type when converting optional conformance to Python: %s", entity.EntityType().String())
		}
	case nil:
		return fmt.Errorf("unexpected nil entity when converting optional conformance to Python: %T", entity)
	default:
		return fmt.Errorf("unexpected entity type when converting optional conformance to Python: %T", entity)
	}
	if o.Expression == nil {
		return nil
	}
	builder.WriteString(" and ")
	return buildPythonConformanceExpression(o.Expression, entity, builder)
}

func buildPythonConformanceExpression(e conformance.Expression, entity types.Entity, builder *strings.Builder) error {
	switch e := e.(type) {
	case *conformance.LogicalExpression:
		builder.WriteRune('(')
		switch e.Operand {
		case "|":
			if e.Not {
				builder.WriteRune('!')
				if err := buildPythonConformanceExpression(e.Left, entity, builder); err != nil {
					return err
				}
				for _, r := range e.Right {
					builder.WriteString(" and !")
					if err := buildPythonConformanceExpression(r, entity, builder); err != nil {
						return err
					}
				}
			} else {
				if err := buildPythonConformanceExpression(e.Left, entity, builder); err != nil {
					return err
				}
				for _, r := range e.Right {
					builder.WriteString(" or ")
					if err := buildPythonConformanceExpression(r, entity, builder); err != nil {
						return err
					}
				}

			}
		case "&":
			if err := buildPythonConformanceExpression(e.Left, entity, builder); err != nil {
				return err
			}
			for _, r := range e.Right {
				if e.Not {
					builder.WriteString(" or ")
				} else {
					builder.WriteString(" and ")
				}
				if err := buildPythonConformanceExpression(r, entity, builder); err != nil {
					return err
				}
			}
		case "^":
			if e.Not {
				builder.WriteString("!")
			}
			builder.WriteRune('(')
			if err := buildPythonConformanceExpression(e.Left, entity, builder); err != nil {
				return err
			}
			for _, r := range e.Right {

				builder.WriteString(" ^ ")
				if err := buildPythonConformanceExpression(r, entity, builder); err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("unknown operator: %s", e.Operand)
		}
		builder.WriteRune(')')
		return nil
	case *conformance.FeatureExpression:
		var feature *matter.Feature
		switch fe := e.Entity.(type) {
		case *matter.Feature:
			feature = fe
		case nil:
			return fmt.Errorf("conformance feature expression missing entity: %s", e.Feature)
		default:
			return fmt.Errorf("unexpected entity type when converting feature expression to Python: %T", fe)
		}

		builder.WriteString("await self.feature_guard(endpoint=endpoint, cluster=cluster, feature_int=cluster.Bitmaps.Feature.k")
		builder.WriteString(feature.Name())
		builder.WriteString(")")
		return nil
	default:
		return fmt.Errorf("unimplemented conformance expression converting to Python: %T", e)

	}
}

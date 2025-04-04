package python

import (
	"fmt"
	"log/slog"
	"strconv"
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
		if needsConformanceCheck(action.Attribute.Conformance) {
			return options.Fn()
		}
		return options.Inverse()
	case *testscript.WriteAttribute:
		if needsConformanceCheck(action.Attribute.Conformance) {
			return options.Fn()
		}
		return options.Inverse()
	default:
		return options.Inverse()
	}
}

func needsConformanceCheck(c conformance.Conformance) bool {
	if c == nil {
		return false
	}
	switch c {
	default:
		if conformance.IsDeprecated(c) {
			return false
		}
		if conformance.IsMandatory(c) {
			return false
		}
		if conformance.IsDescribed(c) {
			return false
		}
		return true
	}
}

func conformanceGuardHelper(action testscript.TestAction) raymond.SafeString {
	var sb strings.Builder
	switch action := action.(type) {
	case *testscript.ReadAttribute:
		err := buildPythonConformance(action.Cluster, action.Conformance, action.Attribute, &sb)
		if err != nil {
			slog.Error("Error building conformance", slog.Any("error", err))
			return raymond.SafeString("True")
		}
	case *testscript.WriteAttribute:
		err := buildPythonConformance(action.Cluster, action.Conformance, action.Attribute, &sb)
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

func buildPythonConformance(cluster *matter.Cluster, c conformance.Conformance, entity types.Entity, builder *strings.Builder) error {
	switch c := c.(type) {
	case *conformance.Deprecated:
		return fmt.Errorf("deprecated conformance cannot be converted to Python")
	case *conformance.Described:
		return fmt.Errorf("described conformance cannot be converted to Python")
	case *conformance.Generic:
		return fmt.Errorf("generic conformance cannot be converted to Python")
	case *conformance.Mandatory:
		return buildPythonMandatoryConformance(cluster, c, entity, builder)
	case *conformance.Optional:
		return buildPythonOptionalConformance(cluster, c, entity, builder)
	case conformance.Set:
		var count int
		for _, c := range c {
			switch c := c.(type) {
			case *conformance.Provisional:
				continue

			default:
				if conformance.IsDeprecated(c) {
					continue
				}
				if count > 0 {
					builder.WriteString(" and ")
				}
				if err := buildPythonConformance(cluster, c, entity, builder); err != nil {
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

func buildPythonMandatoryConformance(cluster *matter.Cluster, o *conformance.Mandatory, entity types.Entity, builder *strings.Builder) error {
	if o.Expression == nil {
		builder.WriteString("True")
		return nil
	}
	return buildPythonConformanceExpression(cluster, o.Expression, entity, builder)
}

func buildPythonOptionalConformance(cluster *matter.Cluster, o *conformance.Optional, entity types.Entity, builder *strings.Builder) error {
	switch entity := entity.(type) {
	case *matter.Field:
		switch entity.EntityType() {
		case types.EntityTypeAttribute:
			builder.WriteString("await self.attribute_guard(endpoint=endpoint, attribute=attributes.")
			builder.WriteString(entity.Name)
			builder.WriteString(")")
		case types.EntityTypeStructField:
			builder.WriteString("struct.")
			builder.WriteString(matter.CamelCase(entity.Name))
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
	return buildPythonConformanceExpression(cluster, o.Expression, entity, builder)
}

func buildPythonConformanceExpression(cluster *matter.Cluster, e conformance.Expression, entity types.Entity, builder *strings.Builder) error {
	switch e := e.(type) {
	case *conformance.ComparisonExpression:
		if err := buildPythonConformanceCompareValue(e.Left, entity, builder); err != nil {
			return err
		}
		builder.WriteString(" ")
		builder.WriteString(e.Op.String())
		builder.WriteString(" ")
		if err := buildPythonConformanceCompareValue(e.Right, entity, builder); err != nil {
			return err
		}
		return nil
	case *conformance.LogicalExpression:
		builder.WriteRune('(')
		switch e.Operand {
		case "|":
			if e.Not {
				builder.WriteRune('!')
				if err := buildPythonConformanceExpression(cluster, e.Left, entity, builder); err != nil {
					return err
				}
				for _, r := range e.Right {
					builder.WriteString(" and !")
					if err := buildPythonConformanceExpression(cluster, r, entity, builder); err != nil {
						return err
					}
				}
			} else {
				if err := buildPythonConformanceExpression(cluster, e.Left, entity, builder); err != nil {
					return err
				}
				for _, r := range e.Right {
					builder.WriteString(" or ")
					if err := buildPythonConformanceExpression(cluster, r, entity, builder); err != nil {
						return err
					}
				}

			}
		case "&":
			if err := buildPythonConformanceExpression(cluster, e.Left, entity, builder); err != nil {
				return err
			}
			for _, r := range e.Right {
				if e.Not {
					builder.WriteString(" or ")
				} else {
					builder.WriteString(" and ")
				}
				if err := buildPythonConformanceExpression(cluster, r, entity, builder); err != nil {
					return err
				}
			}
		case "^":
			if e.Not {
				builder.WriteString("!")
			}
			builder.WriteRune('(')
			if err := buildPythonConformanceExpression(cluster, e.Left, entity, builder); err != nil {
				return err
			}
			for _, r := range e.Right {

				builder.WriteString(" ^ ")
				if err := buildPythonConformanceExpression(cluster, r, entity, builder); err != nil {
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
			slog.Error("Unexpected entity type when converting feature expression to Python", log.Type("type", fe), matter.LogEntity("entity", fe))
			return fmt.Errorf("unexpected entity type when converting feature expression to Python: %T", fe)
		}

		builder.WriteString("await self.feature_guard(endpoint=endpoint, cluster=cluster, feature_int=cluster.Bitmaps.Feature.k")
		builder.WriteString(feature.Name())
		builder.WriteString(")")
		return nil
	case *conformance.IdentifierExpression:
		switch ie := e.Entity.(type) {
		case *matter.Field:
			switch ie.EntityType() {
			case types.EntityTypeAttribute:
				builder.WriteString("await self.attribute_guard(endpoint=endpoint, attribute=attributes.")
				builder.WriteString(ie.Name)
				builder.WriteString(")")
				return nil
			default:
				return fmt.Errorf("unexpected field entity type when converting identifier expression to Python: %s", ie.EntityType().String())
			}
		case *matter.Condition:
			slog.Error("Condition identifier conformance expression not supported", slog.String("id", e.ID), log.Path("source", ie))
			return fmt.Errorf("condition identifier expression not supported: %s", e.ID)
		case nil:
			return fmt.Errorf("conformance identifier expression missing entity: %s", e.ID)
		default:
			return fmt.Errorf("unexpected entity type when converting identifier expression to Python: %T", ie)
		}
	default:
		return fmt.Errorf("unimplemented conformance expression converting to Python: %T", e)

	}
}

func buildPythonConformanceCompareValue(cv conformance.ComparisonValue, entity types.Entity, builder *strings.Builder) error {
	switch cv := cv.(type) {
	case *conformance.IdentifierValue:
		builder.WriteString(cv.ID)
		if cv.Field != nil {
			builder.WriteString(".")
			buildPythonConformanceCompareValue(cv.Field, cv.Entity, builder)
		}
	case *conformance.IntValue:
		builder.WriteString(strconv.FormatInt(cv.Int, 10))
	default:
		return fmt.Errorf("unimplemented conformance compare value converting to Python: %T", cv)
	}
	return nil
}

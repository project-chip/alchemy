package render

import (
	"fmt"
	"strconv"

	"github.com/iancoleman/strcase"
	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func ifIsConstraintHelper(c constraint.Constraint, cons string, options *raymond.Options) raymond.SafeString {
	var isConstraint bool
	switch cons {
	case "all":
		_, isConstraint = c.(*constraint.AllConstraint)
	case "described":
		_, isConstraint = c.(*constraint.DescribedConstraint)
	case "exact":
		_, isConstraint = c.(*constraint.ExactConstraint)
	case "generic":
		_, isConstraint = c.(*constraint.GenericConstraint)
	case "list":
		_, isConstraint = c.(*constraint.ListConstraint)
	case "max":
		_, isConstraint = c.(*constraint.MaxConstraint)
	case "min":
		_, isConstraint = c.(*constraint.MinConstraint)
	case "range":
		_, isConstraint = c.(*constraint.RangeConstraint)
	}
	if isConstraint {
		return raymond.SafeString(options.Fn())
	}
	return raymond.SafeString(options.Inverse())
}

func constraintHelper(c constraint.Constraint, dataType types.DataType) raymond.SafeString {
	if c == nil {
		return raymond.SafeString("")
	}
	switch c.(type) {
	case *constraint.AllConstraint, *constraint.ExactConstraint:
		return raymond.SafeString("")
	case *constraint.RangeConstraint:
		return raymond.SafeString("{valrange} " + c.ASCIIDocString(&dataType))
	}
	return raymond.SafeString("")
	//return raymond.SafeString(fmt.Sprintf("unimplemented constraint type: %T", c))
}

func listConstraintHelper(c constraint.Constraint) raymond.SafeString {
	if c == nil {
		return raymond.SafeString("")
	}
	switch c.(type) {
	case *constraint.AllConstraint, *constraint.ExactConstraint:
		return raymond.SafeString("")
	case *constraint.RangeConstraint:
		return raymond.SafeString("{valrange}")
	}
	return raymond.SafeString("")
	//return raymond.SafeString(fmt.Sprintf("unimplemented constraint type: %T", c))
}

func limitString(cluster *matter.Cluster, limit constraint.Limit) string {
	switch limit := limit.(type) {
	case *constraint.BooleanLimit:
		return strconv.FormatBool(limit.Value)
	case *constraint.ExpLimit:
		return fmt.Sprintf("%d^%d^", limit.Value, limit.Exp)
	case *constraint.HexLimit:
		return fmt.Sprintf("0x%X", limit.Value)
	case *constraint.IntLimit:
		return strconv.FormatInt(limit.Value, 10)
	case *constraint.TemperatureLimit:
		return fmt.Sprintf("%s°C", limit.Value.String())
	case *constraint.ReferenceLimit:
		switch ref := limit.Entity.(type) {
		case *matter.Field:
			return fmt.Sprintf("{A_%s}", strcase.ToScreamingSnake(ref.Name))
		default:
			return fmt.Sprintf("ERR: unknown reference type %T (%s)", ref, limit.Reference)
		}
	default:
		return fmt.Sprintf("unknown limit type: %T", limit)
	}
}

func limitHelper(cluster *matter.Cluster, l constraint.Limit) raymond.SafeString {
	return raymond.SafeString(limitString(cluster, l))
}

func minimumHelper(cluster *matter.Cluster, c constraint.Constraint) raymond.SafeString {
	switch c := c.(type) {
	case *constraint.RangeConstraint:
		return raymond.SafeString(limitString(cluster, c.Minimum))
	case *constraint.MinConstraint:
		return raymond.SafeString(limitString(cluster, c.Minimum))
	default:
		return raymond.SafeString(fmt.Sprintf("unknown minimum constraint: %T", c))
	}
}

func maximumHelper(cluster *matter.Cluster, c constraint.Constraint) raymond.SafeString {
	switch c := c.(type) {
	case *constraint.RangeConstraint:
		return raymond.SafeString(limitString(cluster, c.Maximum))
	case *constraint.MaxConstraint:
		return raymond.SafeString(limitString(cluster, c.Maximum))
	default:
		return raymond.SafeString(fmt.Sprintf("unknown maximum constraint: %T", c))
	}
}

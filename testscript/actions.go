package testscript

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
)

type TestAction interface {
}

type action struct {
	Comments    []string
	Description []string

	Conformance conformance.Conformance
}

type remoteAction struct {
	action

	Cluster       *matter.Cluster
	Endpoint      uint64
	ExpectedError string
}

type ReadAttribute struct {
	remoteAction

	AttributeName string
	Attribute     *matter.Field
	Attributes    matter.FieldSet

	Variable string

	Validations []TestAction
}

type WriteAttribute struct {
	remoteAction

	AttributeName string
	Attribute     *matter.Field

	Value any
}

type SubscribeAttribute struct {
	remoteAction

	AttributeName string
	Attribute     *matter.Field

	MinInterval uint64
	MaxInterval uint64
	Timeout     uint64
}

type TestEventTrigger struct {
	remoteAction

	EventTrigger string
	EnableKey    string
}

type constraintAction struct {
	action

	Field    *matter.Field
	FieldSet matter.FieldSet

	Variable string
}

type CheckType struct {
	constraintAction
}

type CheckMinConstraint struct {
	constraintAction

	Constraint *constraint.MinConstraint
}

type CheckMaxConstraint struct {
	constraintAction

	Constraint *constraint.MaxConstraint
}

type CheckRangeConstraint struct {
	constraintAction

	Constraint *constraint.RangeConstraint
}

type CheckAnyOfConstraint struct {
	constraintAction

	Values any
}

type CheckValueConstraint struct {
	constraintAction

	Value any
}

type CheckNotValueConstraint struct {
	constraintAction

	Value any
}

type CheckStruct struct {
	action

	Struct *matter.Struct

	Fields []*CheckStructField
}

type CheckStructField struct {
	action

	Field *matter.Field

	Validations []TestAction
}

type CheckListEntries struct {
	constraintAction

	Validations []TestAction
}

type CallCommand struct {
	remoteAction

	Cluster   *matter.Cluster
	Command   *matter.Command
	Arguments []*CommandArgument

	Variable string

	Validations []TestAction
}

type CommandArgument struct {
	Field *matter.Field
	Value any
}
type WaitForCommissionee struct {
}

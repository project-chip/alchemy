package spec

import (
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

type ErrorType uint16

const (
	ErrorTypeUnknown ErrorType = iota
	ErrorTypeDuplicateEntityID
	ErrorTypeDuplicateEntityName
	ErrorTypeUnknownConstraintIdentifier
	ErrorTypeUnknownConstraintReference
	ErrorTypeUnknownCustomDataType
	ErrorTypeUnknownClusterRequirement
	ErrorTypeUnknownElementRequirementCluster
	ErrorTypeUnknownBaseCluster
	ErrorTypeUnknownConformanceIdentifier
	ErrorTypeUnknownConformanceReference
)

type Error interface {
	Type() ErrorType
}

func (s *Specification) addError(e Error) {
	if s == nil {
		return
	}
	s.Errors = append(s.Errors, e)
}

type DuplicateEntityIDError struct {
	Entity   types.Entity
	Previous types.Entity
}

func (ddt DuplicateEntityIDError) Type() ErrorType {
	return ErrorTypeDuplicateEntityID
}

type DuplicateEntityNameError struct {
	Entity   types.Entity
	Previous types.Entity
}

func (ddt DuplicateEntityNameError) Type() ErrorType {
	return ErrorTypeDuplicateEntityName
}

type UnknownConstraintIdentifierError struct {
	Identifier *constraint.IdentifierLimit
	Source     log.Source
}

func (ddt UnknownConstraintIdentifierError) Type() ErrorType {
	return ErrorTypeUnknownConstraintIdentifier
}

type UnknownConstraintReferenceError struct {
	Reference *constraint.ReferenceLimit
	Source    log.Source
}

func (ddt UnknownConstraintReferenceError) Type() ErrorType {
	return ErrorTypeUnknownConstraintReference
}

type UnknownCustomDataTypeError struct {
	Field    *matter.Field
	DataType *types.DataType
}

func (ddt UnknownCustomDataTypeError) Type() ErrorType {
	return ErrorTypeUnknownCustomDataType
}

type UnknownClusterRequirementError struct {
	Requirement *matter.ClusterRequirement
}

func (ddt UnknownClusterRequirementError) Type() ErrorType {
	return ErrorTypeUnknownClusterRequirement
}

type UnknownElementRequirementClusterError struct {
	Requirement *matter.ElementRequirement
}

func (ddt UnknownElementRequirementClusterError) Type() ErrorType {
	return ErrorTypeUnknownElementRequirementCluster
}

type UnknownBaseClusterError struct {
	Cluster *matter.Cluster
}

func (ddt UnknownBaseClusterError) Type() ErrorType {
	return ErrorTypeUnknownBaseCluster
}

type UnknownConformanceIdentifierError struct {
	Cluster *matter.Cluster
}

func (cf UnknownConformanceIdentifierError) Type() ErrorType {
	return ErrorTypeUnknownConformanceIdentifier
}

type UnknownConformanceReferenceError struct {
	Cluster *matter.Cluster
}

func (cf UnknownConformanceReferenceError) Type() ErrorType {
	return ErrorTypeUnknownConformanceReference
}

package spec

import (
	"fmt"
	"strings"

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
	ErrorTypeFabricScopingNotAllowed
	ErrorTypeFabricSensitivityNotAllowed
	ErrorTypeFabricScopedStructNotAllowed
	ErrorTypeInvalidConformance
	ErrorTypeInvalidConstraint
	ErrorTypeInvalidFallback
)

type Error interface {
	Type() ErrorType
	Error() string
	log.Source
}

type ParseErrors struct {
	Errors []Error
}

func (pe ParseErrors) Error() string {
	var sb strings.Builder
	for _, e := range pe.Errors {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(e.Error())
	}
	return fmt.Sprintf("spec parsing errors:%s", sb.String())
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

func (ddt DuplicateEntityIDError) Origin() (path string, line int) {
	return ddt.Entity.Origin()
}

func (ddt DuplicateEntityIDError) Error() string {
	return fmt.Sprintf("duplicate %s ID: %s", ddt.Entity.EntityType().String(), matter.EntityName(ddt.Entity))
}

type DuplicateEntityNameError struct {
	Entity   types.Entity
	Previous types.Entity
}

func (ddt DuplicateEntityNameError) Type() ErrorType {
	return ErrorTypeDuplicateEntityName
}

func (ddt DuplicateEntityNameError) Origin() (path string, line int) {
	return ddt.Entity.Origin()
}

func (ddt DuplicateEntityNameError) Error() string {
	return fmt.Sprintf("duplicate %s name: %s", ddt.Entity.EntityType().String(), matter.EntityName(ddt.Entity))
}

type UnknownConstraintIdentifierError struct {
	Identifier *constraint.IdentifierLimit
	Source     log.Source
}

func (ddt UnknownConstraintIdentifierError) Type() ErrorType {
	return ErrorTypeUnknownConstraintIdentifier
}

func (ddt UnknownConstraintIdentifierError) Origin() (path string, line int) {
	return ddt.Source.Origin()
}

func (ddt UnknownConstraintIdentifierError) Error() string {
	return fmt.Sprintf("unknown constraint identifier: %s", ddt.Identifier.ID)
}

type UnknownConstraintReferenceError struct {
	Reference *constraint.ReferenceLimit
	Source    log.Source
}

func (ddt UnknownConstraintReferenceError) Type() ErrorType {
	return ErrorTypeUnknownConstraintReference
}

func (ddt UnknownConstraintReferenceError) Origin() (path string, line int) {
	return ddt.Source.Origin()
}

func (ddt UnknownConstraintReferenceError) Error() string {
	return fmt.Sprintf("unknown constraint reference: %s", ddt.Reference.Reference)
}

type UnknownCustomDataTypeError struct {
	Field    *matter.Field
	DataType *types.DataType
}

func (ddt UnknownCustomDataTypeError) Type() ErrorType {
	return ErrorTypeUnknownCustomDataType
}

func (ddt UnknownCustomDataTypeError) Origin() (path string, line int) {
	return ddt.Field.Origin()
}

func (ddt UnknownCustomDataTypeError) Error() string {
	return fmt.Sprintf("unknown custom data type: %s", ddt.DataType.Name)
}

type UnknownClusterRequirementError struct {
	Requirement *matter.ClusterRequirement
}

func (ddt UnknownClusterRequirementError) Type() ErrorType {
	return ErrorTypeUnknownClusterRequirement
}

func (ddt UnknownClusterRequirementError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt UnknownClusterRequirementError) Error() string {
	return fmt.Sprintf("unknown cluster requirement: %s", ddt.Requirement.ClusterName)
}

type UnknownElementRequirementClusterError struct {
	Requirement *matter.ElementRequirement
}

func (ddt UnknownElementRequirementClusterError) Type() ErrorType {
	return ErrorTypeUnknownElementRequirementCluster
}

func (ddt UnknownElementRequirementClusterError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt UnknownElementRequirementClusterError) Error() string {
	return fmt.Sprintf("unknown element requirement cluster: %s", ddt.Requirement.ClusterName)
}

type UnknownBaseClusterError struct {
	Cluster *matter.Cluster
}

func (ddt UnknownBaseClusterError) Type() ErrorType {
	return ErrorTypeUnknownBaseCluster
}

func (ddt UnknownBaseClusterError) Origin() (path string, line int) {
	return ddt.Cluster.Origin()
}

func (ddt UnknownBaseClusterError) Error() string {
	return fmt.Sprintf("unknown base cluster: %s", ddt.Cluster.Hierarchy)
}

type UnknownConformanceIdentifierError struct {
	Entity     types.Entity
	Identifier string
}

func (cf UnknownConformanceIdentifierError) Type() ErrorType {
	return ErrorTypeUnknownConformanceIdentifier
}

func (cf UnknownConformanceIdentifierError) Origin() (path string, line int) {
	return cf.Entity.Origin()
}

func (ddt UnknownConformanceIdentifierError) Error() string {
	return fmt.Sprintf("unknown conformance identifier: %s", ddt.Identifier)
}

type UnknownConformanceReferenceError struct {
	Entity    types.Entity
	Reference string
}

func (cf UnknownConformanceReferenceError) Type() ErrorType {
	return ErrorTypeUnknownConformanceReference
}

func (cf UnknownConformanceReferenceError) Origin() (path string, line int) {
	return cf.Entity.Origin()
}

func (ddt UnknownConformanceReferenceError) Error() string {
	return fmt.Sprintf("unknown conformance reference: %s", ddt.Reference)
}

type FabricScopingNotAllowedError struct {
	Entity types.Entity
}

func (cf FabricScopingNotAllowedError) Type() ErrorType {
	return ErrorTypeFabricScopingNotAllowed
}

func (cf FabricScopingNotAllowedError) Origin() (path string, line int) {
	return cf.Entity.Origin()
}

func (ddt FabricScopingNotAllowedError) Error() string {
	return fmt.Sprintf("fabric scoping not allowed: %s", matter.EntityName(ddt.Entity))
}

type FabricSensitivityNotAllowedError struct {
	Entity types.Entity
}

func (cf FabricSensitivityNotAllowedError) Type() ErrorType {
	return ErrorTypeFabricSensitivityNotAllowed
}

func (cf FabricSensitivityNotAllowedError) Origin() (path string, line int) {
	return cf.Entity.Origin()
}

func (ddt FabricSensitivityNotAllowedError) Error() string {
	return fmt.Sprintf("fabric sensitivity not allowed: %s", matter.EntityName(ddt.Entity))
}

type FabricScopedStructNotAllowedError struct {
	Entity types.Entity
}

func (cf FabricScopedStructNotAllowedError) Type() ErrorType {
	return ErrorTypeFabricScopedStructNotAllowed
}

func (cf FabricScopedStructNotAllowedError) Origin() (path string, line int) {
	return cf.Entity.Origin()
}

func (ddt FabricScopedStructNotAllowedError) Error() string {
	return fmt.Sprintf("fabric scoped struct not allowed: %s", matter.EntityName(ddt.Entity))
}

type InvalidConformanceError struct {
	Conformance string
	Source      log.Source
}

func (cf InvalidConformanceError) Type() ErrorType {
	return ErrorTypeInvalidConformance
}

func (cf InvalidConformanceError) Origin() (path string, line int) {
	return cf.Source.Origin()
}

func (ddt InvalidConformanceError) Error() string {
	return fmt.Sprintf("invalid conformance: \"%s\"", ddt.Conformance)
}

type InvalidConstraintError struct {
	Constraint string
	Source     log.Source
}

func (cf InvalidConstraintError) Type() ErrorType {
	return ErrorTypeInvalidConstraint
}

func (cf InvalidConstraintError) Origin() (path string, line int) {
	return cf.Source.Origin()
}

func (ddt InvalidConstraintError) Error() string {
	return fmt.Sprintf("invalid constraint: \"%s\"", ddt.Constraint)
}

type InvalidFallbackError struct {
	Fallback string
	Source   log.Source
}

func (cf InvalidFallbackError) Type() ErrorType {
	return ErrorTypeInvalidConstraint
}

func (cf InvalidFallbackError) Origin() (path string, line int) {
	return cf.Source.Origin()
}

func (ddt InvalidFallbackError) Error() string {
	return fmt.Sprintf("invalid fallback: \"%s\"", ddt.Fallback)
}

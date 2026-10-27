package spec

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

type ErrorType uint16

const (
	ErrorTypeUnknown ErrorType = iota
	ErrorTypeGenericParse
	ErrorTypeDuplicateEntityID
	ErrorTypeDuplicateEntityName
	ErrorTypeUnknownConstraintIdentifier
	ErrorTypeUnknownConstraintReference
	ErrorTypeUnknownCustomDataType
	ErrorTypeUnknownSuperset
	ErrorTypeUnknownEntity
	ErrorTypeReferenceIDMismatch
	ErrorTypeReferenceNameMismatch
	ErrorTypeReferenceTypeMismatch
	ErrorTypeNamespaceUnknownTag
	ErrorTypeUnknownBaseCluster
	ErrorTypeUnreferencedEntity

	ErrorTypeUnknownConformanceIdentifier
	ErrorTypeUnknownConformanceReference

	ErrorTypeFabricScopingNotAllowed
	ErrorTypeFabricSensitivityNotAllowed
	ErrorTypeFabricScopedStructNotAllowed

	ErrorTypeInvalidConformance

	ErrorTypeInvalidConstraint
	ErrorTypeInvalidConstraintLimit

	ErrorTypeInvalidFallback
	ErrorTypeInvalidFallbackLimit
	ErrorTypeInvalidQuality

	ErrorTypeConformanceChoiceOrphan
	ErrorTypeConformanceChoiceMismatch
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

type GenericParseError struct {
	error
	Source log.Source
}

func (gpe GenericParseError) Type() ErrorType {
	return ErrorTypeGenericParse
}

func (gpe GenericParseError) Origin() (path string, line int) {
	return gpe.Source.Origin()
}

func newGenericParseError(source log.Source, format string, a ...any) *GenericParseError {
	return &GenericParseError{error: fmt.Errorf(format, a...), Source: source}
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

type UnknownSupersetError struct {
	DeviceType *matter.DeviceType
}

func (ddt UnknownSupersetError) Type() ErrorType {
	return ErrorTypeUnknownSuperset
}

func (ddt UnknownSupersetError) Origin() (path string, line int) {
	return ddt.DeviceType.Origin()
}

func (ddt UnknownSupersetError) Error() string {
	return fmt.Sprintf("unknown superset: %s", ddt.DeviceType.SupersetOf)
}

type UnknownEntityError[SourceT types.Entity] struct {
	EntityType       types.EntityType
	EntityReference  *matter.EntityReference
	ElementReference *matter.ElementReference
	Source           SourceT
}

func (uee UnknownEntityError[SourceT]) Type() ErrorType {
	return ErrorTypeUnknownEntity
}

func (uee UnknownEntityError[SourceT]) Origin() (path string, line int) {
	return uee.Source.Origin()
}

func (uee UnknownEntityError[SourceT]) Error() string {
	if uee.EntityReference != nil {
		return fmt.Sprintf("unknown %s: \"%s\" (%s) for %s", uee.EntityType.String(), uee.EntityReference.Name, uee.EntityReference.ID.HexString(), matter.EntityName(uee.Source))
	}
	if uee.ElementReference != nil {
		return fmt.Sprintf("unknown %s: \"%s\" (%s) for %s", uee.EntityType.String(), uee.ElementReference.Name, uee.ElementReference.Name, matter.EntityName(uee.Source))
	}
	return fmt.Sprintf("unknown %s on %s", uee.EntityType.String(), matter.EntityName(uee.Source))
}

type UnreferencedEntityError[EntityT, SourceT types.Entity] struct {
	Entity EntityT
	Source SourceT
}

func (uee UnreferencedEntityError[EntityT, SourceT]) Type() ErrorType {
	return ErrorTypeUnknownEntity
}

func (uee UnreferencedEntityError[EntityT, SourceT]) Origin() (path string, line int) {
	return uee.Source.Origin()
}

func (uee UnreferencedEntityError[EntityT, SourceT]) Error() string {
	return fmt.Sprintf("%s unreferenced on %s: %s (%s)", uee.Entity.EntityType().String(), matter.EntityName(uee.Source), matter.EntityName(uee.Entity), matter.EntityID(uee.Entity).HexString())
}

type NamespaceUnknownTagError struct {
	Requirement *matter.TagRequirement
	Namespace   *matter.Namespace
}

func (ddt NamespaceUnknownTagError) Type() ErrorType {
	return ErrorTypeNamespaceUnknownTag
}

func (ddt NamespaceUnknownTagError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt NamespaceUnknownTagError) Error() string {
	return fmt.Sprintf("unknown tag in namespace: %s for namespace %s", ddt.Requirement.SemanticTagRef.Name, ddt.Namespace.Name)
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

type ReferenceTypeMismatch struct {
	Entity  types.Entity
	Element types.EntityType
	Source  log.Source
}

func (rtm ReferenceTypeMismatch) Type() ErrorType {
	return ErrorTypeReferenceTypeMismatch
}

func (rtm ReferenceTypeMismatch) Origin() (path string, line int) {
	return rtm.Source.Origin()
}

func (rtm ReferenceTypeMismatch) Error() string {
	return fmt.Sprintf("reference points to %s but %s was expected", rtm.Entity.EntityType().String(), rtm.Element.String())
}

type ReferenceIDMismatchError[EntityT types.Entity, SourceT types.Entity] struct {
	Source SourceT
	Entity EntityT
	ID     *matter.Number
}

func (rim ReferenceIDMismatchError[EntityT, SourceT]) Type() ErrorType {
	return ErrorTypeReferenceIDMismatch
}

func (rim ReferenceIDMismatchError[EntityT, SourceT]) Origin() (path string, line int) {
	return rim.Source.Origin()
}

func (rim ReferenceIDMismatchError[EntityT, SourceT]) Error() string {
	return fmt.Sprintf("reference id mismatch: %s %s vs. %s (%s) on %s", rim.Entity.EntityType().String(), rim.ID.HexString(), matter.EntityID(rim.Entity).HexString(), matter.EntityName(rim.Entity), rim.Source.EntityType().String())
}

type ReferenceNameMismatchError[EntityT types.Entity, SourceT types.Entity] struct {
	Source SourceT
	Entity EntityT
	Name   string
}

func (rim ReferenceNameMismatchError[EntityT, SourceT]) Type() ErrorType {
	return ErrorTypeReferenceNameMismatch
}

func (rim ReferenceNameMismatchError[EntityT, SourceT]) Origin() (path string, line int) {
	return rim.Source.Origin()
}

func (rim ReferenceNameMismatchError[EntityT, SourceT]) Error() string {
	return fmt.Sprintf("reference name mismatch: %s vs. %s on %s", rim.Name, matter.EntityName(rim.Entity), rim.Source.EntityType().String())
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

func (cf UnknownConformanceIdentifierError) ComparableEntity() types.Entity {
	return cf.Entity
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

func (cf UnknownConformanceReferenceError) ComparableEntity() types.Entity {
	return cf.Entity
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

func (cf FabricScopingNotAllowedError) ComparableEntity() types.Entity {
	return cf.Entity
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

func (cf FabricSensitivityNotAllowedError) ComparableEntity() types.Entity {
	return cf.Entity
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
	return fmt.Sprintf("fabric scoped struct not allowed: \"%s\"", matter.EntityName(ddt.Entity))
}

func (cf FabricScopedStructNotAllowedError) ComparableEntity() types.Entity {
	return cf.Entity
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

func (ice InvalidConstraintError) Type() ErrorType {
	return ErrorTypeInvalidConstraint
}

func (ice InvalidConstraintError) Origin() (path string, line int) {
	return ice.Source.Origin()
}

func (ice InvalidConstraintError) Error() string {
	return fmt.Sprintf("invalid constraint: \"%s\"", ice.Constraint)
}

type InvalidConstraintLimitError struct {
	Limit  constraint.Limit
	Field  *matter.Field
	Source log.Source
}

func (ice InvalidConstraintLimitError) Type() ErrorType {
	return ErrorTypeInvalidConstraintLimit
}

func (ice InvalidConstraintLimitError) Origin() (path string, line int) {
	return ice.Source.Origin()
}

func (ice InvalidConstraintLimitError) Error() string {
	return fmt.Sprintf("invalid constraint limit: \"%s\" for field with data type %s", ice.Limit.ASCIIDocString(ice.Field.Type), ice.Field.Type.BaseType.String())
}

type InvalidFallbackError struct {
	Fallback string
	Source   log.Source
}

func (ifbe InvalidFallbackError) Type() ErrorType {
	return ErrorTypeInvalidFallback
}

func (ifbe InvalidFallbackError) Origin() (path string, line int) {
	return ifbe.Source.Origin()
}

func (ifbe InvalidFallbackError) Error() string {
	return fmt.Sprintf("invalid fallback: \"%s\"", ifbe.Fallback)
}

type InvalidFallbackLimitError struct {
	Fallback constraint.Limit
	Field    *matter.Field
	Source   log.Source
}

func (ifbe InvalidFallbackLimitError) Type() ErrorType {
	return ErrorTypeInvalidFallbackLimit
}

func (ifbe InvalidFallbackLimitError) Origin() (path string, line int) {
	return ifbe.Source.Origin()
}

func (ifbe InvalidFallbackLimitError) Error() string {
	return fmt.Sprintf("invalid fallback: \"%s\" for field with data type %s", ifbe.Fallback.ASCIIDocString(ifbe.Field.Type), ifbe.Field.Type.BaseType.String())
}

type ConformanceChoiceOrphanError struct {
	Set    string
	Source log.Source
}

func (ccoe ConformanceChoiceOrphanError) Type() ErrorType {
	return ErrorTypeConformanceChoiceOrphan
}

func (ccoe ConformanceChoiceOrphanError) Origin() (path string, line int) {
	return ccoe.Source.Origin()
}

func (ccoe ConformanceChoiceOrphanError) Error() string {
	return fmt.Sprintf("conformance choice set \"%s\" is unused by any other conformances", ccoe.Set)
}

type InvalidQualityError struct {
	Quality    matter.Quality
	EntityType types.EntityType
	Source     types.Entity
}

func (iqe InvalidQualityError) Type() ErrorType {
	return ErrorTypeInvalidQuality
}

func (iqe InvalidQualityError) Origin() (path string, line int) {
	return iqe.Source.Origin()
}

func (iqe InvalidQualityError) Error() string {
	return fmt.Sprintf("invalid quality for %s: %s", iqe.EntityType.String(), iqe.Quality.String())
}

type ConformanceChoiceMismatchError struct {
	Set    string
	Source log.Source

	Entity              types.Entity
	ChoiceLimit         conformance.ChoiceLimit
	Previous            types.Entity
	PreviousChoiceLimit conformance.ChoiceLimit
}

func (ccme ConformanceChoiceMismatchError) Type() ErrorType {
	return ErrorTypeConformanceChoiceMismatch
}

func (ccme ConformanceChoiceMismatchError) Origin() (path string, line int) {
	return ccme.Source.Origin()
}

func (ccme ConformanceChoiceMismatchError) Error() string {
	var limit string
	if ccme.ChoiceLimit != nil {
		limit = ccme.ChoiceLimit.ASCIIDocString()
	}
	var previous string
	if ccme.Previous != nil {
		previous = ccme.PreviousChoiceLimit.ASCIIDocString()
	}

	return fmt.Sprintf("mismatch in choice limit \"%s\": %s vs. %s", ccme.Set, limit, previous)
}

type MissingClonedEntityError struct {
	Entity types.Entity
	Field  *matter.Field
}

func (mcee *MissingClonedEntityError) Type() ErrorType {
	return ErrorTypeUnknown
}

func (mcee *MissingClonedEntityError) Origin() (path string, line int) {
	return mcee.Field.Origin()
}

func (mcee *MissingClonedEntityError) Error() string {
	return fmt.Sprintf("failed to find local clone of entity %s for field %s", matter.EntityName(mcee.Entity), mcee.Field.Name)
}

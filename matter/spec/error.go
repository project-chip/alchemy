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
	ErrorTypeUnknownClusterRequirement
	ErrorTypeUnknownElementRequirementCluster
	ErrorTypeElementRequirementUnreferencedCluster
	ErrorTypeElementRequirementUnknownElement
	ErrorTypeComposingDeviceTypeRequirementUnknownDeviceType
	ErrorTypeComposingDeviceTypeClusterRequirementUnknownCluster
	ErrorTypeComposingDeviceTypeClusterRequirementUnknownDeviceType
	ErrorTypeComposingDeviceTypeClusterRequirementUnreferencedDeviceType
	ErrorTypeComposingDeviceTypeElementRequirementUnknownCluster
	ErrorTypeComposingDeviceTypeElementRequirementUnknownDeviceType
	ErrorTypeComposingDeviceTypeElementRequirementUnreferencedDeviceType
	ErrorTypeConditionRequirementUnknownDeviceType
	ErrorTypeConditionRequirementUnreferencedDeviceType
	ErrorTypeConditionRequirementUnknownCondition
	ErrorTypeTagRequirementUnreferencedDeviceType
	ErrorTypeTagRequirementUnknownNamespace
	ErrorTypeTagRequirementNamespaceNameMismatch
	ErrorTypeTagRequirementUnknownTag
	ErrorTypeTagRequirementTagNameMismatch
	ErrorTypeClusterReferenceNameMismatch
	ErrorTypeDeviceTypeReferenceNameMismatch
	ErrorTypeNamespaceNameMismatch
	ErrorTypeUnknownBaseCluster
	ErrorTypeUnknownConformanceIdentifier
	ErrorTypeUnknownConformanceReference
	ErrorTypeFabricScopingNotAllowed
	ErrorTypeFabricSensitivityNotAllowed
	ErrorTypeFabricScopedStructNotAllowed
	ErrorTypeInvalidConformance
	ErrorTypeInvalidConstraint
	ErrorTypeInvalidFallback
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

type ElementRequirementUnreferencedClusterError struct {
	Requirement *matter.ElementRequirement
}

func (ddt ElementRequirementUnreferencedClusterError) Type() ErrorType {
	return ErrorTypeElementRequirementUnreferencedCluster
}

func (ddt ElementRequirementUnreferencedClusterError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt ElementRequirementUnreferencedClusterError) Error() string {
	return fmt.Sprintf("unreferenced element requirement cluster: %s", ddt.Requirement.ClusterName)
}

type ElementRequirementUnknownElementError struct {
	Requirement *matter.ElementRequirement
}

func (ddt ElementRequirementUnknownElementError) Type() ErrorType {
	return ErrorTypeElementRequirementUnknownElement
}

func (ddt ElementRequirementUnknownElementError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt ElementRequirementUnknownElementError) Error() string {
	return fmt.Sprintf("element requirement references unknown element: %s %s", ddt.Requirement.Element.String(), ddt.Requirement.Name)
}

type UnknownConditionRequirementDeviceTypeError struct {
	Requirement *matter.ConditionRequirement
}

func (ddt UnknownConditionRequirementDeviceTypeError) Type() ErrorType {
	return ErrorTypeConditionRequirementUnknownDeviceType
}

func (ddt UnknownConditionRequirementDeviceTypeError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt UnknownConditionRequirementDeviceTypeError) Error() string {
	return fmt.Sprintf("condition requirement references unknown device type: %s", ddt.Requirement.DeviceTypeName)
}

type UnreferencedConditionRequirementDeviceTypeError struct {
	Requirement *matter.ConditionRequirement
}

func (ddt UnreferencedConditionRequirementDeviceTypeError) Type() ErrorType {
	return ErrorTypeConditionRequirementUnreferencedDeviceType
}

func (ddt UnreferencedConditionRequirementDeviceTypeError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt UnreferencedConditionRequirementDeviceTypeError) Error() string {
	return fmt.Sprintf("unreferenced condition requirement device type: %s", ddt.Requirement.DeviceTypeName)
}

type UnknownConditionRequirementConditionError struct {
	Requirement *matter.ConditionRequirement
}

func (ddt UnknownConditionRequirementConditionError) Type() ErrorType {
	return ErrorTypeConditionRequirementUnknownCondition
}

func (ddt UnknownConditionRequirementConditionError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt UnknownConditionRequirementConditionError) Error() string {
	return fmt.Sprintf("condition requirement references unknown condition: %s", ddt.Requirement.ConditionName)
}

type UnknownComposingDeviceTypeRequirementDeviceTypeError struct {
	Requirement *matter.DeviceTypeRequirement
}

func (ddt UnknownComposingDeviceTypeRequirementDeviceTypeError) Type() ErrorType {
	return ErrorTypeComposingDeviceTypeRequirementUnknownDeviceType
}

func (ddt UnknownComposingDeviceTypeRequirementDeviceTypeError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt UnknownComposingDeviceTypeRequirementDeviceTypeError) Error() string {
	return fmt.Sprintf("unknown composing device device type requirement: %s", ddt.Requirement.DeviceTypeName)
}

type UnknownComposingDeviceTypeRequirementClusterError struct {
	Requirement *matter.DeviceTypeClusterRequirement
}

func (ddt UnknownComposingDeviceTypeRequirementClusterError) Type() ErrorType {
	return ErrorTypeComposingDeviceTypeClusterRequirementUnknownCluster
}

func (ddt UnknownComposingDeviceTypeRequirementClusterError) Origin() (path string, line int) {
	return ddt.Requirement.ClusterRequirement.Origin()
}

func (ddt UnknownComposingDeviceTypeRequirementClusterError) Error() string {
	return fmt.Sprintf("unknown composing device cluster requirement: %s", ddt.Requirement.ClusterRequirement.ClusterName)
}

type UnknownComposingDeviceTypeClusterRequirementDeviceTypeError struct {
	Requirement *matter.DeviceTypeClusterRequirement
}

func (ddt UnknownComposingDeviceTypeClusterRequirementDeviceTypeError) Type() ErrorType {
	return ErrorTypeComposingDeviceTypeClusterRequirementUnknownDeviceType
}

func (ddt UnknownComposingDeviceTypeClusterRequirementDeviceTypeError) Origin() (path string, line int) {
	return ddt.Requirement.ClusterRequirement.Origin()
}

func (ddt UnknownComposingDeviceTypeClusterRequirementDeviceTypeError) Error() string {
	return fmt.Sprintf("unknown composing device cluster requirement device type: %s", ddt.Requirement.DeviceTypeName)
}

type UnreferencedComposingDeviceTypeClusterRequirementDeviceTypeError struct {
	Requirement *matter.DeviceTypeClusterRequirement
}

func (ddt UnreferencedComposingDeviceTypeClusterRequirementDeviceTypeError) Type() ErrorType {
	return ErrorTypeComposingDeviceTypeClusterRequirementUnreferencedDeviceType
}

func (ddt UnreferencedComposingDeviceTypeClusterRequirementDeviceTypeError) Origin() (path string, line int) {
	return ddt.Requirement.ClusterRequirement.Origin()
}

func (ddt UnreferencedComposingDeviceTypeClusterRequirementDeviceTypeError) Error() string {
	return fmt.Sprintf("unreferenced composing device cluster requirement device type: %s", ddt.Requirement.DeviceTypeName)
}

type UnknownComposingElementRequirementClusterError struct {
	Requirement *matter.DeviceTypeElementRequirement
}

func (ddt UnknownComposingElementRequirementClusterError) Type() ErrorType {
	return ErrorTypeComposingDeviceTypeElementRequirementUnknownCluster
}

func (ddt UnknownComposingElementRequirementClusterError) Origin() (path string, line int) {
	return ddt.Requirement.ElementRequirement.Origin()
}

func (ddt UnknownComposingElementRequirementClusterError) Error() string {
	return fmt.Sprintf("unknown composing device element requirement cluster: %s", ddt.Requirement.ElementRequirement.ClusterName)
}

type UnknownComposingDeviceTypeElementRequirementDeviceTypeError struct {
	Requirement *matter.DeviceTypeElementRequirement
}

func (ddt UnknownComposingDeviceTypeElementRequirementDeviceTypeError) Type() ErrorType {
	return ErrorTypeComposingDeviceTypeElementRequirementUnknownDeviceType
}

func (ddt UnknownComposingDeviceTypeElementRequirementDeviceTypeError) Origin() (path string, line int) {
	return ddt.Requirement.ElementRequirement.Origin()
}

func (ddt UnknownComposingDeviceTypeElementRequirementDeviceTypeError) Error() string {
	return fmt.Sprintf("unknown composing device element requirement device type: %s", ddt.Requirement.DeviceTypeName)
}

type UnreferencedComposingDeviceTypeElementRequirementDeviceTypeError struct {
	Requirement *matter.DeviceTypeElementRequirement
}

func (ddt UnreferencedComposingDeviceTypeElementRequirementDeviceTypeError) Type() ErrorType {
	return ErrorTypeComposingDeviceTypeElementRequirementUnreferencedDeviceType
}

func (ddt UnreferencedComposingDeviceTypeElementRequirementDeviceTypeError) Origin() (path string, line int) {
	return ddt.Requirement.ElementRequirement.Origin()
}

func (ddt UnreferencedComposingDeviceTypeElementRequirementDeviceTypeError) Error() string {
	return fmt.Sprintf("unreferenced composing device element requirement device type: %s", ddt.Requirement.DeviceTypeName)
}

type UnknownComposingDeviceTypeTagRequirementDeviceTypeError struct {
	Requirement *matter.DeviceTypeTagRequirement
}

func (ddt UnknownComposingDeviceTypeTagRequirementDeviceTypeError) Type() ErrorType {
	return ErrorTypeTagRequirementUnreferencedDeviceType
}

func (ddt UnknownComposingDeviceTypeTagRequirementDeviceTypeError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt UnknownComposingDeviceTypeTagRequirementDeviceTypeError) Error() string {
	return fmt.Sprintf("unknown device tag requirement device type: %s", ddt.Requirement.DeviceTypeName)
}

type UnreferencedTagRequirementDeviceTypeError struct {
	Requirement *matter.DeviceTypeTagRequirement
}

func (ddt UnreferencedTagRequirementDeviceTypeError) Type() ErrorType {
	return ErrorTypeTagRequirementUnreferencedDeviceType
}

func (ddt UnreferencedTagRequirementDeviceTypeError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt UnreferencedTagRequirementDeviceTypeError) Error() string {
	return fmt.Sprintf("unreferenced composing device cluster requirement device type: %s", ddt.Requirement.DeviceTypeName)
}

type UnknownNamespaceTagRequirementError struct {
	Requirement *matter.TagRequirement
}

func (ddt UnknownNamespaceTagRequirementError) Type() ErrorType {
	return ErrorTypeTagRequirementUnknownNamespace
}

func (ddt UnknownNamespaceTagRequirementError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt UnknownNamespaceTagRequirementError) Error() string {
	return fmt.Sprintf("unrecognized namespace in tag requirement: %s", ddt.Requirement.NamespaceName)
}

type NamespaceNameMismatchTagRequirementError struct {
	Requirement *matter.TagRequirement
	Namespace   *matter.Namespace
}

func (ddt NamespaceNameMismatchTagRequirementError) Type() ErrorType {
	return ErrorTypeTagRequirementNamespaceNameMismatch
}

func (ddt NamespaceNameMismatchTagRequirementError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt NamespaceNameMismatchTagRequirementError) Error() string {
	return fmt.Sprintf("mismatched namespace name in tag requirement: %s vs. %s", ddt.Requirement.NamespaceName, ddt.Namespace.Name)
}

type UnknownTagRequirementError struct {
	Requirement *matter.TagRequirement
}

func (ddt UnknownTagRequirementError) Type() ErrorType {
	return ErrorTypeTagRequirementUnknownTag
}

func (ddt UnknownTagRequirementError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt UnknownTagRequirementError) Error() string {
	return fmt.Sprintf("unrecognized tag in tag requirement: %s", ddt.Requirement.SemanticTagName)
}

type TagNameMismatchTagRequirementError struct {
	Requirement *matter.TagRequirement
	SemanticTag *matter.SemanticTag
}

func (ddt TagNameMismatchTagRequirementError) Type() ErrorType {
	return ErrorTypeTagRequirementTagNameMismatch
}

func (ddt TagNameMismatchTagRequirementError) Origin() (path string, line int) {
	return ddt.Requirement.Origin()
}

func (ddt TagNameMismatchTagRequirementError) Error() string {
	return fmt.Sprintf("mismatched tag name in tag requirement: %s vs. %s", ddt.Requirement.SemanticTagName, ddt.SemanticTag.Name)
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

type ClusterReferenceNameMismatch struct {
	Cluster *matter.Cluster
	Name    string
	Source  log.Source
}

func (ddt ClusterReferenceNameMismatch) Type() ErrorType {
	return ErrorTypeClusterReferenceNameMismatch
}

func (ddt ClusterReferenceNameMismatch) Origin() (path string, line int) {
	return ddt.Source.Origin()
}

func (ddt ClusterReferenceNameMismatch) Error() string {
	return fmt.Sprintf("cluster reference has mismatched name: %s vs. %s", ddt.Cluster.Name, ddt.Name)
}

type DeviceTypeReferenceNameMismatch struct {
	DeviceType *matter.DeviceType
	Name       string
	Source     log.Source
}

func (ddt DeviceTypeReferenceNameMismatch) Type() ErrorType {
	return ErrorTypeDeviceTypeReferenceNameMismatch
}

func (ddt DeviceTypeReferenceNameMismatch) Origin() (path string, line int) {
	return ddt.Source.Origin()
}

func (ddt DeviceTypeReferenceNameMismatch) Error() string {
	return fmt.Sprintf("device type reference has mismatched name: %s vs. %s", ddt.DeviceType.Name, ddt.Name)
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

type InvalidFallbackError struct {
	Fallback string
	Source   log.Source
}

func (ifbe InvalidFallbackError) Type() ErrorType {
	return ErrorTypeInvalidConstraint
}

func (ifbe InvalidFallbackError) Origin() (path string, line int) {
	return ifbe.Source.Origin()
}

func (ifbe InvalidFallbackError) Error() string {
	return fmt.Sprintf("invalid fallback: \"%s\"", ifbe.Fallback)
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

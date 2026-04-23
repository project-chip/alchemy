package render

import (
	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/provisional"
	"github.com/project-chip/alchemy/zap"
)

func (cr *configuratorRenderer) setFieldAttributes(fieldElement *etree.Element, parentEntityType types.EntityType, field *matter.Field, fieldSet matter.FieldSet, newlyAdded bool) {
	switch parentEntityType {
	case types.EntityTypeEvent:
		// Remove incorrect attributes from legacy XML, but leave "id" on event for now, as ZAP expects it; once ZAP is reading "fieldId", remove this
		fieldElement.RemoveAttr("code")
	default:
		// Remove incorrect attributes from legacy XML
		fieldElement.RemoveAttr("code")
		fieldElement.RemoveAttr("id")
	}
	xml.PrependAttribute(fieldElement, "fieldId", field.ID.IntString(), "id") // once ZAP is reading "fieldId", remove "id"
	fieldName := zap.CleanName(field.Name)
	fieldElement.CreateAttr("name", fieldName)
	cr.writeDataType(fieldElement, fieldSet, field)
	if field.Quality.Has(matter.QualityNullable) && !cr.generator.options.ExtendedQuality {
		fieldElement.CreateAttr("isNullable", "true")
	} else {
		fieldElement.RemoveAttr("isNullable")
	}
	if !conformance.IsMandatory(field.Conformance) {
		fieldElement.CreateAttr("optional", "true")
	} else {
		fieldElement.RemoveAttr("optional")
	}
	if field.Access.IsFabricSensitive() {
		fieldElement.CreateAttr("isFabricSensitive", "true")
	} else {
		fieldElement.RemoveAttr("isFabricSensitive")
	}
	cr.setFieldFallback(fieldElement, field, fieldSet)
	cr.setQuality(fieldElement, field.EntityType(), field.Quality)
	cr.renderConstraint(fieldElement, fieldSet, field)
	if parentEntityType == types.EntityTypeStruct && provisional.Policy(cr.generator.options.ProvisionalPolicy) == provisional.PolicyNone {
		if newlyAdded {
			fieldElement.CreateAttr("apiMaturity", "provisional")
		}
	} else {
		cr.setProvisional(fieldElement, field)
	}
}

func (cr *configuratorRenderer) writeDataType(element *etree.Element, fieldSet matter.FieldSet, field *matter.Field) {
	if field.Type == nil {
		return
	}
	dts := cr.getDataTypeString(fieldSet, field)
	if field.Type.IsArray() {
		element.CreateAttr("array", "true")
		element.CreateAttr("type", dts)
	} else {
		element.CreateAttr("type", dts)
		element.RemoveAttr("array")
	}
}

func (cr *configuratorRenderer) getDataTypeString(fs matter.FieldSet, f *matter.Field) string {
	switch f.Type.BaseType {
	case types.BaseDataTypeTag:
		if f.Type.Entity != nil {
			if namespace, ok := f.Type.Entity.(*matter.Namespace); ok {
				return matterNamespaceName(namespace)
			}
		} else {
			return "enum8"
		}
	case types.BaseDataTypeNamespaceID:
		return "enum8"
	}
	return zap.FieldToZapDataType(fs, f, f.Constraint)
}

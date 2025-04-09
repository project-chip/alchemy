package generate

import (
	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func (cr *configuratorRenderer) setFieldAttributes(fieldElement *etree.Element, parentEntityType types.EntityType, parentTypeName string, field *matter.Field, fieldSet matter.FieldSet) {
	mandatory := conformance.IsMandatory(field.Conformance)
	fieldName := cr.configurator.Errata.FieldName(parentEntityType, parentTypeName, field.Name)
	fieldElement.CreateAttr("name", fieldName)
	cr.writeDataType(fieldElement, parentEntityType, parentTypeName, fieldSet, field)
	if !mandatory {
		fieldElement.CreateAttr("optional", "true")
	} else {
		fieldElement.RemoveAttr("optional")
	}
	if field.Quality.Has(matter.QualityNullable) && !cr.generator.generateExtendedQualityElement {
		fieldElement.CreateAttr("isNullable", "true")
	} else {
		fieldElement.RemoveAttr("isNullable")
	}
	if field.Access.IsFabricSensitive() {
		fieldElement.CreateAttr("isFabricSensitive", "true")
	} else {
		fieldElement.RemoveAttr("isFabricSensitive")
	}
	setFieldFallback(fieldElement, field, fieldSet)
	cr.setQuality(fieldElement, field.EntityType(), field.Quality)
	renderConstraint(fieldElement, fieldSet, field)
}

func (cr *configuratorRenderer) writeDataType(element *etree.Element, parentEntityType types.EntityType, parentTypeName string, fieldSet matter.FieldSet, field *matter.Field) {
	if field.Type == nil {
		return
	}
	dts := getDataTypeString(fieldSet, field)
	dts = cr.configurator.Errata.TypeName(parentEntityType, dts)
	dts = cr.configurator.Errata.FieldTypeName(parentEntityType, parentTypeName, field.Name, dts)
	if field.Type.IsArray() {
		element.CreateAttr("array", "true")
		element.CreateAttr("type", dts)
	} else {
		element.CreateAttr("type", dts)
		element.RemoveAttr("array")
	}
}

func getDataTypeString(fs matter.FieldSet, f *matter.Field) string {
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
	return zap.FieldToZapDataType(fs, f)
}

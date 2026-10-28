package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func validateDataType(spec *Specification, cluster *matter.Cluster, field *matter.Field, dataType *types.DataType) {
	switch dataType.BaseType {
	case types.BaseDataTypeTag:
	case types.BaseDataTypeList:
		validateDataType(spec, cluster, field, dataType.EntryType)
	case types.BaseDataTypeCustom:
		if dataType.Entity == nil {
			spec.addError(&UnknownCustomDataTypeError{Field: field, DataType: dataType})
		}
	}
	if field.Type.Entity == nil {
		return
	}
	if cluster != nil && cluster.Hierarchy != "Base" && !spec.ignoredHierarchy {
		// If this data type's parent is not the cluster it belongs to, then we should have
		// redirected it to the local copy of this data type in resolveFieldDataTypes
		parentCluster, hasParentCluster := field.Type.Entity.Parent().(*matter.Cluster)
		if hasParentCluster && parentCluster != cluster {
			slog.Error("failed to find local cloned entity for data type reference", slog.String("name", matter.EntityName(field.Type.Entity)), slog.String("field", field.Name), slog.String("cluster", cluster.Name))
			spec.addError(&MissingClonedEntityError{Entity: field.Type.Entity, Field: field})
		}
	}
}

func validateFieldDataType(spec *Specification, cluster *matter.Cluster, field *matter.Field) {
	if field.Type == nil {
		if !conformance.IsDeprecated(field.Conformance) && !conformance.IsDisallowed(field.Conformance) && !spec.ignoredHierarchy && (cluster == nil || cluster.Hierarchy == "Base") {
			var clusterName string
			if cluster != nil {
				clusterName = cluster.Name
			}
			slog.Warn("missing type on field", log.Path("source", field), slog.String("id", field.ID.HexString()), slog.String("name", field.Name), slog.String("cluster", clusterName))
		}
		return
	}
	validateDataType(spec, cluster, field, field.Type)
}

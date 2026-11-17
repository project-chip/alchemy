package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func validateDataType(spec *Specification, cluster *matter.Cluster, field *matter.Field, dataType *types.DataType) {
	if dataType == nil {
		return
	}
	switch dataType.BaseType {
	case types.BaseDataTypeTag:
	case types.BaseDataTypeList:
		validateDataType(spec, cluster, field, dataType.EntryType)
	case types.BaseDataTypeCustom:
		if dataType.Entity == nil {
			spec.addError(&UnknownCustomDataTypeError{Field: field, DataType: dataType})
		}
	}
	if dataType.Entity == nil {
		return
	}
	if cluster == nil || cluster.Hierarchy == "Base" || spec.ignoredHierarchy {
		return
	}
	// If this data type's parent is not the cluster it belongs to, then we should have
	// redirected it to the local copy of this data type in resolveFieldDataTypes
	if _, isXRef := dataType.Source.(*asciidoc.CrossReference); isXRef {
		// If the data type was resolved through a cross-reference, then it's shared
		// with the original cluster, so we don't need to validate it here.
		return
	}
	parentCluster, hasParentCluster := dataType.Entity.Parent().(*matter.Cluster)
	if hasParentCluster && parentCluster != cluster {
		slog.Error("failed to find local cloned entity for data type reference", slog.String("name", matter.EntityName(dataType.Entity)), slog.String("field", field.Name), slog.String("cluster", cluster.Name))
		spec.addError(&MissingClonedEntityError{Entity: dataType.Entity, Field: field})
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

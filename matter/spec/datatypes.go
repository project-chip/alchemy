package spec

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toDataTypes(d *Doc, pc *parseContext, parentEntity types.Entity) (dataTypes []types.Entity, err error) {

	traverseSections(d, s, errata.SpecPurposeDataTypes, func(s *Section, parent parse.HasElements, index int) parse.SearchShould {
		switch s.SecType {
		case matter.SectionDataTypeBitmap:
			var mb *matter.Bitmap
			mb, err = s.toBitmap(d, pc, parentEntity)
			if err != nil {
				if err != ErrNotEnoughRowsInTable || d.parsed {
					slog.Warn("Error converting section to bitmap", log.Element("source", d.Path, s.Base), slog.Any("error", err))
				}
				err = nil
			} else {
				dataTypes = append(dataTypes, mb)
			}
		case matter.SectionDataTypeEnum:
			var me *matter.Enum
			me, err = s.toEnum(d, pc, parentEntity)
			if err != nil {
				if err != ErrNotEnoughRowsInTable || d.parsed {
					slog.Warn("Error converting section to enum", log.Element("source", d.Path, s.Base), slog.Any("error", err))
				}
				err = nil
			} else {
				dataTypes = append(dataTypes, me)
			}
		case matter.SectionDataTypeStruct:
			var me *matter.Struct
			me, err = s.toStruct(d, pc, parentEntity)
			if err != nil {
				if err != ErrNotEnoughRowsInTable || d.parsed {
					slog.Warn("Error converting section to struct", log.Element("source", d.Path, s.Base), slog.Any("error", err))
				}
				err = nil
			} else {
				dataTypes = append(dataTypes, me)
			}
		case matter.SectionDataTypeDef:
			var me *matter.TypeDef
			me, err = s.toTypeDef(d, pc, parentEntity)
			if err != nil {
				if err != ErrNotEnoughRowsInTable || d.parsed {
					slog.Warn("Error converting section to typedef", log.Element("source", d.Path, s.Base), slog.Any("error", err))
				}
				err = nil
			} else {
				dataTypes = append(dataTypes, me)
			}
		case matter.SectionDataTypeConstant:
			id, _ := getAnchorElements(s.Base, nil)
			switch id {
			case "ref_RespMaxConstant":
				c := matter.NewConstant(s)
				c.Name = "RESP_MAX"
				c.Value = 900
				pc.orderedEntities = append(pc.orderedEntities, c)
				pc.entitiesByElement[s.Base] = append(pc.entitiesByElement[s.Base], c)
				dataTypes = append(dataTypes, c)
			}
		default:
		}
		return parse.SearchShouldContinue
	})

	return
}

func (sp *Builder) resolveClusterDataTypeReferences(onlyBaseClusters bool) {
	specEntityFinder := newSpecEntityFinder(sp.Spec, nil, nil)
	for cluster := range sp.Spec.Clusters {
		inheritedCluster := cluster.Hierarchy != "Base"
		if (onlyBaseClusters && inheritedCluster) || (!onlyBaseClusters && !inheritedCluster) {
			continue
		}

		specEntityFinder.cluster = cluster
		clusterFinder := newClusterEntityFinder(cluster, specEntityFinder)

		for _, a := range cluster.Attributes {
			clusterFinder.setIdentity(a)
			sp.resolveFieldDataTypes(cluster, cluster.Attributes, a, a.Type, clusterFinder)
		}
		for _, s := range cluster.Structs {
			for _, f := range s.Fields {
				clusterFinder.setIdentity(f)
				sp.resolveFieldDataTypes(cluster, s.Fields, f, f.Type, clusterFinder)
			}
		}
		for _, event := range cluster.Events {
			for _, f := range event.Fields {
				clusterFinder.setIdentity(f)
				sp.resolveFieldDataTypes(cluster, event.Fields, f, f.Type, clusterFinder)
			}
		}
		for _, command := range cluster.Commands {
			for _, f := range command.Fields {
				clusterFinder.setIdentity(f)
				sp.resolveFieldDataTypes(cluster, command.Fields, f, f.Type, clusterFinder)
			}
			clusterFinder.setIdentity(command)
			sp.resolveCommandResponseDataType(cluster, command, clusterFinder)
		}
	}
}

func (sp *Builder) resolveGlobalDataTypeReferences() {
	specEntityFinder := newSpecEntityFinder(sp.Spec, nil, nil)
	for _, s := range sp.Spec.structIndex {
		for _, f := range s.Fields {
			specEntityFinder.setIdentity(f)
			sp.resolveFieldDataTypes(nil, s.Fields, f, f.Type, specEntityFinder)
		}
	}
}

func (sp *Builder) resolveFieldDataTypes(cluster *matter.Cluster, fieldSet matter.FieldSet, field *matter.Field, dataType *types.DataType, finder entityFinder) {
	if dataType == nil {
		if !conformance.IsDeprecated(field.Conformance) && !conformance.IsDisallowed(field.Conformance) && !sp.ignoreHierarchy && (cluster == nil || cluster.Hierarchy == "Base") {
			var clusterName string
			if cluster != nil {
				clusterName = cluster.Name
			}
			slog.Warn("missing type on field", log.Path("path", field), slog.String("id", field.ID.HexString()), slog.String("name", field.Name), slog.String("cluster", clusterName))
		}
		return
	}
	if dataType.Entity != nil {
		// This has already been resolved by some other process
		if cluster != nil {
			sp.Spec.DataTypeRefs.Add(field, dataType.Entity)
			sp.Spec.ClusterRefs.Add(cluster, dataType.Entity)
			sp.Spec.addEntity(dataType.Entity, cluster)
		}
		return
	}
	switch dataType.BaseType {
	case types.BaseDataTypeTag:
		sp.getTagNamespace(field)
	case types.BaseDataTypeList:
		sp.resolveFieldDataTypes(cluster, fieldSet, field, dataType.EntryType, finder)
	case types.BaseDataTypeCustom:
		if dataType.Entity == nil {
			finder.setIdentity(field)
			sp.getCustomDataType(dataType, cluster, field, finder)
			if dataType.Entity == nil {
				slog.Error("unknown custom data type", slog.String("cluster", clusterName(cluster)), slog.String("field", field.Name), slog.String("type", dataType.Name), log.Path("source", field), log.Type("element", dataType.Source))
			} else {
				dataType.Name = matter.EntityName(dataType.Entity)
			}
		}
		if cluster == nil || dataType.Entity == nil {
			return
		}
		sp.Spec.DataTypeRefs.Add(field, dataType.Entity)
		sp.Spec.ClusterRefs.Add(cluster, dataType.Entity)
		sp.Spec.addEntity(dataType.Entity, cluster)
		switch e := dataType.Entity.(type) {
		case *matter.Struct:
			// If this data type is a struct, we need to resolve all data types on its fields
			for _, f := range e.Fields {
				sp.resolveFieldDataTypes(cluster, fieldSet, f, f.Type, finder)
			}
		case *matter.TypeDef:
			switch e.Name {
			case "SignedTemperature":
				dataType.BaseType = types.BaseDataTypeSignedTemperature
			case "UnsignedTemperature":
				dataType.BaseType = types.BaseDataTypeUnsignedTemperature
			case "TemperatureDifference":
				dataType.BaseType = types.BaseDataTypeTemperatureDifference
			}
		}

	}
}

func (sp *Builder) resolveCommandResponseDataType(cluster *matter.Cluster, command *matter.Command, finder entityFinder) {
	if command.Response == nil {
		return
	}
	if command.Response.Entity != nil {
		return
	}
	switch source := command.Response.Source.(type) {
	case *asciidoc.CrossReference:
		command.Response.Entity = sp.getCustomDataTypeFromFieldReference(cluster, source, newCommandFinder(cluster.Commands, finder))
	case nil:

	}
	if command.Response.Entity != nil && command.Response.Name == "" {
		if cluster != nil {
			sp.Spec.ClusterRefs.Add(cluster, command.Response.Entity)
		}
		return
	}
	var desiredDirection matter.Interface
	switch command.Direction {
	case matter.InterfaceServer:
		desiredDirection = matter.InterfaceClient
	case matter.InterfaceClient:
		desiredDirection = matter.InterfaceServer
	}
	for _, cmd := range cluster.Commands {
		if cmd.Direction == desiredDirection && cmd.Name == command.Response.Name {
			if cmd.Response == nil {
				break
			}
			command.Response.Entity = cmd.Response.Entity
			if cluster != nil && command.Response.Entity != nil {
				sp.Spec.ClusterRefs.Add(cluster, command.Response.Entity)
			}
			return
		}
	}
}

func (sp *Builder) getCustomDataType(dataType *types.DataType, cluster *matter.Cluster, field *matter.Field, finder entityFinder) {
	//slog.Info("getCustomDataType", log.Type("refType", dataType.Source))
	switch ref := dataType.Source.(type) {
	case *asciidoc.CrossReference:
		dataType.Entity = sp.getCustomDataTypeFromFieldReference(cluster, ref, finder)
		if dataType.Entity != nil {
			// We have a reference to a data type; use that
			return
		}
	}

	dataType.Entity = finder.findEntityByIdentifier(dataType.Name, field)

}

func (sp *Builder) getCustomDataTypeFromFieldReference(cluster *matter.Cluster, reference *asciidoc.CrossReference, finder entityFinder) (e types.Entity) {
	label := strings.TrimSpace(asciidoc.ValueToString(reference.Elements()))
	//slog.Info("getCustomDataTypeFromFieldReference", "id", source.ID, "label", label, log.Path("source", source))
	e = finder.findEntityByReference(reference.ID, label, reference)
	if e != nil {
		return
	}
	if label != "" {
		e = finder.findEntityByIdentifier(label, reference)
	}
	return
}

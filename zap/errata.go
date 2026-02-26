package zap

import (
	"log/slog"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (c *Configurator) addExtraTypes(errata *errata.SDK, entities []types.Entity) {
	if errata.ExtraTypes == nil {
		return
	}
	// Extra types added by Errata do not have any references from any other data types in the spec,
	// and thus need to be added manually to their associated clusters
	for _, m := range entities {
		switch v := m.(type) {
		case *matter.ClusterGroup:
			for _, cl := range v.Clusters {
				c.addExtraTypesForCluster(cl, errata.ExtraTypes)
			}
		case *matter.Cluster:
			c.addExtraTypesForCluster(v, errata.ExtraTypes)
		}
	}
}

func (c *Configurator) addExtraTypesForCluster(cluster *matter.Cluster, extraTypes *errata.SDKTypes) {
	for _, bm := range cluster.Bitmaps {
		if _, ok := extraTypes.Bitmaps[bm.Name]; ok {
			c.addEntityType(cluster, bm)
		}
	}
	for _, en := range cluster.Enums {
		if _, ok := extraTypes.Enums[en.Name]; ok {
			c.addEntityType(cluster, en)
		}
	}
	for _, s := range cluster.Structs {
		if _, ok := extraTypes.Structs[s.Name]; ok {
			c.addEntityType(cluster, s)
		}
	}
	if override, ok := extraTypes.Clusters[cluster.Name]; ok {
		c.addExtraAttributes(cluster, override)
		c.addExtraCommands(cluster, override)
	}
	if len(extraTypes.Attributes) > 0 {
		c.addExtraAttributes(cluster, &errata.SDKType{Attributes: extraTypes.Attributes})
	}
	if len(extraTypes.Commands) > 0 {
		c.addExtraCommands(cluster, &errata.SDKType{Commands: extraTypes.Commands})
	}
}

func (c *Configurator) addExtraAttributes(cluster *matter.Cluster, extra *errata.SDKType) {
	existingAttributes := make(map[string]struct{}, len(cluster.Attributes))
	for _, attr := range cluster.Attributes {
		existingAttributes[attr.Name] = struct{}{}
	}

	for name, a := range extra.Attributes {
		if _, ok := existingAttributes[name]; ok {
			slog.Warn("extra attribute already exists", slog.String("cluster", cluster.Name), slog.String("attribute", name))
			continue
		}
		field := matter.NewAttribute(nil, cluster)
		field.Name = name
		if a.Value != "" {
			field.ID = matter.ParseNumber(a.Value)
		}
		if a.Type != "" {
			bt := ToBaseDataType(a.Type)
			if bt != types.BaseDataTypeUnknown {
				field.Type = types.NewDataType(bt, types.DataTypeRankScalar)
			}
		}
		if a.Access != "" {
			var parsed bool
			field.Access, parsed = spec.ParseAccess(a.Access, types.EntityTypeAttribute)
			if !parsed {
				slog.Warn("failed to parse access string for extra attribute", slog.String("cluster", cluster.Name), slog.String("attribute", name), slog.String("access", a.Access))
			}
		}
		if a.Conformance != "" {
			field.Conformance = conformance.ParseConformance(a.Conformance)
		}
		if a.Quality != "" {
			field.Quality = matter.ParseQuality(a.Quality)
		}
		cluster.Attributes = append(cluster.Attributes, field)
		c.addType(cluster, field.Type)
	}
}

func (c *Configurator) addExtraCommands(cluster *matter.Cluster, extra *errata.SDKType) {
	existingCommands := make(map[string]struct{}, len(cluster.Commands))
	for _, cmd := range cluster.Commands {
		existingCommands[cmd.Name] = struct{}{}
	}

	for name, cmd := range extra.Commands {
		if _, ok := existingCommands[name]; ok {
			slog.Warn("extra command already exists", slog.String("cluster", cluster.Name), slog.String("command", name))
			continue
		}
		command := matter.NewCommand(nil, cluster)
		command.Name = name
		if cmd.Value != "" {
			command.ID = matter.ParseNumber(cmd.Value)
		}
		if cmd.Access != "" {
			var parsed bool
			command.Access, parsed = spec.ParseAccess(cmd.Access, types.EntityTypeCommand)
			if !parsed {
				slog.Warn("failed to parse access string for extra command", slog.String("cluster", cluster.Name), slog.String("command", name), slog.String("access", cmd.Access))
			}
		}
		if cmd.Conformance != "" {
			command.Conformance = conformance.ParseConformance(cmd.Conformance)
		}
		if cmd.Direction != "" {
			switch cmd.Direction {
			case "client":
				command.Direction = matter.InterfaceClient
			case "server":
				command.Direction = matter.InterfaceServer
			default:
				slog.Warn("unknown direction for extra command", slog.String("cluster", cluster.Name), slog.String("command", name), slog.String("direction", cmd.Direction))
			}
		}
		cluster.Commands = append(cluster.Commands, command)
		c.addTypes(cluster, command.Fields)
	}
}

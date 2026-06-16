package sdk

import (
	"fmt"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func addExtraTypes(extraTypes *errata.SDKTypes, entities []types.Entity) error {
	if extraTypes == nil {
		return nil
	}

	var extraEntities []types.Entity
	for name, eb := range extraTypes.Bitmaps {
		bm := matter.NewBitmap(nil, nil)
		bm.Name = name
		typeName := eb.Type
		// This is a workaround for the errata file using the ZAP type name for bitmaps
		switch typeName {
		case "bitmap8":
			typeName = "map8"
		case "bitmap16":
			typeName = "map16"
		case "bitmap32":
			typeName = "map32"
		case "bitmap64":
			typeName = "map64"
		}
		bm.Type = types.ParseDataType(typeName, types.DataTypeRankScalar)
		bm.Description = eb.Description
		for _, ef := range eb.Fields {
			b := matter.NewBitmapBit(nil, bm, ef.Bit, ef.Name, "", nil)
			bm.Bits = append(bm.Bits, b)
		}
		extraEntities = append(extraEntities, bm)
	}
	for name, ee := range extraTypes.Enums {
		e := matter.NewEnum(nil, nil)
		e.Name = name
		e.Type = types.ParseDataType(ee.Type, types.DataTypeRankScalar)
		e.Description = ee.Description
		for _, ef := range ee.Fields {
			ev := matter.NewEnumValue(nil, e)
			ev.Name = ef.Name
			ev.Value = matter.ParseNumber(ef.Value)
			e.Values = append(e.Values, ev)
		}
		extraEntities = append(extraEntities, e)
	}
	for name, es := range extraTypes.Structs {
		s := matter.NewStruct(nil, nil)
		s.Name = name
		s.Description = es.Description
		switch es.FabricScoping {
		case "unscoped":
			s.FabricScoping = matter.FabricScopingUnscoped
		case "scoped":
			s.FabricScoping = matter.FabricScopingScoped
		}
		for i, ef := range es.Fields {
			f := matter.NewField(nil, s, types.EntityTypeStructField)
			if ef.Value != "" {
				f.ID = matter.ParseNumber(ef.Value)
			} else {
				f.ID = matter.NewNumber(uint64(i))
			}
			f.Name = ef.Name
			var rank types.DataTypeRank
			if ef.List {
				rank = types.DataTypeRankList
			}
			f.Type = types.ParseDataType(ef.Type, rank)
			if ef.Constraint != "" {
				f.Constraint = constraint.ParseString(ef.Constraint)
			}
			if ef.Conformance != "" {
				f.Conformance = conformance.ParseConformance(ef.Conformance)
			} else {
				f.Conformance = conformance.Set{&conformance.Mandatory{}}
			}
			if ef.Fallback != "" {
				f.Fallback = constraint.ParseLimit(ef.Fallback)
			}
			if ef.Quality != "" {
				f.Quality = matter.ParseQuality(ef.Quality)
			}
			if ef.Access != "" {
				var parsed bool
				f.Access, parsed = spec.ParseAccess(ef.Access, types.EntityTypeStructField)
				if !parsed {
					fmt.Printf("failed to parse access string %q for extra struct field %s\n", ef.Access, ef.Name)
				}
			}
			s.Fields = append(s.Fields, f)
		}
		extraEntities = append(extraEntities, s)
	}
	for _, m := range entities {
		switch v := m.(type) {
		case *matter.ClusterGroup:
			for _, cl := range v.Clusters {
				err := addExtraAttributesAndCommandsToCluster(cl, extraTypes)
				if err != nil {
					return err
				}
				for _, e := range extraEntities {
					addExtraEntity(cl, e)
				}
			}
		case *matter.Cluster:
			err := addExtraAttributesAndCommandsToCluster(v, extraTypes)
			if err != nil {
				return err
			}
			for _, e := range extraEntities {
				addExtraEntity(v, e)
			}
		}
	}
	return nil
}

func addExtraEntity(cluster *matter.Cluster, e types.Entity) {
	switch e := e.(type) {
	case *matter.Bitmap:
		for _, bm := range cluster.Bitmaps {
			if bm.Name == e.Name {
				return
			}
		}
		e.SetParent(cluster)
		cluster.AddBitmaps(e)
	case *matter.Enum:
		for _, en := range cluster.Enums {
			if en.Name == e.Name {
				return
			}
		}
		e.SetParent(cluster)
		cluster.AddEnums(e)
	case *matter.Struct:
		for _, s := range cluster.Structs {
			if s.Name == e.Name {
				return
			}
		}
		e.SetParent(cluster)
		cluster.AddStructs(e)
	}
}

func addExtraAttributesAndCommandsToCluster(cluster *matter.Cluster, extraTypes *errata.SDKTypes) error {
	if override, ok := extraTypes.Clusters[cluster.Name]; ok {
		err := addExtraAttributes(cluster, override)
		if err != nil {
			return err
		}
		err = addExtraCommands(cluster, override)
		if err != nil {
			return err
		}
		err = addExtraEvents(cluster, override)
		if err != nil {
			return err
		}
	}
	if len(extraTypes.Attributes) > 0 {
		err := addExtraAttributes(cluster, &errata.SDKType{Attributes: extraTypes.Attributes})
		if err != nil {
			return err
		}
	}
	if len(extraTypes.Commands) > 0 {
		err := addExtraCommands(cluster, &errata.SDKType{Commands: extraTypes.Commands})
		if err != nil {
			return err
		}
	}
	if len(extraTypes.Events) > 0 {
		err := addExtraEvents(cluster, &errata.SDKType{Events: extraTypes.Events})
		if err != nil {
			return err
		}
	}
	return nil
}

func addExtraEvents(cluster *matter.Cluster, extra *errata.SDKType) error {
	existingEvents := make(map[string]struct{}, len(cluster.Events))
	for _, ev := range cluster.Events {
		existingEvents[ev.Name] = struct{}{}
	}

	for name, ev := range extra.Events {
		if _, ok := existingEvents[name]; ok {
			continue
		}
		event := matter.NewEvent(nil, cluster)
		event.Name = name
		if ev.Value != "" {
			event.ID = matter.ParseNumber(ev.Value)
		}
		if ev.Priority != "" {
			event.Priority = ev.Priority
		}
		if ev.Access != "" {
			var parsed bool
			event.Access, parsed = spec.ParseAccess(ev.Access, types.EntityTypeEvent)
			if !parsed {
				return fmt.Errorf("failed to parse access string %q for extra event %s in cluster %s from errata", ev.Access, name, cluster.Name)
			}
		}
		if ev.Conformance != "" {
			event.Conformance = conformance.ParseConformance(ev.Conformance)
			resolveExtraConformance(cluster, event.Conformance)
		}
		for i, f := range ev.Fields {
			field := matter.NewField(nil, event, types.EntityTypeEventField)
			field.Name = f.Name
			if f.Value != "" {
				field.ID = matter.ParseNumber(f.Value)
			} else {
				field.ID = matter.NewNumber(uint64(i))
			}
			var rank types.DataTypeRank
			if f.List {
				rank = types.DataTypeRankList
			}
			field.Type = types.ParseDataType(f.Type, rank)
			if f.Constraint != "" {
				field.Constraint = constraint.ParseString(f.Constraint)
			}
			if f.Conformance != "" {
				field.Conformance = conformance.ParseConformance(f.Conformance)
				resolveExtraConformance(cluster, field.Conformance)
			}
			if f.Fallback != "" {
				field.Fallback = constraint.ParseLimit(f.Fallback)
			}
			if f.Quality != "" {
				field.Quality = matter.ParseQuality(f.Quality)
			}
			if f.Access != "" {
				var parsed bool
				field.Access, parsed = spec.ParseAccess(f.Access, types.EntityTypeEventField)
				if !parsed {
					fmt.Printf("failed to parse access string %q for extra event field %s\n", f.Access, f.Name)
				}
			}
			event.Fields = append(event.Fields, field)
		}
		event.SetParent(cluster)
		cluster.Events = append(cluster.Events, event)
	}
	return nil
}


func addExtraAttributes(cluster *matter.Cluster, extra *errata.SDKType) error {
	existingAttributes := make(map[string]struct{}, len(cluster.Attributes))
	for _, f := range cluster.Attributes {
		existingAttributes[f.Name] = struct{}{}
	}

	for name, a := range extra.Attributes {
		if _, exists := existingAttributes[name]; exists {
			continue
		}
		field := matter.NewAttribute(nil, cluster)
		field.Name = name
		if a.Value != "" {
			field.ID = matter.ParseNumber(a.Value)
		}
		if a.Type != "" {
			var rank types.DataTypeRank = types.DataTypeRankScalar
			if a.List {
				rank = types.DataTypeRankList
			}
			field.Type = types.ParseDataType(a.Type, rank)
		}
		if a.Access != "" {
			var parsed bool
			field.Access, parsed = spec.ParseAccess(a.Access, types.EntityTypeAttribute)
			if !parsed {
				return fmt.Errorf("failed to parse access string %q for extra attribute %s in cluster %s from errata", a.Access, name, cluster.Name)
			}
		}
		if a.Conformance != "" {
			field.Conformance = conformance.ParseConformance(a.Conformance)
			resolveExtraConformance(cluster, field.Conformance)
		}
		if a.Constraint != "" {
			field.Constraint = constraint.ParseString(a.Constraint)
		}
		if a.Fallback != "" {
			field.Fallback = constraint.ParseLimit(a.Fallback)
		}
		if a.Quality != "" {
			field.Quality = matter.ParseQuality(a.Quality)
		}
		field.SetParent(cluster)
		cluster.Attributes = append(cluster.Attributes, field)
	}
	return nil
}

func addExtraCommands(cluster *matter.Cluster, extra *errata.SDKType) error {
	existingCommands := make(map[string]struct{}, len(cluster.Commands))
	for _, cmd := range cluster.Commands {
		existingCommands[cmd.Name] = struct{}{}
	}

	for name, cmd := range extra.Commands {
		if _, ok := existingCommands[name]; ok {
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
				return fmt.Errorf("failed to parse access string %q for extra command %s in cluster %s from errata", cmd.Access, name, cluster.Name)
			}
		}
		if cmd.Conformance != "" {
			command.Conformance = conformance.ParseConformance(cmd.Conformance)
			resolveExtraConformance(cluster, command.Conformance)
		}
		if cmd.Direction != "" {
			switch cmd.Direction {
			case "client":
				command.Direction = matter.InterfaceClient
			case "server":
				command.Direction = matter.InterfaceServer
			}
		}
		if cmd.Response != "" {
			command.Response = types.ParseDataType(cmd.Response, types.DataTypeRankScalar)
		}
		for i, f := range cmd.Fields {
			field := matter.NewField(nil, command, types.EntityTypeCommandField)
			field.Name = f.Name
			if f.Value != "" {
				field.ID = matter.ParseNumber(f.Value)
			} else {
				field.ID = matter.NewNumber(uint64(i))
			}
			var rank types.DataTypeRank
			if f.List {
				rank = types.DataTypeRankList
			}
			field.Type = types.ParseDataType(f.Type, rank)
			if f.Constraint != "" {
				field.Constraint = constraint.ParseString(f.Constraint)
			}
			if f.Conformance != "" {
				field.Conformance = conformance.ParseConformance(f.Conformance)
				resolveExtraConformance(cluster, field.Conformance)
			}
			if f.Fallback != "" {
				field.Fallback = constraint.ParseLimit(f.Fallback)
			}
			if f.Quality != "" {
				field.Quality = matter.ParseQuality(f.Quality)
			}
			if f.Access != "" {
				var parsed bool
				field.Access, parsed = spec.ParseAccess(f.Access, types.EntityTypeCommandField)
				if !parsed {
					fmt.Printf("failed to parse access string %q for extra command field %s\n", f.Access, f.Name)
				}
			}
			command.Fields = append(command.Fields, field)
		}
		command.SetParent(cluster)
		cluster.Commands = append(cluster.Commands, command)
	}

	// Since addExtraCommands runs after the main type resolution pass, we must
	// manually resolve the response command references to their corresponding command entities.
	for _, command := range cluster.Commands {
		if command.Response != nil && command.Response.Entity == nil && command.Response.Name != "" {
			var desiredDirection matter.Interface
			switch command.Direction {
			case matter.InterfaceServer:
				desiredDirection = matter.InterfaceClient
			case matter.InterfaceClient:
				desiredDirection = matter.InterfaceServer
			}
			for _, cmd := range cluster.Commands {
				if cmd.Direction == desiredDirection && cmd.Name == command.Response.Name {
					command.Response.Entity = cmd
					break
				}
			}
		}
	}
	return nil
}

func resolveExtraConformance(cluster *matter.Cluster, conf conformance.Conformance) {
	if conf == nil {
		return
	}
	switch c := conf.(type) {
	case conformance.Set:
		for _, cx := range c {
			resolveExtraConformance(cluster, cx)
		}
	case *conformance.Mandatory:
		resolveExtraConformanceExpression(cluster, c.Expression)
	case *conformance.Optional:
		resolveExtraConformanceExpression(cluster, c.Expression)
	}
}

func resolveExtraConformanceExpression(cluster *matter.Cluster, expr conformance.Expression) {
	if expr == nil {
		return
	}
	switch e := expr.(type) {
	case *conformance.EqualityExpression:
		resolveExtraConformanceExpression(cluster, e.Left)
		resolveExtraConformanceExpression(cluster, e.Right)
	case *conformance.LogicalExpression:
		resolveExtraConformanceExpression(cluster, e.Left)
		for _, re := range e.Right {
			resolveExtraConformanceExpression(cluster, re)
		}
	case *conformance.IdentifierExpression:
		if e.Entity == nil && cluster.Features != nil {
			for f := range cluster.Features.FeatureBits() {
				if f.Code == e.ID || f.Name() == e.ID {
					e.Entity = f
					break
				}
			}
		}
	}
}

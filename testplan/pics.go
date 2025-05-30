package testplan

import (
	"fmt"
	"log/slog"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/testplan/pics"
)

type PicsAlias struct {
	EntityType types.EntityType
	Pics       string
	Alias      string
	Comments   []string
}

func BuildPicsMap(spec *spec.Specification, t *Test) (aliases map[string]string) {
	aliases = make(map[string]string)
	clusters := make(map[string]*matter.Cluster)
	for _, tp := range t.PICSList {
		buildMapForPics(spec, tp, clusters, aliases)
	}
	for _, ts := range t.Groups {
		for _, ts := range ts.Steps {
			buildMapForPics(spec, ts.PICSSet, clusters, aliases)
		}
	}
	return
}

func buildMapForPics(spec *spec.Specification, exp pics.Expression, clusters map[string]*matter.Cluster, aliases map[string]string) {
	switch exp := exp.(type) {
	case *pics.PICSExpression:
		_, ok := aliases[exp.PICS]
		if ok {
			return
		}
		aliases[exp.PICS] = makePicsAlias(spec, clusters, exp.PICS)
	case *pics.LogicalExpression:
		buildMapForPics(spec, exp.Left, clusters, aliases)
		for _, re := range exp.Right {
			buildMapForPics(spec, re, clusters, aliases)
		}
	}
}

var picsPattern = regexp.MustCompile(`^([A-Z]+)\.([SC])\.([FACE])([0-9a-fA-F]+)(?:\.(Tx|Rsp))?$`)

func makePicsAlias(spec *spec.Specification, clusters map[string]*matter.Cluster, p string) string {
	parts := picsPattern.FindStringSubmatch(p)
	if len(parts) < 5 {
		return ""
	}
	iface := parts[2]
	if !strings.EqualFold(iface, "S") {
		// We don't handle aliases for client PICS
		return ""
	}
	clusterPics := parts[1]
	cluster, ok := clusters[clusterPics]
	if !ok {
		for c := range spec.Clusters {
			if strings.EqualFold(c.PICS, clusterPics) {
				cluster = c
				clusters[c.PICS] = c
				break
			}
		}
	}
	if cluster == nil {
		slog.Warn("Unable to find matching cluster for PICS", slog.String("pics", clusterPics))
		return ""
	}
	entityType := parts[3]
	id, err := strconv.ParseUint(parts[4], 16, 64)
	if err != nil {
		slog.Warn("Error parsing PICS id", slog.String("id", parts[4]), slog.Any("error", err))
		return ""
	}
	switch entityType {
	case "F":
		if cluster.Features == nil {
			slog.Warn("Cluster missing features", slog.String("clusterName", cluster.Name), slog.Uint64("Feature Bit", id))
			return ""
		}
		for _, b := range cluster.Features.Bits {
			f := b.(*matter.Feature)
			from, to, err := f.Bits()
			if err != nil {
				slog.Warn("Error parsing feature bits", slog.String("feature", f.Code), slog.Any("error", err))
				return ""
			}
			if from != to {
				continue
			}
			if from == id {
				return fmt.Sprintf("has%sFeature", matter.Case(f.Name()))
			}
		}
		slog.Warn("Unable to find matching feature for PICS", slog.String("clusterName", cluster.Name), slog.String("pics", p))
	case "A":
		for _, a := range cluster.Attributes {
			if !a.ID.Valid() {
				continue
			}
			if a.ID.Value() == id {
				return fmt.Sprintf("has%sAttribute", matter.Case(a.Name))
			}
		}
		slog.Warn("Unable to find matching attribute for PICS", slog.String("clusterName", cluster.Name), slog.String("pics", p))
	case "C":
		var direction matter.Interface
		switch parts[5] {
		case "Rsp":
			direction = matter.InterfaceServer
		case "Tx":
			direction = matter.InterfaceClient
		default:
			slog.Warn("Unknown command direction while building aliases", slog.String("direction", parts[5]))
		}
		for _, c := range cluster.Commands {
			if !c.ID.Valid() {
				continue
			}
			if c.ID.Value() == id && c.Direction == direction {
				return fmt.Sprintf("has%sCommand", matter.Case(c.Name))
			}
		}
		var commands []any
		for _, c := range cluster.Commands {
			commands = append(commands, slog.Group("command", slog.Uint64("id", c.ID.Value()), slog.String("direction", c.Direction.String())))
		}
		slog.Warn("Unable to find matching command for PICS", slog.String("clusterName", cluster.Name), slog.String("pics", p), slog.Uint64("id", id), slog.String("direction", direction.String()), slog.Group("checked", commands...))

	case "E":
		for _, e := range cluster.Events {
			if !e.ID.Valid() {
				continue
			}
			if e.ID.Value() == id {
				return fmt.Sprintf("has%sEvent", matter.Case(e.Name))
			}
		}
		slog.Warn("Unable to find matching event for PICS", slog.String("clusterName", cluster.Name), slog.String("pics", p))
	default:
		slog.Warn("Unknown entity type while building aliases", slog.String("entityType", entityType))
	}
	slog.Info("pics parts", slog.String("cluster", parts[1]), slog.String("interface", parts[2]), slog.String("entity", parts[3]), slog.String("id", parts[4]))
	return ""
}

func BuildPicsAliasList(picsAliases map[string]string) (aliases []*PicsAlias) {
	for pics, alias := range picsAliases {
		if alias == "" {
			continue
		}
		pa := &PicsAlias{Pics: pics, Alias: alias}
		if strings.HasSuffix(alias, "Feature") {
			pa.EntityType = types.EntityTypeFeature
		} else if strings.HasSuffix(alias, "Command") {
			pa.EntityType = types.EntityTypeCommand
		} else if strings.HasSuffix(alias, "Attribute") {
			pa.EntityType = types.EntityTypeAttribute
		} else if strings.HasSuffix(alias, "Event") {
			pa.EntityType = types.EntityTypeEvent
		}
		aliases = append(aliases, pa)
	}
	if len(aliases) == 0 {
		return
	}
	slices.SortStableFunc(aliases, func(a, b *PicsAlias) int {
		if a.EntityType == b.EntityType {
			return strings.Compare(a.Alias, b.Alias)
		}
		switch a.EntityType {
		case types.EntityTypeFeature:
			return -1
		case types.EntityTypeAttribute:
			switch b.EntityType {
			case types.EntityTypeFeature:
				return 1
			default:
				return -1
			}
		case types.EntityTypeCommand:
			switch b.EntityType {
			case types.EntityTypeFeature, types.EntityTypeAttribute:
				return 1
			default:
				return -1
			}
		case types.EntityTypeEvent:
			return 1
		}
		if a.EntityType < b.EntityType {
			return -1
		}
		return 1
	})
	return
}

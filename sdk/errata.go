package sdk

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

// patchMiscellaneousErrata is a grab bag of oddities in the spec that need to be corrected for use in the SDK
func patchMiscellaneousErrata(spec *spec.Specification) {
	patchLabelCluster(spec)
	patchLevelControlCluster(spec)
}

func patchContentLaunchCluster(spec *spec.Specification) {
	/*
		Another hacky workaround: the spec defines CharacteristicEnum under Media Playback, and Content Launcher
		just has a reference to it; by default, we would share the enums, but the hand-coded ZAP XML has its own copy of the
		enum. This function makes a copy of the enum and updates references
	*/
	contentLaunchCluster, ok := spec.ClustersByName["Content Launcher"]
	if !ok {
		slog.Warn("Could not find Content Launch cluster for patching")
		return
	}
	mediaPlaybackCluster, ok := spec.ClustersByName["Media Playback"]
	if !ok {
		slog.Warn("Could not find Media Playback cluster for patching")
		return
	}
	for _, s := range mediaPlaybackCluster.Enums {
		if s.Name != "CharacteristicEnum" {
			continue
		}
		clone := s.CloneTo(contentLaunchCluster)
		contentLaunchCluster.MoveEnum(clone)
		spec.ClusterRefs.Add(contentLaunchCluster, clone)
		// Update any references to the enum on the Content Launch cluster
		contentLaunchCluster.TraverseDataTypes(func(parent, entity types.Entity) parse.SearchShould {
			if entity == s {
				switch parent := parent.(type) {
				case *matter.Field:
					fieldType := parent.Type
					if fieldType == nil {
						return parse.SearchShouldContinue
					}
					if fieldType.IsArray() {
						fieldType = fieldType.EntryType
					}
					if fieldType == nil {
						return parse.SearchShouldContinue
					}
					fieldType.Entity = clone
				}
			}
			return parse.SearchShouldContinue
		})
		break

	}
}

func patchLabelCluster(spec *spec.Specification) {
	/*
		Another hacky workaround: the spec defines LabelStruct under a base cluster called Label Cluster, but the
		ZAP XML has this struct under Fixed Label and User Label
	*/
	fixedLabelCluster, ok := spec.ClustersByName["Fixed Label"]
	if !ok {
		slog.Warn("Could not find Fixed Label cluster")
		return
	}
	userLabelCluster, ok := spec.ClustersByName["User Label"]
	if !ok {
		slog.Warn("Could not find User Label cluster")
		return
	}
	labelCluster, ok := spec.ClustersByName["Label"]
	if !ok {
		slog.Warn("Could not find Label cluster")
		return
	}
	for i, s := range labelCluster.Structs {
		if s.Name == "LabelStruct" {
			fixedLabelCluster.MoveStruct(s)
			spec.DataTypeRefs.Add(fixedLabelCluster, s)

			clone := s.CloneTo(userLabelCluster)
			userLabelCluster.MoveStruct(clone)
			spec.DataTypeRefs.Add(userLabelCluster, clone)

			userLabelCluster.TraverseDataTypes(func(parent, entity types.Entity) parse.SearchShould {
				if entity == s {
					switch parent := parent.(type) {
					case *matter.Field:
						fieldType := parent.Type
						if fieldType == nil {
							return parse.SearchShouldContinue
						}
						if fieldType.IsArray() {
							fieldType = fieldType.EntryType
						}
						if fieldType == nil {
							return parse.SearchShouldContinue
						}
						fieldType.Entity = clone
					}
				}
				return parse.SearchShouldContinue
			})

			labelCluster.Structs = slices.Delete(labelCluster.Structs, i, i+1)
			break
		}
	}
}

func patchLevelControlCluster(spec *spec.Specification) {
	levelControlCluster, ok := spec.ClustersByID[0x0008]
	if !ok {
		slog.Warn("Unable to patch Level Control cluster; not found")
		return
	}
	// Level Control cluster has a series of OnOff commands that are spec'd to have the same parameters as their non-OnOff versions,
	// but those parameters aren't explicitly listed in the spec, so we copy them over
	for _, c := range levelControlCluster.Commands {
		if !text.HasCaseInsensitiveSuffix(c.Name, "WithOnOff") {
			continue
		}
		if len(c.Fields) > 0 {
			// Assume that if we have fields already, this has been fixed in the spec
			continue
		}
		baseName := text.TrimCaseInsensitiveSuffix(c.Name, "WithOnOff")
		var matchingCommand *matter.Command
		for _, mc := range levelControlCluster.Commands {
			if strings.EqualFold(mc.Name, baseName) {
				matchingCommand = mc
				break
			}
		}
		if matchingCommand == nil {
			slog.Warn("Unable to find matching command for Level Control cluster", slog.String("commandName", c.Name))
			continue
		}
		for _, f := range matchingCommand.Fields {
			clone := f.CloneTo(c)
			clone.Type.Entity = f.Type.Entity
			c.Fields = append(c.Fields, clone)
		}
	}
}

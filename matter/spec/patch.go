package spec

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

// patchSpecForSdk is a grab bag of oddities in the spec that need to be corrected for use in the SDK
func patchSpecForSdk(spec *Specification) error {
	patchScenesCluster(spec)
	patchLabelCluster(spec)
	patchLevelControlCluster(spec)
	patchContentLaunchCluster(spec)

	return nil
}

func resolveAtomicOperations(spec *Specification) {
	for m := range spec.DocRefs {
		switch v := m.(type) {
		case *matter.ClusterGroup:
			for _, cl := range v.Clusters {
				if hasAtomicAttributes(cl) {
					addAtomicOperations(spec, cl)
				}
			}
		case *matter.Cluster:
			if hasAtomicAttributes(v) {
				addAtomicOperations(spec, v)
			}
		}
	}
}

func patchScenesCluster(spec *Specification) {
	scenesCluster, ok := spec.ClustersByName["Scenes Management"]
	if !ok {
		slog.Warn("Could not find Scenes cluster")
		return
	}
	var addSceneCommand *matter.Command
	for _, c := range scenesCluster.Commands {
		if strings.EqualFold(c.Name, "AddScene") {
			addSceneCommand = c
			break
		}
	}
	if addSceneCommand == nil {
		slog.Warn("Could not find AddScene command on Scenes cluster")
		return
	}
	for _, f := range addSceneCommand.Fields {
		if strings.EqualFold(f.Name, "ExtensionFieldSetStructs") {
			f.Name = "ExtensionFieldSets"
			break
		}
	}
}

func patchContentLaunchCluster(spec *Specification) {
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
		spec.ClusterRefs.Remove(contentLaunchCluster, s)
		clone := s.Clone()
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

func patchLabelCluster(spec *Specification) {
	/*
		Another hacky workaround: the spec defines LabelStruct under a base cluster called Label Cluster, but the
		ZAP XML has this struct under Fixed Label
	*/
	fixedLabelCluster, ok := spec.ClustersByName["Fixed Label"]
	if !ok {
		slog.Warn("Could not find Fixed Label cluster")
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
			labelCluster.Structs = slices.Delete(labelCluster.Structs, i, i+1)
			break
		}
	}
}

func patchLevelControlCluster(spec *Specification) {
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
			clone := f.Clone()
			clone.Type.Entity = f.Type.Entity
			c.Fields = append(c.Fields, clone)
		}
	}
}

func hasAtomicAttributes(cluster *matter.Cluster) bool {
	for _, f := range cluster.Attributes {
		if f.Quality.Has(matter.QualityAtomicWrite) {
			return true
		}
	}
	return false
}

func addAtomicOperations(spec *Specification, cluster *matter.Cluster) {
	var atomicRequest, atomicResponse *matter.Command
	for o := range spec.GlobalObjects {
		switch o := o.(type) {
		case *matter.Command:
			switch o.Name {
			case "AtomicRequest":
				atomicRequest = o
			case "AtomicResponse":
				atomicResponse = o
			}
		}
	}

	if atomicRequest == nil {
		slog.Warn("Could not find AtomicRequest command")
		return
	}
	if atomicResponse == nil {
		slog.Warn("Could not find AtomicResponse command")
		return
	}
	cluster.Commands = append(cluster.Commands, atomicRequest)
	cluster.Commands = append(cluster.Commands, atomicResponse)
	spec.ClusterRefs.Add(cluster, atomicRequest)
	spec.ClusterRefs.Add(cluster, atomicResponse)
}

package spec

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type Builder struct {
	specRoot string

	Spec *Specification

	ignoreHierarchy bool

	conformanceFailures map[any]referenceFailure
	constraintFailures  map[any]referenceFailure
}

func NewBuilder(specRoot string, options ...BuilderOption) Builder {
	b := Builder{
		specRoot:            specRoot,
		conformanceFailures: make(map[any]referenceFailure),
		constraintFailures:  make(map[any]referenceFailure),
	}
	for _, o := range options {
		o(&b)
	}
	return b
}

func (sp Builder) Name() string {
	return "Building spec"
}

func (sp *Builder) Process(cxt context.Context, inputs []*pipeline.Data[*Library]) (outputs []*pipeline.Data[*asciidoc.Document], err error) {
	docs := make([]*Library, 0, len(inputs))
	for _, i := range inputs {
		docs = append(docs, i.Content)
	}
	var referencedDocs []*asciidoc.Document
	referencedDocs, err = sp.buildSpec(cxt, docs)
	for _, d := range referencedDocs {
		outputs = append(outputs, pipeline.NewData(d.Path.Absolute, d))
	}
	return
}

func (sp *Builder) buildSpec(cxt context.Context, libraries []*Library) (referencedDocs []*asciidoc.Document, err error) {

	sp.Spec = newSpec(sp.specRoot)
	spec := sp.Spec

	slices.SortStableFunc(libraries, func(a *Library, b *Library) int {
		return strings.Compare(a.Root.Path.Relative, b.Root.Path.Relative)
	})

	docs := make(map[string][]*asciidoc.Document)

	for _, l := range libraries {
		l.Spec = spec
		for _, d := range l.Docs {
			l.DocType(d)
			existing, ok := spec.libraryIndex[d]
			if ok {
				slog.Warn("Document referenced by multiple libraries", "source", d.Path.Relative, "other", existing.Root.Path.Relative)
			} else {
				spec.libraryIndex[d] = l
			}
			docs[d.Path.Absolute] = append(docs[d.Path.Absolute], d)
		}
		for top := range parse.Skim[*asciidoc.Section](l, l.Root, l.Children(l.Root)) {
			AssignSectionTypes(l, l, l, l.Root, top)
			break
		}
		dumpLibrary(l)
		if slog.Default().Enabled(cxt, slog.LevelDebug) {
			dumpLibrary(l)
		}
	}

	for _, docs := range docs {
		if len(docs) == 1 {
			referencedDocs = append(referencedDocs, docs[0])
			spec.Docs[docs[0].Path.Relative] = docs[0]
		}
	}

	var basicInformationCluster, bridgedDeviceBasicInformationCluster *matter.Cluster
	basicInformationCluster, bridgedDeviceBasicInformationCluster, err = sp.readEntities(spec, libraries)
	if err != nil {
		return
	}

	if basicInformationCluster == nil {
		err = fmt.Errorf("missing Basic Information Cluster in spec")
		return
	}
	if bridgedDeviceBasicInformationCluster == nil {
		err = fmt.Errorf("missing Bridged Device Basic Information Cluster in spec")
		return
	}
	if spec.RootNodeDeviceType == nil {
		err = fmt.Errorf("missing Root Node Device Type in spec")
		return
	}
	if spec.BaseDeviceType == nil {
		err = fmt.Errorf("missing Base Device Type in spec")
		return
	}

	sp.resolveClusterDataTypeReferences(true)
	sp.resolveGlobalDataTypeReferences()
	if !sp.ignoreHierarchy {
		sp.resolveHierarchy()
	}
	err = spec.associateDeviceTypeRequirements()
	if err != nil {
		return
	}

	sp.resolveClusterDataTypeReferences(false)

	sp.resolveConformances()
	sp.resolveConstraints()
	err = updateBridgedBasicInformationCluster(spec, basicInformationCluster, bridgedDeviceBasicInformationCluster)
	if err != nil {
		return
	}

	spec.BuildClusterReferences()
	spec.BuildDataTypeReferences()

	sp.noteConformanceResolutionFailures(spec)
	sp.noteConstraintResolutionFailures(spec)

	Validate(spec)

	return
}

func (sp *Builder) readEntities(spec *Specification, libraries []*Library) (basicInformationCluster *matter.Cluster, bridgedDeviceBasicInformationCluster *matter.Cluster, err error) {

	for _, library := range libraries {
		library.indexCrossReferences()
		err = library.indexAnchors()
		if err != nil {
			return
		}
		for doc, result := range library.parseEntities(spec) {
			if e, ok := result.(error); ok {
				if pe, isParseError := e.(Error); isParseError {
					slog.Error("parse error parsing entities", "err", e, log.Path("source", pe))
					spec.addError(pe)
				} else {
					slog.Error("error parsing entities", "err", e)
				}
				continue
			}
			switch entity := result.(type) {
			case *matter.ClusterGroup:
				for _, c := range entity.Clusters {
					sp.addCluster(doc, c)
				}
			case *matter.Cluster:
				switch entity.Name {
				case "Basic Information":
					basicInformationCluster = entity
				case "Bridged Device Basic Information":
					bridgedDeviceBasicInformationCluster = entity
				}
				sp.addCluster(doc, entity)
			case *matter.DeviceType:
				spec.DeviceTypes = append(spec.DeviceTypes, entity)
				if entity.ID.Valid() {
					if existing, ok := spec.DeviceTypesByID[entity.ID.Value()]; ok {
						slog.Error("duplicate device type ID", slog.String("deviceTypeId", entity.ID.HexString()), log.Path("previousSource", existing), log.Path("newSource", entity))
						spec.addError(&DuplicateEntityIDError{Entity: entity, Previous: existing})
					} else {
						spec.DeviceTypesByID[entity.ID.Value()] = entity
					}

				}
				existing, ok := spec.DeviceTypesByName[entity.Name]
				if ok {
					slog.Warn("Duplicate Device Type Name", slog.String("deviceTypeId", entity.ID.HexString()), slog.String("deviceTypeName", entity.Name), slog.String("existingDeviceTypeId", existing.ID.HexString()))
					spec.addError(&DuplicateEntityNameError{Entity: entity, Previous: existing})
				}
				spec.DeviceTypesByName[entity.Name] = entity
				switch entity.Name {
				case "Root Node":
					spec.RootNodeDeviceType = entity
				case "Base Device Type":
					spec.BaseDeviceType = entity
				}
			case *matter.Namespace:
				spec.Namespaces = append(spec.Namespaces, entity)
			case *matter.Bitmap:
				slog.Debug("Found global bitmap", "name", entity.Name, "path", doc.Path)
				spec.addEntityByName(entity.Name, entity, nil)
				spec.GlobalObjects[entity] = doc
			case *matter.Enum:
				slog.Debug("Found global enum", "name", entity.Name, "path", doc.Path)
				spec.addEntityByName(entity.Name, entity, nil)
				spec.GlobalObjects[entity] = doc
			case *matter.Struct:
				slog.Debug("Found global struct", "name", entity.Name, "path", doc.Path)
				spec.addEntityByName(entity.Name, entity, nil)
				spec.GlobalObjects[entity] = doc
			case *matter.TypeDef:
				slog.Debug("Found global typedef", "name", entity.Name, "path", doc.Path)
				spec.addEntityByName(entity.Name, entity, nil)
				spec.GlobalObjects[entity] = doc
			case *matter.Command:
				spec.addEntityByName(entity.Name, entity, nil)
				spec.GlobalObjects[entity] = doc
			case *matter.Event:
				spec.addEntityByName(entity.Name, entity, nil)
				spec.GlobalObjects[entity] = doc

			default:
				slog.Warn("unknown entity type", "path", doc.Path, "type", fmt.Sprintf("%T", entity))

			}
			switch entity := result.(type) {
			case *matter.ClusterGroup:
				spec.entityRefs[doc] = append(spec.entityRefs[doc], entity)
				for _, c := range entity.Clusters {
					spec.DocRefs[c] = doc
					spec.LibraryRefs[c] = library
				}
			case types.Entity:
				spec.entityRefs[doc] = append(spec.entityRefs[doc], entity)
				spec.DocRefs[entity] = doc
				spec.LibraryRefs[entity] = library
			}
		}
	}
	return
}

func (spec *Specification) BuildClusterReferences() {
	iterateOverDataTypes(spec, func(cluster *matter.Cluster, parent, entity types.Entity) {
		if cluster != nil {
			spec.ClusterRefs.Add(cluster, entity)
		}
	})
}

func (sp *Builder) addCluster(doc *asciidoc.Document, cluster *matter.Cluster) {
	sp.Spec.Clusters[cluster] = struct{}{}
	if cluster.ID.Valid() {
		existing, ok := sp.Spec.ClustersByID[cluster.ID.Value()]
		if ok {
			slog.Error("Duplicate cluster ID", slog.String("clusterId", cluster.ID.HexString()), slog.String("clusterName", cluster.Name), slog.String("existingClusterName", existing.Name))
			sp.Spec.addError(&DuplicateEntityIDError{Entity: cluster, Previous: existing})
		}
		sp.Spec.ClustersByID[cluster.ID.Value()] = cluster
	} else {
		idText := cluster.ID.Text()
		if !strings.EqualFold(idText, "n/a") {
			if strings.EqualFold(idText, "ID-TBD") {
				slog.Warn("Cluster has not yet been assigned an ID; this may cause issues with generated code", slog.String("clusterName", cluster.Name))
			} else {
				slog.Warn("Cluster has invalid ID; this may cause issues with generated code", slog.String("clusterId", idText), slog.String("clusterName", cluster.Name))

			}
		}
	}
	existing, ok := sp.Spec.ClustersByName[cluster.Name]
	if ok {
		slog.Error("Duplicate cluster Name", slog.String("clusterId", cluster.ID.HexString()), slog.String("clusterName", cluster.Name), slog.String("existingClusterId", existing.ID.HexString()))
		sp.Spec.addError(&DuplicateEntityNameError{Entity: cluster, Previous: existing})

	}
	sp.Spec.ClustersByName[cluster.Name] = cluster

	for _, en := range cluster.Bitmaps {
		sp.Spec.addEntityByName(en.Name, en, cluster)
	}
	for _, en := range cluster.Enums {
		sp.Spec.addEntityByName(en.Name, en, cluster)
	}
	for _, en := range cluster.Structs {
		sp.Spec.addEntityByName(en.Name, en, cluster)
	}
	for _, en := range cluster.TypeDefs {
		sp.Spec.addEntityByName(en.Name, en, cluster)
	}
	sp.noteDocRefs(doc, cluster)
}

func (sp *Builder) noteDocRefs(doc *asciidoc.Document, cluster *matter.Cluster) {
	for _, bm := range cluster.Bitmaps {
		sp.Spec.DocRefs[bm] = doc
	}
	for _, e := range cluster.Enums {
		sp.Spec.DocRefs[e] = doc
	}
	for _, s := range cluster.Structs {
		sp.Spec.DocRefs[s] = doc
	}
	for _, td := range cluster.TypeDefs {
		sp.Spec.DocRefs[td] = doc
	}
	for _, a := range cluster.Attributes {
		sp.Spec.DocRefs[a] = doc
	}
	for _, e := range cluster.Events {
		sp.Spec.DocRefs[e] = doc
	}
	for _, cmd := range cluster.Commands {
		sp.Spec.DocRefs[cmd] = doc
	}
}

func (sp *Builder) getTagNamespace(field *matter.Field) {
	for _, ns := range sp.Spec.Namespaces {
		if strings.EqualFold(ns.Name, field.Type.Name) {
			field.Type.Entity = ns
			return
		}
	}
	if field.Type.Name != "tag" {
		// Warn on unknown tag namespace, except for the example namespace "tag"
		slog.Warn("failed to match tag name space", slog.String("name", field.Name), log.Path("field", field), slog.String("namespace", field.Type.Name))
	}
}

func clusterName(cluster *matter.Cluster) string {
	if cluster != nil {
		return cluster.Name
	}
	return "none"
}

func (sp *Builder) resolveHierarchy() {
	for _, c := range sp.Spec.ClustersByName {
		if c.Hierarchy == "Base" {
			continue
		}
		base, ok := sp.Spec.ClustersByName[c.Hierarchy]
		if !ok {
			slog.Warn("Failed to find base cluster", "cluster", c.Name, "baseCluster", c.Hierarchy)
			sp.Spec.addError(&UnknownBaseClusterError{Cluster: c})
			continue
		}
		_, err := c.Inherit(base)
		if err != nil {
			slog.Warn("Failed to inherit from base cluster", "cluster", c.Name, "baseCluster", c.Hierarchy, "error", err)
		}
		doc, ok := sp.Spec.DocRefs[c]
		if ok {
			// We may have created some new entities during the inherit, so make sure their doc refs are set
			sp.noteDocRefs(doc, c)
		}
	}
}

func updateBridgedBasicInformationCluster(spec *Specification, basicInformationCluster *matter.Cluster, bridgedBasicInformationCluster *matter.Cluster) error {

	linkedEntities, err := bridgedBasicInformationCluster.Inherit(basicInformationCluster)
	if err != nil {
		return err
	}
	for _, linkedEntity := range linkedEntities {
		spec.addEntity(linkedEntity, bridgedBasicInformationCluster)
	}
	return nil
}

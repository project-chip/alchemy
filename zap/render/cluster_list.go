package render

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/tidwall/pretty"
)

type ClusterListPatcher struct {
	spec    *spec.Specification
	sdkRoot string
}

func NewClusterListPatcher(spec *spec.Specification, sdkRoot string) *ClusterListPatcher {
	return &ClusterListPatcher{spec: spec, sdkRoot: sdkRoot}
}

func (p ClusterListPatcher) Name() string {
	return "Patching files with cluster list"
}

func (p ClusterListPatcher) Process(cxt context.Context, inputs []*pipeline.Data[*asciidoc.Document]) (outputs []*pipeline.Data[[]byte], err error) {

	clusterListPath := path.Join(p.sdkRoot, "/src/app/zap_cluster_list.json")
	var clusterListBytes []byte
	clusterListBytes, err = os.ReadFile(clusterListPath)
	if err != nil {
		return
	}

	var o internal.JSONMap
	err = json.Unmarshal(clusterListBytes, &o)
	if err != nil {
		err = fmt.Errorf("error unmarshaling %s: %w", clusterListPath, err)
		return
	}

	var names []string
	for _, input := range inputs {
		doc := input.Content

		library, ok := p.spec.LibraryForDocument(input.Content)
		if !ok {
			err = fmt.Errorf("unable to find library for doc %s", doc.Path)
			return
		}
		entities := p.spec.EntitiesForDocument(input.Content)
		errata := library.ErrataForPath(doc.Path.Relative)
		for _, e := range entities {
			switch e := e.(type) {
			case *matter.Cluster:
				if !e.ID.Valid() {
					slog.Debug("Skipping adding cluster, no valid ID", slog.String("clusterName", e.Name))
					continue
				}
				name := e.Name
				if !text.HasCaseInsensitiveSuffix(name, " cluster") {
					name += " Cluster"
				}
				name = strings.ToUpper(matter.CaseWithSeparator(name, '_'))
				if errata.SDK.ClusterListKeys != nil {
					if alias, ok := errata.SDK.ClusterListKeys[name]; ok {
						name = alias
					}

				}
				names = append(names, name)
			}
		}
	}

	err = insertClusterName(&o, "ClientDirectories", names)
	if err != nil {
		return
	}

	err = insertClusterName(&o, "ServerDirectories", names)
	if err != nil {
		return
	}

	clusterListBytes, err = json.MarshalIndent(&o, "", "    ")
	if err != nil {
		return
	}
	clusterListBytes = pretty.PrettyOptions(clusterListBytes, &prettyOptions)
	outputs = append(outputs, pipeline.NewData(clusterListPath, clusterListBytes))
	return
}

func insertClusterName(o *internal.JSONMap, key string, names []string) error {
	val, ok := o.Get(key)
	if !ok {
		return fmt.Errorf("no %s field in zap_cluster_list.json", key)
	}
	is, ok := val.(*internal.JSONMap)
	if !ok {
		return fmt.Errorf("%s not a map in zap_cluster_list.json; %T", key, val)
	}
	var insertedNames []string
	for _, name := range names {
		if _, ok := is.Get(name); !ok {
			is.Set(name, []string{})
			insertedNames = append(insertedNames, name)
		}
	}

	is.SortKeys(func(keys []string) {
		reorderLinesSemiAlphabetically(keys, insertedNames, 0)
	})
	return nil
}

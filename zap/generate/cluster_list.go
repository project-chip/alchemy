package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/iancoleman/orderedmap"
	"github.com/iancoleman/strcase"
)

type ClusterListPatcher struct {
	sdkRoot string
}

func NewClusterListPatcher(sdkRoot string) *ClusterListPatcher {
	return &ClusterListPatcher{sdkRoot: sdkRoot}
}

func (p ClusterListPatcher) Name() string {
	return "Patching files with cluster list"
}

func (p ClusterListPatcher) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeCollective
}

func (p ClusterListPatcher) Process(cxt context.Context, inputs []*pipeline.Data[*ascii.Doc]) (outputs []*pipeline.Data[[]byte], err error) {

	clusterListPath := path.Join(p.sdkRoot, "/src/app/zap_cluster_list.json")
	var clusterListBytes []byte
	clusterListBytes, err = os.ReadFile(clusterListPath)
	if err != nil {
		return
	}

	o := orderedmap.New()
	err = json.Unmarshal(clusterListBytes, &o)
	if err != nil {
		return
	}

	var names []string
	for _, input := range inputs {
		doc := input.Content
		path := doc.Path
		name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)) + " Cluster"
		name = strcase.ToScreamingSnake(name)
		names = append(names, name)
	}

	err = insertClusterName(o, "ClientDirectories", names)
	if err != nil {
		return
	}

	err = insertClusterName(o, "ServerDirectories", names)
	if err != nil {
		return
	}

	clusterListBytes, err = json.MarshalIndent(o, "", "    ")
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData[[]byte](clusterListPath, clusterListBytes))
	return
}

func insertClusterName(o *orderedmap.OrderedMap, key string, names []string) error {
	val, ok := o.Get(key)
	if !ok {
		return fmt.Errorf("no %s field in zap_cluster_list.json", key)
	}
	is, ok := val.(orderedmap.OrderedMap)
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

	o.Set(key, is)
	return nil
}

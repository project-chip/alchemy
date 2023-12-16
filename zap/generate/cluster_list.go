package generate

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/hasty/alchemy/ascii"
	"github.com/iancoleman/orderedmap"
	"github.com/iancoleman/strcase"
)

func patchClusterList(zclRoot string, docs []*ascii.Doc) error {
	clusterListPath := path.Join(zclRoot, "/src/app/zap_cluster_list.json")
	clusterListBytes, err := os.ReadFile(clusterListPath)
	if err != nil {
		return err
	}

	o := orderedmap.New()
	err = json.Unmarshal(clusterListBytes, &o)
	if err != nil {
		return err
	}
	val, ok := o.Get("ClientDirectories")
	if !ok {
		return fmt.Errorf("no ClientDirectories field in zap_cluster_list.json")
	}
	is, ok := val.(orderedmap.OrderedMap)
	if !ok {
		return fmt.Errorf("ClientDirectories not a map in zap_cluster_list.json; %T", val)
	}
	for _, doc := range docs {
		path := doc.Path
		name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)) + " Cluster"
		name = strcase.ToScreamingSnake(name)
		if _, ok := is.Get(name); !ok {
			is.Set(name, []string{})
		}
	}
	is.SortKeys(func(keys []string) {
		slices.Sort(keys)
	})
	o.Set("ClientDirectories", is)

	clusterListBytes, err = json.MarshalIndent(o, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(clusterListPath, []byte(clusterListBytes), os.ModeAppend|0644)
}

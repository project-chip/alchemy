package generate

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/hasty/alchemy/ascii"
	"github.com/iancoleman/orderedmap"
	"github.com/iancoleman/strcase"
)

func patchClusterList(sdkRoot string, docs []*ascii.Doc) error {
	clusterListPath := path.Join(sdkRoot, "/src/app/zap_cluster_list.json")
	clusterListBytes, err := os.ReadFile(clusterListPath)
	if err != nil {
		return err
	}

	o := orderedmap.New()
	err = json.Unmarshal(clusterListBytes, &o)
	if err != nil {
		return err
	}

	var names []string
	for _, doc := range docs {
		path := doc.Path
		name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)) + " Cluster"
		name = strcase.ToScreamingSnake(name)
		names = append(names, name)
	}

	err = insertClusterName(o, "ClientDirectories", names)
	if err != nil {
		return err
	}

	err = insertClusterName(o, "ServerDirectories", names)
	if err != nil {
		return err
	}

	clusterListBytes, err = json.MarshalIndent(o, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(clusterListPath, []byte(clusterListBytes), os.ModeAppend|0644)
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

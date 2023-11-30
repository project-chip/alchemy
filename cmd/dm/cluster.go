package dm

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/matter"
)

type Result struct {
	XML    string
	Path   string
	Doc    *ascii.Doc
	Models []interface{}
}

func renderAppClusters(cxt context.Context, zclRoot string, appClusters []*ascii.Doc, filesOptions files.Options) error {
	var lock sync.Mutex
	outputs := make(map[string]string)
	err := files.ProcessDocs(cxt, appClusters, func(cxt context.Context, doc *ascii.Doc, index, total int) error {
		slog.Info("App cluster doc", "name", doc.Path)

		models, err := doc.ToModel()
		if err != nil {
			return err
		}
		var clusters []*matter.Cluster
		for _, m := range models {
			slog.Debug("model", "type", m)
			switch m := m.(type) {
			case *matter.Cluster:
				clusters = append(clusters, m)
			}
		}
		s, err := renderAppCluster(cxt, clusters)
		if err != nil {
			return fmt.Errorf("failed rendering %s: %w", doc.Path, err)
		}
		lock.Lock()
		outputs[doc.Path] = s
		lock.Unlock()
		return nil
	}, filesOptions)

	if err != nil {
		return err
	}

	if !filesOptions.DryRun {
		for path, result := range outputs {
			path := filepath.Base(path)
			newPath := filepath.Join(zclRoot, fmt.Sprintf("/data_model/clusters/%s.xml", strings.TrimSuffix(path, filepath.Ext(path))))
			err = os.WriteFile(newPath, []byte(result), os.ModeAppend|0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func renderAppCluster(cxt context.Context, clusters []*matter.Cluster) (output string, err error) {
	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(license)
	for _, cluster := range clusters {
		c := x.CreateElement("cluster")
		c.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
		c.CreateAttr("xsi:schemaLocation", "types types.xsd cluster cluster.xsd")
		c.CreateAttr("id", cluster.ID.HexString())
		c.CreateAttr("name", cluster.Name)

		revs := c.CreateElement("revisionHistory")
		var latestRev uint64 = 0
		for _, r := range cluster.Revisions {
			id := matter.ParseID(r.Number)
			if id.Valid() {
				rev := revs.CreateElement("revision")
				rev.CreateAttr("revision", id.IntString())
				rev.CreateAttr("summary", r.Description)
				latestRev = max(id.Value(), latestRev)
			}
		}
		c.CreateAttr("revision", strconv.FormatUint(latestRev, 10))
		class := c.CreateElement("classification")
		class.CreateAttr("hierarchy", strings.ToLower(cluster.Hierarchy))
		class.CreateAttr("role", strings.ToLower(cluster.Role))
		class.CreateAttr("picsCode", cluster.PICS)
		class.CreateAttr("scope", cluster.Scope)

		err = renderFeatures(cluster, c)
		if err != nil {
			return
		}
		err = renderDataTypes(cluster, c)
		if err != nil {
			return
		}
		err = renderAttributes(cluster, c)
		if err != nil {
			return
		}
		err = renderCommands(cluster, c)
		if err != nil {
			return
		}
		err = renderEvents(cluster, c)
		if err != nil {
			return
		}

	}
	x.Indent(2)

	var b bytes.Buffer
	x.WriteTo(&b)
	output = b.String()
	return
}

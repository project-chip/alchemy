package dm

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func getAppClusterPath(dmRoot string, path spec.Path, clusterName string) string {
	p := path.Base()
	file := strings.TrimSuffix(p, path.Ext())
	if len(clusterName) > 0 {
		file += "-" + clusterName
	}
	return filepath.Join(dmRoot, fmt.Sprintf("/clusters/%s.xml", file))
}

func (p *Renderer) renderAppCluster(doc *spec.Doc, clusters ...*matter.Cluster) (output string, err error) {
	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(getLicense())

	root := &x.Element

	cluster := clusters[0]

	c := root.CreateElement("cluster")
	c.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	c.CreateAttr("xsi:schemaLocation", "types types.xsd cluster cluster.xsd")
	if cluster.ID.Valid() {
		c.CreateAttr("id", cluster.ID.HexString())
	}
	name := cluster.Name
	if !strings.HasSuffix(name, " Cluster") {
		name += " Cluster"
	}
	c.CreateAttr("name", name)

	revs := c.CreateElement("revisionHistory")
	var latestRev uint64 = 0
	for _, r := range cluster.Revisions {
		id := matter.ParseNumber(r.Number)
		if id.Valid() {
			rev := revs.CreateElement("revision")
			rev.CreateAttr("revision", id.IntString())
			if len(r.Description) > 0 {
				rev.CreateAttr("summary", r.Description)
			}
			latestRev = max(id.Value(), latestRev)
		}
	}
	ids := c.CreateElement("clusterIds")
	for _, cluster := range clusters {
		clusterID := ids.CreateElement("clusterId")
		if cluster.ID.Valid() {
			clusterID.CreateAttr("id", cluster.ID.HexString())
		}
		clusterID.CreateAttr("name", cluster.Name)
		err = renderConformanceString(doc, cluster, cluster.Conformance, clusterID)
		if err != nil {
			return
		}
	}
	c.CreateAttr("revision", strconv.FormatUint(latestRev, 10))
	class := c.CreateElement("classification")
	switch cluster.Hierarchy {
	case "Base":
		class.CreateAttr("hierarchy", strings.ToLower(cluster.Hierarchy))
	default:
		class.CreateAttr("hierarchy", "derived")
		class.CreateAttr("baseCluster", cluster.Hierarchy)
	}
	class.CreateAttr("role", strings.ToLower(cluster.Role))
	class.CreateAttr("picsCode", cluster.PICS)
	class.CreateAttr("scope", cluster.Scope)

	err = renderFeatures(doc, cluster, c)
	if err != nil {
		return
	}
	err = renderDataTypes(doc, cluster, c)
	if err != nil {
		return
	}
	err = renderAttributes(doc, cluster, c)
	if err != nil {
		return
	}
	err = renderCommands(doc, cluster, c)
	if err != nil {
		return
	}
	err = renderEvents(doc, cluster, c)
	if err != nil {
		return
	}

	x.Indent(2)

	var b bytes.Buffer
	_, err = x.WriteTo(&b)
	output = b.String()

	p.clustersLock.Lock()
	p.clusters = append(p.clusters, clusters...)
	p.clustersLock.Unlock()
	return
}

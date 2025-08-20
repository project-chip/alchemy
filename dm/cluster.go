package dm

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func getAppClusterPath(dmRoot string, path asciidoc.Path, clusterName string) string {
	p := path.Base()
	file := strings.TrimSuffix(p, path.Ext())
	if len(clusterName) > 0 {
		file += "-" + clusterName
	}
	return filepath.Join(dmRoot, fmt.Sprintf("/clusters/%s.xml", file))
}

func (p *Renderer) renderAppCluster(doc *asciidoc.Document, entity types.Entity) (output string, clusterName string, err error) {
	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(getLicense())

	root := &x.Element

	var cluster *matter.Cluster
	var clusters []*matter.Cluster
	var clusterID *matter.Number
	var clusterClassification *matter.ClusterClassification
	var isClusterGroup bool
	switch entity := entity.(type) {
	case *matter.Cluster:
		cluster = entity
		clusterClassification = &cluster.ClusterClassification
		clusters = []*matter.Cluster{cluster}
		clusterName = cluster.Name
		if !strings.HasSuffix(clusterName, " Cluster") {
			clusterName += " Cluster"
		}
		clusterID = cluster.ID
	case *matter.ClusterGroup:
		cluster = entity.Clusters[0]
		clusterClassification = &entity.ClusterClassification
		clusters = entity.Clusters
		clusterName = entity.Name
		if !strings.HasSuffix(clusterName, " Clusters") {
			clusterName += " Clusters"
		}
		isClusterGroup = true
	default:
		err = fmt.Errorf("unexpected app cluster type: %T", entity)
		return
	}

	c := root.CreateElement("cluster")
	c.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	c.CreateAttr("xsi:schemaLocation", "types types.xsd cluster cluster.xsd")
	if clusterID.Valid() {
		c.CreateAttr("id", cluster.ID.HexString())
	}

	c.CreateAttr("name", clusterName)

	revs := c.CreateElement("revisionHistory")
	var latestRev uint64 = 0
	for _, r := range cluster.Revisions {
		id := matter.ParseNumber(r.Number)
		if id.Valid() {
			rev := revs.CreateElement("revision")
			rev.CreateAttr("revision", id.IntString())
			if len(r.Description) > 0 {
				rev.CreateAttr("summary", scrubDescription(r.Description))
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
		if isClusterGroup && len(cluster.PICS) > 0 {
			clusterID.CreateAttr("picsCode", cluster.PICS)
		}
		if cluster.Conformance != nil && !conformance.IsMandatory(cluster.Conformance) {
			err = renderConformanceElement(cluster.Conformance, clusterID, nil)
			if err != nil {
				return
			}
		}
	}
	c.CreateAttr("revision", strconv.FormatUint(latestRev, 10))
	class := c.CreateElement("classification")
	switch clusterClassification.Hierarchy {
	case "Base":
		class.CreateAttr("hierarchy", strings.ToLower(clusterClassification.Hierarchy))
	default:
		if clusterClassification.Hierarchy != "" {
			class.CreateAttr("hierarchy", "derived")
			class.CreateAttr("baseCluster", clusterClassification.Hierarchy)
		}
	}
	class.CreateAttr("role", strings.ToLower(clusterClassification.Role))
	class.CreateAttr("picsCode", clusterClassification.PICS)
	class.CreateAttr("scope", clusterClassification.Scope)
	renderQuality(class, clusterClassification.Quality)

	err = renderFeatures(doc, cluster, c)
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
	err = renderCommands(cluster.Commands, c)
	if err != nil {
		return
	}
	err = renderEvents(cluster.Events, c)
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

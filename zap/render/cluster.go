package render

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/find"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/zap"
)

func (cr *configuratorRenderer) renderClusters(ce *etree.Element) (err error) {
	configurator := cr.configurator
	for _, clusterElement := range ce.SelectElements("cluster") {
		var cluster *matter.Cluster
		var skip bool

		code, ok := xml.ReadSimpleElement(clusterElement, "code")
		if ok {
			clusterID := matter.ParseNumber(code)
			if !clusterID.Valid() {
				slog.Warn("invalid code ID in cluster, attempting to match by name", slog.String("path", configurator.OutPath), slog.String("code", clusterID.Text()))
			} else {
				for c, handled := range configurator.Clusters {
					if c.ID.Equals(clusterID) {
						cluster = c
						skip = handled
						configurator.Clusters[c] = true
					}
				}

				if cluster == nil {
					// We don't have this cluster in the spec; leave it here for now
					slog.Warn("unknown code ID in cluster", slog.String("path", configurator.OutPath), slog.String("clusterId", clusterID.Text()))
					continue
				}
			}

		}
		if cluster == nil {
			name, ok := xml.ReadSimpleElement(clusterElement, "name")
			if !ok {
				slog.Warn("invalid code ID in cluster and no name backup", slog.String("path", configurator.OutPath))
				continue
			}
			for c, handled := range configurator.Clusters {
				if c.Name == name {
					cluster = c
					skip = handled
					configurator.Clusters[c] = true
				}
			}
			if cluster == nil {
				// We don't have this cluster in the spec; leave it here for now
				slog.Warn("unknown name in cluster", slog.String("path", configurator.OutPath), slog.String("name", name))
				continue
			} else {
				slog.Warn("matched cluster by name; please allocate an ID for this cluster", slog.String("clusterName", cluster.Name))
			}
		}

		if skip {
			continue
		}

		err = cr.populateCluster(clusterElement, cluster)
		if err != nil {
			return
		}
	}

	var remainingClusters []*matter.Cluster
	for cluster, handled := range configurator.Clusters {
		if handled {
			continue
		}
		remainingClusters = append(remainingClusters, cluster)
	}

	slices.SortStableFunc(remainingClusters, func(a, b *matter.Cluster) int {
		if a.ID.Valid() && b.ID.Valid() {
			cmp := a.ID.Compare(b.ID)
			if cmp != 0 {
				return cmp
			}
		}
		return strings.Compare(a.Name, b.Name)
	})

	for _, cluster := range remainingClusters {
		if cr.isProvisionalViolation(cluster) {
			err = fmt.Errorf("new cluster added without provisional conformance: %s", cluster.Name)
			return
		}
		cle := etree.NewElement("cluster")
		xml.AppendElement(ce, cle, "struct", "enum", "bitmap", "domain")
		err = cr.populateCluster(cle, cluster)
		if err != nil {
			return
		}
	}
	return
}

func (cr *configuratorRenderer) populateCluster(clusterElement *etree.Element, cluster *matter.Cluster) (err error) {
	cr.elementMap[clusterElement] = cluster
	var define string
	var clusterPrefix string

	define = getDefine(cluster.Name+" Cluster", "", cr.configurator.Errata)
	if len(cr.configurator.Errata.ClusterDefinePrefix) > 0 {
		clusterPrefix = cr.configurator.Errata.ClusterDefinePrefix
	}

	cr.setProvisional(clusterElement, cluster)

	attributes := find.ToMap(cluster.Attributes)
	events := find.ToMap(slices.Collect(find.Filter(cluster.Events, func(e *matter.Event) bool {
		return !conformance.IsZigbee(e.Conformance) && !zap.IsDisallowed(e, e.Conformance)
	})))
	commands := find.ToMap(slices.Collect(find.Filter(cluster.Commands, func(c *matter.Command) bool {
		return !conformance.IsZigbee(c.Conformance) && !zap.IsDisallowed(c, c.Conformance)
	})))

	xml.SetOrCreateSimpleElement(clusterElement, "domain", cluster.Domain)
	clusterName := cluster.Name
	if cr.configurator.Errata.ClusterName != "" {
		clusterName = cr.configurator.Errata.ClusterName
	}
	xml.SetOrCreateSimpleElement(clusterElement, "name", clusterName, "domain")
	patchNumberElement(xml.SetOrCreateSimpleElement(clusterElement, "code", "", "name", "domain"), cluster.ID)
	xml.CreateSimpleElementIfNotExists(clusterElement, "define", define, "code", "name", "domain")

	descriptionElement := clusterElement.SelectElement("description")
	if descriptionElement == nil || descriptionElement.Text() == "" {
		xml.SetOrCreateSimpleElement(clusterElement, "description", cluster.Description, "define", "code", "name", "domain")
	}

	if client := clusterElement.SelectElement("client"); client == nil {
		client = xml.SetOrCreateSimpleElement(clusterElement, "client", "true", "description", "define", "code", "name", "domain")
		client.CreateAttr("init", "false")
		client.CreateAttr("tick", "false")
		client.SetText("true")
	}
	if server := clusterElement.SelectElement("server"); server == nil {
		server = xml.SetOrCreateSimpleElement(clusterElement, "server", "true", "client", "description", "define", "code", "name", "domain")
		server.CreateAttr("init", "false")
		server.CreateAttr("tick", "false")
		server.SetText("true")
	}
	if cr.generator.options.FeatureXML {
		err = cr.generateFeaturesXML(clusterElement, cluster)
		if err != nil {
			return
		}
	}
	err = generateClusterGlobalAttributes(cr.configurator, clusterElement, cluster)
	if err != nil {
		return
	}
	err = cr.generateAttributes(clusterElement, cluster, attributes, clusterPrefix)
	if err != nil {
		return
	}
	err = cr.generateCommands(commands, clusterElement, cluster)
	if err != nil {
		return
	}
	err = cr.generateEvents(clusterElement, cluster, events)
	if err != nil {
		return
	}
	return
}

func generateClusterGlobalAttributes(configurator *zap.Configurator, cle *etree.Element, cluster *matter.Cluster) (err error) {
	globalAttributes := cle.SelectElements("globalAttribute")
	var setClusterRevision bool
	for _, globalAttribute := range globalAttributes {
		code := globalAttribute.SelectAttr("code")
		if code == nil {
			slog.Warn("globalAttribute element with no code attribute", slog.String("path", configurator.OutPath))
			continue
		}
		id := matter.ParseNumber(code.Value)
		if !id.Valid() {
			slog.Warn("globalAttribute element with invalid code attribute", slog.String("path", configurator.OutPath), slog.String("code", code.Value))
			continue
		}
		setClusterGlobalAttribute(cle, globalAttribute, cluster, id)
		if id.Value() == 0xFFFD {
			setClusterRevision = true
		}
	}
	if !setClusterRevision {
		globalAttribute := etree.NewElement("globalAttribute")
		id := matter.NewNumber(0xFFFD)
		globalAttribute.CreateAttr("code", id.HexString())
		setClusterGlobalAttribute(cle, globalAttribute, cluster, id)
		xml.AppendElement(cle, globalAttribute, "server", "client", "description", "define")
	}
	return
}

func setClusterGlobalAttribute(parent *etree.Element, globalAttribute *etree.Element, cluster *matter.Cluster, id *matter.Number) {
	switch id.Value() {
	case 0xFFFD:
		globalAttribute.CreateAttr("side", "either")
		mostRecentRevision := cluster.Revisions.MostRecent()
		if mostRecentRevision != nil {
			globalAttribute.CreateAttr("value", mostRecentRevision.Number.IntString())
		}
	case 0xFFFC:
		slog.Warn("Removing redundant feature global attribute", slog.String("clusterName", cluster.Name))
		parent.RemoveChild(globalAttribute)
	default:
		slog.Warn("Unrecognized global attribute", slog.String("clusterId", id.HexString()))
	}
}

func (tg *TemplateGenerator) buildClusterAliases(configurator *zap.Configurator) {
	for c := range configurator.Clusters {
		if c != nil {

			if len(configurator.Errata.ClusterAliases) > 0 {
				if aliases, ok := configurator.Errata.ClusterAliases[c.Name]; ok {
					tg.ClusterAliases.Store(c.Name, aliases)
				}
			}
		}
	}
}

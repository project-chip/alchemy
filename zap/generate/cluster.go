package generate

import (
	"log/slog"
	"strconv"

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
				slog.Warn("invalid code ID in cluster", slog.String("path", configurator.OutPath), slog.String("id", clusterID.Text()))
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
					slog.Warn("unknown code ID in cluster", slog.String("path", configurator.OutPath), slog.String("id", clusterID.Text()))
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

	for cluster, handled := range configurator.Clusters {
		if handled {
			continue
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

	var define string
	var clusterPrefix string

	define = getDefine(cluster.Name+" Cluster", "", cr.configurator.Errata)
	if len(cr.configurator.Errata.ClusterDefinePrefix) > 0 {
		clusterPrefix = cr.configurator.Errata.ClusterDefinePrefix
	}

	if cluster.Conformance != nil {
		if conformance.IsProvisional(cluster.Conformance) {
			clusterElement.CreateAttr("apiMaturity", "provisional")
		} else {
			clusterElement.RemoveAttr("apiMaturity")
		}
	}

	attributes := find.ToMap(cluster.Attributes)
	events := find.ToMap(cluster.Events)
	commands := find.ToMap(find.ToList(find.Filter(cluster.Commands, func(c *matter.Command) bool {
		return !conformance.IsZigbee(cluster.Commands, c.Conformance) && !conformance.IsDisallowed(c.Conformance)
	})))

	xml.SetOrCreateSimpleElement(clusterElement, "domain", cr.configurator.Domain)
	clusterName := cluster.Name
	if cr.configurator.Errata.ClusterName != "" {
		clusterName = cr.configurator.Errata.ClusterName
	}
	xml.SetOrCreateSimpleElement(clusterElement, "name", clusterName, "domain")
	patchNumberElement(xml.SetOrCreateSimpleElement(clusterElement, "code", "", "name", "domain"), cluster.ID)
	xml.CreateSimpleElementIfNotExists(clusterElement, "define", define, "code", "name", "domain")

	if clusterElement.SelectElement("description") == nil {
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
	if cr.generator.generateFeaturesXML {
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
		setClusterGlobalAttribute(globalAttribute, cluster, id)
		if id.Value() == 0xFFFD {
			setClusterRevision = true
		}
	}
	if !setClusterRevision {
		globalAttribute := etree.NewElement("globalAttribute")
		id := matter.NewNumber(0xFFFD)
		globalAttribute.CreateAttr("code", id.HexString())
		setClusterGlobalAttribute(globalAttribute, cluster, id)
		xml.AppendElement(cle, globalAttribute, "server", "client", "description", "define")
	}
	return
}

func setClusterGlobalAttribute(globalAttribute *etree.Element, cluster *matter.Cluster, id *matter.Number) {
	switch id.Value() {
	case 0xFFFD:
		var lastRevision uint64
		for _, rev := range cluster.Revisions {
			revNumber := matter.ParseNumber(rev.Number)
			if revNumber.Valid() && revNumber.Value() > lastRevision {
				lastRevision = revNumber.Value()
			}
		}
		globalAttribute.CreateAttr("side", "either")
		globalAttribute.CreateAttr("value", strconv.FormatUint(lastRevision, 10))
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

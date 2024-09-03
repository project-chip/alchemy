package generate

import (
	"log/slog"
	"strconv"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/zap"
)

func (tg *TemplateGenerator) renderClusters(configurator *zap.Configurator, ce *etree.Element, errata *errata.ZAP) (err error) {

	for _, cle := range ce.SelectElements("cluster") {
		var cluster *matter.Cluster
		var skip bool

		code, ok := xml.ReadSimpleElement(cle, "code")
		if ok {
			clusterID := matter.ParseNumber(code)
			if !clusterID.Valid() {
				slog.Warn("invalid code ID in cluster", log.Path("path", configurator.Doc.Path), slog.String("id", clusterID.Text()))
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
					slog.Warn("unknown code ID in cluster", log.Path("path", configurator.Doc.Path), slog.String("id", clusterID.Text()))
					continue
				}
			}

		}
		if cluster == nil {
			name, ok := xml.ReadSimpleElement(cle, "name")
			if !ok {
				slog.Warn("invalid code ID in cluster and no name backup", log.Path("path", configurator.Doc.Path))
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
				slog.Warn("unknown name in cluster", log.Path("path", configurator.Doc.Path), slog.String("name", name))
				continue
			}
		}

		if skip {
			continue
		}

		err = tg.populateCluster(configurator, cle, cluster, errata)
		if err != nil {
			return
		}
	}

	for cluster, handled := range configurator.Clusters {
		if handled {
			continue
		}
		if !cluster.ID.Valid() {
			continue
		}
		cle := etree.NewElement("cluster")
		xml.AppendElement(ce, cle, "struct", "enum", "bitmap", "domain")
		err = tg.populateCluster(configurator, cle, cluster, errata)
		if err != nil {
			return
		}
	}
	return
}

func (tg *TemplateGenerator) populateCluster(configurator *zap.Configurator, cle *etree.Element, cluster *matter.Cluster, errata *errata.ZAP) (err error) {

	var define string
	var clusterPrefix string

	define = getDefine(cluster.Name+" Cluster", "", errata)
	if len(errata.ClusterDefinePrefix) > 0 {
		clusterPrefix = errata.ClusterDefinePrefix
	}

	if cluster.Conformance != nil {
		if conformance.IsProvisional(cluster.Conformance) {
			cle.CreateAttr("apiMaturity", "provisional")
		} else {
			cle.RemoveAttr("apiMaturity")
		}
	}

	attributes := make(map[*matter.Field]struct{})
	events := make(map[*matter.Event]struct{})
	commands := make(map[*matter.Command][]*matter.Number)

	for _, a := range cluster.Attributes {
		attributes[a] = struct{}{}
	}

	for _, e := range cluster.Events {
		events[e] = struct{}{}
	}

	for _, c := range cluster.Commands {
		if conformance.IsZigbee(cluster.Commands, c.Conformance) || conformance.IsDisallowed(c.Conformance) {
			continue
		}
		commands[c] = []*matter.Number{}
	}

	xml.SetOrCreateSimpleElement(cle, "domain", matter.DomainNames[configurator.Doc.Domain])
	xml.SetOrCreateSimpleElement(cle, "name", cluster.Name, "domain")
	patchNumberElement(xml.SetOrCreateSimpleElement(cle, "code", "", "name", "domain"), cluster.ID)
	xml.CreateSimpleElementIfNotExists(cle, "define", define, "code", "name", "domain")

	if cle.SelectElement("description") == nil {
		xml.SetOrCreateSimpleElement(cle, "description", cluster.Description, "define", "code", "name", "domain")
	}

	if client := cle.SelectElement("client"); client == nil {
		client = xml.SetOrCreateSimpleElement(cle, "client", "true", "description", "define", "code", "name", "domain")
		client.CreateAttr("init", "false")
		client.CreateAttr("tick", "false")
		client.SetText("true")
	}
	if server := cle.SelectElement("server"); server == nil {
		server = xml.SetOrCreateSimpleElement(cle, "server", "true", "client", "description", "define", "code", "name", "domain")
		server.CreateAttr("init", "false")
		server.CreateAttr("tick", "false")
		server.SetText("true")
	}
	if tg.generateFeaturesXML {
		err = generateFeaturesXML(configurator, cle, cluster)
		if err != nil {
			return
		}
	}
	err = generateClusterGlobalAttributes(configurator, cle, cluster)
	if err != nil {
		return
	}
	err = generateAttributes(configurator, cle, cluster, attributes, clusterPrefix, errata)
	if err != nil {
		return
	}
	err = generateCommands(commands, configurator.Doc.Path.Relative, cle, errata)
	if err != nil {
		return
	}
	err = generateEvents(configurator, cle, cluster, events, errata)
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
			slog.Warn("globalAttribute element with no code attribute", log.Path("path", configurator.Doc.Path))
			continue
		}
		id := matter.ParseNumber(code.Value)
		if !id.Valid() {
			slog.Warn("globalAttribute element with invalid code attribute", log.Path("path", configurator.Doc.Path), slog.String("code", code.Value))
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

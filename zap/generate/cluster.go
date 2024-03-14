package generate

import (
	"log/slog"
	"strconv"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func renderClusters(configurator *zap.Configurator, ce *etree.Element, errata *zap.Errata) (err error) {

	for _, cle := range ce.SelectElements("cluster") {
		code, ok := readSimpleElement(cle, "code")
		if !ok {
			slog.Warn("missing code element in cluster", slog.String("path", configurator.Doc.Path))
			continue
		}
		clusterId := matter.ParseNumber(code)
		if !clusterId.Valid() {
			slog.Warn("invalid code ID in cluster", slog.String("path", configurator.Doc.Path), slog.String("id", clusterId.Text()))
			continue
		}

		var cluster *matter.Cluster
		var skip bool
		for c, handled := range configurator.Clusters {
			if c.ID.Equals(clusterId) {
				cluster = c
				skip = handled
				configurator.Clusters[c] = true
			}
		}

		if skip {
			continue
		}

		if cluster == nil {
			// We don't have this cluster in the spec; leave it here for now
			slog.Warn("unknown code ID in cluster", slog.String("path", configurator.Doc.Path), slog.String("id", clusterId.Text()))
			continue
		}
		err = populateCluster(configurator, cle, cluster, errata)
		if err != nil {
			return
		}
	}

	for cluster, handled := range configurator.Clusters {
		if handled {
			continue
		}
		cle := etree.NewElement("cluster")
		cle.CreateAttr("code", cluster.ID.HexString())
		appendElement(ce, cle, "struct", "enum", "bitmap", "domain")
		err = populateCluster(configurator, cle, cluster, errata)
		if err != nil {
			return
		}
	}
	return
}

func populateCluster(configurator *zap.Configurator, cle *etree.Element, cluster *matter.Cluster, errata *zap.Errata) (err error) {

	var define string
	var clusterPrefix string

	define = getDefine(cluster.Name+" Cluster", "", errata)
	if len(errata.ClusterDefinePrefix) > 0 {
		clusterPrefix = errata.ClusterDefinePrefix
	}

	attributes := make(map[*matter.Field]struct{})
	events := make(map[*matter.Event]struct{})
	commands := make(map[*matter.Command]struct{})

	for _, a := range cluster.Attributes {
		attributes[a] = struct{}{}
	}

	for _, e := range cluster.Events {
		events[e] = struct{}{}
	}

	for _, c := range cluster.Commands {
		commands[c] = struct{}{}
	}

	setOrCreateSimpleElement(cle, "domain", matter.DomainNames[configurator.Doc.Domain])
	setOrCreateSimpleElement(cle, "name", cluster.Name, "domain")
	patchNumberElement(setOrCreateSimpleElement(cle, "code", "", "name", "domain"), cluster.ID)
	setOrCreateSimpleElement(cle, "define", define, "code", "name", "domain")

	if cle.SelectElement("description") == nil {
		setOrCreateSimpleElement(cle, "description", cluster.Description, "define", "code", "name", "domain")
	}

	if client := cle.SelectElement("client"); client == nil {
		client = setOrCreateSimpleElement(cle, "client", "true", "description", "define", "code", "name", "domain")
		client.CreateAttr("init", "false")
		client.CreateAttr("tick", "false")
		client.SetText("true")
	}
	if server := cle.SelectElement("server"); server == nil {
		server = setOrCreateSimpleElement(cle, "server", "true", "client", "description", "define", "code", "name", "domain")
		server.CreateAttr("init", "false")
		server.CreateAttr("tick", "false")
		server.SetText("true")
	}
	err = generateClusterGlobalAttributes(configurator, cle, cluster, errata)
	if err != nil {
		return
	}
	err = generateAttributes(configurator, cle, cluster, attributes, clusterPrefix, errata)
	if err != nil {
		return
	}
	err = generateCommands(configurator, cle, cluster, commands, errata)
	if err != nil {
		return
	}
	err = generateEvents(configurator, cle, cluster, events, errata)
	if err != nil {
		return
	}
	return
}

func generateClusterGlobalAttributes(configurator *zap.Configurator, cle *etree.Element, cluster *matter.Cluster, errata *zap.Errata) (err error) {
	globalAttributes := cle.SelectElements("globalAttribute")
	var setClusterRevision bool
	for _, globalAttribute := range globalAttributes {
		code := globalAttribute.SelectAttr("code")
		if code == nil {
			slog.Warn("globalAttribute element with no code attribute", slog.String("path", configurator.Doc.Path))
			continue
		}
		id := matter.ParseNumber(code.Value)
		if !id.Valid() {
			slog.Warn("globalAttribute element with invalid code attribute", slog.String("path", configurator.Doc.Path), slog.String("code", code.Value))
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
		appendElement(cle, globalAttribute, "server", "client", "description", "define")
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

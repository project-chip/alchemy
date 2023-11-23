package amend

import (
	"encoding/xml"
	"io"

	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/iancoleman/strcase"
)

func (r *renderer) amendCluster(d xmlDecoder, e xmlEncoder, el xml.StartElement, clusters map[*matter.Cluster]struct{}) (err error) {
	var clusterTokens []xml.Token
	clusterTokens, err = Extract(d, el)
	if err != nil {
		return
	}
	var clusterID *matter.ID
	for i, tok := range clusterTokens {
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "code":
				var cid string
				cid, err = getSimpleElement(clusterTokens[i+1:], "code")
				if err != nil {
					return
				}
				if len(cid) > 0 {
					clusterID = matter.ParseID(cid)
					break
				}
			}
		}
	}

	if clusterID == nil {
		// Can't find cluster ID, so dump
		err = writeTokens(e, clusterTokens)
		return
	}

	var cluster *matter.Cluster
	for c := range clusters {
		if c.ID.Equals(clusterID) {
			cluster = c
			delete(clusters, c)
		}
	}

	if cluster == nil {
		// We don't have this cluster in the spec; leave it here for now

		err = writeTokens(e, clusterTokens)
		return
	}

	var define string
	var clusterPrefix string

	define = strcase.ToScreamingDelimited(cluster.Name+" Cluster", '_', "", true)
	if !r.errata.SuppressClusterDefinePrefix {
		clusterPrefix = strcase.ToScreamingDelimited(cluster.Name, '_', "", true) + "_"
		if len(r.errata.ClusterDefinePrefix) > 0 {
			clusterPrefix = r.errata.ClusterDefinePrefix
		}
	}

	clusterValues := map[string]string{
		"name":        cluster.Name,
		"description": cluster.Description,
		"define":      define,
		"domain":      matter.DomainNames[r.doc.Domain],
		"code":        cluster.ID.HexString(),
	}
	clusterValuesWritten := make(map[string]bool)

	attributes := make(map[*matter.Field]struct{})
	events := make(map[*matter.Event]struct{})
	commands := make(map[*matter.Command]struct{})

	for _, a := range cluster.Attributes {
		if conformance.IsZigbee(a.Conformance) {
			continue
		}
		attributes[a] = struct{}{}
	}

	for _, e := range cluster.Events {

		events[e] = struct{}{}
	}

	for _, c := range cluster.Commands {
		commands[c] = struct{}{}
	}

	var lastSection = matter.SectionUnknown
	var nextCharData string
	var hasCharDataPending bool
	var hasCommentCharDataPending bool
	var lastIgnoredCharData xml.CharData

	ts := &tokenSet{tokens: clusterTokens}

	for {
		var tok xml.Token
		tok, err = ts.Token()
		if tok == nil || err == io.EOF {
			err = io.EOF
			return
		} else if err != nil {
			return
		}

		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "attribute":
				if lastSection != matter.SectionAttribute {
					err = r.flushUnusedClusterElements(cluster, e, lastSection, clusterValues, attributes, events, commands, clusterPrefix)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionAttribute
				err = r.amendAttribute(cluster, ts, e, t, attributes, clusterPrefix)
			case "command":
				if lastSection != matter.SectionCommand {
					err = r.flushUnusedClusterElements(cluster, e, lastSection, clusterValues, attributes, events, commands, clusterPrefix)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionCommand
				err = r.amendCommand(ts, e, t, commands)
			case "event":
				if lastSection != matter.SectionEvent {
					err = r.flushUnusedClusterElements(cluster, e, lastSection, clusterValues, attributes, events, commands, clusterPrefix)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionEvent
				err = r.amendEvent(ts, e, t, events)
			case "client", "server":
				e.EncodeToken(t)
				hasCommentCharDataPending = true
			default:
				if lastSection != matter.SectionCluster {
					err = r.flushUnusedClusterElements(cluster, e, lastSection, clusterValues, attributes, events, commands, clusterPrefix)
					if err != nil {
						return
					}
				}
				lastSection = matter.SectionCluster
				val, ok := clusterValues[t.Name.Local]
				if ok {
					nextCharData = val
					hasCharDataPending = true
					delete(clusterValues, t.Name.Local)
				}
				if !clusterValuesWritten[t.Name.Local] {
					err = e.EncodeToken(tok)
				}
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "server", "client":
				hasCommentCharDataPending = false
				err = e.EncodeToken(t)
			case "attribute", "command", "event":
			case "cluster":
				err = r.flushClusterValues(e, clusterValues)
				if err != nil {
					return
				}
				err = r.flushAttributes(cluster, e, attributes, clusterPrefix)
				if err != nil {
					return
				}
				err = r.flushCommands(e, commands, clusterPrefix)
				if err != nil {
					return
				}
				err = r.flushEvents(e, events, clusterPrefix)
				if err != nil {
					return
				}
				err = e.EncodeToken(t)
				return
			default:
				if !clusterValuesWritten[t.Name.Local] {
					err = e.EncodeToken(tok)
					clusterValuesWritten[t.Name.Local] = true
				}
			}
		case xml.Comment:
			if lastIgnoredCharData != nil {
				err = e.EncodeToken(lastIgnoredCharData)
				lastIgnoredCharData = nil
			}
			err = e.EncodeToken(t)
			hasCommentCharDataPending = true
		case xml.CharData:
			if hasCharDataPending {
				err = e.EncodeToken(xml.CharData(nextCharData))
				hasCharDataPending = false
				lastIgnoredCharData = nil
			} else if hasCommentCharDataPending {
				err = e.EncodeToken(t)
				hasCommentCharDataPending = false
				lastIgnoredCharData = nil
			} else {
				lastIgnoredCharData = t
			}
		default:
			err = e.EncodeToken(t)
			lastIgnoredCharData = nil
		}
		if err != nil {
			return
		}
	}
}

func (r *renderer) flushUnusedClusterElements(cluster *matter.Cluster, e xmlEncoder, lastSection matter.Section, clusterValues map[string]string, attributes map[*matter.Field]struct{}, events map[*matter.Event]struct{}, commands map[*matter.Command]struct{}, clusterPrefix string) (err error) {
	switch lastSection {
	case matter.SectionAttribute:
		err = r.flushAttributes(cluster, e, attributes, clusterPrefix)
	case matter.SectionCommand:
		err = r.flushCommands(e, commands, clusterPrefix)
	case matter.SectionEvent:
		err = r.flushEvents(e, events, clusterPrefix)
	}
	return
}

func (*renderer) flushClusterValues(e xmlEncoder, clusterValues map[string]string) (err error) {
	for k, v := range clusterValues {
		err = e.EncodeToken(xml.StartElement{Name: xml.Name{Local: k}})
		if err != nil {
			return
		}
		err = e.EncodeToken(xml.CharData(v))
		if err != nil {
			return
		}
		err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: k}})
		if err != nil {
			return
		}
		delete(clusterValues, k)
	}
	return
}

func (r *renderer) flushAttributes(cluster *matter.Cluster, e xmlEncoder, attributes map[*matter.Field]struct{}, clusterPrefix string) (err error) {
	for a := range attributes {
		if conformance.IsDeprecated(a.Conformance) {
			delete(attributes, a)
			continue
		}
		err = r.writeAttribute(cluster, e, xml.StartElement{Name: xml.Name{Local: "attribute"}}, a, clusterPrefix)
		delete(attributes, a)
		if err != nil {
			return
		}
	}
	return
}

func (r *renderer) flushCommands(e xmlEncoder, commands map[*matter.Command]struct{}, clusterPrefix string) (err error) {
	for c := range commands {
		err = r.writeCommand(e, xml.StartElement{Name: xml.Name{Local: "command"}}, c)
		delete(commands, c)
		if err != nil {
			return
		}
	}
	return
}

func (r *renderer) flushEvents(e xmlEncoder, events map[*matter.Event]struct{}, clusterPrefix string) (err error) {
	for ev := range events {
		err = r.writeEvent(e, xml.StartElement{Name: xml.Name{Local: "event"}}, ev, true)
		delete(events, ev)
		if err != nil {
			return
		}
	}
	return
}

func (*renderer) renderClusterCodes(e xmlEncoder, clusterIDs []string) (err error) {
	for _, clusterID := range clusterIDs {
		elName := xml.Name{Local: "cluster"}
		xcs := xml.StartElement{Name: elName, Attr: []xml.Attr{{Name: xml.Name{Local: "code"}, Value: clusterID}}}
		err = e.EncodeToken(xcs)
		if err != nil {
			return
		}
		xce := xml.EndElement{Name: elName}
		err = e.EncodeToken(xce)
		if err != nil {
			return
		}
	}
	return
}

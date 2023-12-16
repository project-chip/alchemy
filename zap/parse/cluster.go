package parse

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/hasty/alchemy/matter"
)

func readCluster(d *xml.Decoder, e xml.StartElement) (cluster *matter.Cluster, err error) {
	cluster = &matter.Cluster{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "singleton":
		case "apiMaturity":
		default:
			return nil, fmt.Errorf("unexpected cluster attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return nil, fmt.Errorf("EOF before end of cluster")
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "attribute":
				var attribute *matter.Field
				attribute, err = readAttribute(d, t)
				if err == nil {
					cluster.Attributes = append(cluster.Attributes, attribute)
				}
			case "client":
				err = readClient(d, t)
			case "server":
				err = readServer(d, t)
			case "code":
				var cid string
				cid, err = readSimpleElement(d, t.Name.Local)
				if err == nil {
					cluster.ID = matter.ParseNumber(cid)
				}
			case "define":
				_, err = readSimpleElement(d, t.Name.Local)
			case "description":
				cluster.Description, err = readSimpleElement(d, t.Name.Local)
			case "domain":
				_, err = readSimpleElement(d, t.Name.Local)
			case "event":
				var event *matter.Event
				event, err = readEvent(d, t)
				if err == nil {
					cluster.Events = append(cluster.Events, event)
				}
			case "name":
				cluster.Name, err = readSimpleElement(d, t.Name.Local)
			case "command":
				var command *matter.Command
				command, err = readCommand(d, t)
				if err == nil {
					cluster.Commands = append(cluster.Commands, command)
				}
			case "globalAttribute":
				err = Ignore(d, t.Name.Local)
			case "tag":
				_, err = readTag(d, t)
			default:
				err = fmt.Errorf("unexpected cluster level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "cluster":
				return
			default:
				err = fmt.Errorf("unexpected cluster end element: %s", t.Name.Local)
			}
		case xml.Comment:
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected cluster level type: %T", t)
		}
		if err != nil {
			err = fmt.Errorf("error parsing cluster: %w", err)
			return
		}
	}
}

func readClusterCode(d *xml.Decoder, e xml.StartElement) (id *matter.Number, err error) {
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "code":
			id = matter.ParseNumber(a.Value)
		default:
			return nil, fmt.Errorf("unexpected cluster code attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of cluster code")

		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "cluster":
				return
			default:
				err = fmt.Errorf("unexpected cluster code end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected cluster code level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

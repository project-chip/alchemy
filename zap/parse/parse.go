package parse

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"sync"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type ZapParser struct {
	lock sync.Mutex

	clusterReferences map[uint64]*matter.Cluster
	bitmapReferences  map[uint64][]*matter.Bitmap
	enumReferences    map[uint64][]*matter.Enum
	structReferences  map[uint64][]*matter.Struct
}

func NewZapParser() *ZapParser {
	return &ZapParser{
		clusterReferences: make(map[uint64]*matter.Cluster),
		bitmapReferences:  make(map[uint64][]*matter.Bitmap),
		enumReferences:    make(map[uint64][]*matter.Enum),
		structReferences:  make(map[uint64][]*matter.Struct),
	}
}

func (sp *ZapParser) Name() string {
	return "Parsing ZAP templates"
}

func (sp *ZapParser) Process(cxt context.Context, input *pipeline.Data[[]byte], index int32, total int32) (outputs []*pipeline.Data[[]types.Entity], extras []*pipeline.Data[[]byte], err error) {
	if filepath.Base(input.Path) == "matter-devices.xml" {
		return
	}
	d := xml.NewDecoder(bytes.NewReader(input.Content))
	var entities []types.Entity
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = nil
			outputs = append(outputs, pipeline.NewData[[]types.Entity](input.Path, entities))
			return
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.ProcInst:
		case xml.CharData:
		case xml.StartElement:
			switch t.Name.Local {
			case "configurator":
				var cm []types.Entity
				cm, err = sp.readConfigurator(input.Path, d)
				if err == nil {
					entities = append(entities, cm...)
					sp.lock.Lock()
					for _, e := range cm {
						switch e := e.(type) {
						case *matter.ClusterGroup:
							for _, c := range e.Clusters {
								sp.clusterReferences[c.ID.Value()] = c
							}
						case *matter.Cluster:
							sp.clusterReferences[e.ID.Value()] = e
						}
					}
					sp.lock.Unlock()
				}
			default:
				err = fmt.Errorf("unexpected top level element: %s", t.Name.Local)
			}
		case xml.Comment:
		default:
			err = fmt.Errorf("unexpected top level type: %T", t)
		}
		if err != nil {
			err = fmt.Errorf("error parsing %s: %w", input.Path, err)
			return
		}
	}
}

func (sp *ZapParser) ResolveReferences() {
	for cid, e := range sp.enumReferences {
		c, ok := sp.clusterReferences[cid]
		if !ok {
			slog.Warn("unknown cluster reference for enum", "clusterId", cid)
			continue
		}
		c.Enums = append(c.Enums, e...)
	}
	for cid, e := range sp.structReferences {
		c, ok := sp.clusterReferences[cid]
		if !ok {
			slog.Warn("unknown cluster reference for struct", "clusterId", cid)
			continue
		}
		c.Structs = append(c.Structs, e...)
	}
}

func Privilege(a string) matter.Privilege {
	switch a {
	case "view":
		return matter.PrivilegeView
	case "manage":
		return matter.PrivilegeManage
	case "administer":
		return matter.PrivilegeAdminister
	case "operate":
		return matter.PrivilegeOperate
	default:
		return matter.PrivilegeUnknown
	}
}

func readSimpleElement(d *xml.Decoder, name string) (val string, err error) {
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of %s", name)
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case name:
				return
			default:
				err = fmt.Errorf("unexpected %s end element: %s", name, t.Name.Local)
			}
		case xml.CharData:
			val = string(t)
		default:
			err = fmt.Errorf("unexpected %s level type: %T", name, t)
		}
		if err != nil {
			return
		}
	}
}

func Ignore(d *xml.Decoder, name string) (err error) {
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of %s", name)
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case name:
				return nil
			default:
			}
		default:
		}
	}
}

func Extract(d *xml.Decoder, el xml.StartElement) (tokens []xml.Token, err error) {
	tokens = append(tokens, el)
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of %s", el.Name.Local)
		}
		if err != nil {
			return
		}
		tokens = append(tokens, tok)
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case el.Name.Local:
				return
			default:
			}
		default:
		}
	}
}

func readTag(d *xml.Decoder, e xml.StartElement) (c *matter.Bitmap, err error) {
	c = &matter.Bitmap{Name: "Feature", Type: types.NewDataType(types.BaseDataTypeMap32, types.DataTypeRankScalar)}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			c.Name = a.Value
		case "description":
			c.Description = a.Value
		default:
			return nil, fmt.Errorf("unexpected tag attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "tag":
				return
			default:
				err = fmt.Errorf("unexpected tag end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected tag level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

package parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (sp *ZapParser) readConfigurator(path string, d *xml.Decoder) (entities []types.Entity, err error) {
	enums := make(map[uint64][]*matter.Enum)
	bitmaps := make(map[uint64][]*matter.Bitmap)
	structs := make(map[uint64][]*matter.Struct)
	clusters := make(map[uint64]*matter.Cluster)
	var features *matter.Features
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of configurator")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "cluster":
				var cluster *matter.Cluster
				cluster, err = readCluster(path, d, t)
				if err == nil {
					clusters[cluster.ID.Value()] = cluster
					entities = append(entities, cluster)
				}
			case "domain":
				_, err = readSimpleElement(d, t.Name.Local)
			case "enum":
				var en *matter.Enum
				var clusterIDs []*matter.Number
				en, clusterIDs, err = sp.readEnum(d, t)
				if err == nil {
					for _, cid := range clusterIDs {
						enums[cid.Value()] = append(enums[cid.Value()], en)
					}
				}
			case "struct":
				var s *matter.Struct
				var clusterIDs []*matter.Number
				s, clusterIDs, err = sp.readStruct(path, d, t)
				if err == nil {
					for _, cid := range clusterIDs {
						structs[cid.Value()] = append(structs[cid.Value()], s)
					}
				}
			case "bitmap":
				var bitmap *matter.Bitmap
				var clusterIDs []*matter.Number
				bitmap, clusterIDs, err = sp.readBitmap(d, t)
				if err == nil {
					if bitmap.Name == "Feature" {
						features = &matter.Features{Bitmap: *bitmap}
					} else {
						for _, cid := range clusterIDs {
							bitmaps[cid.Value()] = append(bitmaps[cid.Value()], bitmap)
						}
					}
				}
			case "accessControl", "atomic", "clusterExtension", "global", "deviceType":
				err = Ignore(d, t.Name.Local)
			default:
				err = fmt.Errorf("unexpected configurator level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "configurator":
				for cid, b := range bitmaps {
					if c, ok := clusters[cid]; ok {
						c.AddBitmaps(b...)
					} else {
						slog.Debug("orphan bitmaps", "clusterId", cid)
						sp.lock.Lock()
						sp.bitmapReferences[cid] = append(sp.bitmapReferences[cid], b...)
						sp.lock.Unlock()
					}
				}
				for cid, e := range enums {
					if c, ok := clusters[cid]; ok {
						c.AddEnums(e...)
					} else {
						slog.Debug("orphan enums", "clusterId", cid)
						sp.lock.Lock()
						sp.enumReferences[cid] = append(sp.enumReferences[cid], e...)
						sp.lock.Unlock()
					}
				}
				for cid, s := range structs {
					if c, ok := clusters[cid]; ok {
						c.AddStructs(s...)
					} else {
						slog.Debug("orphan structs", "clusterId", cid)
						sp.lock.Lock()
						sp.structReferences[cid] = append(sp.structReferences[cid], s...)
						sp.lock.Unlock()
					}
				}
				for _, c := range clusters {
					c.Features = features
				}
				return
			default:
				err = fmt.Errorf("unexpected configurator end element: %s", t.Name.Local)
			}
		case xml.CharData:
		case xml.Comment:
		default:
			err = fmt.Errorf("unexpected configurator level type: %T", t)
		}
		if err != nil {
			err = fmt.Errorf("error parsing configurator: %w", err)
			return
		}
	}
}

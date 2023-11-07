package parse

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/hasty/matterfmt/matter"
)

func readConfigurator(d *xml.Decoder) (models []any, err error) {
	enums := make(map[uint64][]*matter.Enum)
	bitmaps := make(map[uint64][]*matter.Bitmap)
	structs := make(map[uint64][]*matter.Struct)
	clusters := make(map[uint64]*matter.Cluster)
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
				cluster, err = readCluster(d, t)
				if err == nil {
					clusters[cluster.ID.Value()] = cluster
					models = append(models, cluster)
				}
			case "domain":
				_, err = readSimpleElement(d, t.Name.Local)
			case "enum":
				var en *matter.Enum
				var clusterIDs []*matter.ID
				en, clusterIDs, err = readEnum(d, t)
				if err == nil {
					for _, cid := range clusterIDs {
						enums[cid.Value()] = append(enums[cid.Value()], en)
					}
				}
			case "struct":
				var s *matter.Struct
				var clusterIDs []*matter.ID
				s, clusterIDs, err = readStruct(d, t)
				if err == nil {
					for _, cid := range clusterIDs {
						structs[cid.Value()] = append(structs[cid.Value()], s)
					}
				}
			case "bitmap":
				var bitmap *matter.Bitmap
				var clusterIDs []*matter.ID
				bitmap, clusterIDs, err = readBitmap(d, t)
				if err == nil {
					for _, cid := range clusterIDs {
						bitmaps[cid.Value()] = append(bitmaps[cid.Value()], bitmap)
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
						c.Bitmaps = append(c.Bitmaps, b...)
					} else {
						fmt.Printf("orphan bitmaps: %v\n", cid)
					}
				}
				for cid, e := range enums {
					if c, ok := clusters[cid]; ok {
						c.Enums = append(c.Enums, e...)
					} else {
						fmt.Printf("orphan enums: %v\n", cid)
					}
				}
				for cid, s := range structs {
					if c, ok := clusters[cid]; ok {
						c.Structs = append(c.Structs, s...)
					} else {
						fmt.Printf("orphan structs: %v\n", cid)
					}
				}

				return
			default:
				err = fmt.Errorf("unexpected configurator end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected configurator level type: %T", t)
		}
		if err != nil {
			err = fmt.Errorf("error parsing configurator: %w", err)
			return
		}
	}
}

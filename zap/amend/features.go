package amend

import (
	"encoding/xml"
	"fmt"

	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (r *renderer) writeFeatures(d xmlDecoder, e xmlEncoder, el xml.StartElement, cluster *matter.Cluster, clusterIDs []string) (err error) {
	Ignore(d, "bitmap")

	el = el.Copy()

	el.Attr = setAttributeValue(el.Attr, "type", "bitmap32")

	err = e.EncodeToken(el)
	if err != nil {
		return err
	}

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
	for _, f := range cluster.Features {
		if conformance.IsZigbee(f.Conformance) {
			continue
		}
		bit, er := parse.HexOrDec(f.Bit)
		if er != nil {
			continue
		}
		bit = (1 << bit)
		elName := xml.Name{Local: "field"}
		xfs := xml.StartElement{Name: elName}
		xfs.Attr = setAttributeValue(xfs.Attr, "name", f.Name)
		xfs.Attr = setAttributeValue(xfs.Attr, "mask", fmt.Sprintf("%#x", bit))
		err = e.EncodeToken(xfs)
		if err != nil {
			return
		}
		xfe := xml.EndElement{Name: elName}
		err = e.EncodeToken(xfe)
		if err != nil {
			return
		}
	}
	xfe := xml.EndElement{Name: xml.Name{Local: "bitmap"}}
	return e.EncodeToken(xfe)
}

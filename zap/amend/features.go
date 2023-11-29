package amend

import (
	"encoding/xml"
	"fmt"
	"io"
	"slices"

	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (r *renderer) amendFeatures(d xmlDecoder, e xmlEncoder, el xml.StartElement, cluster *matter.Cluster, clusterIDs []string) (err error) {

	el.Attr = setAttributeValue(el.Attr, "type", "bitmap32")

	err = e.EncodeToken(el)
	if err != nil {
		return
	}

	remainingClusterIDs := make([]string, len(clusterIDs))
	copy(remainingClusterIDs, clusterIDs)

	var featureIndex int

	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = io.EOF
			return
		} else if err != nil {
			return
		}

		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "description":
				writeThrough(d, e, t)
			case "cluster":
				code := getAttributeValue(t.Attr, "code")
				id := matter.ParseID(code)
				if id.Valid() {
					ids := id.HexString()
					remainingClusterIDs = slices.DeleteFunc(remainingClusterIDs, func(s string) bool {
						return ids == s
					})
				}
				writeThrough(d, e, t)
			case "field":
				if len(remainingClusterIDs) > 0 {
					err = r.renderClusterCodes(e, remainingClusterIDs)
					if err != nil {
						return
					}
				}
				for {
					if featureIndex >= len(cluster.Features) {
						Ignore(d, "field")
						break
					} else {
						f := cluster.Features[featureIndex]
						featureIndex++
						if conformance.IsZigbee(f.Conformance) {
							continue
						}
						t.Attr, err = r.setFeatureAttributes(t.Attr, f)
						if err != nil {
							err = fmt.Errorf("failed setting feature attributes on feature %s: %w", f.Name, err)
							return
						}
						writeThrough(d, e, t)
						break
					}
				}

			default:

			}
		case xml.EndElement:
			switch t.Name.Local {
			case "bitmap":
				if len(remainingClusterIDs) > 0 {
					err = r.renderClusterCodes(e, remainingClusterIDs)
					if err != nil {
						return
					}
				}
				for featureIndex < len(cluster.Features) {
					f := cluster.Features[featureIndex]
					featureIndex++
					if conformance.IsZigbee(f.Conformance) {
						continue
					}

					elName := xml.Name{Local: "field"}
					xfs := xml.StartElement{Name: elName}
					xfs.Attr, err = r.setFeatureAttributes(xfs.Attr, f)
					if err != nil {
						err = fmt.Errorf("failed setting bitmap attributes on bitmap %s: %w", f.Name, err)
						return
					}
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
				err = e.EncodeToken(t)
				return
			default:
				err = e.EncodeToken(tok)

			}
		case xml.CharData:
		default:
			err = e.EncodeToken(t)
		}
		if err != nil {
			return
		}
	}
}

func (r *renderer) writeFeatures(d xmlDecoder, e xmlEncoder, el xml.StartElement, cluster *matter.Cluster, clusterIDs []string) (err error) {

	el = el.Copy()

	el.Attr = setAttributeValue(el.Attr, "name", "Feature")
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
		elName := xml.Name{Local: "field"}
		xfs := xml.StartElement{Name: elName}
		xfs.Attr, err = r.setFeatureAttributes(xfs.Attr, f)
		if err != nil {
			return
		}
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

func (*renderer) setFeatureAttributes(xfs []xml.Attr, f *matter.Feature) ([]xml.Attr, error) {
	bit, err := parse.HexOrDec(f.Bit)
	if err != nil {
		return nil, err
	}
	bit = (1 << bit)
	xfs = setAttributeValue(xfs, "name", f.Name)
	xfs = setAttributeValue(xfs, "mask", fmt.Sprintf("%#x", bit))
	return xfs, nil
}

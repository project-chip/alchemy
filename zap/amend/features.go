package amend

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"slices"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

func (r *renderer) amendFeatures(d xmlDecoder, e xmlEncoder, el xml.StartElement, cluster *matter.Cluster, clusterIDs []string) (err error) {

	if cluster.Features == nil {
		err = Ignore(d, el.Name.Local)
		return
	}
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
				err = writeThrough(d, e, t)
			case "cluster":
				code := getAttributeValue(t.Attr, "code")
				id := matter.ParseNumber(code)
				if id.Valid() {
					ids := id.HexString()
					remainingClusterIDs = slices.DeleteFunc(remainingClusterIDs, func(s string) bool {
						return ids == s
					})
				}
				err = writeThrough(d, e, t)
			case "field":
				if len(remainingClusterIDs) > 0 {
					err = r.renderClusterCodes(e, remainingClusterIDs)
					if err != nil {
						return
					}
					remainingClusterIDs = nil
				}
				for {
					if featureIndex >= len(cluster.Features.Bits) {
						err = Ignore(d, t.Name.Local)
						break
					} else {
						f := cluster.Features.Bits[featureIndex]
						featureIndex++
						if conformance.IsZigbee(cluster.Features, f.Conformance) {
							continue
						}
						t.Attr, err = r.setFeatureAttributes(t.Attr, f)
						if err != nil {
							err = fmt.Errorf("failed setting feature attributes on feature %s: %w", f.Name, err)
							return
						}
						err = writeThrough(d, e, t)
						if err != nil {
							return
						}
						break
					}
				}

			default:
				slog.Warn("unexpected element in features", "name", t.Name.Local)
				err = Ignore(d, t.Name.Local)
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
				for featureIndex < len(cluster.Features.Bits) {
					f := cluster.Features.Bits[featureIndex]
					featureIndex++
					if conformance.IsZigbee(cluster.Features, f.Conformance) {
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
						err = fmt.Errorf("failed closing %s element on cluster %s: %w", elName, cluster.Name, err)
						return
					}
				}
				err = e.EncodeToken(t)
				if err != nil {
					return
				}
				err = newLine(e)
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
		return fmt.Errorf("error opening feature element on cluster %s: %w", cluster.Name, err)
	}

	for _, clusterID := range clusterIDs {
		elName := xml.Name{Local: "cluster"}
		xcs := xml.StartElement{Name: elName, Attr: []xml.Attr{{Name: xml.Name{Local: "code"}, Value: clusterID}}}
		err = e.EncodeToken(xcs)
		if err != nil {
			return fmt.Errorf("error opening feature cluster code element on cluster %s: %w", cluster.Name, err)
		}
		xce := xml.EndElement{Name: elName}
		err = e.EncodeToken(xce)
		if err != nil {
			return fmt.Errorf("error closing feature cluster element on cluster %s: %w", cluster.Name, err)
		}
	}
	for _, f := range cluster.Features.Bits {
		if conformance.IsZigbee(cluster.Features, f.Conformance) {
			continue
		}

		_, parseErr := parse.HexOrDec(f.Bit)
		if parseErr != nil {
			slog.Debug("skipping feature with non-parsable bit", "bit", f.Bit)
			continue
		}

		elName := xml.Name{Local: "field"}
		xfs := xml.StartElement{Name: elName}
		xfs.Attr, err = r.setFeatureAttributes(xfs.Attr, f)
		if err != nil {
			return fmt.Errorf("error setting feature field element attributes on cluster %s: %w", cluster.Name, err)
		}
		err = e.EncodeToken(xfs)
		if err != nil {
			return fmt.Errorf("error opening feature field element on cluster %s: %w", cluster.Name, err)
		}
		xfe := xml.EndElement{Name: elName}
		err = e.EncodeToken(xfe)
		if err != nil {
			return fmt.Errorf("error closing feature field element on cluster %s: %w", cluster.Name, err)
		}
	}
	xfe := xml.EndElement{Name: xml.Name{Local: "bitmap"}}
	err = e.EncodeToken(xfe)
	if err != nil {
		return fmt.Errorf("error closing feature element on cluster %s: %w", cluster.Name, err)
	}
	return newLine(e)
}

func (*renderer) setFeatureAttributes(xfs []xml.Attr, f *matter.Bit) ([]xml.Attr, error) {
	mask, err := f.Mask()
	if err != nil {
		err = fmt.Errorf("error parsing feature bit %s: %w", f.Bit, err)
		return nil, err
	}
	xfs = setAttributeValue(xfs, "name", zap.CleanName(f.Name))
	xfs = setAttributeValue(xfs, "mask", fmt.Sprintf("0x%02X", mask))
	return xfs, nil
}

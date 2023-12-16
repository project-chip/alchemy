package amend

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"slices"
	"strings"

	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func (r *renderer) amendBitmap(d xmlDecoder, e xmlEncoder, el xml.StartElement, cluster *matter.Cluster) (err error) {
	name := getAttributeValue(el.Attr, "name")

	var matchingBitmap *matter.Bitmap
	var skip bool
	for bm, handled := range r.bitmaps {
		if bm.Name == name || strings.TrimSuffix(bm.Name, "Bitmap") == name {
			matchingBitmap = bm
			skip = handled
			r.bitmaps[bm] = true
			break
		}
	}

	if matchingBitmap == nil || skip {
		Ignore(d, "bitmap")
		return nil
	}

	var valFormat string
	el.Attr, valFormat = r.setBitmapAttributes(el.Attr, matchingBitmap)
	err = e.EncodeToken(el)
	if err != nil {
		return
	}

	remainingClusterIDs := r.getClusterCodes(matchingBitmap)

	var bitIndex int

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
					if bitIndex >= len(matchingBitmap.Bits) {
						Ignore(d, "field")
						break
					} else {
						b := matchingBitmap.Bits[bitIndex]
						bitIndex++

						if conformance.IsZigbee(matchingBitmap.Bits, b.Conformance) {
							continue
						}

						t.Attr, err = r.setBitmapFieldAttributes(t.Attr, b, valFormat)
						if err != nil {
							err = fmt.Errorf("failed setting bitmap attributes on bitmap %s: %w", b.Name, err)
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
				slog.Warn("unexpected element in bitmap", "name", t.Name.Local)

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
				for bitIndex < len(matchingBitmap.Bits) {
					b := matchingBitmap.Bits[bitIndex]
					bitIndex++
					if conformance.IsZigbee(matchingBitmap.Bits, b.Conformance) {
						continue
					}

					elName := xml.Name{Local: "field"}
					xfs := xml.StartElement{Name: elName}
					xfs.Attr, err = r.setBitmapFieldAttributes(xfs.Attr, b, valFormat)
					if err != nil {
						err = fmt.Errorf("failed setting bitmap attributes on bitmap %s: %w", b.Name, err)
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

func (r *renderer) writeBitmap(e xmlEncoder, xfb xml.StartElement, bitmap *matter.Bitmap, provisional bool) (err error) {
	var valFormat string
	xfb.Attr, valFormat = r.setBitmapAttributes(xfb.Attr, bitmap)

	err = e.EncodeToken(xfb)
	if err != nil {
		return
	}

	err = r.renderClusterCodes(e, r.getClusterCodes(bitmap))
	if err != nil {
		return
	}

	for _, b := range bitmap.Bits {

		if conformance.IsZigbee(bitmap.Bits, b.Conformance) {
			continue
		}

		elName := xml.Name{Local: "field"}
		xfs := xml.StartElement{Name: elName}
		xfs.Attr, err = r.setBitmapFieldAttributes(xfs.Attr, b, valFormat)
		if err != nil {
			err = fmt.Errorf("failed setting bitmap attributes on bitmap %s: %w", b.Name, err)
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
	err = e.EncodeToken(xml.EndElement{Name: xfb.Name})
	if err != nil {
		return
	}
	err = newLine(e)
	return
}

func (*renderer) setBitmapFieldAttributes(xfs []xml.Attr, b *matter.Bit, valFormat string) ([]xml.Attr, error) {
	mask, err := b.Mask()
	if err != nil {
		return nil, err
	}

	name := zap.CleanName(b.Name)
	xfs = setAttributeValue(xfs, "name", name)
	xfs = setAttributeValue(xfs, "mask", fmt.Sprintf(valFormat, mask))
	return xfs, nil
}

func (*renderer) setBitmapAttributes(xfb []xml.Attr, bitmap *matter.Bitmap) ([]xml.Attr, string) {
	var valFormat string
	switch bitmap.Type.BaseType {
	case matter.BaseDataTypeMap64:
		valFormat = "0x%016X"
	case matter.BaseDataTypeMap32:
		valFormat = "0x%08X"
	case matter.BaseDataTypeMap16:
		valFormat = "0x%04X"
	default:
		valFormat = "0x%02X"
	}

	xfb = setAttributeValue(xfb, "name", bitmap.Name)
	if bitmap.Type != nil {
		xfb = setAttributeValue(xfb, "type", zap.ConvertDataTypeNameToZap(bitmap.Type.Name))
	} else {
		xfb = setAttributeValue(xfb, "type", "bitmap8")
	}
	return xfb, valFormat
}

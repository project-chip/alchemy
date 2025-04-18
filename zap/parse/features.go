package parse

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
)

func readFeatures(path string, d *xml.Decoder, e xml.StartElement) (features *matter.Features, err error) {
	features = matter.NewFeatures(nil, nil)
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of struct")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "feature":
				var feature *matter.Feature
				feature, err = readFeature(path, d, t)
				if err != nil {
					break
				}
				features.AddFeatureBit(feature)
			default:
				err = fmt.Errorf("unexpected features level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "features":
				return
			default:
				err = fmt.Errorf("unexpected features end element: %s", t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected features level type: %T", t)
		}
		if err != nil {
			err = fmt.Errorf("error parsing features: %w", err)
			return
		}
	}
}

func readFeature(path string, d *xml.Decoder, e xml.StartElement) (feature *matter.Feature, err error) {
	var bit, code, name, summary string
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			name = a.Value
		case "bit":
			bit = a.Value
		case "code":
			code = a.Value
		case "summary":
			summary = a.Value
		case "apiMaturity":
		default:
			err = fmt.Errorf("unexpected feature attribute: %s", a.Name.Local)
			return
		}
	}
	var con conformance.Conformance
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
		case xml.StartElement:
			if isConformanceElement(t) {
				con, err = parseConformance(d, t)
			} else {
				err = fmt.Errorf("unexpected feature start element: %s", t.Name.Local)
			}

		case xml.EndElement:
			switch t.Name.Local {
			case "feature":
				var cs conformance.Set
				if con != nil {
					cs = append(cs, con)
				}
				feature = matter.NewFeature(nil, bit, name, code, summary, cs)
				return
			default:
				err = fmt.Errorf("unexpected feature end element: %s", t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected %s level type: %T", name, t)
		}
		if err != nil {
			return
		}
	}
}

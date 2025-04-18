package parse

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/project-chip/alchemy/matter"
)

func parseQuality(d *xml.Decoder, e xml.StartElement) (q matter.Quality, err error) {
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "changeOmitted":
			q |= matter.QualityChangedOmitted
		case "largeMessage":
			q |= matter.QualityLargeMessage
		case "nullable":
			q |= matter.QualityNullable
		case "scene":
			q |= matter.QualityScene
		case "fixed":
			q |= matter.QualityFixed
		case "diagnostics":
			q |= matter.QualityDiagnostics
		case "singleton":
			q |= matter.QualitySingleton
		case "sourceAttribution":
			q |= matter.QualitySourceAttribution
		case "atomicWrite":
			q |= matter.QualityAtomicWrite
		case "persistence":
			switch a.Value {
			case "nonVolatile":
				q |= matter.QualityNonVolatile
			}
		default:
			err = fmt.Errorf("unexpected quality attribute: %s", a.Name.Local)
			return
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
		case xml.StartElement:
			err = fmt.Errorf("unexpected quality start element: %s", t.Name.Local)
		case xml.EndElement:
			switch t.Name.Local {
			case "quality":
				return
			default:
				err = fmt.Errorf("unexpected quality end element: %s", t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected quality type: %T", t)
		}
		if err != nil {
			return
		}
	}
	return
}

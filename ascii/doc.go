package ascii

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

type Doc struct {
	Path string

	Base     *types.Document
	Elements []interface{}

	docType matter.DocType

	Domain matter.Domain

	anchors map[string]*Anchor
}

func NewDoc(d *types.Document) (*Doc, error) {
	doc := &Doc{
		Base: d,
	}
	for _, e := range d.Elements {
		switch el := e.(type) {
		case *types.Section:
			s, err := NewSection(doc, el)
			if err != nil {
				return nil, err
			}
			doc.Elements = append(doc.Elements, s)
		default:
			doc.Elements = append(doc.Elements, NewElement(doc, e))
		}
	}
	return doc, nil
}

func (doc *Doc) DocType() (matter.DocType, error) {
	if doc.docType != matter.DocTypeUnknown {
		return doc.docType, nil
	}
	if len(doc.Path) == 0 {
		return matter.DocTypeUnknown, fmt.Errorf("missing path")
	}
	return GetDocType(doc.Path)
}

func GetDocType(docPath string) (matter.DocType, error) {
	path, err := filepath.Abs(docPath)
	if err != nil {
		return matter.DocTypeUnknown, err
	}
	dir, file := filepath.Split(path)
	pathParts := strings.Split(dir, string(os.PathSeparator))

	for i := len(pathParts) - 1; i >= 0; i-- {
		part := pathParts[i]
		switch part {
		case "app_clusters":
			if firstLetterIsLower(file) {
				return matter.DocTypeAppClusterIndex, nil
			}
			return matter.DocTypeAppCluster, nil
		case "common_protocol":
			return matter.DocTypeCommonProtocol, nil
		case "data_model":
			return matter.DocTypeDataModel, nil
		case "device_attestation":
			return matter.DocTypeDeviceAttestation, nil
		case "device_types":
			if firstLetterIsLower(file) {
				return matter.DocTypeDeviceTypeIndex, nil
			}
			return matter.DocTypeDeviceType, nil
		case "multi_admin":
			return matter.DocTypeMultiAdmin, nil
		case "namespaces":
			return matter.DocTypeNamespaces, nil
		case "qr_code":
			return matter.DocTypeQRCode, nil
		case "rendezvous":
			return matter.DocTypeRendezvous, nil
		case "secure_channel":
			return matter.DocTypeSecureChannel, nil
		case "service_device_management":
			return matter.DocTypeServiceDeviceManagement, nil
		case "softAp":
			return matter.DocTypeSoftAP, nil
		}
	}
	slog.Warn("could not determine doc type", "path", docPath)
	return matter.DocTypeUnknown, nil
}

func firstLetterIsLower(s string) bool {
	firstLetter, _ := utf8.DecodeRuneInString(s)
	return unicode.IsLower(firstLetter)
}

func (d *Doc) GetElements() []interface{} {
	return d.Elements
}

func (d *Doc) SetElements(elements []interface{}) error {
	d.Elements = elements
	return nil
}

func (d *Doc) Footnotes() []*types.Footnote {
	return d.Base.Footnotes
}

func (d *Doc) ToModel() (models []interface{}, err error) {
	dt, err := d.DocType()
	if err != nil {
		return nil, err
	}

	crossReferences := d.CrossReferences()

	d.anchors, err = d.Anchors(crossReferences)
	if err != nil {
		return nil, err
	}

	for _, top := range parse.Skim[*Section](d.Elements) {
		AssignSectionTypes(dt, top)

		var m []interface{}
		m, err = top.ToModels(d)
		if err != nil {
			return
		}
		models = append(models, m...)
	}
	return
}

func Open(path string, settings ...configuration.Setting) (*Doc, error) {

	baseConfig := []configuration.Setting{configuration.WithFilename(path)}

	baseConfig = append(baseConfig, settings...)

	config := configuration.NewConfiguration(baseConfig...)

	file, err := os.ReadFile(config.Filename)
	if err != nil {
		return nil, err
	}

	contents := string(file)

	if len(config.Attributes) > 2 { // By default, there are two attributes in the renderer; if there are more, we need to pre-process
		contents, err = parser.Preprocess(strings.NewReader(contents), config)
		if err != nil {
			return nil, err
		}
	}

	d, err := parse.ParseDocument(strings.NewReader(contents), config, parser.MaxExpressions(2000000))

	if err != nil {
		return nil, fmt.Errorf("failed parse: %w", err)
	}

	doc, err := NewDoc(d)
	if err != nil {
		return nil, err
	}
	doc.Path = path
	return doc, nil
}

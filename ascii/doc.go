package ascii

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

type Doc struct {
	lock sync.Mutex
	Path string

	Base     *types.Document
	Elements []interface{}

	docType matter.DocType

	Domain  matter.Domain
	parents []*Doc

	anchors         map[string]*Anchor
	crossReferences map[string][]*types.InternalCrossReference
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

	switch filepath.Base(docPath) {
	case "appclusters.adoc":
		return matter.DocTypeAppClusters, nil
	case "standard_namespaces.adoc":
		return matter.DocTypeNamespaces, nil
	case "device_library.adoc":
		return matter.DocTypeDeviceTypes, nil
	}

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
			return matter.DocTypeNamespace, nil
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

func (d *Doc) Parents() []*Doc {
	return d.parents
}

func (d *Doc) addParent(parent *Doc) {
	d.lock.Lock()
	d.parents = append(d.parents, parent)
	d.lock.Unlock()
}

func (d *Doc) ToModel() (models []interface{}, err error) {
	dt, err := d.DocType()
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

	baseConfig := make([]configuration.Setting, len(settings)+1)
	baseConfig[0] = configuration.WithFilename(path)
	copy(baseConfig[1:], settings)

	//baseConfig = append(baseConfig, settings...)

	config := configuration.NewConfiguration(baseConfig...)

	file, err := os.ReadFile(config.Filename)
	if err != nil {
		wd, _ := os.Getwd()
		panic(fmt.Errorf("failed opening %s: %v wd: %s", config.Filename, err, wd))
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

func GithubSettings() []configuration.Setting {
	return []configuration.Setting{configuration.WithAttribute("env-github", true)}
}

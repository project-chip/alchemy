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
	mattertypes "github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/parse"
)

type Doc struct {
	sync.RWMutex

	Path string

	Base     *types.Document
	Elements []interface{}

	docType matter.DocType

	Domain  matter.Domain
	parents []*Doc

	anchors         map[string]*Anchor
	crossReferences map[string][]*types.InternalCrossReference

	entities []mattertypes.Entity

	entitiesBySection map[types.WithAttributes][]mattertypes.Entity
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
			return matter.DocTypeCluster, nil
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
			if strings.HasSuffix(file, "Cluster.adoc") {
				return matter.DocTypeCluster, nil
			}
			return matter.DocTypeServiceDeviceManagement, nil
		case "softAp":
			return matter.DocTypeSoftAP, nil
		}
	}
	slog.Debug("could not determine doc type", "path", docPath)
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
	d.RLock()
	p := make([]*Doc, len(d.parents))
	copy(p, d.parents)
	d.RUnlock()
	return p
}

func (d *Doc) addParent(parent *Doc) {
	d.Lock()
	d.parents = append(d.parents, parent)
	d.Unlock()
}

func (d *Doc) Entities() (entities []mattertypes.Entity, err error) {
	if len(d.entities) > 0 {
		return d.entities, nil
	}
	dt, err := d.DocType()
	if err != nil {
		return nil, err
	}

	var entitiesBySection = make(map[types.WithAttributes][]mattertypes.Entity)
	for _, top := range parse.Skim[*Section](d.Elements) {
		AssignSectionTypes(dt, top)

		var m []mattertypes.Entity
		m, err = top.toEntities(d, entitiesBySection)
		if err != nil {
			return nil, fmt.Errorf("failed converting doc to entities: %w", err)
		}
		entities = append(entities, m...)

	}
	d.entities = entities
	d.entitiesBySection = entitiesBySection
	return
}

func (d *Doc) Reference(ref string) (mattertypes.Entity, bool) {

	a, err := d.getAnchor(ref)

	if err != nil {
		slog.Warn("failed getting anchor", slog.String("path", d.Path), slog.String("reference", ref), slog.Any("error", err))
		return nil, false
	}
	if a == nil {
		slog.Warn("unknown reference", slog.String("path", d.Path), slog.String("reference", ref))
		return nil, false
	}
	entities, ok := d.entitiesBySection[a.Element]
	if !ok {
		slog.Warn("unknown reference entity", slog.String("path", d.Path), slog.String("reference", ref), slog.Any("count", len(d.entitiesBySection)))
		for sec, e := range d.entitiesBySection {
			slog.Warn("reference", slog.String("path", d.Path), slog.String("reference", ref), slog.Any("sec", sec), slog.Any("entity", e))

		}
	}
	if len(entities) == 0 {
		slog.Warn("unknown reference entity", slog.String("path", d.Path), slog.String("reference", ref))
		return nil, false
	}
	if len(entities) > 1 {
		slog.Warn("ambiguous reference", slog.String("path", d.Path), slog.String("reference", ref))
		for _, e := range entities {
			slog.Warn("reference", slog.String("path", d.Path), slog.String("reference", ref), slog.Any("entity", e))

		}
		return nil, false
	}
	return entities[0], true
	/*entities, err := d.Entities()
	if err != nil {
		slog.Warn("failed generating entities", slog.String("path", d.Path), slog.String("reference", ref), slog.Any("error", err))
	}
	var matches []mattertypes.Entity
	for _, e := range entities {
		if identifierStore, ok := e.(conformance.IdentifierStore); ok {
			slog.Info("querying anchor", slog.String("name", a.Name))

			entity, ok := identifierStore.Identifier(a.Name)
			if ok {
				matches = append(matches, entity)
			}
		}
	}
	if len(matches) == 0 {
		slog.Warn("unknown reference entity", slog.String("path", d.Path), slog.String("reference", ref))
		return nil, false
	}
	if len(matches) > 1 {
		slog.Warn("ambiguous reference", slog.String("path", d.Path), slog.String("reference", ref))
		for _, e := range entities {
			slog.Warn("reference", slog.String("path", d.Path), slog.String("reference", ref), slog.Any("entity", e))

		}
		return nil, false
	}
	return matches[0], true*/
}

func OpenFile(path string, settings ...configuration.Setting) (*Doc, error) {

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	contents := string(file)

	return Read(contents, path, settings...)
}

func Read(contents string, path string, settings ...configuration.Setting) (doc *Doc, err error) {
	baseConfig := make([]configuration.Setting, len(settings)+1)
	baseConfig[0] = configuration.WithFilename(path)
	copy(baseConfig[1:], settings)

	config := configuration.NewConfiguration(baseConfig...)
	config.IgnoreIncludes = true

	// By default, there are two attributes in the renderer; if there are more, we need to pre-process
	if len(config.Attributes) > 2 {
		contents, err = parser.Preprocess(strings.NewReader(contents), config)
		if err != nil {
			return nil, err
		}
	}

	var d *types.Document
	d, err = ParseDocument(strings.NewReader(contents), config, parser.MaxExpressions(2000000))

	if err != nil {
		return nil, fmt.Errorf("failed parse: %w", err)
	}

	doc, err = NewDoc(d)
	if err != nil {
		return nil, err
	}
	doc.Path = path

	PatchUnrecognizedReferences(doc)

	return doc, nil
}

func GithubSettings() []configuration.Setting {
	return []configuration.Setting{configuration.WithAttribute("env-github", true)}
}

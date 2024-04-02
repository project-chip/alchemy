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
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

type Doc struct {
	sync.RWMutex

	Path string

	Base     *types.Document
	Elements []any

	docType matter.DocType

	Domain  matter.Domain
	parents []*Doc

	anchors         map[string]*Anchor
	crossReferences map[string][]*types.InternalCrossReference
	attributes      map[string]any

	entities       []mattertypes.Entity
	entitiesParsed bool

	entitiesBySection map[types.WithAttributes][]mattertypes.Entity
}

func NewDoc(d *types.Document) (*Doc, error) {
	doc := &Doc{
		Base:       d,
		attributes: make(map[string]any),
	}
	for _, e := range d.Elements {
		switch el := e.(type) {
		case *types.AttributeDeclaration:
			doc.attributes[el.Name] = el.Value
			doc.Elements = append(doc.Elements, NewElement(doc, e))
		case *types.Section:
			s, err := NewSection(doc, doc, el)
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

func (doc *Doc) GetElements() []any {
	return doc.Elements
}

func (doc *Doc) SetElements(elements []any) error {
	doc.Elements = elements
	return nil
}

func (doc *Doc) Footnotes() []*types.Footnote {
	return doc.Base.Footnotes
}

func (doc *Doc) Parents() []*Doc {
	doc.RLock()
	p := make([]*Doc, len(doc.parents))
	copy(p, doc.parents)
	doc.RUnlock()
	return p
}

func (doc *Doc) addParent(parent *Doc) {
	doc.Lock()
	doc.parents = append(doc.parents, parent)
	doc.Unlock()
}

func (doc *Doc) Entities() (entities []mattertypes.Entity, err error) {
	if doc.entitiesParsed {
		return doc.entities, nil
	}
	doc.entitiesParsed = true

	var entitiesBySection = make(map[types.WithAttributes][]mattertypes.Entity)
	for _, top := range parse.Skim[*Section](doc.Elements) {
		err := AssignSectionTypes(doc, top)
		if err != nil {
			return nil, err
		}

		var m []mattertypes.Entity
		m, err = top.toEntities(doc, entitiesBySection)
		if err != nil {
			return nil, fmt.Errorf("failed converting doc %s to entities: %w", doc.Path, err)
		}
		entities = append(entities, m...)
	}
	doc.entities = entities
	doc.entitiesBySection = entitiesBySection
	return
}

func (doc *Doc) Reference(ref string) (mattertypes.Entity, bool) {

	a, err := doc.getAnchor(ref)

	if err != nil {
		slog.Warn("failed getting anchor", slog.String("path", doc.Path), slog.String("reference", ref), slog.Any("error", err))
		return nil, false
	}
	if a == nil {
		slog.Warn("unknown reference", slog.String("path", doc.Path), slog.String("reference", ref))
		return nil, false
	}
	entities, ok := doc.entitiesBySection[a.Element]
	if !ok {
		slog.Warn("unknown reference entity", slog.String("path", doc.Path), slog.String("reference", ref), slog.Any("count", len(doc.entitiesBySection)))
		for sec, e := range doc.entitiesBySection {
			slog.Warn("reference", slog.String("path", doc.Path), slog.String("reference", ref), slog.Any("sec", sec), slog.Any("entity", e))

		}
	}
	if len(entities) == 0 {
		slog.Warn("unknown reference entity", slog.String("path", doc.Path), slog.String("reference", ref))
		return nil, false
	}
	if len(entities) > 1 {
		slog.Warn("ambiguous reference", slog.String("path", doc.Path), slog.String("reference", ref))
		for _, e := range entities {
			slog.Warn("reference", slog.String("path", doc.Path), slog.String("reference", ref), slog.Any("entity", e))

		}
		return nil, false
	}
	return entities[0], true
}

func GithubSettings() []configuration.Setting {
	return []configuration.Setting{configuration.WithAttribute("env-github", true)}
}

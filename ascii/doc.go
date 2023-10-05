package ascii

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/matter"
)

type Doc struct {
	Path string

	Base     *types.Document
	Elements []interface{}

	docType matter.DocType
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

	path, err := filepath.Abs(doc.Path)
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
	slog.Warn("could not determine doc type", "path", doc.Path)
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

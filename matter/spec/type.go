package spec

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
)

func (library *Library) DocType(doc *asciidoc.Document) (matter.DocType, error) {
	dt, ok := library.docTypes[doc]
	if ok {
		return dt, nil
	}
	dt, err := determineDocType(library, library, asciidoc.RawReader, doc)
	if err == nil {
		library.docTypes[doc] = dt
	}
	return dt, err
}

func (library *Library) SetDocType(doc *asciidoc.Document, dt matter.DocType) {
	library.docTypes[doc] = dt
}

func AssignDocType(sectionInfoCache SectionInfoCache, docInfoCache DocumentInfoCache, reader asciidoc.Reader, doc *asciidoc.Document) error {
	dt, err := determineDocType(sectionInfoCache, docInfoCache, reader, doc)
	if err == nil {
		docInfoCache.SetDocType(doc, dt)
	}
	return err
}

func determineDocType(sectionInfoCache SectionInfoCache, docInfoCache DocumentInfoCache, reader asciidoc.Reader, doc *asciidoc.Document) (matter.DocType, error) {
	if len(doc.Path.Absolute) == 0 {
		return matter.DocTypeUnknown, fmt.Errorf("missing path")
	}

	switch doc.Path.Base() {
	case "appclusters.adoc":
		return matter.DocTypeAppClusters, nil
	case "standard_namespaces.adoc":
		return matter.DocTypeNamespaces, nil
	case "device_library.adoc":
		return matter.DocTypeDeviceTypes, nil
	case "BaseDeviceType.adoc":
		return matter.DocTypeBaseDeviceType, nil
	}

	path := doc.Path.Absolute
	dir, file := filepath.Split(path)
	pathParts := strings.Split(dir, string(os.PathSeparator))

	for i := len(pathParts) - 1; i >= 0; i-- {
		part := pathParts[i]
		switch part {
		case "app_clusters":
			if firstLetterIsLower(file) {
				return matter.DocTypeAppClusterIndex, nil
			}
			for _, p := range docInfoCache.Parents(doc) {
				if filepath.Base(p.Path.Relative) == "appclusters.adoc" {
					return matter.DocTypeAppClusterIndex, nil
				}
			}
			return matter.DocTypeCluster, nil
		case "common_protocol":
			return matter.DocTypeCommonProtocol, nil
		case "data_model":
			name := text.TrimCaseInsensitiveSuffix(path, filepath.Ext(path))
			if strings.Contains(strings.ToLower(name), "cluster") {
				return matter.DocTypeCluster, nil
			}
			if dt := guessDocType(sectionInfoCache, reader, doc); dt != matter.DocTypeUnknown {
				return dt, nil
			}
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
			name := text.TrimCaseInsensitiveSuffix(path, filepath.Ext(path))
			if strings.Contains(strings.ToLower(name), "cluster") {
				return matter.DocTypeCluster, nil
			}
			if dt := guessDocType(sectionInfoCache, reader, doc); dt != matter.DocTypeUnknown {
				return dt, nil
			}
			return matter.DocTypeServiceDeviceManagement, nil
		case "softAp":
			return matter.DocTypeSoftAP, nil
		}
	}
	guess := guessDocType(sectionInfoCache, reader, doc)
	slog.Debug("could not determine doc type", "path", doc.Path, slog.String("guessing", guess.String()))
	return guess, nil
}

func guessDocType(sectionInfoCache SectionInfoCache, reader asciidoc.Reader, doc *asciidoc.Document) matter.DocType {
	firstSection := parse.FindFirst[*asciidoc.Section](doc, reader, doc)
	if firstSection != nil {
		if text.HasCaseInsensitiveSuffix(sectionInfoCache.SectionName(firstSection), " cluster") {
			return matter.DocTypeCluster
		}
	}
	return matter.DocTypeUnknown
}

func firstLetterIsLower(s string) bool {
	firstLetter, _ := utf8.DecodeRuneInString(s)
	return unicode.IsLower(firstLetter)
}

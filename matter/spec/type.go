package spec

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
)

func (doc *Doc) DocType() (matter.DocType, error) {
	if doc.docType != matter.DocTypeUnknown {
		return doc.docType, nil
	}
	dt, err := doc.determineDocType()
	if err == nil {
		doc.docType = dt
	}
	return dt, err
}

func (doc *Doc) determineDocType() (matter.DocType, error) {
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
			for _, p := range doc.Parents() {
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
			if dt := doc.guessDocType(); dt != matter.DocTypeUnknown {
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
			if dt := doc.guessDocType(); dt != matter.DocTypeUnknown {
				return dt, nil
			}
			return matter.DocTypeServiceDeviceManagement, nil
		case "softAp":
			return matter.DocTypeSoftAP, nil
		}
	}
	guess := doc.guessDocType()
	slog.Debug("could not determine doc type", "path", doc.Path, slog.String("guessing", guess.String()))
	return guess, nil
}

func (doc *Doc) guessDocType() matter.DocType {
	firstSection := parse.FindFirst[*Section](doc.Elements())
	if firstSection != nil {
		if text.HasCaseInsensitiveSuffix(firstSection.Name, " cluster") {
			return matter.DocTypeCluster
		}
	}
	return matter.DocTypeUnknown
}

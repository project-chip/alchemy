package spec

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
)

func (doc *Doc) DocType() (matter.DocType, error) {
	if doc.docType != matter.DocTypeUnknown {
		return doc.docType, nil
	}
	if len(doc.Path) == 0 {
		return matter.DocTypeUnknown, fmt.Errorf("missing path")
	}

	switch filepath.Base(doc.Path) {
	case "appclusters.adoc":
		return matter.DocTypeAppClusters, nil
	case "standard_namespaces.adoc":
		return matter.DocTypeNamespaces, nil
	case "device_library.adoc":
		return matter.DocTypeDeviceTypes, nil
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
			return matter.DocTypeCluster, nil
		case "common_protocol":
			return matter.DocTypeCommonProtocol, nil
		case "data_model":
			name := strings.TrimSuffix(path, filepath.Ext(path))
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
			name := strings.TrimSuffix(path, filepath.Ext(path))
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
	slog.Debug("could not determine doc type", "path", doc.Path)
	return matter.DocTypeUnknown, nil
}

func (doc *Doc) guessDocType() matter.DocType {
	firstSection := parse.FindFirst[*Section](doc.Elements())
	if firstSection != nil {
		if strings.HasSuffix(strings.ToLower(firstSection.Name), " cluster") {
			return matter.DocTypeCluster
		}
	}
	return matter.DocTypeUnknown
}

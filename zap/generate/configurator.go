package generate

import (
	"bytes"
	"fmt"
	"log/slog"
	"slices"
	"time"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/zap"
)

func newZapTemplate() (x *etree.Document) {
	x = etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(fmt.Sprintf(license, time.Now().Year()))
	return
}

func renderZapTemplate(configurator *zap.Configurator, x *etree.Document, errata *zap.Errata) (result string, err error) {

	var exampleCluster *matter.Cluster
	for c := range configurator.Clusters {
		if c != nil {
			exampleCluster = c
			break
		}
	}

	ce := x.SelectElement("configurator")
	if ce == nil {
		ce = x.CreateElement("configurator")
	}

	de := ce.SelectElement("domain")
	if de == nil {
		de = ce.CreateElement("domain")
		de.CreateAttr("name", matter.DomainNames[configurator.Doc.Domain])
	}

	if exampleCluster != nil {
		err = generateFeatures(configurator, ce, exampleCluster.Features, errata)
		if err != nil {
			return
		}
	}

	err = generateBitmaps(configurator, ce, exampleCluster, errata)
	if err != nil {
		return
	}

	err = generateEnums(configurator, ce, exampleCluster, errata)
	if err != nil {
		return
	}

	err = generateStructs(configurator, ce, exampleCluster, errata)
	if err != nil {
		return
	}

	err = renderClusters(configurator, ce, errata)
	if err != nil {
		return
	}

	x.Indent(2)
	var b bytes.Buffer
	x.WriteTo(&b)
	return b.String(), nil

}

func generateFeatures(configurator *zap.Configurator, configuratorElement *etree.Element, features *matter.Bitmap, errata *zap.Errata) (err error) {
	needFeatures := features != nil && len(features.Bits) > 0

	bitmaps := configuratorElement.SelectElements("bitmap")
	var clusterIds []string
	for c := range configurator.Clusters {
		clusterIds = append(clusterIds, c.ID.HexString())
	}
	for _, bm := range bitmaps {
		nameAttr := bm.SelectAttr("name")
		if nameAttr == nil || nameAttr.Value != "Feature" {
			continue
		}
		if needFeatures {

			err = populateBitmap(configurator, bm, features, clusterIds, errata)
			needFeatures = false
		} else {
			configuratorElement.RemoveChild(bm)
		}
	}
	if needFeatures {
		fe := etree.NewElement("bitmap")
		err = populateBitmap(configurator, fe, features, clusterIds, errata)
		if err != nil {
			return
		}
		appendElement(configuratorElement, fe, "domain")
	}
	return
}

func generateClusterCodes(spec *matter.Spec, entity types.Entity, parent *etree.Element) {
	clusterIDs := clusterIdsForEntity(spec, entity)
	_, remainingClusterIDs := amendExistingClusterCodes(parent, entity, clusterIDs)
	flushClusterCodes(parent, remainingClusterIDs)
}

func amendExistingClusterCodes(parent *etree.Element, entity types.Entity, clusterIDs []string) (amendedCodes []string, remainingCodes []string) {
	clusterCodeElements := parent.SelectElements("cluster")
	remainingCodes = make([]string, len(clusterIDs))
	copy(remainingCodes, clusterIDs)
	for _, cce := range clusterCodeElements {
		ca := cce.SelectAttr("code")
		if ca == nil {
			slog.Warn("missing cluster code", "val", entity)
			continue
		}
		id := matter.ParseNumber(ca.Value)

		if !id.Valid() || !slices.Contains(clusterIDs, ca.Value) {
			parent.RemoveChild(cce)
			continue
		}
		ids := id.HexString()
		amendedCodes = append(amendedCodes, ids)
		remainingCodes = slices.DeleteFunc(remainingCodes, func(s string) bool { return s == ids })
	}
	return
}

func flushClusterCodes(parent *etree.Element, clusterIDs []string) {
	for _, clusterID := range clusterIDs {
		ce := etree.NewElement("cluster")
		ce.CreateAttr("code", clusterID)
		appendElement(parent, ce, "cluster")
	}
}

func clusterIdsForEntity(spec *matter.Spec, entity types.Entity) (clusterIDs []string) {
	refs, ok := spec.ClusterRefs[entity]
	if !ok {
		slog.Warn("unknown cluster ref", "val", entity)
		return
	}
	for ref := range refs {
		clusterIDs = append(clusterIDs, ref.ID.HexString())
	}
	slices.Sort(clusterIDs)
	return
}

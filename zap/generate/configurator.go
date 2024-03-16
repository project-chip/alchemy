package generate

import (
	"bytes"
	"fmt"
	"log/slog"
	"regexp"
	"slices"
	"time"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/internal/xml"
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
	_, err = x.WriteTo(&b)
	s := b.String()
	s = postProcessTemplate(s)
	return s, err
}

var tagClosePattern = regexp.MustCompile(`(?m)/(?P<Tag>bitmap|cluster|command|enum|event|struct)>\n(\s+)<`)

func postProcessTemplate(s string) string {
	// etree removes extraneous whitespace between tags, so this restores it for commonly separated tags in ZAP templates
	s = tagClosePattern.ReplaceAllString(s, "/$1>\n\n$2<")
	return s
}

func generateFeatures(configurator *zap.Configurator, configuratorElement *etree.Element, features *matter.Features, errata *zap.Errata) (err error) {
	needFeatures := features != nil && len(features.Bits) > 0

	bitmaps := configuratorElement.SelectElements("bitmap")

	var clusterIds []*matter.Number
	for c := range configurator.Clusters {
		clusterIds = append(clusterIds, c.ID)
	}
	for _, bm := range bitmaps {
		nameAttr := bm.SelectAttr("name")
		if nameAttr == nil || nameAttr.Value != "Feature" {
			continue
		}
		if needFeatures {

			err = populateBitmap(configurator, bm, &features.Bitmap, clusterIds, errata)
			needFeatures = false
		} else {
			configuratorElement.RemoveChild(bm)
		}
	}
	if needFeatures {
		fe := etree.NewElement("bitmap")
		err = populateBitmap(configurator, fe, &features.Bitmap, clusterIds, errata)
		if err != nil {
			return
		}
		xml.AppendElement(configuratorElement, fe, "domain")
	}
	return
}

func amendExistingClusterCodes(parent *etree.Element, entity types.Entity, clusterIDs []*matter.Number) (amendedCodes []*matter.Number, remainingCodes []*matter.Number) {
	clusterCodeElements := parent.SelectElements("cluster")
	remainingCodes = make([]*matter.Number, len(clusterIDs))
	copy(remainingCodes, clusterIDs)
	for _, cce := range clusterCodeElements {
		ca := cce.SelectAttr("code")
		if ca == nil {
			slog.Debug("missing cluster code", "val", entity)
			continue
		}
		id := matter.ParseNumber(ca.Value)

		if !id.Valid() || !matter.ContainsNumber(clusterIDs, id) {
			parent.RemoveChild(cce)
			continue
		}
		amendedCodes = append(amendedCodes, id)
		remainingCodes = slices.DeleteFunc(remainingCodes, func(s *matter.Number) bool { return s.Equals(id) })
	}
	return
}

func flushClusterCodes(parent *etree.Element, clusterIDs []*matter.Number) {
	for _, clusterID := range clusterIDs {
		ce := etree.NewElement("cluster")
		patchNumberAttribute(ce, clusterID, "code")
		xml.AppendElement(parent, ce, "cluster")
	}
}

func clusterIdsForEntity(spec *matter.Spec, entity types.Entity) (clusterIDs []*matter.Number) {
	refs, ok := spec.ClusterRefs[entity]
	if !ok {
		slog.Warn("unknown cluster ref", "val", entity)
		return
	}
	for ref := range refs {
		clusterIDs = append(clusterIDs, ref.ID)
	}
	matter.SortNumbers(clusterIDs)
	return
}

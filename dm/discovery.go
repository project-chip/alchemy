package dm

import (
	"bytes"
	"path/filepath"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
)

func renderDiscoveryBitmap(bm *matter.Bitmap, dmRoot string, globalFiles pipeline.StringSet) (err error) {
	x := etree.NewDocument()
	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(getLicense())

	root := x.CreateElement("discovery")
	bmEl := root.CreateElement("bitmap")
	bmEl.CreateAttr("name", bm.Name)
	size := bm.Size() / 4
	for _, v := range bm.Bits {
		err = renderBit(bmEl, v, size, bm)
		if err != nil {
			return
		}
	}

	x.Indent(2)
	var b bytes.Buffer
	_, err = x.WriteTo(&b)
	if err != nil {
		return
	}

	outPath := getDiscoveryBitmapPath(dmRoot, bm)
	globalFiles.Store(outPath, pipeline.NewData(outPath, b.String()))
	return
}

func getDiscoveryBitmapPath(dmRoot string, bm *matter.Bitmap) string {
	name := bm.Name
	if strings.HasSuffix(name, "Bitmap") {
		name = name[:len(name)-len("Bitmap")]
	}
	return filepath.Join(dmRoot, "discovery", name+".xml")
}

func isDiscoveryDoc(doc *asciidoc.Document) bool {
	return strings.HasSuffix(doc.Path.Relative, "secure_channel/Discovery.adoc")
}

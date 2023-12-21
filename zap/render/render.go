package render

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

type Result struct {
	ZCL    string
	Doc    *ascii.Doc
	Models []matter.Model
}

type renderer struct {
	spec   *matter.Spec
	doc    *ascii.Doc
	errata *zap.Errata
}

func Render(cxt context.Context, spec *matter.Spec, doc *ascii.Doc, models []matter.Model, errata *zap.Errata) (*Result, error) {

	r := &renderer{
		spec:   spec,
		doc:    doc,
		errata: errata,
	}
	return r.render(cxt, models)

}

func (r *renderer) render(cxt context.Context, models []matter.Model) (*Result, error) {

	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(fmt.Sprintf(license, time.Now().Year()))
	err := r.renderModels(cxt, &x.Element, models)
	if err != nil {
		return nil, err
	}
	x.Indent(2)

	var b bytes.Buffer
	x.WriteTo(&b)
	return &Result{ZCL: b.String(), Doc: r.doc, Models: models}, nil
}

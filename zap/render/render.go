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

type renderer struct {
	spec         *matter.Spec
	configurator *zap.Configurator
	errata       *zap.Errata
}

func Render(cxt context.Context, spec *matter.Spec, doc *ascii.Doc, configurator *zap.Configurator, errata *zap.Errata) (string, error) {

	r := &renderer{
		spec:         spec,
		configurator: configurator,
		errata:       errata,
	}

	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(fmt.Sprintf(license, time.Now().Year()))
	err := r.renderModels(cxt, doc, &x.Element)
	if err != nil {
		return "", err
	}
	x.Indent(2)

	var b bytes.Buffer
	x.WriteTo(&b)
	return b.String(), nil
}

package zcl

import (
	"bytes"
	"context"
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

type Result struct {
	ZCL    string
	Doc    *ascii.Doc
	Models []interface{}
}

func Render(cxt context.Context, doc *ascii.Doc) (*Result, error) {
	docType, err := doc.DocType()
	if err != nil {
		return nil, err
	}

	models, err := doc.ToModel()
	if err != nil {
		return nil, err
	}

	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)

	errata, ok := Erratas[filepath.Base(doc.Path)]
	if !ok {
		errata = DefaultErrata
	}

	//x.CreateComment(license)
	c := x.CreateElement("configurator")
	dom := c.CreateElement("domain")
	dom.CreateAttr("name", "CHIP")
	switch docType {
	case matter.DocTypeAppCluster:

		err = renderAppCluster(cxt, doc, models, c, errata)
	}
	if err != nil {
		return nil, err
	}
	x.Indent(2)

	var b bytes.Buffer
	x.WriteTo(&b)
	return &Result{ZCL: b.String(), Doc: doc, Models: models}, nil
}

var license = `
Copyright (c) 2021 Project CHIP Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
`

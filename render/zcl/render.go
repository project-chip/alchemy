package zcl

import (
	"bytes"
	"context"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func Render(cxt context.Context, doc *ascii.Doc) (string, error) {
	docType, err := doc.DocType()
	if err != nil {
		return "", err
	}

	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	switch docType {
	case matter.DocTypeAppCluster:
		err = renderAppCluster(cxt, doc, x)
	}
	if err != nil {
		return "", err
	}
	x.Indent(4)
	var b bytes.Buffer
	x.WriteTo(&b)
	return b.String(), nil
}

var license = `Copyright (c) 2021 Project CHIP Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.`

package dm

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/constraint"
	"github.com/hasty/alchemy/matter"
)

type Result struct {
	XML    string
	Path   string
	Doc    *ascii.Doc
	Models []interface{}
}

func renderAppClusters(cxt context.Context, zclRoot string, appClusters []*ascii.Doc, filesOptions files.Options) error {
	var lock sync.Mutex
	outputs := make(map[string]string)
	err := files.ProcessDocs(cxt, appClusters, func(cxt context.Context, doc *ascii.Doc, index, total int) error {
		slog.Info("App cluster doc", "name", doc.Path)

		models, err := doc.ToModel()
		if err != nil {
			return err
		}
		var clusters []*matter.Cluster
		for _, m := range models {
			slog.Info("model", "type", m)
			switch m := m.(type) {
			case *matter.Cluster:
				clusters = append(clusters, m)
			}
		}
		s, err := renderAppCluster(cxt, clusters)
		if err != nil {
			return fmt.Errorf("failed rendering %s: %w", doc.Path, err)
		}
		lock.Lock()
		outputs[doc.Path] = s
		lock.Unlock()
		return nil
	}, filesOptions)

	if err != nil {
		return err
	}

	if !filesOptions.DryRun {
		for path, result := range outputs {
			path := filepath.Base(path)
			newPath := filepath.Join(zclRoot, fmt.Sprintf("/data_model/clusters/%s.xml", strings.TrimSuffix(path, filepath.Ext(path))))
			err = os.WriteFile(newPath, []byte(result), os.ModeAppend|0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func renderAppCluster(cxt context.Context, clusters []*matter.Cluster) (output string, err error) {
	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(license)
	for _, cluster := range clusters {
		c := x.CreateElement("cluster")
		c.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
		c.CreateAttr("xsi:schemaLocation", "types types.xsd cluster cluster.xsd")
		c.CreateAttr("id", cluster.ID.HexString())
		c.CreateAttr("name", cluster.Name)

		revs := c.CreateElement("revisionHistory")
		var latestRev uint64 = 0
		for _, r := range cluster.Revisions {
			id := matter.ParseID(r.Number)
			if id.Valid() {
				rev := revs.CreateElement("revision")
				rev.CreateAttr("revision", id.IntString())
				rev.CreateAttr("summary", r.Description)
				latestRev = max(id.Value(), latestRev)
			}
		}
		c.CreateAttr("revision", strconv.FormatUint(latestRev, 10))
		class := c.CreateElement("classification")
		class.CreateAttr("hierarchy", strings.ToLower(cluster.Hierarchy))
		class.CreateAttr("role", strings.ToLower(cluster.Role))
		class.CreateAttr("picsCode", cluster.PICS)
		class.CreateAttr("scope", cluster.Scope)

		if len(cluster.Features) > 0 {
			features := c.CreateElement("features")
			for _, f := range cluster.Features {
				bit := matter.ParseID(f.Bit)
				if !bit.Valid() {
					continue
				}
				feature := features.CreateElement("feature")
				feature.CreateAttr("bit", bit.IntString())
				feature.CreateAttr("code", f.Code)
				feature.CreateAttr("name", f.Name)
				if len(f.Summary) > 0 {
					feature.CreateAttr("summary", f.Summary)
				}
				err = renderConformanceString(f.Conformance, feature)
				if err != nil {
					return "", err
				}
			}
		}
		if len(cluster.Enums) > 0 || len(cluster.Bitmaps) > 0 || len(cluster.Structs) > 0 {
			dt := c.CreateElement("dataTypes")
			for _, e := range cluster.Enums {
				en := dt.CreateElement("enum")
				en.CreateAttr("name", e.Name)
				for _, v := range e.Values {
					val := matter.ParseID(v.Value)
					if !val.Valid() {
						continue
					}
					i := en.CreateElement("item")
					i.CreateAttr("value", val.IntString())
					i.CreateAttr("name", v.Name)
					if len(v.Summary) > 0 {
						i.CreateAttr("summary", v.Summary)
					}
					err = renderConformanceString(v.Conformance, i)
					if err != nil {
						return "", err
					}

				}
			}

			for _, bm := range cluster.Bitmaps {
				en := dt.CreateElement("bitmap")
				en.CreateAttr("name", bm.Name)
				for _, v := range bm.Bits {
					val := matter.ParseID(v.Bit)
					if !val.Valid() {
						continue
					}
					i := en.CreateElement("bitfield")
					i.CreateAttr("name", v.Name)
					i.CreateAttr("bit", val.IntString())
					i.CreateAttr("summary", v.Summary)
					err = renderConformanceString(v.Conformance, i)
					if err != nil {
						return "", err
					}

				}
			}

			for _, s := range cluster.Structs {
				en := dt.CreateElement("struct")
				en.CreateAttr("name", s.Name)
				for _, f := range s.Fields {
					if !f.ID.Valid() {
						continue
					}
					i := en.CreateElement("field")
					i.CreateAttr("id", f.ID.IntString())
					i.CreateAttr("name", f.Name)
					if f.Type != nil {
						i.CreateAttr("type", f.Type.Name)
					}
					if f.Access.Read != matter.PrivilegeUnknown {
						i.CreateAttr("read", "true")
					}
					if f.Access.Write != matter.PrivilegeUnknown {
						i.CreateAttr("write", "true")
					}
					err = renderConformanceString(f.Conformance, i)
					if err != nil {
						return "", err
					}
					err = renderConstraint(f.Constraint, f.Type, i)
					if err != nil {
						return "", err
					}
				}
			}
		}
		if len(cluster.Attributes) > 0 {
			attributes := c.CreateElement("attributes")
			for _, a := range cluster.Attributes {
				ax := attributes.CreateElement("attribute")
				ax.CreateAttr("id", a.ID.HexString())
				ax.CreateAttr("name", a.Name)
				renderDataType(a, ax)
				if len(a.Default) > 0 {
					ax.CreateAttr("default", a.Default)
				}
				renderAccess(ax, a)
				renderQuality(ax, a)
				err = renderConformanceString(a.Conformance, ax)
				if err != nil {
					return "", err
				}

				err = renderConstraint(a.Constraint, a.Type, ax)
				if err != nil {
					return "", err
				}

			}
		}
		if len(cluster.Commands) > 0 {
			commands := c.CreateElement("commands")
			for _, cmd := range cluster.Commands {
				cx := commands.CreateElement("command")
				cx.CreateAttr("id", cmd.ID.ShortHexString())
				cx.CreateAttr("name", cmd.Name)
				if cmd.Access.Invoke != matter.PrivilegeUnknown {
					a := cx.CreateElement("access")
					a.CreateAttr("invokePrivilege", strings.ToLower(matter.PrivilegeNames[cmd.Access.Invoke]))
					if cmd.Access.FabricScoped {
						a.CreateAttr("fabricScoped", "true")
					}
					if cmd.Access.Timed {
						a.CreateAttr("timed", "true")
					}
				}
				if cmd.Response != "" {
					cx.CreateAttr("response", cmd.Response)
				}
				err = renderConformanceString(cmd.Conformance, cx)
				if err != nil {
					return "", err
				}

				for _, f := range cmd.Fields {
					if !f.ID.Valid() {
						continue
					}
					i := cx.CreateElement("field")
					i.CreateAttr("id", f.ID.IntString())
					i.CreateAttr("name", f.Name)
					renderDataType(f, i)
					if len(f.Default) > 0 {
						i.CreateAttr("default", f.Default)
					}
					err = renderConformanceString(f.Conformance, i)
					if err != nil {
						return "", err
					}

					err = renderConstraint(f.Constraint, f.Type, i)
					if err != nil {
						return "", err
					}

				}
			}
		}
	}
	x.Indent(2)

	var b bytes.Buffer
	x.WriteTo(&b)
	output = b.String()
	return
}

func renderDataType(f *matter.Field, i *etree.Element) {
	if f.Type != nil {
		if !f.Type.IsArray {
			i.CreateAttr("type", f.Type.Name)
		} else {
			i.CreateAttr("type", "list")
			e := i.CreateElement("entry")
			e.CreateAttr("type", f.Type.Name)
			if lc, ok := f.Constraint.(*constraint.ListConstraint); ok {
				renderConstraint(lc.EntryConstraint, f.Type, e)
			}
		}
	}
}

func renderAccess(ax *etree.Element, a *matter.Field) {
	acx := ax.CreateElement("access")
	if a.Access.Read != matter.PrivilegeUnknown {
		if a.Access.Read == matter.PrivilegeView {
			acx.CreateAttr("read", "true")
		} else {
			acx.CreateAttr("read", "optional")
		}
	}
	if a.Access.Write != matter.PrivilegeUnknown {
		if a.Access.Write == matter.PrivilegeOperate {
			acx.CreateAttr("write", "true")
		} else {
			acx.CreateAttr("write", "optional")
		}
	}
	if a.Access.Read != matter.PrivilegeUnknown {
		acx.CreateAttr("readPrivilege", strings.ToLower(matter.PrivilegeNames[a.Access.Read]))
	}
	if a.Access.Write != matter.PrivilegeUnknown {
		acx.CreateAttr("writePrivilege", strings.ToLower(matter.PrivilegeNames[a.Access.Write]))
	}
}

func renderQuality(parent *etree.Element, a *matter.Field) {
	changeOmitted := a.Quality.Has(matter.QualityChangedOmitted)
	nullable := a.Quality.Has(matter.QualityNullable)
	scene := a.Quality.Has(matter.QualityScene)
	fixed := a.Quality.Has(matter.QualityFixed)
	nonvolatile := a.Quality.Has(matter.QualityNonVolatile)
	reportable := a.Quality.Has(matter.QualityReportable)
	if !changeOmitted && !nullable && !scene && !fixed && !nonvolatile && !reportable {
		return
	}
	qx := parent.CreateElement("quality")
	qx.CreateAttr("changeOmitted", strconv.FormatBool(changeOmitted))
	qx.CreateAttr("nullable", strconv.FormatBool(nullable))
	qx.CreateAttr("scene", strconv.FormatBool(scene))
	if fixed {
		qx.CreateAttr("persistence", "fixed")
	} else if nonvolatile {
		qx.CreateAttr("persistence", "nonVolatile")
	} else {
		qx.CreateAttr("persistence", "volatile")
	}
	qx.CreateAttr("reportable", strconv.FormatBool(reportable))
}

var license = `
Copyright (C) Connectivity Standards Alliance (2021). All rights reserved.
The information within this document is the property of the Connectivity
Standards Alliance and its use and disclosure are restricted, except as
expressly set forth herein.

Connectivity Standards Alliance hereby grants you a fully-paid, non-exclusive,
nontransferable, worldwide, limited and revocable license (without the right to
sublicense), under Connectivity Standards Alliance's applicable copyright
rights, to view, download, save, reproduce and use the document solely for your
own internal purposes and in accordance with the terms of the license set forth
herein. This license does not authorize you to, and you expressly warrant that
you shall not: (a) permit others (outside your organization) to use this
document; (b) post or publish this document; (c) modify, adapt, translate, or
otherwise change this document in any manner or create any derivative work
based on this document; (d) remove or modify any notice or label on this
document, including this Copyright Notice, License and Disclaimer. The
Connectivity Standards Alliance does not grant you any license hereunder other
than as expressly stated herein.

Elements of this document may be subject to third party intellectual property
rights, including without limitation, patent, copyright or trademark rights,
and any such third party may or may not be a member of the Connectivity
Standards Alliance. Connectivity Standards Alliance members grant other
Connectivity Standards Alliance members certain intellectual property rights as
set forth in the Connectivity Standards Alliance IPR Policy. Connectivity
Standards Alliance members do not grant you any rights under this license. The
Connectivity Standards Alliance is not responsible for, and shall not be held
responsible in any manner for, identifying or failing to identify any or all
such third party intellectual property rights. Please visit www.csa-iot.org for
more information on how to become a member of the Connectivity Standards
Alliance.

This document and the information contained herein are provided on an “AS IS”
basis and the Connectivity Standards Alliance DISCLAIMS ALL WARRANTIES EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO (A) ANY WARRANTY THAT THE USE OF THE
INFORMATION HEREIN WILL NOT INFRINGE ANY RIGHTS OF THIRD PARTIES (INCLUDING
WITHOUT LIMITATION ANY INTELLECTUAL PROPERTY RIGHTS INCLUDING PATENT, COPYRIGHT
OR TRADEMARK RIGHTS); OR (B) ANY IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE, TITLE OR NONINFRINGEMENT. IN NO EVENT WILL THE
CONNECTIVITY STANDARDS ALLIANCE BE LIABLE FOR ANY LOSS OF PROFITS, LOSS OF
BUSINESS, LOSS OF USE OF DATA, INTERRUPTION OF BUSINESS, OR FOR ANY OTHER
DIRECT, INDIRECT, SPECIAL OR EXEMPLARY, INCIDENTAL, PUNITIVE OR CONSEQUENTIAL
DAMAGES OF ANY KIND, IN CONTRACT OR IN TORT, IN CONNECTION WITH THIS DOCUMENT
OR THE INFORMATION CONTAINED HEREIN, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH
LOSS OR DAMAGE.

All company, brand and product names in this document may be trademarks that
are the sole property of their respective owners.

This notice and disclaimer must be included on all copies of this document.

Connectivity Standards Alliance
508 Second Street, Suite 206
Davis, CA 95616, USA
`

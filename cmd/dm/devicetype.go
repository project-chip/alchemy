package dm

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
)

func renderDeviceTypes(cxt context.Context, sdkRoot string, deviceTypes []*ascii.Doc, filesOptions files.Options) error {
	var lock sync.Mutex
	outputs := make(map[string]string)
	err := files.ProcessDocs(cxt, deviceTypes, func(cxt context.Context, doc *ascii.Doc, index, total int) error {
		slog.Info("Device Type doc", "name", doc.Path)

		entities, err := doc.Entities()
		if err != nil {
			slog.ErrorContext(cxt, "error converting doc to entities", "doc", doc.Path, "error", err)
			return nil
		}
		var deviceTypes []*matter.DeviceType
		for _, m := range entities {
			slog.Debug("entity", "type", m)
			switch m := m.(type) {
			case *matter.DeviceType:
				deviceTypes = append(deviceTypes, m)
			}
		}
		s, err := renderDeviceType(cxt, deviceTypes)
		if err != nil {
			slog.ErrorContext(cxt, "error rendering entities", "doc", doc.Path, "error", err)
			return nil
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
			newPath := filepath.Join(sdkRoot, fmt.Sprintf("/data_model/device_types/%s.xml", strings.TrimSuffix(path, filepath.Ext(path))))
			result, err = patchLicense(result, newPath)
			if err != nil {
				return fmt.Errorf("error patching license for %s: %w", newPath, err)
			}
			err = os.WriteFile(newPath, []byte(result), os.ModeAppend|0644)
			if err != nil {
				return fmt.Errorf("error writing %s: %w", newPath, err)
			}
		}
	}
	return nil
}

func renderDeviceType(cxt context.Context, deviceTypes []*matter.DeviceType) (output string, err error) {
	x := etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(getLicense())
	for _, deviceType := range deviceTypes {
		c := x.CreateElement("deviceType")
		c.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
		c.CreateAttr("xsi:schemaLocation", "types types.xsd devicetype devicetype.xsd")
		c.CreateAttr("id", deviceType.ID.HexString())
		c.CreateAttr("name", deviceType.Name)

		revs := c.CreateElement("revisionHistory")
		var latestRev uint64 = 0
		for _, r := range deviceType.Revisions {
			id := matter.ParseNumber(r.Number)
			if id.Valid() {
				rev := revs.CreateElement("revision")
				rev.CreateAttr("revision", id.IntString())
				rev.CreateAttr("summary", r.Description)
				latestRev = max(id.Value(), latestRev)
			}
		}
		c.CreateAttr("revision", strconv.FormatUint(latestRev, 10))
		class := c.CreateElement("classification")
		if deviceType.Superset != "" {
			class.CreateAttr("superset", deviceType.Superset)
		}
		class.CreateAttr("class", strings.ToLower(deviceType.Class))
		class.CreateAttr("scope", strings.ToLower(deviceType.Scope))

		conditions := c.CreateElement("conditions")
		if len(deviceType.Conditions) > 0 {
			for _, condition := range deviceType.Conditions {
				cx := conditions.CreateElement("condition")
				cx.CreateAttr("name", condition.Feature)
				cx.CreateAttr("summary", condition.Description)
			}
		}

		if len(deviceType.ClusterRequirements) > 0 {
			cx := c.CreateElement("clusters")
			reqs := make([]*matter.ClusterRequirement, len(deviceType.ClusterRequirements))
			copy(reqs, deviceType.ClusterRequirements)
			slices.SortFunc(reqs, func(a, b *matter.ClusterRequirement) int {
				return a.ID.Compare(b.ID)
			})
			for _, cr := range reqs {
				clx := cx.CreateElement("cluster")
				clx.CreateAttr("id", cr.ID.HexString())
				clx.CreateAttr("name", cr.ClusterName)
				switch cr.Interface {
				case matter.InterfaceClient:
					clx.CreateAttr("side", "client")
				case matter.InterfaceServer:
					clx.CreateAttr("side", "server")
				}
				err = renderConformanceString(deviceType, cr.Conformance, clx)
				if err != nil {
					return
				}
				err = renderElementRequirements(deviceType, cr, clx)
				if err != nil {
					return
				}

			}
		}
	}
	x.Indent(2)

	var b bytes.Buffer
	_, err = x.WriteTo(&b)
	output = b.String()
	return
}

func renderElementRequirements(deviceType *matter.DeviceType, cr *matter.ClusterRequirement, clx *etree.Element) (err error) {
	erMap := make(map[types.EntityType][]*matter.ElementRequirement)
	for _, er := range deviceType.ElementRequirements {
		if er.ID.Equals(cr.ID) {
			erMap[er.Element] = append(erMap[er.Element], er)
		}
	}
	for _, entity := range []types.EntityType{types.EntityTypeFeature, types.EntityTypeAttribute, types.EntityTypeCommand, types.EntityTypeEvent} {
		ers, ok := erMap[entity]
		if !ok || len(ers) == 0 {
			continue
		}
		var erx *etree.Element
		switch entity {
		case types.EntityTypeAttribute:
			erx = clx.CreateElement("attributes")
		case types.EntityTypeFeature:
			erx = clx.CreateElement("features")
		case types.EntityTypeEvent:
			erx = clx.CreateElement("events")
		case types.EntityTypeCommand:
			erx = clx.CreateElement("commands")
		}
		for _, er := range ers {
			switch entity {
			case types.EntityTypeAttribute:
				err = renderAttributeRequirement(deviceType, er, erx)
			case types.EntityTypeCommand:
				ex := erx.CreateElement("command")
				var code string
				if er.Cluster != nil {
					for _, cmd := range er.Cluster.Commands {
						if cmd.ID.Equals(er.ID) {
							code = cmd.ID.HexString()
							break
						}
					}
				}
				if code != "" {
					ex.CreateAttr("code", code)

				}
				ex.CreateAttr("name", er.Name)
				err = renderConformanceString(deviceType, er.Conformance, ex)
				if err != nil {
					return
				}
			case types.EntityTypeFeature:
				ex := erx.CreateElement("feature")
				var code string

				featureCode := er.Name
				if strings.ContainsRune(featureCode, ' ') {
					featureCode = matter.Case(featureCode)
				}

				if er.Cluster != nil {
					for _, b := range er.Cluster.Features.Bits {
						f := b.(*matter.Feature)
						if f.Code == featureCode {
							code = f.Code
							break
						}
					}
				}
				ex.CreateAttr("code", code)
				ex.CreateAttr("name", er.Name)
				err = renderConformanceString(deviceType, er.Conformance, ex)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func renderAttributeRequirement(deviceType *matter.DeviceType, er *matter.ElementRequirement, parent *etree.Element) (err error) {
	var code string
	var attribute *matter.Field
	var dataType *types.DataType
	if er.Cluster != nil {
		for _, a := range er.Cluster.Attributes {
			if a.ID.Equals(er.ID) {
				attribute = a
				dataType = a.Type
				break
			}
		}
	}
	if attribute != nil {
		code = attribute.ID.HexString()
	}
	ex := parent.CreateElement("attribute")
	ex.CreateAttr("code", code)
	ex.CreateAttr("name", er.Name)

	renderAttributeAccess(ex, er.Access)
	renderQuality(ex, er.Quality, matter.QualityAll^matter.QualitySingleton)
	err = renderConformanceString(deviceType, er.Conformance, ex)
	if err != nil {
		return
	}
	err = renderConstraint(er.Constraint, dataType, ex)
	return
}

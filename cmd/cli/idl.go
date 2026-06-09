package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/idl"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/sdk"
)

type IDL struct {
	IDLRegen           IDLRegen              `cmd:"" name:"regen"`
	IDLControllerClusters IDLControllerClusters `cmd:"" name:"controller-clusters"`
}

type IDLRegen struct {
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
	spec.ParserOptions         `embed:""`
	spec.FilterOptions         `embed:""`
	sdk.SDKOptions             `embed:""`
	SuppressProvisional        string `name:"suppress-provisional" help:"Suppress rendering of provisional elements" default:"none" enum:"none,all,keep-existing"`
}

func (z *IDLRegen) Run(cc *Context) (err error) {

	var specification *spec.Specification
	specification, _, err = spec.Parse(cc, z.ParserOptions, z.ProcessingOptions, []spec.BuilderOption{spec.PatchForSdk(true)}, z.ASCIIDocAttributes.ToList())
	if err != nil {
		return
	}

	err = sdk.ApplyErrata(specification, sdk.WithSkipSharedEntities(true))
	if err != nil {
		return
	}

	zapTargeter := idl.Targeter(z.SdkRoot)

	var zapPaths pipeline.Paths
	zapPaths, err = pipeline.Start(cc, zapTargeter)
	if err != nil {
		return
	}

	var reader idl.Reader
	reader, err = idl.NewReader()
	if err != nil {
		return
	}

	var zapFiles pipeline.Map[string, *pipeline.Data[*idl.File]]
	zapFiles, err = pipeline.Parallel(cc, z.ProcessingOptions, reader, zapPaths)
	if err != nil {
		return
	}

	var renderer idl.IdlRenderer
	renderer, err = idl.NewIdlRenderer(specification)
	if err != nil {
		return
	}
	renderer.SuppressProvisional = z.SuppressProvisional

	var matterFiles pipeline.StringSet
	matterFiles, err = pipeline.Parallel(cc, z.ProcessingOptions, renderer, zapFiles)
	if err != nil {
		return
	}

	writer := files.NewWriter[string]("Writing .matter files", z.OutputOptions)
	err = writer.Write(cc, matterFiles, z.ProcessingOptions)

	return
}

type IDLControllerClusters struct {
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
	spec.ParserOptions         `embed:""`
	Output                     string `name:"output" placeholder:"path" help:"Output file or directory for controller-clusters.matter" optional:"" default:"controller-clusters.matter"`
	SuppressProvisional        string `name:"suppress-provisional" help:"Suppress rendering of provisional elements" default:"none" enum:"none,all,keep-existing"`
}

func (z *IDLControllerClusters) Run(cc *Context) (err error) {
	var specification *spec.Specification
	specification, _, err = spec.Parse(cc, z.ParserOptions, z.ProcessingOptions, []spec.BuilderOption{spec.PatchForSdk(true)}, z.ASCIIDocAttributes.ToList())
	if err != nil {
		return
	}

	err = sdk.ApplyErrata(specification, sdk.WithSkipSharedEntities(true))
	if err != nil {
		return
	}

	var clusterRefs []idl.ClusterRef
	for code, cluster := range specification.ClustersByID {
		clusterRefs = append(clusterRefs, idl.ClusterRef{
			Code: int(code),
			Name: cluster.Name,
			Side: "server",
		})
	}

	slices.SortFunc(clusterRefs, func(a, b idl.ClusterRef) int {
		return a.Code - b.Code
	})

	if _, ok := specification.DeviceTypesByID[22]; !ok {
		return fmt.Errorf("device type 22 (Root Node) not found in specification")
	}
	validDeviceTypeCode := uint64(22)

	syntheticFile := &idl.File{
		EndpointTypes: []idl.EndpointType{
			{
				ID:             0,
				Name:           "Synthetic Endpoint",
				DeviceTypeCode: int(validDeviceTypeCode),
				Clusters:       clusterRefs,
			},
		},
		Endpoints: []idl.JSONEndpoint{
			{
				EndpointId:        0,
				EndpointTypeIndex: 0,
			},
		},
	}

	renderer, err := idl.NewIdlRenderer(specification)
	if err != nil {
		return err
	}
	renderer.SuppressEndpoints = true
	renderer.SuppressProvisional = z.SuppressProvisional

	var zapPath string
	outPath := filepath.Clean(z.Output)
	isDir := false
	if fi, err := os.Stat(outPath); err == nil {
		isDir = fi.IsDir()
	} else if strings.HasSuffix(z.Output, "/") || strings.HasSuffix(z.Output, string(filepath.Separator)) {
		isDir = true
	}

	if isDir {
		zapPath = filepath.Join(outPath, "controller-clusters.zap")
	} else {
		ext := filepath.Ext(outPath)
		if ext == "" {
			zapPath = outPath + ".zap"
		} else {
			zapPath = strings.TrimSuffix(outPath, ext) + ".zap"
		}
	}
	zapData := pipeline.NewData[*idl.File](zapPath, syntheticFile)

	outputs, _, err := renderer.Process(cc, zapData, 0, 1)
	if err != nil {
		return err
	}

	if len(outputs) == 0 {
		return fmt.Errorf("no output generated")
	}

	writer := files.NewWriter[string]("Writing controller-clusters.matter", z.OutputOptions)

	outputSet := pipeline.StringSet(pipeline.NewMap[string, *pipeline.Data[string]]())
	outputSet.Store(outputs[0].Path, outputs[0])

	return writer.Write(cc, outputSet, z.ProcessingOptions)
}

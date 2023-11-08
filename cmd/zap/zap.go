package zap

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"
	"sync"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/render/zcl"
	"github.com/hasty/alchemy/zap"
	zaprender "github.com/hasty/alchemy/zap/render"
	"github.com/iancoleman/orderedmap"
	"github.com/iancoleman/strcase"
)

var selfClosingTags = regexp.MustCompile("></[^>]+>")

func Migrate(cxt context.Context, specRoot string, zclRoot string, filesOptions files.Options, asciiSettings []configuration.Setting) error {

	specPaths, err := getSpecPaths(specRoot)
	if err != nil {
		return err
	}

	var lock sync.Mutex

	var docs []*ascii.Doc

	asciiSettings = append(ascii.GithubSettings, asciiSettings...)

	files.Process(cxt, specPaths, func(cxt context.Context, file string, index, total int) error {
		if strings.HasSuffix(file, "-draft.adoc") {
			return nil
		}
		doc, err := ascii.Open(file, asciiSettings...)
		if err != nil {
			return err
		}
		lock.Lock()
		docs = append(docs, doc)
		lock.Unlock()
		slog.InfoContext(cxt, "Parsed adoc", "file", file)
		return nil
	}, filesOptions)

	slog.InfoContext(cxt, "Building spec tree...")
	ascii.BuildTree(docs)

	var appClusters, appClusterIndexes []*ascii.Doc
	for _, d := range docs {
		docType, err := d.DocType()
		if err != nil {
			return err
		}
		switch docType {
		case matter.DocTypeAppCluster:
			appClusters = append(appClusters, d)
		case matter.DocTypeAppClusterIndex:
			appClusterIndexes = append(appClusterIndexes, d)

		}
	}

	files.ProcessDocs(cxt, appClusterIndexes, func(cxt context.Context, doc *ascii.Doc, index, total int) error {
		slog.InfoContext(cxt, "Building spec tree...")

		top := parse.FindFirst[*ascii.Section](doc.Elements)
		if top == nil {
			return nil
		}
		doc.Domain = zap.StringToDomain(top.Name)
		slog.InfoContext(cxt, "Assigned domain", "file", top.Name, "domain", doc.Domain)
		return nil
	}, filesOptions)

	outputs := make(map[string]*zcl.Result)
	var provisionalZclFiles []string
	files.ProcessDocs(cxt, appClusters, func(cxt context.Context, doc *ascii.Doc, index, total int) error {
		path := doc.Path
		newFile := filepath.Base(path)
		newFile = zap.ZAPName(strings.TrimSuffix(newFile, filepath.Ext(path)))
		newFile = strcase.ToKebab(newFile)
		newPath := filepath.Join(zclRoot, "app/zap-templates/zcl/data-model/chip", newFile+".xml")

		existing, err := os.ReadFile(newPath)
		if errors.Is(err, os.ErrNotExist) {
			slog.InfoContext(cxt, "Rendering new ZAP template", "from", path, "to", newPath, "index", index, "count", total)
			var result *zcl.Result
			result, err = zcl.Render(cxt, doc)
			if err != nil {
				err = fmt.Errorf("failed rendering %s: %w", path, err)
				return err
			}
			lock.Lock()
			outputs[newPath] = result
			provisionalZclFiles = append(provisionalZclFiles, filepath.Base(newPath))
			lock.Unlock()
			slog.InfoContext(cxt, "Rendered new ZAP template", "from", path, "to", newPath, "index", index, "count", total)
		} else if err != nil {
			return err
		} else {
			slog.InfoContext(cxt, "Rendering existing ZAP template", "from", path, "to", newPath, "index", index, "count", total)
			var buf bytes.Buffer
			models, err := doc.ToModel()
			if err != nil {
				return err
			}
			var clusters []any
			for _, m := range models {
				switch m := m.(type) {
				case *matter.Cluster:
					if m.Hierarchy == "Base" && m.Name != "Mode Base" && m.Name != "Scenes" && m.ID.Valid() {
						clusters = append(clusters, m)
					}
				}
			}
			if len(clusters) == 0 {
				slog.InfoContext(cxt, "Skipped existing ZAP template", "from", path, "to", newPath, "index", index, "count", total)
				return nil
			}
			err = zaprender.Render(doc, bytes.NewReader(existing), &buf, clusters)
			if err != nil {
				return err
			}
			out := selfClosingTags.ReplaceAllString(buf.String(), "/>") // Lame limitation of Go's XML encoder
			lock.Lock()
			if false {
				outputs[newPath] = &zcl.Result{ZCL: out, Doc: doc, Models: models}
			}
			lock.Unlock()
			slog.InfoContext(cxt, "Rendered existing ZAP template", "from", path, "to", newPath, "index", index, "count", total)
		}
		return nil
	}, filesOptions)

	if err != nil {
		return err
	}
	for path, result := range outputs {
		if len(result.Models) == 0 {
			continue
		}

		if !filesOptions.DryRun {
			err = os.WriteFile(path, []byte(result.ZCL), os.ModeAppend|0644)
			if err != nil {
				return err
			}
		}
	}

	if !filesOptions.DryRun {
		err = patchZCLJson(zclRoot, provisionalZclFiles)
	}

	/*appClusterPaths, appClusterIndexPaths, err := splitAppClusterDocs(specRoot)
	if err != nil {
		return err
	}

	var domains map[string]matter.Domain
	_, domains, err = findIndexes(cxt, appClusterIndexPaths, asciiSettings, filesOptions)

	if err != nil {
		return err
	}

	var provisionalZclFiles []string

	outputs := make(map[string]*zcl.Result)

	err = files.Process(cxt, appClusterPaths, func(cxt context.Context, file string, index int, total int) error {

		doc, err := ascii.Open(file, append(ascii.GithubSettings, asciiSettings...)...)
		if err != nil {
			return err
		}
		if domain, ok := domains[file]; ok {
			doc.Domain = domain
		} else {
			doc.Domain = matter.DomainCHIP
		}

		newFile := filepath.Base(file)
		newFile = zap.ZAPName(strings.TrimSuffix(newFile, filepath.Ext(file)))
		newFile = strcase.ToKebab(newFile)
		newPath := filepath.Join(zclRoot, "app/zap-templates/zcl/data-model/chip", newFile+".xml")

		existing, err := os.ReadFile(newPath)
		if errors.Is(err, os.ErrNotExist) {
			var result *zcl.Result
			result, err = zcl.Render(cxt, doc)
			if err != nil {
				err = fmt.Errorf("failed rendering %s: %w", file, err)
				return err
			}
			lock.Lock()
			outputs[newPath] = result
			lock.Unlock()
			fmt.Fprintf(os.Stderr, "ZCL'd %s to %s (%d of %d)...\n", file, newPath, index, total)
			provisionalZclFiles = append(provisionalZclFiles, filepath.Base(newPath))

		} else if err != nil {
			return err
		} else {
			fmt.Fprintf(os.Stderr, "ZCLing existing %s to %s (%d of %d)...\n", file, newPath, index, total)
			var buf bytes.Buffer
			models, err := doc.ToModel()
			if err != nil {
				return err
			}
			var clusters []any
			for _, m := range models {
				switch m := m.(type) {
				case *matter.Cluster:
					if m.Hierarchy == "Base" && m.Name != "Mode Base" && m.Name != "Scenes" && m.ID.Valid() {
						clusters = append(clusters, m)
					}
				}
			}
			if len(clusters) == 0 {
				fmt.Fprintf(os.Stderr, "Skipped ZCLing %s to %s (%d of %d)...\n", file, newPath, index, total)
				return nil
			}
			err = zaprender.Render(doc, bytes.NewReader(existing), &buf, clusters)
			if err != nil {
				return err
			}
			out := selfClosingTags.ReplaceAllString(buf.String(), "/>") // Lame limitation of Go's XML encoder
			lock.Lock()
			if false {
				outputs[newPath] = &zcl.Result{ZCL: out, Doc: doc, Models: models}
			}
			lock.Unlock()
			fmt.Fprintf(os.Stderr, "ZCL'd existing %s to %s (%d of %d)...\n", file, newPath, index, total)
		}

		return nil
	}, filesOptions)
	if err != nil {
		return err
	}
	for path, result := range outputs {
		if len(result.Models) == 0 {
			continue
		}

		if !filesOptions.DryRun {
			err = os.WriteFile(path, []byte(result.ZCL), os.ModeAppend|0644)
			if err != nil {
				return err
			}
		}
	}

	if !filesOptions.DryRun {
		err = patchZCLJson(zclRoot, provisionalZclFiles)
	}

	/*zclJSON := make(map[string]any)
	err = json.Unmarshal(zclJSONBytes, &zclJSON)
	if err != nil {
		return err
	}

	xm, ok := zclJSON["xmlFile"]
	if !ok {
		return fmt.Errorf("could not find xmlFile in zcl.json")
	}
	xms, ok := xm.([]interface{})
	if !ok {
		return fmt.Errorf("could not cast xmlFile in zcl.json; %T", xm)
	}
	paths := make([]string, 0, len(xms)+len(provisionalZclFiles))
	for _, p := range xms {
		paths = append(paths, p.(string))
	}
	for _, p := range provisionalZclFiles {
		paths = append(paths, p)
	}
	slices.Sort(paths)
	slices.Compact(paths)
	zclJSON["xmlFile"] = paths

	zclJSONBytes, err = json.Marshal(zclJSON)
	if err != nil {
		return err
	}
	err = os.WriteFile(zclJSONPath, zclJSONBytes, os.ModeAppend|0644)*/
	if err != nil {
		return err
	}

	return nil
}

func findIndexes(cxt context.Context, appClusterIndexPaths []string, asciiSettings []configuration.Setting, filesOptions files.Options) (indexes map[string]*ascii.Doc, domains map[string]matter.Domain, err error) {
	indexes = make(map[string]*ascii.Doc)
	domains = make(map[string]matter.Domain)

	err = files.Process(cxt, appClusterIndexPaths, func(cxt context.Context, file string, index, total int) error {
		fmt.Fprintf(os.Stderr, "ZCLing index %s (%d of %d)...\n", file, index, total)
		doc, err := ascii.Open(file, append(ascii.GithubSettings, asciiSettings...)...)
		if err != nil {
			return err
		}

		top := parse.FindFirst[*ascii.Section](doc.Elements)
		if top == nil {
			return nil
		}

		domain := zap.StringToDomain(top.Name)
		fmt.Printf("Index Name: %s domain: %v\n", top.Name, domain)

		parse.Search[*types.Section](top.Base.Elements, func(t *types.Section) bool {
			link := parse.FindFirst[*types.InlineLink](t.Title)
			if link != nil {
				linkPath, ok := link.Location.Path.(string)
				if ok {
					linkPath = filepath.Join(filepath.Dir(file), linkPath)
					fmt.Printf("file link %s -> %s\n", file, linkPath)
					domains[linkPath] = domain
					indexes[linkPath] = doc
				}
			}
			return false
		})
		return nil
	}, filesOptions)
	return
}

func splitAppClusterDocs(specRoot string) (appClusterPaths []string, appClusterIndexPaths []string, err error) {
	err = filepath.WalkDir(specRoot, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".adoc" {
			docType, err := ascii.GetDocType(path)
			if err != nil {
				return err
			}
			switch docType {
			case matter.DocTypeAppCluster:
				appClusterPaths = append(appClusterPaths, path)
			case matter.DocTypeAppClusterIndex:
				appClusterIndexPaths = append(appClusterIndexPaths, path)
			}
		}
		return nil
	})
	return
}

func getSpecPaths(specRoot string) (paths []string, err error) {
	err = filepath.WalkDir(specRoot, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".adoc" {
			paths = append(paths, path)
		}
		return nil
	})
	return
}

func loadSpecDocs(specRoot string, asciiSettings []configuration.Setting) (docs []*ascii.Doc, err error) {
	err = filepath.WalkDir(specRoot, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".adoc" {
			var doc *ascii.Doc
			doc, err = ascii.Open(path, append(ascii.GithubSettings, asciiSettings...)...)
			if err != nil {
				return err
			}
			docs = append(docs, doc)
		}
		return nil
	})
	return
}

func patchZCLJson(zclRoot string, files []string) error {
	zclJSONPath := path.Join(zclRoot, "app/zap-templates/zcl/zcl.json")
	zclJSONBytes, err := os.ReadFile(zclJSONPath)
	if err != nil {
		return err
	}

	o := orderedmap.New()
	err = json.Unmarshal(zclJSONBytes, &o)
	if err != nil {
		return err
	}
	val, ok := o.Get("xmlFile")
	if !ok {
		return fmt.Errorf("no xmlField field in zcl.json")
	}
	is, ok := val.([]interface{})
	if !ok {
		return fmt.Errorf("xmlField not a string array in zcl.json; %T", val)
	}
	xmls := make([]string, 0, len(is)+len(files))
	for _, i := range is {
		if s, ok := i.(string); ok {
			xmls = append(xmls, s)
		}
	}
	xmls = append(xmls, files...)
	sort.Slice(xmls, func(i, j int) bool {
		a := xmls[i]
		b := xmls[j]
		if a == "access-control-definitions.xml" {
			return true
		}
		if b == "access-control-definitions.xml" {
			return false
		}
		if a < b {
			return true
		}
		return false
	})
	//slices.Sort(xmls)
	slices.Compact(xmls)
	fmt.Printf("xml: %v\n", xmls)
	o.Set("xmlFile", xmls)

	zclJSONBytes, err = json.MarshalIndent(o, "", "    ")
	if err != nil {
		return err
	}

	/*v := gjson.GetBytes(zclJSONBytes, "xmlFile").Array()

	xmls := make([]string, 0, len(v)+len(files))
	for _, s := range v {
		xmls = append(xmls, s.Str)
	}
	xmls = append(xmls, files...)
	slices.Sort(xmls)
	slices.Compact(xmls)
	fmt.Printf("xml: %v\n", xmls)
	zclJSONBytes, err = sjson.SetBytes(zclJSONBytes, "xmlFile", xmls)
	if err != nil {
		return err
	}*/
	err = os.WriteFile(zclJSONPath, []byte(zclJSONBytes), os.ModeAppend|0644)
	return err
}

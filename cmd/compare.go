package cmd

import (
	"context"

	"github.com/hasty/matterfmt/compare"
)

type zclConparer struct {
	processor
	asciiParser

	serial bool
	dryRun bool
}

func Compare(cxt context.Context, specRoot string, zclRoot string, options ...Option) error {
	z := &zclConparer{}
	for _, o := range options {
		err := o(z)
		if err != nil {
			return err
		}
	}
	return compare.Compare(cxt, specRoot, zclRoot, z.settings)
}

/*
func (z *zclConparer) run(cxt context.Context, specRoot string, zclRoot string) error {

	return
	zapPath := filepath.Join(zclRoot, "app/zap-templates/zcl/data-model/chip/*.xml")
	xmlFiles, err := filepath.Glob(zapPath)
	if err != nil {
		return err
	}
	zapModels := make(map[string][]any)
	for _, f := range xmlFiles {
		fmt.Printf("ZAP file: %s\n", f)

		file, err := os.Open(f)
		if err != nil {
			return err
		}
		defer file.Close()
		var models []any
		models, err = zap.Parse(file)
		if err != nil {
			return err
		}
		zapModels[f] = models
	}
	var appClusterPaths []string
	var appClusterIndexPaths []string
	filepath.WalkDir(specRoot, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".adoc" {
			docType, e := ascii.GetDocType(path)
			if e != nil {
				return nil
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

	domains := make(map[string]matter.Domain)

	err = z.processFiles(cxt, appClusterIndexPaths, func(cxt context.Context, file string, index, total int) error {
		fmt.Fprintf(os.Stderr, "ZCLing index %s (%d of %d)...\n", file, index, total)
		doc, err := ascii.Open(file, z.settings...)
		if err != nil {
			return err
		}

		top := parse.FindFirst[*ascii.Section](doc.Elements)
		if top == nil {
			return nil
		}

		domain := getDomain(top.Name)
		fmt.Printf("Index Name: %s domain: %v\n", top.Name, domain)

		parse.Search[*types.Section](top.Base.Elements, func(t *types.Section) bool {
			link := parse.FindFirst[*types.InlineLink](t.Title)
			if link != nil {
				linkPath, ok := link.Location.Path.(string)
				if ok {
					linkPath = filepath.Join(filepath.Dir(file), linkPath)
					fmt.Printf("file link %s -> %s\n", file, linkPath)
					domains[linkPath] = domain
				}
			}
			return false
		})
		return nil
	})

	if err != nil {
		return err
	}

	specModels := make(map[string][]any)
	var lock sync.Mutex
	err = z.processFiles(cxt, appClusterPaths, func(cxt context.Context, file string, index int, total int) error {
		doc, err := ascii.Open(file, z.settings...)
		if err != nil {
			return err
		}
		if domain, ok := domains[file]; ok {
			doc.Domain = domain
		} else {
			doc.Domain = matter.DomainCHIP
		}

		models, err := doc.ToModel()
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stderr, "ZCL'd %s (%d of %d)...\n", file, index, total)

		newFile := filepath.Base(file)
		newFile = zap.ZAPName(strings.TrimSuffix(newFile, filepath.Ext(file)))
		newFile = strcase.ToKebab(newFile)

		newPath := filepath.Join(zclRoot, "app/zap-templates/zcl/data-model/chip", newFile+".xml")

		lock.Lock()
		specModels[newPath] = models
		lock.Unlock()
		return nil
	})
	if err != nil {
		return err
	}
	for path, sm := range specModels {
		zm, ok := zapModels[path]
		if !ok {
			fmt.Printf("path %s missing from ZAP models\n", path)
			continue
		}
		fmt.Printf("path %s found in ZAP models\n", path)

		specClusters := make(map[string]*matter.Cluster)
		for _, m := range sm {
			switch v := m.(type) {
			case *matter.Cluster:
				fmt.Printf("adding spec cluster: %s\n", v.ID)
				specClusters[strings.ToLower(v.ID)] = v
			default:
				fmt.Printf("unexpected spec model: %T\n", m)
			}
		}
		zapClusters := make(map[string]*matter.Cluster)
		for _, m := range zm {
			switch v := m.(type) {
			case *matter.Cluster:
				fmt.Printf("adding ZAP cluster: %s\n", v.ID)
				zapClusters[strings.ToLower(v.ID)] = v
			default:
				fmt.Printf("unexpected ZAP model: %T\n", m)
			}

		}
		delete(zapModels, path)
		for cid, sc := range specClusters {
			if zc, ok := zapClusters[cid]; ok {
				sc.Compare(zc)
				delete(zapClusters, cid)
			} else {
				fmt.Printf("ZAP cluster %s missing from %s; ", cid, path)
				for zid := range zapClusters {
					fmt.Printf("have %s,", zid)
				}
				fmt.Println()
			}
		}
		for cid := range zapClusters {
			fmt.Printf("Spec cluster %s missing from %s\n", cid, path)
		}
	}

	for path := range zapModels {
		fmt.Printf("path %s missing from Spec models\n", path)
	}
	return nil
}
*/

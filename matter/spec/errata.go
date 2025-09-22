package spec

import "github.com/project-chip/alchemy/errata"

func (library *Library) DiscoErrata(path string) *errata.Disco {
	return &library.Errata.Get(path).Disco
}

func (library *Library) SpecErrata(path string) *errata.Spec {
	return &library.Errata.Get(path).Spec
}

func (library *Library) SdkErrata(path string) *errata.SDK {
	return &library.Errata.Get(path).SDK
}

func applyErrata(spec *Specification, library *Library) {
	for _, doc := range library.Docs {
		se := library.SdkErrata(doc.Path.Relative)
		if se == nil {
			continue
		}
	}
}

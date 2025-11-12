package spec

import (
	"iter"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (library *Library) parseEntities(spec *Specification) iter.Seq2[*asciidoc.Document, any] {

	return func(yield func(*asciidoc.Document, any) bool) {
		var currentDomain string
		currentDomain = "General"
		var currentDoc *asciidoc.Document
		parse.Search(library.Root, library, library.Root, library.Children(library.Root), func(doc *asciidoc.Document, section *asciidoc.Section, parent asciidoc.ParentElement, index int) parse.SearchShould {
			var err error
			var skip bool

			if doc != currentDoc {
				if libraryDoc, ok := library.config.Documents[doc.Path.Relative]; ok {
					if libraryDoc.Domain != "" {
						currentDomain = libraryDoc.Domain
					}
				}
				currentDoc = doc
			}
			sectionType := library.SectionType(section)
			switch sectionType {
			case matter.SectionTop:
				dt, _ := library.DocType(doc)
				switch dt {
				case matter.DocTypeAppClusterIndex:
					//currentDomain = ParseDomain(library.SectionName(section))
				}
			case matter.SectionCluster:
				//slog.Info("parsing cluster in library doc", "path", library.Root.Path.Relative, "doc", doc.Path.Relative, "sectionName", library.SectionName(section), "sectionType", sectionType.String(), log.Path("source", section))
				var clusterOrGroup types.Entity
				clusterOrGroup, err = library.toClusters(spec, library, doc, section, currentDomain)
				if err == nil && !yield(doc, clusterOrGroup) {
					return parse.SearchShouldStop
				}
				skip = true
			case matter.SectionDataTypeBitmap:
				//slog.Info("parsing bitmap in library doc", "path", library.Root.Path.Relative, "doc", doc.Path.Relative, "sectionName", library.SectionName(section), "sectionType", sectionType.String(), log.Path("source", section))
				var bm *matter.Bitmap
				bm, err = library.toBitmap(library, doc, section, nil)
				if err == nil && !yield(doc, bm) {
					return parse.SearchShouldStop
				}
				skip = true
			case matter.SectionDataTypeEnum:
				//	slog.Info("parsing enum in library doc", "path", library.Root.Path.Relative, "doc", doc.Path.Relative, "sectionName", library.SectionName(section), "sectionType", sectionType.String(), log.Path("source", section))
				var s *matter.Enum
				s, err = library.toEnum(library, doc, section, nil)
				if err == nil && !yield(doc, s) {
					return parse.SearchShouldStop
				}
				skip = true
			case matter.SectionDataTypeStruct:
				var s *matter.Struct
				s, err = library.toStruct(spec, library, doc, section, nil)
				if err == nil && !yield(doc, s) {
					return parse.SearchShouldStop
				}
			case matter.SectionDataTypeDef:
				var td *matter.TypeDef
				td, err = library.toTypeDef(library, doc, section, nil)
				if err == nil && !yield(doc, td) {
					return parse.SearchShouldStop
				}
			case matter.SectionGlobalElements:
				var ges []types.Entity
				ges, err = library.toGlobalElements(spec, library, doc, section, nil)
				if err == nil {
					for _, dt := range ges {
						if !yield(doc, dt) {
							return parse.SearchShouldStop
						}
					}
				}
				skip = true

			case matter.SectionDeviceType:
				//	slog.Info("parsing device type in library doc", "path", library.Root.Path.Relative, "doc", doc.Path.Relative, "sectionName", library.SectionName(section), "sectionType", sectionType.String(), log.Path("source", section))
				var deviceTypes []*matter.DeviceType
				var docType matter.DocType
				docType, err = library.DocType(doc)
				if err == nil {
					switch docType {
					case matter.DocTypeBaseDeviceType:
						var baseDeviceType *matter.DeviceType
						baseDeviceType, err = library.toBaseDeviceType(library, section)
						if err == nil {
							deviceTypes = append(deviceTypes, baseDeviceType)
						}
					default:
						deviceTypes, err = library.toDeviceTypes(library, doc, section)
					}
				}
				if err == nil {
					for _, dt := range deviceTypes {
						if !yield(doc, dt) {
							return parse.SearchShouldStop
						}
					}
				}
				skip = true
			case matter.SectionNamespace:
				var ns *matter.Namespace
				ns, err = library.toNamespace(library, doc, section)
				if err == nil && !yield(doc, ns) {
					return parse.SearchShouldStop
				}
				skip = true
			case matter.SectionRevisionHistory,
				matter.SectionIntroduction,
				matter.SectionCommand,
				matter.SectionField,
				matter.SectionUnknown:

			default:
				slog.Debug("Unexpected section type", "st", sectionType.String(), log.Path("source", section))
			}
			if err != nil {
				if !yield(doc, err) {
					return parse.SearchShouldStop
				}
			}
			if skip {
				return parse.SearchShouldSkip
			}
			return parse.SearchShouldContinue
		})

	}

}

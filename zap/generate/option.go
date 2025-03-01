package generate

import "github.com/project-chip/alchemy/asciidoc"

type TemplateOption func(tg *TemplateGenerator)

func GenerateFeatureXML(generate bool) TemplateOption {
	return func(tg *TemplateGenerator) {
		tg.generateFeaturesXML = generate
	}
}

func GenerateConformanceXML(generate bool) TemplateOption {
	return func(tg *TemplateGenerator) {
		tg.generateConformanceXML = generate
	}
}

func SpecOrder(specOrder bool) TemplateOption {
	return func(tg *TemplateGenerator) {
		tg.specOrder = specOrder
	}
}

func AsciiAttributes(attributes []asciidoc.AttributeName) TemplateOption {
	return func(tg *TemplateGenerator) {
		tg.attributes = attributes
	}
}

func SpecRoot(specRoot string) TemplateOption {
	return func(tg *TemplateGenerator) {
		tg.specRoot = specRoot
	}
}

func ExtendedQuality(extendedQuality bool) TemplateOption {
	return func(tg *TemplateGenerator) {
		tg.generateExtendedQualityElement = extendedQuality
	}
}

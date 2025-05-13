package render

type TemplateOptions struct {
	FeatureXML             bool   `default:"true" aliases:"featureXML" help:"write new style feature XML" group:"ZAP:"`
	ConformanceXML         bool   `default:"true" aliases:"conformanceXML" help:"write new style conformance XML" group:"ZAP:"`
	EndpointCompositionXML bool   `default:"false" aliases:"endpointCompositionXML" help:"write new style endpoint composition XML" group:"ZAP:"`
	SpecOrder              bool   `default:"false" aliases:"specOrder" help:"write ZAP template XML in spec order" group:"ZAP:"`
	ExtendedQuality        bool   `default:"false" aliases:"extendedQuality" help:"write quality element with all qualities, suppressing redundant attributes" group:"ZAP:"`
	ProvisionalPolicy      string `enum:"none,loose,strict" default:"none" help:"enforce a provisional policy for generating ZAP XML" group:"ZAP:"`
}

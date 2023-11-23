package matter

type Feature struct {
	Bit         string      `json:"bit,omitempty"`
	Code        string      `json:"code,omitempty"`
	Name        string      `json:"name,omitempty"`
	Summary     string      `json:"summary,omitempty"`
	Conformance Conformance `json:"conformance,omitempty"`
}

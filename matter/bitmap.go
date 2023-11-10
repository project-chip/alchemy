package matter

type Bitmap struct {
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Type        string         `json:"type,omitempty"`
	Bits        []*BitmapValue `json:"bits,omitempty"`
}

type BitmapValue struct {
	Bit         string `json:"bit,omitempty"`
	Name        string `json:"name,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Conformance string `json:"conformance,omitempty"`
}

func (c *Bitmap) ModelType() ModelType {
	return ModelTypeBitmap
}

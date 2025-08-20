package common

/*
type DocTypeFilter struct {
	docType matter.DocType
}

func (sp *DocTypeFilter) Name() string {
	return ""
}

func (sp *DocTypeFilter) Process(cxt context.Context, inputs []*pipeline.Data[*asciidoc.Document]) (outputs []*pipeline.Data[*asciidoc.Document], err error) {
	for _, i := range inputs {
		var docType matter.DocType
		docType, err = i.Content.DocType()
		if err != nil {
			return
		}
		if docType == sp.docType {
			outputs = append(outputs, i)
		}
	}
	return
}

func NewDocTypeFilter(docType matter.DocType) *DocTypeFilter {
	return &DocTypeFilter{docType: docType}
}
*/

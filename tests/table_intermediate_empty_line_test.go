package tests

import "github.com/project-chip/alchemy/asciidoc"

var tableIntermediateEmptyLine = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.Table{
			AttributeList: nil,
			ColumnCount:   1,
			Set: asciidoc.Set{
				&asciidoc.TableRow{
					Set: asciidoc.Set{
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
								VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
									Value: 0,
									IsSet: false,
								},
								Style: asciidoc.Optional[asciidoc.TableCellStyle]{
									Value: 0,
									IsSet: false,
								},
							},
							Set: asciidoc.Set{
								&asciidoc.String{
									Value: "Cell in column 1, row 1",
								},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Some extra content",
								},
								&asciidoc.NewLine{},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Some other extra content",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.TableRow{
					Set: asciidoc.Set{
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
								VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
									Value: 0,
									IsSet: false,
								},
								Style: asciidoc.Optional[asciidoc.TableCellStyle]{
									Value: 0,
									IsSet: false,
								},
							},
							Set: asciidoc.Set{
								&asciidoc.String{
									Value: "Cell in column 1, row 2",
								},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Some extra content",
								},
							},
							Blank: false,
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "\n",
				},
				&asciidoc.TableRow{
					Set: asciidoc.Set{
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
								VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
									Value: 0,
									IsSet: false,
								},
								Style: asciidoc.Optional[asciidoc.TableCellStyle]{
									Value: 0,
									IsSet: false,
								},
							},
							Set: asciidoc.Set{
								&asciidoc.String{
									Value: "Cell in column 1, row 3",
								},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Some extra content",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.TableRow{
					Set: asciidoc.Set{
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
								VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
									Value: 0,
									IsSet: false,
								},
								Style: asciidoc.Optional[asciidoc.TableCellStyle]{
									Value: 0,
									IsSet: false,
								},
							},
							Set: asciidoc.Set{
								&asciidoc.String{
									Value: "Cell in column 1, row 4",
								},
								&asciidoc.NewLine{},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Some extra content",
								},
								&asciidoc.NewLine{},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Some other extra content",
								},
							},
							Blank: false,
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "\n    ",
				},
				&asciidoc.TableRow{
					Set: asciidoc.Set{
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
								VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
									Value: 0,
									IsSet: false,
								},
								Style: asciidoc.Optional[asciidoc.TableCellStyle]{
									Value: 0,
									IsSet: false,
								},
							},
							Set: asciidoc.Set{
								&asciidoc.String{
									Value: "Cell in column 1, row 5",
								},
								&asciidoc.NewLine{},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Some extra content",
								},
								&asciidoc.NewLine{},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Some other extra content",
								},
								&asciidoc.NewLine{},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Even more extra content",
								},
							},
							Blank: false,
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "\n",
				},
				&asciidoc.TableRow{
					Set: asciidoc.Set{
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
								VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
									Value: 0,
									IsSet: false,
								},
								Style: asciidoc.Optional[asciidoc.TableCellStyle]{
									Value: 0,
									IsSet: false,
								},
							},
							Set: asciidoc.Set{
								&asciidoc.String{
									Value: "Cell in column 1, row 6, with a xref: ",
								},
								&asciidoc.CrossReference{
									Set: nil,
									ID:  "ref_Ref",
								},
							},
							Blank: false,
						},
					},
				},
			},
		},
	},
}

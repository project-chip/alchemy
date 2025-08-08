package parse

import "github.com/project-chip/alchemy/asciidoc"

func buildDoc(els asciidoc.Elements) (d *asciidoc.Document) {
	d = &asciidoc.Document{}
	var current asciidoc.ParentElement
	current = d
	var lastSection *asciidoc.Section
	for _, el := range els {
		switch el := el.(type) {
		case *asciidoc.Section:
			if lastSection != nil {
				if el.Level > lastSection.Level {
					lastSection.AddChildSection(el)
					current.Append(el)
				} else if el.Level <= lastSection.Level {
					parent := lastSection.ParentSection()
					var found bool
					for parent != nil {
						if el.Level > parent.Level {
							parent.Append(el)
							parent.AddChildSection(el)
							found = true
							break
						}
						parent = parent.ParentSection()
					}
					if !found { // No parent smaller
						d.Append(el)
					}
				}
			} else {
				current.Append(el)
			}

			lastSection = el
			current = el
		default:
			current.Append(el)
		}
	}
	return
}

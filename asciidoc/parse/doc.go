package parse

import "github.com/project-chip/alchemy/asciidoc"

func buildDoc(els asciidoc.Set) (d *asciidoc.Document, err error) {
	d = &asciidoc.Document{}
	var current asciidoc.HasElements
	current = d
	var lastSection *asciidoc.Section
	//	var lastUnorderedList *asciidoc.UnorderedList

	for _, el := range els {
		switch el := el.(type) {
		case *asciidoc.Section:
			if lastSection != nil {
				if el.Level > lastSection.Level {
					lastSection.AddChild(el)
					err = current.Append(el)
				} else if el.Level <= lastSection.Level {
					parent := lastSection.Parent()
					var found bool
					for parent != nil {
						if el.Level > parent.Level {
							err = parent.Append(el)
							if err != nil {
								return
							}
							parent.AddChild(el)
							found = true
							break
						}
						parent = parent.Parent()
					}
					if !found { // No parent smaller
						err = d.Append(el)
					}
				}
			} else {
				err = current.Append(el)
			}

			lastSection = el
			current = el
		default:
			err = current.Append(el)
		}
		if err != nil {
			return
		}
	}
	return
}

package zapdiff

import (
	"fmt"

	"github.com/beevik/etree"
)

func checkMismatches(ep elementPair, baseName string, n1, n2 string) (mm []XmlMismatch) {
	e1Children := make(map[string]*etree.Element)
	e2Children := make(map[string]*etree.Element)
	mm = make([]XmlMismatch, 0)

	for _, c1 := range ep.e1.ChildElements() {
		id := getElementID(c1)
		e1Children[id] = c1
	}

	for _, c2 := range ep.e2.ChildElements() {
		id := getElementID(c2)
		e2Children[id] = c2
	}

	for id, e1 := range e1Children {
		if _, ok := e2Children[id]; !ok {
			m := XmlMismatch{
				Path:      baseName,
				Type:      getMismatchMissingType(e1),
				Details:   fmt.Sprintf("Only found in %s", n1),
				ElementID: id,
			}
			mm = append(mm, m)
		}
	}

	for id, e2 := range e2Children {
		if _, ok := e1Children[id]; !ok {
			m := XmlMismatch{
				Path:      baseName,
				Type:      getMismatchMissingType(e2),
				Details:   fmt.Sprintf("Only found in %s", n2),
				ElementID: id,
			}
			mm = append(mm, m)
		}
	}

	// Recurse into common tags
	for id, e1 := range e1Children {
		if e2, ok := e2Children[id]; ok {
			// Check attributes
			attrMM := checkAttributes(elementPair{e1: e1, e2: e2}, id, baseName, n1, n2)
			mm = append(mm, attrMM...)

			// Recurse
			subMM := checkMismatches(elementPair{e1: e1, e2: e2}, baseName, n1, n2)
			mm = append(mm, subMM...)
		}
	}

	return
}

func checkAttributes(ep elementPair, id string, baseName string, n1, n2 string) (mm []XmlMismatch) {
	mm = make([]XmlMismatch, 0)
	e1Attrs := make(map[string]string)
	e2Attrs := make(map[string]string)

	for _, a := range ep.e1.Attr {
		e1Attrs[a.Key] = a.Value
	}
	for _, a := range ep.e2.Attr {
		e2Attrs[a.Key] = a.Value
	}

	for k, v1 := range e1Attrs {
		if v2, ok := e2Attrs[k]; !ok {
			m := XmlMismatch{
				Path:      baseName,
				Type:      getMismatchMissingAttrType(ep.e1),
				Details:   fmt.Sprintf("Attribute [%s] only found in %s", k, n1),
				ElementID: id,
			}
			mm = append(mm, m)
		} else if v1 != v2 {
			m := XmlMismatch{
				Path:      baseName,
				Type:      getMismatchAttrValueType(ep.e1),
				Details:   fmt.Sprintf("Attribute [%s] has different values: '%s' in %s, '%s' in %s", k, v1, n1, v2, n2),
				ElementID: id,
			}
			mm = append(mm, m)
		}
	}

	for k := range e2Attrs {
		if _, ok := e1Attrs[k]; !ok {
			m := XmlMismatch{
				Path:      baseName,
				Type:      getMismatchMissingAttrType(ep.e2),
				Details:   fmt.Sprintf("Attribute [%s] only found in %s", k, n2),
				ElementID: id,
			}
			mm = append(mm, m)
		}
	}
	return
}

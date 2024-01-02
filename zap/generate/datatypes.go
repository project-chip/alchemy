package generate

/*
func patchMissingTypes(doc *ascii.Doc, models []interface{}, missingStructs *concurrentMap) {
	for {
		missingCount := restoreMissingTypes(doc, models, missingStructs)

		if missingCount == 0 {
			break
		}

		foundCount := findMissingTypes(doc, missingStructs)
		if foundCount == 0 {
			break
		}
		missingStructs.Lock()
		for n, v := range missingStructs.Map {
			switch v := v.(type) {
			case *matter.Bitmap:
				slog.Debug("adding missing bitmap", "bitmap", v)
				for _, m := range models {
					if c, ok := m.(*matter.Cluster); ok {
						c.Bitmaps = append(c.Bitmaps, v)
					}
				}
				missingStructs.Map[n] = false
			case *matter.Enum:
				slog.Debug("adding missing enum", "enum", v)
				for _, m := range models {
					if c, ok := m.(*matter.Cluster); ok {
						c.Enums = append(c.Enums, v)
					}
				}
				missingStructs.Map[n] = false
			case *matter.Struct:
				slog.Debug("adding missing struct", "struct", v)
				for _, m := range models {
					if c, ok := m.(*matter.Cluster); ok {
						c.Structs = append(c.Structs, v)
					}
				}
				missingStructs.Map[n] = false
			case bool:
				slog.Debug("missing type already handled", "name", n)
			}
		}
		missingStructs.Unlock()
	}
}

func restoreMissingTypes(doc *ascii.Doc, models []any, missingStructs *concurrentMap) (missingCount int) {

	for _, m := range models {
		switch m := m.(type) {
		case *matter.Cluster:
			var unknownTypes []string
			for _, f := range m.Attributes {
				if f.Type != nil && f.Type.BaseType == matter.BaseDataTypeUnknown {
					slog.Debug("unknown attribute data type", "name", f.Type.Name)
					unknownTypes = append(unknownTypes, f.Type.Name)
				}
			}
			for _, s := range m.Structs {
				for _, f := range s.Fields {
					if f.Type != nil && f.Type.BaseType == matter.BaseDataTypeUnknown {
						slog.Debug("unknown struct data type", "name", f.Type.Name)
						unknownTypes = append(unknownTypes, f.Type.Name)
					}
				}
			}
			for _, e := range m.Events {
				for _, f := range e.Fields {
					if f.Type != nil && f.Type.BaseType == matter.BaseDataTypeUnknown {
						slog.Debug("unknown event data type", "name", f.Type.Name)
						unknownTypes = append(unknownTypes, f.Type.Name)
					}
				}
			}
			for _, ut := range unknownTypes {
				var found bool
				if strings.HasSuffix(ut, "Struct") {
					for _, s := range m.Structs {
						if s.Name == ut {
							found = true
							break
						}
					}

				} else if strings.HasSuffix(ut, "Enum") {
					for _, s := range m.Enums {
						if s.Name == ut {
							found = true
							break
						}
					}
				} else if strings.HasSuffix(ut, "Bitmap") {
					for _, s := range m.Bitmaps {
						if s.Name == ut {
							found = true
							break
						}
					}
				}
				if !found {
					missingStructs.Lock()
					if _, ok := missingStructs.Map[ut]; !ok {
						slog.Debug("missing type", "name", ut)
						missingStructs.Map[ut] = nil
						missingCount++
					}
					missingStructs.Unlock()
				}
			}

		}
	}
	return
}

func findMissingTypes(d *ascii.Doc, missingStructs *concurrentMap) (foundCount int) {
	for _, p := range d.Parents() {
		models, err := p.Entities()
		if err != nil {
			slog.Error("error getting models", "err", err)
			continue
		}
		for _, m := range models {
			var name string
			switch m := m.(type) {
			case *matter.Cluster:
				for _, b := range m.Bitmaps {
					missingStructs.Lock()
					if v, ok := missingStructs.Map[b.Name]; ok && v == nil {
						slog.Debug("Found bitmap!", "name", name)
						missingStructs.Map[b.Name] = m
					}
					missingStructs.Unlock()
				}
				for _, b := range m.Enums {
					missingStructs.Lock()
					if v, ok := missingStructs.Map[b.Name]; ok && v == nil {
						slog.Debug("Found enum!", "name", name)
						missingStructs.Map[b.Name] = m
					}
					missingStructs.Unlock()
				}
				for _, b := range m.Structs {
					missingStructs.Lock()
					if v, ok := missingStructs.Map[b.Name]; ok && v == nil {
						slog.Debug("Found struct!", "name", name)
						missingStructs.Map[b.Name] = m
					}
					missingStructs.Unlock()
				}
			case *matter.Bitmap:
				name = m.Name
			case *matter.Enum:
				name = m.Name
			case *matter.Struct:
				name = m.Name
			}
			if len(name) > 0 {
				missingStructs.Lock()
				if v, ok := missingStructs.Map[name]; ok && v == nil {
					slog.Debug("Found!", "name", name)
					missingStructs.Map[name] = m
					foundCount++
				}
				missingStructs.Unlock()

			}
		}
		foundCount += findMissingTypes(p, missingStructs)
	}
	return
}
*/

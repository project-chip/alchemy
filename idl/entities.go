package idl

import (
	"os"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/provisional"
	"github.com/project-chip/alchemy/zap"
)

type ProvisionalFilter struct {
	Mode             string
	ExistingElements map[string]bool
}

func entityShouldBeIncluded(spec *spec.Specification, filter ProvisionalFilter, e types.Entity) bool {
	conf := matter.EntityConformance(e)
	if conformance.IsZigbee(conf) || zap.IsDisallowed(e, conf) || conformance.IsDeprecated(conf) {
		return false
	}

	if filter.Mode != "none" && isProvisional(spec, e) {
		if filter.Mode == "keep-existing" && isElementPresent(filter.ExistingElements, e) {
			return true
		}
		return false
	}
	return true
}

func filterEntities[T types.Entity](spec *spec.Specification, filter ProvisionalFilter, sets ...[]T) (set []T) {
	for _, s := range sets {
		for _, e := range s {
			if !entityShouldBeIncluded(spec, filter, e) {
				continue
			}
			set = append(set, e)
		}
	}
	return
}

func enumerateEntitiesHelper[T types.Entity](list []T, spec *spec.Specification, filter ProvisionalFilter, options *raymond.Options) raymond.SafeString {
	var result strings.Builder
	for i, en := range filterEntities(spec, filter, list) {
		df := options.DataFrame().Copy()
		df.Set("index", i)
		df.Set("key", nil)
		df.Set("first", i == 0)
		df.Set("last", i == len(list)-1)
		if spec != nil {
			if isShared(spec, en) {
				df.Set("shared", true)
			}
			df.Set("provisional", isProvisional(spec, en))
			dr, ok := spec.DataTypeRefs.Get(en)
			if ok {
				df.Set("refCount", dr.Size())
			}
		}
		result.WriteString(options.FnCtxData(en, df))
	}
	return raymond.SafeString(result.String())
}

func isShared(spec *spec.Specification, en types.Entity) bool {
	refs, ok := spec.ClusterRefs.Get(en)
	if ok && refs.Size() > 1 {
		return true
	}
	var cluster *matter.Cluster
	var name string
	var isTargetType bool
	switch entity := any(en).(type) {
	case *matter.Struct:
		cluster = entity.Cluster()
		name = entity.Name
		isTargetType = true
	case *matter.Enum:
		cluster = entity.Cluster()
		name = entity.Name
		isTargetType = true
	case *matter.Bitmap:
		cluster = entity.Cluster()
		name = entity.Name
		isTargetType = true
	}
	if !isTargetType || cluster == nil {
		return false
	}
	doc, ok := spec.DocRefs[cluster]
	if !ok {
		return false
	}
	if spec.Errata == nil {
		return false
	}
	errata := spec.Errata.Get(doc.Path.Relative)
	if errata == nil {
		return false
	}
	switch any(en).(type) {
	case *matter.Struct:
		_, ok = errata.SDK.SharedStructs[name]
		return ok
	case *matter.Enum:
		_, ok = errata.SDK.SharedEnums[name]
		return ok
	case *matter.Bitmap:
		_, ok = errata.SDK.SharedBitmaps[name]
		return ok
	}
	return false
}

func isProvisional(spec *spec.Specification, entity types.Entity) bool {
	switch entity := entity.(type) {
	case *matter.Bitmap:
		if entity.Name == "Feature" || entity.Name == "Features" {
			return false
		}
	case *matter.Enum:
		if strings.HasSuffix(entity.Name, "Tag") {
			// This is a hacky workaround for namespaces being represented as enumerations in .matter files, but not being provisional
			return false
		}
	case *matter.Field:
		if entity.Name == "FabricIndex" {
			return false
		}
		return conformance.IsProvisional(matter.EntityConformance(entity))
	case nil:
		return false
	}
	is := provisional.Check(spec, entity, entity)

	switch is {
	case provisional.StateAllClustersProvisional,
		provisional.StateAllDataTypeReferencesProvisional,
		provisional.StateExplicitlyProvisional,
		provisional.StateUnreferenced:
		return true
	default:
		return false
	}
}

func entityPath(e types.Entity) string {
	var parts []string
	curr := e
	for curr != nil {
		switch node := curr.(type) {
		case *matter.Cluster:
			parts = append(parts, caseify(node.Name, false, false))
			curr = nil // Stop at cluster level
		case *matter.Enum:
			parts = append(parts, node.Name)
			curr = node.Parent()
		case *matter.Bitmap:
			if node.Name == "Features" {
				parts = append(parts, "Feature")
			} else {
				parts = append(parts, node.Name)
			}
			curr = node.Parent()
		case *matter.Features:
			parts = append(parts, "Feature")
			curr = node.Parent()
		case *matter.Struct:
			parts = append(parts, node.Name)
			curr = node.Parent()
		case *matter.Event:
			parts = append(parts, node.Name, "event")
			curr = node.Parent()
		case *matter.Command:
			parts = append(parts, node.Name, "command")
			curr = node.Parent()
		case *matter.Field:
			if _, isCluster := node.Parent().(*matter.Cluster); isCluster {
				parts = append(parts, node.Name, "attribute")
			} else {
				parts = append(parts, node.Name)
			}
			curr = node.Parent()
		case *matter.EnumValue:
			parts = append(parts, node.Name)
			curr = node.Parent()
		case *matter.BitmapBit:
			parts = append(parts, node.Name())
			curr = node.Parent()
		case *matter.Feature:
			parts = append(parts, node.Name())
			curr = node.Parent()
		default:
			curr = curr.Parent()
		}
	}
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	return strings.Join(parts, ".")
}

func isElementPresent(existing map[string]bool, e types.Entity) bool {
	if len(existing) == 0 {
		return false
	}
	path := strings.ToLower(entityPath(e))
	return existing[path]
}

func parseExistingMatterElements(path string) (map[string]bool, error) {
	elements := make(map[string]bool)
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return elements, nil
		}
		return nil, err
	}
	content := string(contentBytes)
	lines := strings.Split(content, "\n")

	var currentCluster string

	type blockContext struct {
		blockType string
		blockName string
	}
	var stack []blockContext

	getCurrentPath := func(name string) string {
		var pathParts []string
		if currentCluster != "" {
			pathParts = append(pathParts, currentCluster)
		}
		for _, ctx := range stack {
			if ctx.blockType != "cluster" && ctx.blockName != "" {
				if ctx.blockType == "event" || ctx.blockType == "command" || ctx.blockType == "attribute" {
					pathParts = append(pathParts, ctx.blockType, ctx.blockName)
				} else {
					pathParts = append(pathParts, ctx.blockName)
				}
			}
		}
		if name != "" {
			pathParts = append(pathParts, name)
		}
		return strings.ToLower(strings.Join(pathParts, "."))
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "/*") || strings.HasPrefix(trimmed, "*") {
			continue
		}
		for {
			accessIdx := strings.Index(trimmed, "access(")
			if accessIdx == -1 {
				break
			}
			closeIdx := strings.Index(trimmed[accessIdx:], ")")
			if closeIdx == -1 {
				break
			}
			trimmed = trimmed[:accessIdx] + " " + trimmed[accessIdx+closeIdx+1:]
		}

		if strings.HasPrefix(trimmed, "cluster ") || strings.Contains(trimmed, " cluster ") {
			parts := strings.Fields(trimmed)
			for idx, word := range parts {
				if word == "cluster" && idx+1 < len(parts) {
					name := strings.TrimSuffix(parts[idx+1], "{")
					currentCluster = strings.TrimSpace(name)
					elements[strings.ToLower(currentCluster)] = true
					break
				}
			}
		}

		var declType, declName string

		if strings.HasPrefix(trimmed, "enum ") || strings.Contains(trimmed, " enum ") {
			parts := strings.Fields(trimmed)
			for idx, word := range parts {
				if word == "enum" && idx+1 < len(parts) {
					name := parts[idx+1]
					name = strings.Split(name, ":")[0]
					name = strings.TrimSuffix(name, "{")
					declName = strings.TrimSpace(name)
					declType = "enum"
					break
				}
			}
		} else if strings.HasPrefix(trimmed, "bitmap ") || strings.Contains(trimmed, " bitmap ") {
			parts := strings.Fields(trimmed)
			for idx, word := range parts {
				if word == "bitmap" && idx+1 < len(parts) {
					name := parts[idx+1]
					name = strings.Split(name, ":")[0]
					name = strings.TrimSuffix(name, "{")
					declName = strings.TrimSpace(name)
					declType = "bitmap"
					break
				}
			}
		} else if (strings.HasPrefix(trimmed, "struct ") || strings.Contains(trimmed, " struct ")) && !strings.Contains(trimmed, "attribute ") {
			parts := strings.Fields(trimmed)
			for idx, word := range parts {
				if word == "struct" && idx+1 < len(parts) {
					name := parts[idx+1]
					name = strings.TrimSuffix(name, "{")
					name = strings.Split(name, "=")[0]
					declName = strings.TrimSpace(name)
					declType = "struct"
					break
				}
			}
		} else if strings.Contains(trimmed, " event ") {
			eqIdx := strings.Index(trimmed, "=")
			if eqIdx != -1 {
				left := strings.TrimSpace(trimmed[:eqIdx])
				parts := strings.Fields(left)
				if len(parts) > 0 {
					name := parts[len(parts)-1]
					declName = strings.TrimSpace(name)
					declType = "event"
				}
			}
		} else if strings.Contains(trimmed, "attribute ") {
			eqIdx := strings.Index(trimmed, "=")
			if eqIdx != -1 {
				left := strings.TrimSpace(trimmed[:eqIdx])
				parts := strings.Fields(left)
				if len(parts) > 0 {
					name := parts[len(parts)-1]
					name = strings.TrimSuffix(name, "[]")
					attrName := strings.TrimSpace(name)
					elements[getCurrentPath("attribute."+attrName)] = true
				}
			}
		} else if strings.Contains(trimmed, "command ") {
			parenIdx := strings.Index(trimmed, "(")
			if parenIdx != -1 {
				left := strings.TrimSpace(trimmed[:parenIdx])
				parts := strings.Fields(left)
				if len(parts) > 0 {
					cmdName := strings.TrimSpace(parts[len(parts)-1])
					elements[getCurrentPath("command."+cmdName)] = true
				}
			}
		} else if len(stack) > 0 {
			top := stack[len(stack)-1]
			if top.blockType == "enum" || top.blockType == "bitmap" || top.blockType == "struct" || top.blockType == "event" {
				eqIdx := strings.Index(trimmed, "=")
				if eqIdx != -1 {
					left := strings.TrimSpace(trimmed[:eqIdx])
					parts := strings.Fields(left)
					if len(parts) > 0 {
						name := parts[len(parts)-1]
						name = strings.TrimSuffix(name, "[]")
						fieldName := strings.TrimSpace(name)
						if (top.blockType == "enum" || top.blockType == "bitmap") && len(fieldName) > 1 && fieldName[0] == 'k' && fieldName[1] >= 'A' && fieldName[1] <= 'Z' {
							fieldName = fieldName[1:]
						}
						elements[getCurrentPath(fieldName)] = true
					}
				}
			}
		}

		openBraces := strings.Count(trimmed, "{")
		closeBraces := strings.Count(trimmed, "}")

		for o := 0; o < openBraces; o++ {
			if declType != "" && declName != "" {
				stack = append(stack, blockContext{blockType: declType, blockName: declName})
				elements[getCurrentPath("")] = true
				declType = ""
				declName = ""
			} else {
				stack = append(stack, blockContext{})
			}
		}
		for c := 0; c < closeBraces; c++ {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return elements, nil
}

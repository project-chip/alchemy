package parse

import (
	"regexp"

	"github.com/hasty/matterfmt/matter"
)

var accessPattern = regexp.MustCompile(`^(?:(?:^|\s+)(?:(?P<ReadWrite>(?:R\*W)|(?:R\[W\])|(?:[RW]+))|(?P<Fabric>[FS]+)|(?P<Privileges>[VOMA]+)|(?P<Timed>T)))+\s*$`)
var accessPatternMatchMap map[int]matter.AccessCategory

func init() {
	accessPatternMatchMap = make(map[int]matter.AccessCategory)
	for i, name := range accessPattern.SubexpNames() {
		switch name {
		case "ReadWrite":
			accessPatternMatchMap[i] = matter.AccessCategoryReadWrite
		case "Fabric":
			accessPatternMatchMap[i] = matter.AccessCategoryFabric
		case "Privileges":
			accessPatternMatchMap[i] = matter.AccessCategoryPrivileges
		case "Timed":
			accessPatternMatchMap[i] = matter.AccessCategoryTimed
		}
	}
}

func ParseAccess(vc string) map[matter.AccessCategory]string {
	matches := accessPattern.FindStringSubmatch(vc)
	if matches == nil {
		return nil
	}
	access := make(map[matter.AccessCategory]string)
	for i, s := range matches {
		if s == "" {
			continue
		}
		category, ok := accessPatternMatchMap[i]
		if !ok {
			continue
		}
		access[category] = s
	}
	return access
}

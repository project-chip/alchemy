package matter

import "regexp"

type AccessCategory uint8

const (
	AccessCategoryUnknown AccessCategory = iota
	AccessCategoryReadWrite
	AccessCategoryFabric
	AccessCategoryPrivileges
	AccessCategoryTimed
)

var AccessCategoryOrder = [...]AccessCategory{
	AccessCategoryReadWrite,
	AccessCategoryFabric,
	AccessCategoryPrivileges,
	AccessCategoryTimed,
}

var accessPattern = regexp.MustCompile(`^(?:(?:^|\s+)(?:(?P<ReadWrite>(?:R\*W)|(?:R\[W\])|(?:[RW]+))|(?P<Fabric>[FS]+)|(?P<Privileges>[VOMA]+)|(?P<Timed>T)))+\s*$`)
var accessPatternMatchMap map[int]AccessCategory

func init() {
	accessPatternMatchMap = make(map[int]AccessCategory)
	for i, name := range accessPattern.SubexpNames() {
		switch name {
		case "ReadWrite":
			accessPatternMatchMap[i] = AccessCategoryReadWrite
		case "Fabric":
			accessPatternMatchMap[i] = AccessCategoryFabric
		case "Privileges":
			accessPatternMatchMap[i] = AccessCategoryPrivileges
		case "Timed":
			accessPatternMatchMap[i] = AccessCategoryTimed
		}
	}
}

func ParseAccess(vc string) map[AccessCategory]string {
	matches := accessPattern.FindStringSubmatch(vc)
	if matches == nil {
		return nil
	}
	access := make(map[AccessCategory]string)
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

func ExtractAccessValues(access interface{}) (readAccess string, writeAccess string, fabricScoped int8, fabricSensitive int8, timed int8) {
	s, ok := access.(string)
	if !ok {
		return
	}
	am := ParseAccess(s)
	if am == nil {
		return
	}
	return ParseAccessValues(am)
}

func ParseAccessValues(am map[AccessCategory]string) (readAccess string, writeAccess string, fabricScoped int8, fabricSensitive int8, timed int8) {
	if am[AccessCategoryFabric] == "F" {
		fabricScoped = 1
	} else if am[AccessCategoryFabric] == "S" {
		fabricSensitive = 1
	}
	if am[AccessCategoryTimed] == "T" {
		timed = 1
	}
	rw := am[AccessCategoryReadWrite]
	var hasRead, hasWrite bool
	switch rw {
	case "RW", "R[W]":
		hasRead = true
		hasWrite = true
	case "R":
		hasRead = true
	case "W":
		hasWrite = true
	}
	ps, ok := am[AccessCategoryPrivileges]
	if !ok {
		return
	}
	for _, r := range ps {
		if hasRead {
			readAccess = string(r)
			hasRead = false
		} else if hasWrite {
			writeAccess = string(r)
			break
		}
	}
	return
}

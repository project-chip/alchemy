package ascii

import (
	"regexp"
	"strings"

	"github.com/hasty/matterfmt/matter"
)

type accessCategoryMatch uint8

const (
	accessCategoryMatchUnknown accessCategoryMatch = iota
	accessCategoryMatchReadWrite
	accessCategoryMatchFabric
	accessCategoryMatchPrivileges
	accessCategoryMatchTimed
)

var accessPattern = regexp.MustCompile(`^(?:(?:^|\s+)(?:(?P<ReadWrite>(?:R\*W)|(?:R\[W\])|(?:[RW]+))|(?P<Fabric>[FS]+)|(?P<Privileges>[VOMA]+)|(?P<Timed>T)))+\s*$`)
var accessPatternMatchMap map[int]accessCategoryMatch

func init() {
	accessPatternMatchMap = make(map[int]accessCategoryMatch)
	for i, name := range accessPattern.SubexpNames() {
		switch name {
		case "ReadWrite":
			accessPatternMatchMap[i] = accessCategoryMatchReadWrite
		case "Fabric":
			accessPatternMatchMap[i] = accessCategoryMatchFabric
		case "Privileges":
			accessPatternMatchMap[i] = accessCategoryMatchPrivileges
		case "Timed":
			accessPatternMatchMap[i] = accessCategoryMatchTimed
		}
	}
}

func ParseAccess(vc string) (a matter.Access) {
	matches := accessPattern.FindStringSubmatch(vc)
	if matches == nil {
		return matter.Access{}
	}
	access := make(map[accessCategoryMatch]string)
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
	a = matter.Access{}
	var readAccess, writeAccess, invokeAccess string
	rw := access[accessCategoryMatchReadWrite]
	var hasRead, hasWrite, optionalWrite bool
	switch rw {
	case "RW":
		hasRead = true
		hasWrite = true
	case "R*W", "R[W]":
		hasRead = true
		hasWrite = true
		optionalWrite = true
	case "R":
		hasRead = true
	case "W":
		hasWrite = true
	}
	ps, ok := access[accessCategoryMatchPrivileges]
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
		} else {
			invokeAccess = string(r)
		}
	}
	a.Read = stringToPrivilege(readAccess)
	a.Write = stringToPrivilege(writeAccess)
	a.Invoke = stringToPrivilege(invokeAccess)
	a.OptionalWrite = optionalWrite

	if access[accessCategoryMatchFabric] == "F" {
		a.FabricScoped = true
	} else if access[accessCategoryMatchFabric] == "S" {
		a.FabricSensitive = true
	}
	if access[accessCategoryMatchTimed] == "T" {
		a.Timed = true
	}
	return
}

func AccessToAsciiString(a matter.Access) string {
	var out strings.Builder
	if a.Read != matter.PrivilegeUnknown || a.Write != matter.PrivilegeUnknown {
		if a.Read != matter.PrivilegeUnknown {
			out.WriteRune('R')
		}
		if a.Write != matter.PrivilegeUnknown {
			if a.OptionalWrite {
				out.WriteString("[W]")
			} else {
				out.WriteRune('W')
			}
		}
		out.WriteRune(' ')
		if a.Read != matter.PrivilegeUnknown {
			out.WriteString(privilegeToString(a.Read))
		}
		if a.Write != matter.PrivilegeUnknown {
			out.WriteString(privilegeToString(a.Write))
		}
	} else if a.Invoke != matter.PrivilegeUnknown {
		out.WriteString(privilegeToString(a.Invoke))
	}
	if a.FabricScoped || a.FabricSensitive {
		if out.Len() > 0 {
			out.WriteRune(' ')
		}
		if a.FabricScoped {
			out.WriteRune('F')
		}
		if a.FabricSensitive {
			out.WriteRune('S')
		}
	}
	if a.Timed {
		if out.Len() > 0 {
			out.WriteRune(' ')
		}
		out.WriteRune('T')
	}
	return out.String()
}

func stringToPrivilege(p string) matter.Privilege {
	switch p {
	case "V":
		return matter.PrivilegeView
	case "O":
		return matter.PrivilegeOperate
	case "M":
		return matter.PrivilegeManage
	case "A":
		return matter.PrivilegeAdminister
	}
	return matter.PrivilegeUnknown
}

func privilegeToString(p matter.Privilege) string {
	switch p {
	case matter.PrivilegeView:
		return "V"
	case matter.PrivilegeOperate:
		return "O"
	case matter.PrivilegeManage:
		return "M"
	case matter.PrivilegeAdminister:
		return "A"
	}
	return ""
}

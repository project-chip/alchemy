package mle

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/project-chip/alchemy/matter/spec"
)

type deviceTypeInfo struct {
	DeviceTypeID string
}

func deviceTypeIdTaken(id string, masterDeviceTypeMap map[string]deviceTypeInfo) (taken bool, name string) {
	var dti deviceTypeInfo
	for name, dti = range masterDeviceTypeMap {
		if dti.DeviceTypeID == id {
			taken = true
			return
		}
	}
	taken = false
	return
}

func deviceTypeIdReserved(id string, reserveds []string) (reserved bool) {
	for _, r := range reserveds {
		if r == id {
			reserved = true
			return
		}
	}
	reserved = false
	return
}

func parseMasterDeviceTypeList(filePath string) (dti map[string]deviceTypeInfo, reserveds []string, violations map[string][]spec.Violation, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dti = make(map[string]deviceTypeInfo)
	reserveds = make([]string, 0)
	violations = make(map[string][]spec.Violation)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if !strings.HasPrefix(line, "|") {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) > 2 && strings.Contains(parts[1], "Device Type ID") && strings.Contains(parts[2], "Device Type Name") {
			continue
		}

		if len(parts) >= 3 {
			id := strings.TrimSpace(parts[1])
			name := strings.TrimSpace(parts[2])

			if id == "" || name == "" {
				continue
			}
			if name == "Reserved" {
				reserveds = append(reserveds, id)
				continue
			}
			if taken, _ := deviceTypeIdTaken(id, dti); taken {
				v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Device Type ID is duplicated on Master List. ID='%s'", id)}
				v.Path, v.Line = "MasterDeviceTypeList.adoc", 0
				violations[v.Path] = append(violations[v.Path], v)
				continue
			}
			if _, ok := dti[name]; ok {
				v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Device Type Name is duplicated on Master List. name='%s'", name)}
				v.Path, v.Line = "MasterDeviceTypeList.adoc", 0
				violations[v.Path] = append(violations[v.Path], v)
				continue
			}

			dti[name] = deviceTypeInfo{
				DeviceTypeID: id,
			}
		}
	}

	err = scanner.Err()

	return
}

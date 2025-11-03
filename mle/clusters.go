package mle

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/project-chip/alchemy/matter/spec"
)

type clusterInfo struct {
	ClusterID string
	PICScode  string
}

func clusterIdTaken(id string, masterClusterMap map[string]clusterInfo) (taken bool, name string) {
	var ci clusterInfo
	if id == "n/a" {
		taken = false
		return
	}
	for name, ci = range masterClusterMap {
		if ci.ClusterID == id {
			taken = true
			return
		}
	}
	taken = false
	return
}

func clusterIdReserved(id string, reserveds []string) (reserved bool) {
	for _, r := range reserveds {
		if r == id {
			reserved = true
			return
		}
	}
	reserved = false
	return
}

func parseMasterClusterList(filePath string) (ci map[string]clusterInfo, reserveds []string, violations map[string][]spec.Violation, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ci = make(map[string]clusterInfo)
	reserveds = make([]string, 0)
	violations = make(map[string][]spec.Violation)
	lineNumber := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineNumber++

		if !strings.HasPrefix(line, "|") {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) > 2 && strings.Contains(parts[1], "Cluster ID") && strings.Contains(parts[2], "Cluster Name") {
			continue
		}

		if len(parts) >= 4 {
			id := strings.TrimSpace(parts[1])
			name := strings.TrimSpace(parts[2])
			pics := strings.TrimSpace(parts[3])

			if id == "" || name == "" {
				continue
			}
			if name == "Reserved" {
				reserveds = append(reserveds, id)
				continue
			}
			if taken, _ := clusterIdTaken(id, ci); taken {
				v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster ID is duplicated on Master List. ID='%s'", id)}
				v.Path, v.Line = "MasterClusterList.adoc", lineNumber
				violations[v.Path] = append(violations[v.Path], v)
				continue
			}
			if _, ok := ci[name]; ok {
				v := spec.Violation{Entity: nil, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster Name is duplicated on Master List. name='%s'", name)}
				v.Path, v.Line = "MasterClusterList.adoc", lineNumber
				violations[v.Path] = append(violations[v.Path], v)
				continue
			}

			ci[name] = clusterInfo{
				ClusterID: id,
				PICScode:  pics,
			}
		}
	}

	err = scanner.Err()

	return
}

package zap

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/project-chip/alchemy/cmd/action/github"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/sethvargo/go-githubactions"
	"github.com/spf13/cobra"
)

func CheckZAPXML(cmd *cobra.Command, args []string) error {
	cxt := context.Background()

	action := githubactions.New()

	action.Infof("Alchemy %s", config.Version())

	githubContext, err := githubactions.Context()
	if err != nil {
		return fmt.Errorf("failed on getting GitHub context: %w", err)

	}
	pr, err := github.ReadPullRequest(cxt, githubContext, action)
	if err != nil {
		return fmt.Errorf("failed on reading pull request: %w", err)
	}
	if pr == nil {
		return nil
	}
	var changedFiles []string
	changedFiles, err = github.GetPRChangedFiles(cxt, githubContext, action, pr)
	if err != nil {
		return fmt.Errorf("failed on getting pull request changes: %w", err)
	}
	if len(changedFiles) == 0 {
		action.Infof("No changes found\n")
		return nil
	}

	changedPaths := make(map[alchemyFileType][]string)
	for _, path := range changedFiles {
		fileType := isAlchemyGeneratedXML(path)
		if fileType == alchemyFileTypeNone {
			continue
		}
		changedPaths[fileType] = append(changedPaths[fileType], path)
	}

	if len(changedPaths) == 0 {
		action.Infof("No changed alchemy files found\n")
		return nil
	}

	changedClusters := changedPaths[alchemyFileTypeCluster]
	if len(changedClusters) == 0 {
		action.Infof("No changed alchemy clusters found\n")
		return nil
	}
	return nil
}

type alchemyFileType uint8

const (
	alchemyFileTypeNone alchemyFileType = iota
	alchemyFileTypeCluster
	alchemyFileTypeDeviceTypes
	alchemyFileTypeClusterList
	alchemyFileTypeNamespaces
)

func isAlchemyGeneratedXML(path string) alchemyFileType {
	ext := filepath.Ext(path)
	switch ext {
	case "json", "xml":
	default:
		return alchemyFileTypeNone
	}
	switch path {
	case "src/app/zap-templates/zcl/data-model/chip/matter-devices.xml":
		return alchemyFileTypeDeviceTypes
	case "src/app/zap-templates/zcl/data-model/chip/semantic-tag-namespace-enums.xml":
		return alchemyFileTypeNamespaces
	case "src/app/zap-templates/zcl/zcl.json", "src/app/zap-templates/zcl/zcl-with-test-extensions.json":
		return alchemyFileTypeClusterList
	}
	if text.HasCaseInsensitivePrefix(path, "src/app/zap-templates/zcl/data-model/chip/") {
		return alchemyFileTypeCluster
	}
	return alchemyFileTypeNone
}

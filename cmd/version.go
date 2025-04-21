package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/Masterminds/semver"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/config"
)

type Version struct {
}

func (v Version) Run(*cli.Context) (err error) {
	fmt.Fprintf(os.Stdout, "version: %v\n", config.Version())
	return
}

func checkVersion(ctx context.Context, versionChan chan string) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://raw.githubusercontent.com/project-chip/alchemy/refs/heads/main/.github/version.json", nil)
	if err != nil {
		slog.Debug("error creating version request", slog.Any("error", err))
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Debug("error fetching version", slog.Any("error", err))
		return
	}
	defer resp.Body.Close()

	var versionJson struct {
		Version string `json:"version"`
	}
	err = json.NewDecoder(resp.Body).Decode(&versionJson)
	if err != nil {
		slog.Debug("error decoding version", slog.Any("error", err))
		return
	}
	select {
	case versionChan <- versionJson.Version:
	default:
	}
}

func compareVersion(version string) {
	v, err := semver.NewVersion(version)
	if err != nil {
		slog.Debug("error parsing remote version", slog.Any("error", err))
		return
	}
	bv, err := semver.NewVersion(config.Version())
	if err != nil {
		slog.Debug("error parsing local version", slog.Any("error", err), slog.String("version", config.Version()))
		return
	}
	if v.GreaterThan(bv) {
		fmt.Fprintf(os.Stderr, "\n\n")
		fmt.Fprintf(os.Stderr, "You are running an outdated version of Alchemy (%s).\n", bv.Original())
		fmt.Fprintf(os.Stderr, "Please download the latest version (%s) here: https://github.com/project-chip/alchemy/releases/tag/%s\n", version, version)
	}
}

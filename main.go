package main

import (
	"log/slog"
	"time"

	"github.com/project-chip/alchemy/cmd"
)

var start = time.Now()

func main() {
	cmd.Execute()
	slog.Debug("Complete", "runtime", time.Since(start))
}

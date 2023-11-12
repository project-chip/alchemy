package main

import (
	"log/slog"
	"time"

	"github.com/hasty/alchemy/cmd"
)

var start = time.Now()

func main() {
	cmd.Execute()
	slog.Debug("Complete", "runtime", time.Since(start))
}

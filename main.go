package main

import (
	"log/slog"
	"time"

	"github.com/hasty/alchemy/cmd"
)

var start = time.Now()

func main() {
	cmd.Execute()
	slog.Info("Complete", "runtime", time.Since(start))
}

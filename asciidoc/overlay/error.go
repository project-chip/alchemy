package overlay

import "github.com/project-chip/alchemy/internal/log"

type OverlayError struct {
	error
	Source log.Source
}

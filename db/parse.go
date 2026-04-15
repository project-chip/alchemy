package db

import (
	"fmt"

	"github.com/project-chip/alchemy/internal/log"
)

type sectionInfo struct {
	id     int32
	values *dbRow

	parent *sectionInfo

	children map[string][]*sectionInfo

	source log.Source
}

var errMissingTable = fmt.Errorf("no table found")

func (h *Host) newSectionInfo(table string, parent *sectionInfo, values *dbRow, source log.Source) *sectionInfo {
	return &sectionInfo{
		id:       h.nextID(table),
		values:   values,
		parent:   parent,
		source:   source,
		children: make(map[string][]*sectionInfo),
	}
}

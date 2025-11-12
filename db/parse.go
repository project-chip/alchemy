package db

import (
	"fmt"
)

type sectionInfo struct {
	id     int32
	values *dbRow

	parent *sectionInfo

	children map[string][]*sectionInfo
}

var errMissingTable = fmt.Errorf("no table found")

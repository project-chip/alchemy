package db

import (
	mms "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/types"
	"github.com/project-chip/alchemy/matter"
)

func getQualitySchemaColumns(tableName string) []*mms.Column {
	return []*mms.Column{
		{Name: "nullable", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "non_volatile", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "fixed", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "scene", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "reportable", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "changes_omitted", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "singleton", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "atomic_write", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
		{Name: "large_message", Type: types.Boolean, Nullable: true, Source: tableName, PrimaryKey: false},
	}
}

func getQualitySchemaColumnValues(access any) []any {
	var nullable, nonVolatile, fixed, scene, reportable, changesOmitted, singleton, atomic, largeMessage int8
	switch access := access.(type) {
	case string:
		var val int8 = 1
		for _, r := range access {
			switch r {
			case 'X':
				nullable = val
			case 'N':
				nonVolatile = val
			case 'F':
				fixed = val
			case 'S':
				scene = val
			case 'P':
				reportable = val
			case 'C':
				changesOmitted = val
			case 'I':
				singleton = val
			case 'T':
				atomic = val
			case 'L':
				largeMessage = val
			case '!':
				val = -1
				continue
			}
			if val == -1 {
				val = 1
			}
		}
	case matter.Quality:
		nullable = getQualityValue(access, matter.QualityNullable)
		nonVolatile = getQualityValue(access, matter.QualityNonVolatile)
		fixed = getQualityValue(access, matter.QualityFixed)
		scene = getQualityValue(access, matter.QualityScene)
		reportable = getQualityValue(access, matter.QualityReportable)
		changesOmitted = getQualityValue(access, matter.QualityChangedOmitted)
		singleton = getQualityValue(access, matter.QualitySingleton)
		atomic = getQualityValue(access, matter.QualityAtomicWrite)
		largeMessage = getQualityValue(access, matter.QualityLargeMessage)
	}
	return []any{nullable, nonVolatile, fixed, scene, reportable, changesOmitted, singleton, atomic, largeMessage}
}

func getQualityValue(q matter.Quality, desired matter.Quality) int8 {
	if q.Has(desired) {
		return 1
	}
	return 0
}

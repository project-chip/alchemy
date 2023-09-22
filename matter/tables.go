package matter

import "github.com/bytesparadise/libasciidoc/pkg/types"

type TableColumn uint8

const (
	TableColumnUnknown TableColumn = iota
	TableColumnID                  // Special section type for everything that comes before any known sections
	TableColumnName
	TableColumnType
	TableColumnConstraint
	TableColumnQuality
	TableColumnDefault
	TableColumnAccess
	TableColumnConformance
	TableColumnHierarchy
	TableColumnRole
	TableColumnContext
	TableColumnScope
	TableColumnPICS
)

var BannedTableAttributes = [...]string{"cols", "frame", "width"}

var AttributesTableColumnOrder = [...]TableColumn{
	TableColumnID,
	TableColumnName,
	TableColumnType,
	TableColumnConstraint,
	TableColumnQuality,
	TableColumnDefault,
	TableColumnAccess,
	TableColumnConformance,
}

var ClassificationTableColumnOrder = [...]TableColumn{
	TableColumnHierarchy,
	TableColumnRole,
	TableColumnScope,
	TableColumnContext, // This will get renamed to Scope
	TableColumnPICS,
}

var ClassificationTableColumnNames = map[TableColumn]string{
	TableColumnHierarchy: "Hierarchy",
	TableColumnRole:      "Role",
	TableColumnScope:     "Scope",
	TableColumnContext:   "Scope", // Rename Context to Scope
	TableColumnPICS:      "PICS Code",
}

var ClusterIDSectionName = "Cluster ID"

var ClusterIDTableColumnOrder = [...]TableColumn{
	TableColumnID,
	TableColumnName,
}

var ClusterIDTableColumnNames = map[TableColumn]string{
	TableColumnID:   "ID",
	TableColumnName: "Name",
}

var AllowedTableAttributes = types.Attributes{
	"id":      nil,
	"title":   nil,
	"valign":  "middle",
	"options": types.Options{"header"},
}

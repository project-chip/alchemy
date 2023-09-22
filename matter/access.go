package matter

type AccessCategory uint8

const (
	AccessCategoryUnknown AccessCategory = iota
	AccessCategoryReadWrite
	AccessCategoryFabric
	AccessCategoryPrivileges
	AccessCategoryTimed
)

var AccessCategoryOrder = [...]AccessCategory{
	AccessCategoryReadWrite,
	AccessCategoryFabric,
	AccessCategoryPrivileges,
	AccessCategoryTimed,
}

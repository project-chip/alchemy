package matter

type Role uint8

const (
	RoleUnknown Role = iota
	RoleApplication
	RoleUtility
)

var RoleNames = map[Role]string{
	RoleUnknown:     "Unknown",
	RoleApplication: "Application",
	RoleUtility:     "Utility",
}

func (r Role) String() string {
	return RoleNames[r]
}

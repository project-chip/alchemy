package matter

type Scope uint8

const (
	ScopeUnknown Scope = iota
	ScopeNode
	ScopeEndpoint
)

var ScopeNames = map[Scope]string{
	ScopeUnknown:  "Unknown",
	ScopeNode:     "Node",
	ScopeEndpoint: "Endpoint",
}

func (s Scope) String() string {
	return ScopeNames[s]
}

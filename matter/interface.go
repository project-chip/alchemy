package matter

import (
	"encoding/json"
	"fmt"
)

type Interface uint8

const (
	InterfaceUnknown Interface = iota
	InterfaceClient
	InterfaceServer
)

var (
	interfaceNames = map[Interface]string{
		InterfaceUnknown: "unknown",
		InterfaceClient:  "client",
		InterfaceServer:  "server",
	}
	interfaceValues = map[string]Interface{
		"unknown": InterfaceUnknown,
		"client":  InterfaceClient,
		"server":  InterfaceServer,
	}
)

func (s Interface) MarshalJSON() ([]byte, error) {
	return json.Marshal(interfaceNames[s])
}

func (s *Interface) UnmarshalJSON(data []byte) (err error) {
	var ss string
	if err := json.Unmarshal(data, &ss); err != nil {
		return fmt.Errorf("error parsing interface JSON value %s: %w", string(data), err)
	}
	var ok bool
	if *s, ok = interfaceValues[ss]; !ok {
		return fmt.Errorf("unknown interface: %s", ss)
	}
	return nil
}

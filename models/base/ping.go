package base

import (
	"fmt"
	"runtime"
)

type Ping struct {
	_msgpack     struct{} `msgpack:",as_array"`
	Architecture string   `msgpack:"arch"`
	OS           string   `msgpack:"os"`
}

func NewPing() *Ping {
	return NewPingFromArgs(runtime.GOARCH, runtime.GOOS)
}

func NewPingFromArgs(arch string, os string) *Ping {
	return &Ping{
		Architecture: arch,
		OS:           os,
	}
}

func ConvertFromInterface(args interface{}) (*Ping, error) {
	var ping *Ping
	// We're now focusing on array-based MessagePack only (CASE 2). Remove it if
	//  we're sticking with array-based serialization.
	switch convertVal := args.(type) {
	case map[string]interface{}:
		ping = NewPingFromArgs(convertVal["os"].(string), convertVal["arch"].(string))
	case []interface{}:
		ping = NewPingFromArgs(convertVal[0].(string), convertVal[1].(string))
	default:
		return nil, fmt.Errorf("cannot convert using interface")
	}
	return ping, nil
}

func (p *Ping) String() string {
	return fmt.Sprintf("PING <ARCH: %s, OS: %s>", p.Architecture, p.OS)
}

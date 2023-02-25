package base

import (
	"fmt"
	"runtime"

	"github.com/CrosineEnterprises/ganymede/system"
	"github.com/vmihailenco/msgpack/v5"
)

type Ping struct {
	_msgpack     struct{} `msgpack:",as_array"`
	Architecture string   `msgpack:"arch"`
	OS           string   `msgpack:"os"`
	Capabilities []string `msgpack:"capabilities"`
}

const METHOD_NAME = "init"

func NewPing() *Ping {
	return NewPingFromArgs(system.CheckCapabilities())
}

func NewPingFromArgs(capabilities []string) *Ping {
	return &Ping{
		Architecture: runtime.GOARCH,
		OS:           runtime.GOOS,
		Capabilities: capabilities,
	}
}

func DecodePing(decoder *msgpack.Decoder) (*Ping, error) {
	var ping Ping
	convertErr := decoder.Decode(&ping)
	if convertErr != nil {
		return nil, convertErr
	}
	return &ping, nil
}

func (p *Ping) String() string {
	return fmt.Sprintf("PING <ARCH: %s, OS: %s>", p.Architecture, p.OS)
}

package base

import (
	"fmt"
	"runtime"

	"github.com/CrosineEnterprises/ganymede/system"
	"github.com/vmihailenco/msgpack/v5"
)

type Init struct {
	_msgpack     struct{} `msgpack:",as_array"`
	Architecture string   `msgpack:"arch"`
	OS           string   `msgpack:"os"`
	Capabilities []string `msgpack:"capabilities"`
}

const METHOD_NAME = "init"

func NewInit() *Init {
	return NewInitFromArgs(system.CheckCapabilities())
}

func NewInitFromArgs(capabilities []string) *Init {
	return &Init{
		Architecture: runtime.GOARCH,
		OS:           runtime.GOOS,
		Capabilities: capabilities,
	}
}

func DecodeInit(decoder *msgpack.Decoder) (*Init, error) {
	var ping Init
	convertErr := decoder.Decode(&ping)
	if convertErr != nil {
		return nil, convertErr
	}
	return &ping, nil
}

func (p *Init) String() string {
	return fmt.Sprintf("INIT <ARCH: %s, OS: %s>", p.Architecture, p.OS)
}

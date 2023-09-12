package base

import (
	"fmt"
	"runtime"

	"github.com/Artiqlate/ganymede/models"
	"github.com/Artiqlate/ganymede/system"
	"github.com/vmihailenco/msgpack/v5"
)

// TODO: Remove field names (we use array-style msgpack)
type Init struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack     struct{} `msgpack:",as_array"`
	Architecture string   `msgpack:"arch"`
	OS           string   `msgpack:"os"`
	Capabilities []string `msgpack:"capabilities"`
}

const METHOD_NAME = "init"

func NewInit() *Init {
	return NewInitWithCapabilities(system.CheckCapabilities())
}

func NewInitWithCapabilities(capabilities []string) *Init {
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

func (i *Init) GenMessage(method string) *models.Message {
	return &models.Message{
		Method: method,
		Args:   i,
	}
}

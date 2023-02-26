package base

//lint:file-ignore U1000 `msgpack` options, not for serialization.

type InitModules struct {
	_msgpack       struct{} `msgpack:",as_array"`
	EnabledModules []string `msgpack:"enabledModules"`
}

func NewInitModules(enabledModules []string) *InitModules {
	return &InitModules{
		EnabledModules: enabledModules,
	}
}

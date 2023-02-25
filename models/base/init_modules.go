package base

type InitModules struct {
	_msgpack       struct{} `msgpack:",as_array"`
	EnabledModules []string `msgpack:"enabledModules"`
}

func NewInitModules(enabledModules []string) *InitModules {
	return &InitModules{
		EnabledModules: enabledModules,
	}
}

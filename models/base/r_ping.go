package base

type RPing struct {
	Ping
	EnabledCapabilities []string
}

func NewRPing(enabledCapabilities []string) *RPing {
	return &RPing{
		*NewPing(),
		enabledCapabilities,
	}
}

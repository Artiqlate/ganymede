package base

type RPing struct {
	Init
	EnabledCapabilities []string
}

func NewRPing(enabledCapabilities []string) *RPing {
	return &RPing{
		*NewInit(),
		enabledCapabilities,
	}
}

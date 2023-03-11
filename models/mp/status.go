package mp

type MPlayerStatus struct {
	_msgpack    struct{} `msgpack:",as_array"`
	PlayerIndex int
	PlayStatus  string
}

type Status struct {
	_msgpack struct{}            `msgpack:",as_array"`
	Status   string              `msgpack:"status"`
	Index    int                 `msgpack:"idx"`
	Name     string              `msgpack:"name"`
	Metadata MediaPlayerMetadata `msgpack:"metadata"`
}

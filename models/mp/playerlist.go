package mp

type MPlayerList struct {
	_msgpack struct{} `msgpack:",as_array"`
	Players  []string
}

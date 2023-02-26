package mp

type MPlayerStatus struct {
	_msgpack   struct{} `msgpack:",as_array"`
	PlayStatus string
}

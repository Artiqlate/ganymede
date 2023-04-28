package mp

// TODO: Remove field names (we use array-style msgpack)
type PropertiesChangedSignal struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack      struct{}      `msgpack:",as_array"`
	PlayerName    string        `msgpack:"playerName"`
	PlayerIdx     int           `msgpack:"playerIdx"`
	PlayerDataVal *FullMetadata `msgpack:"playerData"`
}

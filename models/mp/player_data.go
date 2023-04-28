package mp

// TODO: Remove field names (we use array-style msgpack)
type PlayerData struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack       struct{} `msgpack:",as_array"`
	PlayerName     string
	PlaybackStatus string
	Metadata       Metadata
}

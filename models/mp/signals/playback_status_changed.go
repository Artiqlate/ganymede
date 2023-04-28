package mp_signals

// TODO: Remove field names (we use array-style msgpack)
type PlaybackStatusChanged struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack       struct{} `msgpack:",as_array"`
	PlayerIndex    int      `msgpack:"playerIdx"`
	PlayerName     string   `msgpack:"playerName"`
	PlaybackStatus string   `msgpack:"playbackStatus"`
}

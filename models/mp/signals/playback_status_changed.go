package mp_signals

//lint:ignore U1000 `msgpack` options, not for serialization.

type PlaybackStatusChanged struct {
	_msgpack       struct{} `msgpack:",as_array"`
	PlayerIndex    int      `msgpack:"playerIdx"`
	PlayerName     string   `msgpack:"playerName"`
	PlaybackStatus string   `msgpack:"playbackStatus"`
}

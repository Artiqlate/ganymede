package mp_signals

// TODO: Remove field names (we use array-style msgpack)
type Seeked struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack    struct{} `msgpack:",as_array"`
	PlayerIndex int      `msgpack:"playerIdx"`
	PlayerName  string   `msgpack:"playerName"`
	// "seeked" time in microseconds (us).
	SeekedInUs int64 `msgpack:"seekedInUs"`
}

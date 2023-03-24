package mp_signals

type Seeked struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack   struct{} `msgpack:",as_array"`
	PlayerName string   `msgpack:"playerName"`
	PlayerIdx  int      `msgpack:"playerIdx"`
	// "seeked" time in microseconds (us).
	SeekedInUs int64 `msgpack:"seekedInUs"`
}

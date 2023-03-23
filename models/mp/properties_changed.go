package mp

//lint:ignore U1000 `msgpack` options, not for serialization.

type PropertiesChangedSignal struct {
	_msgpack      struct{}    `msgpack:",as_array"`
	PlayerName    string      `msgpack:"playerName"`
	PlayerIdx     int         `msgpack:"playerIdx"`
	PlayerDataVal *PlayerData `msgpack:"playerData"`
}

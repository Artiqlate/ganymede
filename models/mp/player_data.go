package mp

type PlayerData struct {
	_msgpack       struct{} `msgpack:",as_array"`
	PlayerName     string
	PlaybackStatus string
	Metadata       Metadata
}

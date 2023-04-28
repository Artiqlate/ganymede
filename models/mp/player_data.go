//lint:file-ignore U1000 `msgpack` options, not for serialization.
package mp

type PlayerData struct {
	_msgpack       struct{} `msgpack:",as_array"`
	PlayerName     string
	PlaybackStatus string
	Metadata       Metadata
}

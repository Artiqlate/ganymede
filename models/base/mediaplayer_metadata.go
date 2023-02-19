package base

type MediaPlayerMetadata struct {
	_msgpack struct{} `msgpack:",as_array"`
	Title    string   `msgpack:"title"`
	Artist   string   `msgpack:"artist"`
}

func NewMediaPlayerMetadata(title string, artist string) *MediaPlayerMetadata {
	return &MediaPlayerMetadata{}
}

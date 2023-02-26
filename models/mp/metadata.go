/*
Media Player Metadata
REFERENCE: https://www.freedesktop.org/wiki/Specifications/mpris-spec/metadata/
*/

package mp

//lint:file-ignore U1000 `msgpack` options, not for serialization.

import (
	"github.com/godbus/dbus/v5"
)

const (
	TRACKID = "xesam:trackid"
	LENGTH  = "xesam:trackid"
	TITLE   = "xesam:title"
	ARTIST  = "xesam:artist"
	// ALBUM   = "xesam:album"
	// ALBUM_ARTIST = "xesam:albumArtist"
)

type MediaPlayerMetadata struct {
	_msgpack struct{} `msgpack:",as_array"`
	TrackId  string   `msgpack:"trackid"`
	Length   int64    `msgpack:"length"`
	Title    string   `msgpack:"title"`
	Artist   []string `msgpack:"artist"`
}

func NewMediaPlayerMetadata(title string, artist string) *MediaPlayerMetadata {
	return &MediaPlayerMetadata{}
}

// func StoreMetadataValue(metadata map[string]dbus.Variant, key string, value interface{}) error {
// 	metadataParseErr := metadata[key].Store(value)
// 	if metadataParseErr != nil {
// 		return metadataParseErr
// 	}
// 	return nil
// }

func MediaPlayerFromMpris(metadata map[string]dbus.Variant) *MediaPlayerMetadata {
	var trackId string
	var length int64
	var title string
	var artist []string
	var ok bool

	trackId, ok = metadata[TRACKID].Value().(string)
	if !ok {
		trackId = ""
	}
	length, ok = metadata[LENGTH].Value().(int64)
	if !ok {
		length = -1
	}
	title, ok = metadata[TITLE].Value().(string)
	if !ok {
		title = ""
	}
	artist, ok = metadata[ARTIST].Value().([]string)
	if !ok {
		artist = []string{}
	}

	return &MediaPlayerMetadata{
		TrackId: trackId,
		Length:  length,
		Title:   title,
		Artist:  artist,
	}
}

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

type Metadata struct {
	_msgpack struct{} `msgpack:",as_array"`
	TrackId  string   `msgpack:"trackid"`
	Length   int64    `msgpack:"length"`
	Title    string   `msgpack:"title"`
	Artist   []string `msgpack:"artist"`
}

func MetadataFromMPRIS(metadata map[string]dbus.Variant) *Metadata {
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

	return &Metadata{
		TrackId: trackId,
		Length:  length,
		Title:   title,
		Artist:  artist,
	}
}

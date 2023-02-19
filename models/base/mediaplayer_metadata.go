/*
Media Player Metadata
REFERENCE: https://www.freedesktop.org/wiki/Specifications/mpris-spec/metadata/
*/
package base

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
	// var trackId string
	// var length int64
	// var title string
	// var artist []string
	// We have to do this individually because of Go's constraints
	// trackid
	// metaErr := StoreMetadataValue(metadata, TRACKID, &trackId)
	// metaErr := metadata[TRACKID].Store(trackId)
	// if metaErr != nil {
	// 	fmt.Println(metaErr)
	// 	metaErr = nil
	// }
	// metaErr = StoreMetadataValue(metadata, LENGTH, &length)
	// if metaErr != nil {
	// 	fmt.Println(metaErr)
	// 	metaErr = nil
	// }
	// metaErr = StoreMetadataValue(metadata, TITLE, &title)
	// if metaErr != nil {
	// 	fmt.Println(metaErr)
	// 	metaErr = nil
	// }
	// metaErr = StoreMetadataValue(metadata, ARTIST, &artist)
	// if metaErr != nil {
	// 	fmt.Println(metaErr)
	// }

	return &MediaPlayerMetadata{
		TrackId: metadata[TRACKID].Value().(string),
		Length:  metadata[LENGTH].Value().(int64),
		Title:   metadata[TITLE].Value().(string),
		Artist:  metadata[ARTIST].Value().([]string),
	}
}

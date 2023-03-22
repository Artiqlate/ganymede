/*
Media Player Metadata

REFERENCE: https://www.freedesktop.org/wiki/Specifications/mpris-spec/metadata/

NOTE: This is currently linux-only.

TODO: Add Windows/macOS Implementations also.

Copyright (C) 2023 Goutham Krishna K V
*/
package mp

//lint:file-ignore U1000 `msgpack` options, not for serialization.

import (
	"github.com/godbus/dbus/v5"
)

const (
	TRACKID        = "xesam:trackid"
	LENGTH         = "xesam:length"
	LENGTH_SPOTIFY = "mpris:length:@t"
	TITLE          = "xesam:title"
	ARTIST         = "xesam:artist"
	ALBUM          = "xesam:album"
	ALBUM_ARTIST   = "xesam:albumArtist"
	URL            = `xesam:url`
	ART_URL        = "xesam:artUrl"
)

type Metadata struct {
	_msgpack    struct{} `msgpack:",as_array"`
	TrackId     string   `msgpack:"trackid"`
	Length      int64    `msgpack:"length"`
	Title       string   `msgpack:"title"`
	Artist      []string `msgpack:"artist"`
	Album       string   `msgpack:"album"`
	AlbumArtist []string `msgpack:"albumArtist"`
	Url         string   `msgpack:"url"`
	ArtUrl      string   `msgpack:"artUrl"`
}

func MetadataFromMPRIS(metadata map[string]dbus.Variant) *Metadata {
	var trackId string
	var length int64
	var title string
	var artist []string
	var album string
	var albumArtist []string
	var url string
	var artUrl string
	var ok bool

	// -- PARSE THE VALUES
	trackId, ok = metadata[TRACKID].Value().(string)
	if !ok {
		trackId = ""
	}
	length, ok = metadata[LENGTH].Value().(int64)
	if !ok {
		length = -1
	}
	// LENGTH: Spotify Special Case
	if length == -1 {
		if length, ok = metadata[LENGTH_SPOTIFY].Value().(int64); !ok {
			length = -1
		}
	}
	title, ok = metadata[TITLE].Value().(string)
	if !ok {
		title = ""
	}
	artist, ok = metadata[ARTIST].Value().([]string)
	if !ok {
		artist = []string{}
	}
	album, ok = metadata[ALBUM].Value().(string)
	if !ok {
		album = ""
	}
	albumArtist, ok = metadata[ALBUM_ARTIST].Value().([]string)
	if !ok {
		albumArtist = []string{}
	}
	url, ok = metadata[URL].Value().(string)
	if !ok {
		url = ""
	}
	artUrl, ok = metadata[ART_URL].Value().(string)
	if !ok {
		artUrl = ""
	}

	return &Metadata{
		TrackId:     trackId,
		Length:      length,
		Title:       title,
		Artist:      artist,
		Album:       album,
		AlbumArtist: albumArtist,
		Url:         url,
		ArtUrl:      artUrl,
	}
}

/*
Media Player Metadata

REFERENCE: https://www.freedesktop.org/wiki/Specifications/mpris-spec/metadata/

NOTE: This is currently linux-only.

TODO: Add Windows/macOS Implementations also.

Copyright (C) 2023 Goutham Krishna K V
*/
package mp

import (
	"github.com/godbus/dbus/v5"
)

const (
	// MPRIS specific
	TRACKID = "mpris:trackid"
	LENGTH  = "mpris:length"
	ART_URL = "mpris:artUrl"
	// Xesam specific
	TITLE        = "xesam:title"
	ARTIST       = "xesam:artist"
	ALBUM        = "xesam:album"
	ALBUM_ARTIST = "xesam:albumArtist"
	URL          = `xesam:url`
	// TODO: Add other metadata fields also
)

// TODO: Remove field names (we use array-style msgpack)
type Metadata struct {
	// Common Metadata Fields
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack    struct{} `msgpack:",as_array"`
	TrackId     string   `msgpack:"trackid"`
	Length      uint64   `msgpack:"length"`
	Title       string   `msgpack:"title"`
	Artist      []string `msgpack:"artist"`
	Album       string   `msgpack:"album"`
	AlbumArtist []string `msgpack:"albumArtist"`
	Url         string   `msgpack:"url"`
	ArtUrl      string   `msgpack:"artUrl"`
	// TODO: Add other metadata fields also
}

func MetadataFromMPRIS(metadata map[string]dbus.Variant) *Metadata {
	var trackId string
	var length uint64
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
	length, ok = metadata[LENGTH].Value().(uint64)
	if !ok {
		length = 0
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

package mp

/*
Player Data

This includes all the data related to a player (to be played).

Copyright (C) 2023 Goutham Krishna K V
*/

//lint:ignore U1000 `msgpack` options, not for serialization.

type PlayerData struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack       struct{}  `msgpack:",as_array"`
	PlaybackStatus string    `msgpack:"playbackStatus"` // RO
	LoopStatus     *string   `msgpack:"loopStatus"`     // RW (Optional)
	Rate           float64   `msgpack:"rate"`           // RW
	Shuffle        *bool     `msgpack:"shuffleEnabled"` // RW (Optional)
	Metadata       *Metadata `msgpack:"metadata"`       // RO
	Volume         float64   `msgpack:"volume"`         // RW
	Position       int64     `msgpack:"position"`       // RO
	MinimumRate    float64   `msgpack:"minimumRate"`    // RO
	MaximumRate    float64   `msgpack:"maximumRate"`    // RO
	// Capabilities
	// Enable 'next', 'previous', 'play', 'pause', 'seek' capabilities with these.
	CanGoNext     bool `msgpack:"canGoNext"`     // RO
	CanGoPrevious bool `msgpack:"canGoPrevious"` // RO
	CanPlay       bool `msgpack:"canPlay"`       // RO
	CanPause      bool `msgpack:"canPause"`      // RO
	CanSeek       bool `msgpack:"canSeek"`       // RO
	// This will assume all properties are read-only (frozen)
	CanControl bool `msgpack:"canControl"` // RO
}

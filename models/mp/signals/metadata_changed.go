package mp_signals

import "github.com/Artiqlate/ganymede/models/mp"

// TODO: Remove field names (we use array-style msgpack)
type MetadataChanged struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack    struct{}     `msgpack:",as_array"`
	PlayerIndex int          `msgpack:"playerIdx"`
	PlayerName  string       `msgpack:"playerName"`
	Metadata    *mp.Metadata `msgpack:"metadata"`
}

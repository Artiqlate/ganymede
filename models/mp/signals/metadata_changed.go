package mp_signals

import "github.com/CrosineEnterprises/ganymede/models/mp"

//lint:ignore U1000 `msgpack` options, not for serialization.

type MetadataChanged struct {
	_msgpack    struct{}     `msgpack:",as_array"`
	PlayerIndex int          `msgpack:"playerIdx"`
	PlayerName  string       `msgpack:"playerName"`
	Metadata    *mp.Metadata `msgpack:"metadata"`
}

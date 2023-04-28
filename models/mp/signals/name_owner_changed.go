package mp_signals

import "github.com/CrosineEnterprises/ganymede/models/mp"

// TODO: Remove field names (we use array-style msgpack)
type PlayerCreated struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack struct{} `msgpack:",as_array"`
	// This returns the updated player list
	mp.PlayerData
	UpdatedPlayerNames []string
}

// TODO: Remove field names (we use array-style msgpack)
type PlayerUpdated struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack struct{} `msgpack:",as_array"`
	mp.PlayerData
	UpdatedPlayerNames []string
}

// TODO: Remove field names (we use array-style msgpack)
type PlayerRemoved struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack           struct{} `msgpack:",as_array"`
	PlayerName         string
	UpdatedPlayerNames []string
}

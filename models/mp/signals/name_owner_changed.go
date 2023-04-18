package mp_signals

import "github.com/CrosineEnterprises/ganymede/models/mp"

type PlayerCreated struct {
	_msgpack struct{} `msgpack:",as_array"`
	// This returns the updated player list
	mp.PlayerData
	UpdatedPlayerNames []string
}

type PlayerUpdated struct {
	_msgpack struct{} `msgpack:",as_array"`
	mp.PlayerData
	UpdatedPlayerNames []string
}

type PlayerRemoved struct {
	_msgpack           struct{} `msgpack:",as_array"`
	PlayerName         string
	UpdatedPlayerNames []string
}

package mp

// TODO: Remove field names (we use array-style msgpack)
// TODO: Rename it to better reflect it's purpose.
type SetupStatus struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack struct{} `msgpack:",as_array"`
	Statuses []Status `msgpack:"statuses"`
}

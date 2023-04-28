package mp

// TODO: Remove field names (we use array-style msgpack)
type Status struct {
	//lint:ignore U1000 `msgpack` options, not for serialization.
	_msgpack struct{} `msgpack:",as_array"`
	Status   string   `msgpack:"status"`
	Index    int      `msgpack:"idx"`
	Name     string   `msgpack:"name"`
	Metadata Metadata `msgpack:"metadata"`
}

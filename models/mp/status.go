//lint:file-ignore U1000 `msgpack` options, not for serialization.
package mp

type Status struct {
	_msgpack struct{} `msgpack:",as_array"`
	Status   string   `msgpack:"status"`
	Index    int      `msgpack:"idx"`
	Name     string   `msgpack:"name"`
	Metadata Metadata `msgpack:"metadata"`
}

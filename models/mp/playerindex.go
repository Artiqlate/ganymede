//lint:file-ignore U1000 `msgpack` options, not for serialization.
package mp

type PlayerIndex struct {
	_msgpack    struct{} `msgpack:",as_array"`
	PlayerIndex int
}

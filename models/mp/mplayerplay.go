//lint:file-ignore U1000 `msgpack` options, not for serialization.
package mp

type MPlayerPlay struct {
	_msgpack    struct{} `msgpack:",as_array"`
	PlayerIndex int
}

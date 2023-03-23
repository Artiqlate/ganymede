//lint:file-ignore U1000 `msgpack` options, not for serialization.w
package mp

type MPlayerList struct {
	_msgpack struct{} `msgpack:",as_array"`
	Players  []string
}

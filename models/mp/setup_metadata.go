//lint:file-ignore U1000 `msgpack` options, not for serialization.
package mp

type SetupStatus struct {
	_msgpack struct{} `msgpack:",as_array"`
	Statuses []Status `msgpack:"statuses"`
}

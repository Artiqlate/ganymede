package mp

type SetupStatus struct {
	_msgpack struct{} `msgpack:",as_array"`
	Statuses []Status `msgpack:"statuses"`
}

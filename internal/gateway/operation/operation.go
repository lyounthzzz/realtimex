package operation

const (
	Connect int32 = iota + 1
	Disconnect
	Ping
	Pong
	Publish
)

var status = map[int32]string{
	Connect:    "connect",
	Disconnect: "disconnect",
	Ping:       "ping",
	Pong:       "pong",
}

func Text(op int32) string {
	return status[op]
}

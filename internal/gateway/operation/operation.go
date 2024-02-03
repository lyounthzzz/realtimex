package operation

const (
	Connect     = 1
	Connack     = 2
	Publish     = 3
	Puback      = 4
	Pubrec      = 5
	Pubrel      = 6
	Pubcomp     = 7
	Subscribe   = 8
	Suback      = 9
	Unsubscribe = 10
	Unsuback    = 11
	Pingreq     = 12
	Pingresp    = 13
	Disconnect  = 14
)

var status = map[int]string{
	Connect:     "connect",
	Connack:     "connack",
	Publish:     "publish",
	Puback:      "puback",
	Pubrec:      "pubrec",
	Pubrel:      "pubrel",
	Pubcomp:     "pubcomp",
	Subscribe:   "subscribe",
	Suback:      "suback",
	Unsubscribe: "unsubscribe",
	Unsuback:    "unsuback",
	Pingreq:     "pingreq",
	Pingresp:    "pingresp",
	Disconnect:  "disconnect",
}

func Text(op int) string {
	return status[op]
}

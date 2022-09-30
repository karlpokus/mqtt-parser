package main

type Packet struct {
	Header, ControlPacket string
}

// PacketType maps control packet type in hex to string
var PacketType = map[string]string{
	"1": "CONNECT",
	"2": "CONNACK",
	"3": "PUBLISH",
	"4": "PUBACK",
	"5": "PUBREC",
	"6": "PUBREL",
	"7": "PUBCOMP",
	"8": "SUBSCRIBE",
	"9": "SUBACK",
	"a": "UNSUBSCRIBE",
	"b": "UNSUBACK",
	"c": "PINGREQ",
	"d": "PINGRESP",
	"e": "DISCONNECT",
}

// parsePacketType maps single hex string to PacketType
func parsePacketType(s string) string {
	v, ok := PacketType[s]
	if ok {
		return v
	}
	return "undefined"
}

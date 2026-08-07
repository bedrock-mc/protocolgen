// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type WebSocketPacketData struct {
	WebsocketServerURI string
}

// Marshal reads or writes WebSocketPacketData using its canonical wire layout.
func (x *WebSocketPacketData) Marshal(io IO) {
	io.String(&x.WebsocketServerURI)
}

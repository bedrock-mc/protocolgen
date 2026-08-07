// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type WebSocketData struct {
	WebsocketServerURI string
}

// Marshal reads or writes WebSocketData using its canonical wire layout.
func (x *WebSocketData) Marshal(io IO) {
	io.String(&x.WebsocketServerURI)
}

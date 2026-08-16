// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// ToastRequest is a packet sent from the server to the client to display a toast to the top of the
// screen. These toasts are the same as the ones seen when, for example, loading a new resource pack
// or obtaining an achievement.
type ToastRequest struct {
	// Title is the title of the toast.
	Title string
	// Content is the message that the toast may contain alongside the title.
	Content string
}

// Marshal reads or writes ToastRequest using its canonical wire layout.
func (x *ToastRequest) Marshal(io protocol.IO) {
	io.String(&x.Title)
	io.String(&x.Content)
}

// ID returns the protocol ID for ToastRequest.
func (*ToastRequest) ID() uint32 { return IDToastRequest }

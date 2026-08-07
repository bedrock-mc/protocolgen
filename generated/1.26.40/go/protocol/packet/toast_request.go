// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ToastRequest struct {
	Title   string
	Content string
}

// Marshal reads or writes ToastRequest using its canonical wire layout.
func (x *ToastRequest) Marshal(io protocol.IO) {
	io.String(&x.Title)
	io.String(&x.Content)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ServerboundDataDrivenScreenClosed struct {
	FormId      uint32
	CloseReason string
}

// Marshal reads or writes ServerboundDataDrivenScreenClosed using its canonical wire layout.
func (x *ServerboundDataDrivenScreenClosed) Marshal(io protocol.IO) {
	io.Uint32(&x.FormId)
	io.String(&x.CloseReason)
}

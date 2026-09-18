// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type ServerboundDataDrivenScreenClosed struct {
	FormID      uint32
	CloseReason string
}

// ID returns the protocol ID for ServerboundDataDrivenScreenClosed.
func (*ServerboundDataDrivenScreenClosed) ID() uint32 { return IDServerboundDataDrivenScreenClosed }

// Marshal reads or writes ServerboundDataDrivenScreenClosed using its canonical wire layout.
func (pk *ServerboundDataDrivenScreenClosed) Marshal(io protocol.IO) {
	io.Uint32(&pk.FormID)
	io.String(&pk.CloseReason)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type ServerboundDataDrivenScreenClosed struct {
	FormID      uint32
	CloseReason string
}

// ID ...
func (*ServerboundDataDrivenScreenClosed) ID() uint32 {
	return IDServerboundDataDrivenScreenClosed
}

func (pk *ServerboundDataDrivenScreenClosed) Marshal(io protocol.IO) {
	io.Uint32(&pk.FormID)
	io.String(&pk.CloseReason)
}

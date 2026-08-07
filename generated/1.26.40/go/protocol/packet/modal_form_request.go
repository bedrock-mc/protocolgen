// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ModalFormRequest struct {
	FormID     uint32
	FormUIJSON string
}

// Marshal reads or writes ModalFormRequest using its canonical wire layout.
func (x *ModalFormRequest) Marshal(io protocol.IO) {
	io.Varuint32(&x.FormID)
	io.String(&x.FormUIJSON)
}

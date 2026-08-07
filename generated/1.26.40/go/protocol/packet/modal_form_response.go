// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ModalFormResponse struct {
	FormID           uint32
	JSONResponse     protocol.Optional[string]
	FormCancelReason protocol.Optional[protocol.ModalFormCancelReason]
}

// Marshal reads or writes ModalFormResponse using its canonical wire layout.
func (x *ModalFormResponse) Marshal(io protocol.IO) {
	io.Varuint32(&x.FormID)
	protocol.OptionalFunc(io, &x.JSONResponse, io.String)
	protocol.OptionalFunc(io, &x.FormCancelReason, func(value *protocol.ModalFormCancelReason) {
		protocol.IntegerFunc(value, io.Uint8)
	})
}

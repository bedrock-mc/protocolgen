// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// ModalFormRequest is sent by the server to make the client open a form. This form may be either a
// modal form which has two options, a menu form for a selection of options and a custom form for
// properties.
type ModalFormRequest struct {
	// FormID is an ID used to identify the form. The ID is saved by the client and sent back when the
	// player submits the form, so that the server can identify which form was submitted.
	FormID     uint32
	FormUIJSON string
}

// Marshal reads or writes ModalFormRequest using its canonical wire layout.
func (x *ModalFormRequest) Marshal(io protocol.IO) {
	io.Varuint32(&x.FormID)
	protocol.Minimum(io, &x.FormID, 0)
	io.String(&x.FormUIJSON)
}

// ID returns the protocol ID for ModalFormRequest.
func (*ModalFormRequest) ID() uint32 { return IDModalFormRequest }

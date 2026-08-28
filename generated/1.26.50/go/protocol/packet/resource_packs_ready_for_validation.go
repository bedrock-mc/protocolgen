// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// ResourcePacksReadyForValidation is sent by the client to inform the server that the client has
// finished loading resource packs and is ready for validation.
type ResourcePacksReadyForValidation struct {
}

// Marshal reads or writes ResourcePacksReadyForValidation using its canonical wire layout.
func (x *ResourcePacksReadyForValidation) Marshal(io protocol.IO) {
}

// ID returns the protocol ID for ResourcePacksReadyForValidation.
func (*ResourcePacksReadyForValidation) ID() uint32 { return IDResourcePacksReadyForValidation }

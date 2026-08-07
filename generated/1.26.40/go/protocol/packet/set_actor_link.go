// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetActorLink struct {
	Link protocol.EntityLink
}

// Marshal reads or writes SetActorLink using its canonical wire layout.
func (x *SetActorLink) Marshal(io protocol.IO) {
	x.Link.Marshal(io)
}

// ID returns the protocol ID for SetActorLink.
func (*SetActorLink) ID() uint32 { return IDSetActorLink }

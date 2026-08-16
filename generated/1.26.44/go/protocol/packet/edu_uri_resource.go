// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

type EduURIResource struct {
	EduSharedURIResource protocol.EduSharedURIResource
}

// Marshal reads or writes EduURIResource using its canonical wire layout.
func (x *EduURIResource) Marshal(io protocol.IO) {
	x.EduSharedURIResource.Marshal(io)
}

// ID returns the protocol ID for EduURIResource.
func (*EduURIResource) ID() uint32 { return IDEduURIResource }

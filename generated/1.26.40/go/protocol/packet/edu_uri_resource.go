// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type EduUriResource struct {
	EduSharedURIResource protocol.EduSharedUriResource
}

// Marshal reads or writes EduUriResource using its canonical wire layout.
func (x *EduUriResource) Marshal(io protocol.IO) {
	x.EduSharedURIResource.Marshal(io)
}

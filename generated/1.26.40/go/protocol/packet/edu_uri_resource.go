// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type EduURIResource struct {
	EduSharedURIResource protocol.EduSharedURIResource
}

// ID ...
func (*EduURIResource) ID() uint32 {
	return IDEduURIResource
}

func (pk *EduURIResource) Marshal(io protocol.IO) {
	pk.EduSharedURIResource.Marshal(io)
}

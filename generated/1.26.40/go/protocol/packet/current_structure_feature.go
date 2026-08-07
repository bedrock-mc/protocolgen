// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CurrentStructureFeature struct {
	CurrentStructureFeature string
}

// Marshal reads or writes CurrentStructureFeature using its canonical wire layout.
func (x *CurrentStructureFeature) Marshal(io protocol.IO) {
	io.String(&x.CurrentStructureFeature)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// CurrentStructureFeature is sent by the server to let the client know the name of the structure
// feature that the player is currently occupying.
type CurrentStructureFeature struct {
	// CurrentStructureFeature is the identifier of the structure feature that the player is currently
	// occupying. If the player is not occupying any structure feature, this field is empty.
	CurrentStructureFeature string
}

// Marshal reads or writes CurrentStructureFeature using its canonical wire layout.
func (x *CurrentStructureFeature) Marshal(io protocol.IO) {
	io.String(&x.CurrentStructureFeature)
}

// ID returns the protocol ID for CurrentStructureFeature.
func (*CurrentStructureFeature) ID() uint32 { return IDCurrentStructureFeature }

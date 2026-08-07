// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type RemoveObjective struct {
	ObjectiveName string
}

// Marshal reads or writes RemoveObjective using its canonical wire layout.
func (x *RemoveObjective) Marshal(io protocol.IO) {
	io.String(&x.ObjectiveName)
}

// ID returns the protocol ID for RemoveObjective.
func (*RemoveObjective) ID() uint32 { return IDRemoveObjective }

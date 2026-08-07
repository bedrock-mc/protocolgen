// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type MoveActorDelta struct {
	MoveData protocol.MoveActorDeltaData
}

// Marshal reads or writes MoveActorDelta using its canonical wire layout.
func (x *MoveActorDelta) Marshal(io protocol.IO) {
	x.MoveData.Marshal(io)
}

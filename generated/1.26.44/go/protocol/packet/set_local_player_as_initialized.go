// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

type SetLocalPlayerAsInitialized struct {
	PlayerID uint64
}

// ID returns the protocol ID for SetLocalPlayerAsInitialized.
func (*SetLocalPlayerAsInitialized) ID() uint32 { return IDSetLocalPlayerAsInitialized }

// Marshal reads or writes SetLocalPlayerAsInitialized using its canonical wire layout.
func (pk *SetLocalPlayerAsInitialized) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&pk.PlayerID)
}

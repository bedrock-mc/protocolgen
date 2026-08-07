// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetLocalPlayerAsInitialized struct {
	PlayerID uint64
}

// Marshal reads or writes SetLocalPlayerAsInitialized using its canonical wire layout.
func (x *SetLocalPlayerAsInitialized) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.PlayerID)
}

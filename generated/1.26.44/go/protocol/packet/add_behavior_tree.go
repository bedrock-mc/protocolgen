// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// AddBehaviorTree is sent by the server to the client. The packet is currently unused by both client and
// server.
type AddBehaviorTree struct {
	// BehaviourTree is an unused string.
	BehaviorTreeStructureJSON string
}

// ID returns the protocol ID for AddBehaviorTree.
func (*AddBehaviorTree) ID() uint32 { return IDAddBehaviorTree }

// Marshal reads or writes AddBehaviorTree using its canonical wire layout.
func (pk *AddBehaviorTree) Marshal(io protocol.IO) {
	io.String(&pk.BehaviorTreeStructureJSON)
}

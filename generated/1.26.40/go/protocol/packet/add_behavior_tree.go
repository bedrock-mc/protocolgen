// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type AddBehaviorTree struct {
	BehaviorTreeStructureJSON string
}

// ID returns the protocol ID for AddBehaviorTree.
func (*AddBehaviorTree) ID() uint32 { return IDAddBehaviorTree }

// Marshal reads or writes AddBehaviorTree using its canonical wire layout.
func (pk *AddBehaviorTree) Marshal(io protocol.IO) {
	io.String(&pk.BehaviorTreeStructureJSON)
}

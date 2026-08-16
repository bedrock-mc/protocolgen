// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

type AddBehaviorTree struct {
	BehaviorTreeStructureJSON string
}

// Marshal reads or writes AddBehaviorTree using its canonical wire layout.
func (x *AddBehaviorTree) Marshal(io protocol.IO) {
	io.String(&x.BehaviorTreeStructureJSON)
}

// ID returns the protocol ID for AddBehaviorTree.
func (*AddBehaviorTree) ID() uint32 { return IDAddBehaviorTree }

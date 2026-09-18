// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type AddBehaviorTree struct {
	BehaviorTreeStructureJSON string
}

// ID ...
func (*AddBehaviorTree) ID() uint32 {
	return IDAddBehaviorTree
}

func (pk *AddBehaviorTree) Marshal(io protocol.IO) {
	io.String(&pk.BehaviorTreeStructureJSON)
}

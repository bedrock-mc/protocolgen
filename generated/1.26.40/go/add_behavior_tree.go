// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AddBehaviorTree struct {
	BehaviorTreeStructureJSON string
}

// Marshal reads or writes AddBehaviorTree using its canonical wire layout.
func (x *AddBehaviorTree) Marshal(io IO) {
	io.String(&x.BehaviorTreeStructureJSON)
}

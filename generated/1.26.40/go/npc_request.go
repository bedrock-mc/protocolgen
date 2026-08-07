// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type NpcRequest struct {
	NPCRuntimeID uint64
	RequestType  NpcRequestRequestType
	Actions      string
	ActionIndex  uint8
	SceneName    string
}

// Marshal reads or writes NpcRequest using its canonical wire layout.
func (x *NpcRequest) Marshal(io IO) {
	io.ActorRuntimeID(&x.NPCRuntimeID)
	IntegerFunc(&x.RequestType, io.Uint8)
	io.String(&x.Actions)
	io.Uint8(&x.ActionIndex)
	io.String(&x.SceneName)
}

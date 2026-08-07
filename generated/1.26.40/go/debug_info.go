// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type DebugInfo struct {
	ActorId ActorUniqueID
	Data    string
}

// Marshal reads or writes DebugInfo using its canonical wire layout.
func (x *DebugInfo) Marshal(io IO) {
	x.ActorId.Marshal(io)
	io.String(&x.Data)
}

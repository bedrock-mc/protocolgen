// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type DebugInfo struct {
	ActorId int64
	Data    string
}

// Marshal reads or writes DebugInfo using its canonical wire layout.
func (x *DebugInfo) Marshal(io IO) {
	io.ActorUniqueID(&x.ActorId)
	io.String(&x.Data)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MapInfoRequest struct {
	MapUniqueID      int64
	ClientPixelsList []MapInfoRequestPacketAnonClientPixelsProxy
}

// Marshal reads or writes MapInfoRequest using its canonical wire layout.
func (x *MapInfoRequest) Marshal(io IO) {
	io.ActorUniqueID(&x.MapUniqueID)
	FuncSlice(io, &x.ClientPixelsList, io.Uint32, func(value *MapInfoRequestPacketAnonClientPixelsProxy) {
		value.Marshal(io)
	})
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MapInfoRequest struct {
	MapUniqueID      ActorUniqueID
	ClientPixelsList []MapInfoRequestPacketAnonClientPixelsProxy
}

// Marshal reads or writes MapInfoRequest using its canonical wire layout.
func (x *MapInfoRequest) Marshal(io IO) {
	x.MapUniqueID.Marshal(io)
	if !io.Reading() && uint64(len(x.ClientPixelsList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ClientPixelsList), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.ClientPixelsList))
	io.Uint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.ClientPixelsList = make([]MapInfoRequestPacketAnonClientPixelsProxy, int(count1))
	}
	for index2 := range x.ClientPixelsList {
		x.ClientPixelsList[index2].Marshal(io)
	}
}

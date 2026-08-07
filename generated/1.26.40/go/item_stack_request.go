// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequest struct {
	Requests []ItemStackRequestPacketDataRequestData
}

// Marshal reads or writes ItemStackRequest using its canonical wire layout.
func (x *ItemStackRequest) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Requests)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Requests), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.Requests))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.Requests = make([]ItemStackRequestPacketDataRequestData, int(count1))
	}
	for index2 := range x.Requests {
		x.Requests[index2].Marshal(io)
	}
}

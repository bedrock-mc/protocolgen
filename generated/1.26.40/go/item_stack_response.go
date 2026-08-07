// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackResponse struct {
	Responses []ItemStackResponseInfo
}

// Marshal reads or writes ItemStackResponse using its canonical wire layout.
func (x *ItemStackResponse) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Responses)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Responses), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.Responses))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.Responses = make([]ItemStackResponseInfo, int(count1))
	}
	for index2 := range x.Responses {
		x.Responses[index2].Marshal(io)
	}
}

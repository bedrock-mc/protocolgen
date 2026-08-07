// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemRegistry struct {
	ItemData []ItemData
}

// Marshal reads or writes ItemRegistry using its canonical wire layout.
func (x *ItemRegistry) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.ItemData)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ItemData), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.ItemData))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.ItemData = make([]ItemData, int(count1))
	}
	for index2 := range x.ItemData {
		x.ItemData[index2].Marshal(io)
	}
}

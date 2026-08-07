// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerFog struct {
	FogStack []string
}

// Marshal reads or writes PlayerFog using its canonical wire layout.
func (x *PlayerFog) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.FogStack)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.FogStack), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.FogStack))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.FogStack = make([]string, int(count1))
	}
	for index2 := range x.FogStack {
		io.String(&x.FogStack[index2])
	}
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerEnchantOptions struct {
	Options []ItemEnchantOption
}

// Marshal reads or writes PlayerEnchantOptions using its canonical wire layout.
func (x *PlayerEnchantOptions) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Options)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Options), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.Options))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.Options = make([]ItemEnchantOption, int(count1))
	}
	for index2 := range x.Options {
		x.Options[index2].Marshal(io)
	}
}

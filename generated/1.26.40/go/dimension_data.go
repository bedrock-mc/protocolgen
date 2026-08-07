// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type DimensionData struct {
	Definitions []OrderedEntry[string, DimensionDefinition]
}

// Marshal reads or writes DimensionData using its canonical wire layout.
func (x *DimensionData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Definitions)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Definitions), "map length overflows uint32")
		return
	}
	count1 := uint32(len(x.Definitions))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "map length overflows int")
			return
		}
		x.Definitions = make([]OrderedEntry[string, DimensionDefinition], int(count1))
	}
	for index2 := range x.Definitions {
		io.String(&x.Definitions[index2].Key)
		x.Definitions[index2].Value.Marshal(io)
	}
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundDataStore struct {
	Updates []BedrockDDUI
}

// Marshal reads or writes ClientboundDataStore using its canonical wire layout.
func (x *ClientboundDataStore) Marshal(io IO) {
	FuncSlice(io, &x.Updates, io.Varuint32, func(value *BedrockDDUI) {
		item := *value
		marshalBedrockDDUI(io, &item)
		*value = item
	})
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BedrockDDUIDataStoreUpdateDataDouble struct {
	Value float64
}

func (*BedrockDDUIDataStoreUpdateDataDouble) isBedrockDDUIDataStoreUpdateData() {}

// Marshal reads or writes BedrockDDUIDataStoreUpdateDataDouble using its canonical wire layout.
func (x *BedrockDDUIDataStoreUpdateDataDouble) Marshal(io IO) {
	io.Float64(&x.Value)
}

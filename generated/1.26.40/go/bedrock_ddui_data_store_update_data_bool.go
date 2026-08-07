// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BedrockDDUIDataStoreUpdateDataBool struct {
	Value bool
}

func (BedrockDDUIDataStoreUpdateDataBool) isBedrockDDUIDataStoreUpdateData() {}

// Marshal reads or writes BedrockDDUIDataStoreUpdateDataBool using its canonical wire layout.
func (x *BedrockDDUIDataStoreUpdateDataBool) Marshal(io IO) {
	io.Bool(&x.Value)
}

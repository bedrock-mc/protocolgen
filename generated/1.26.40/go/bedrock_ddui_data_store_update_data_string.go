// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BedrockDDUIDataStoreUpdateDataString struct {
	Value string
}

func (BedrockDDUIDataStoreUpdateDataString) isBedrockDDUIDataStoreUpdateData() {}

// Marshal reads or writes BedrockDDUIDataStoreUpdateDataString using its canonical wire layout.
func (x *BedrockDDUIDataStoreUpdateDataString) Marshal(io IO) {
	io.String(&x.Value)
}

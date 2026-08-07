// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BedrockDDUIDataStoreRemoval struct {
	DataStoreName string
}

func (BedrockDDUIDataStoreRemoval) isBedrockDDUI() {}

// Marshal reads or writes BedrockDDUIDataStoreRemoval using its canonical wire layout.
func (x *BedrockDDUIDataStoreRemoval) Marshal(io IO) {
	io.String(&x.DataStoreName)
}

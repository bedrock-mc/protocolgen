// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BedrockDDUIDataStoreChange struct {
	DataStoreName       string
	Property            string
	UpdateCount         uint32
	TheNewPropertyValue CerealDynamicValue
}

func (*BedrockDDUIDataStoreChange) isBedrockDDUI() {}

// Marshal reads or writes BedrockDDUIDataStoreChange using its canonical wire layout.
func (x *BedrockDDUIDataStoreChange) Marshal(io IO) {
	io.String(&x.DataStoreName)
	io.String(&x.Property)
	io.Uint32(&x.UpdateCount)
	MarshalCerealDynamicValue(io, &x.TheNewPropertyValue)
}

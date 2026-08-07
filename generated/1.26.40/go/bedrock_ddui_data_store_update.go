// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BedrockDDUIDataStoreUpdate struct {
	DataStoreName       string
	Property            string
	Path                string
	Data                BedrockDDUIDataStoreUpdateData
	PropertyUpdateCount uint32
	PathUpdateCount     uint32
}

func (BedrockDDUIDataStoreUpdate) isBedrockDDUI() {}

// Marshal reads or writes BedrockDDUIDataStoreUpdate using its canonical wire layout.
func (x *BedrockDDUIDataStoreUpdate) Marshal(io IO) {
	io.String(&x.DataStoreName)
	io.String(&x.Property)
	io.String(&x.Path)
	marshalBedrockDDUIDataStoreUpdateData(io, &x.Data)
	io.Uint32(&x.PropertyUpdateCount)
	io.Uint32(&x.PathUpdateCount)
}

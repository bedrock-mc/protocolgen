// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerboundDataStore struct {
	Update BedrockDDUIDataStoreUpdate
}

// Marshal reads or writes ServerboundDataStore using its canonical wire layout.
func (x *ServerboundDataStore) Marshal(io IO) {
	x.Update.Marshal(io)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ServerConfigurationClientStoreEntryPointConfiguration struct {
	StoreId   string
	StoreName string
}

// Marshal reads or writes ServerConfigurationClientStoreEntryPointConfiguration using its canonical wire layout.
func (x *ServerConfigurationClientStoreEntryPointConfiguration) Marshal(io IO) {
	io.String(&x.StoreId)
	io.String(&x.StoreName)
}

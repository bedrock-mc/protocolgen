// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerStoreInfo struct {
	ClientStoreEntryPointConfiguration Optional[ServerConfigurationClientStoreEntryPointConfiguration]
}

// Marshal reads or writes ServerStoreInfo using its canonical wire layout.
func (x *ServerStoreInfo) Marshal(io IO) {
	OptionalFunc(io, &x.ClientStoreEntryPointConfiguration, func(value *ServerConfigurationClientStoreEntryPointConfiguration) {
		value.Marshal(io)
	})
}

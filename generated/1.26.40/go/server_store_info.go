// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerStoreInfo struct {
	ClientStoreEntryPointConfiguration Optional[ServerConfigurationClientStoreEntryPointConfiguration]
}

// Marshal reads or writes ServerStoreInfo using its canonical wire layout.
func (x *ServerStoreInfo) Marshal(io IO) {
	io.Bool(&x.ClientStoreEntryPointConfiguration.set)
	if x.ClientStoreEntryPointConfiguration.set {
		x.ClientStoreEntryPointConfiguration.val.Marshal(io)
	} else if io.Reading() {
		var zero ServerConfigurationClientStoreEntryPointConfiguration
		x.ClientStoreEntryPointConfiguration.val = zero
	}
}

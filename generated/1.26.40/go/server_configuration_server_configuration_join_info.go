// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerConfigurationServerConfigurationJoinInfo struct {
	Gathering             Optional[ServerConfigurationGatheringsConfigurationJoinInfo]
	ClientStoreEntryPoint Optional[ServerConfigurationClientStoreEntryPointConfiguration]
	Presence              Optional[ServerConfigurationPresenceConfiguration]
}

// Marshal reads or writes ServerConfigurationServerConfigurationJoinInfo using its canonical wire layout.
func (x *ServerConfigurationServerConfigurationJoinInfo) Marshal(io IO) {
	OptionalFunc(io, &x.Gathering, func(value *ServerConfigurationGatheringsConfigurationJoinInfo) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.ClientStoreEntryPoint, func(value *ServerConfigurationClientStoreEntryPointConfiguration) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.Presence, func(value *ServerConfigurationPresenceConfiguration) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}

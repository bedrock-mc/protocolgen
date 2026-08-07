// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ServerConfigurationServerConfigurationJoinInfo struct {
	Gathering             Optional[ServerConfigurationGatheringsConfigurationJoinInfo]
	ClientStoreEntryPoint Optional[ServerConfigurationClientStoreEntryPointConfiguration]
	Presence              Optional[ServerConfigurationPresenceConfiguration]
}

// Marshal reads or writes ServerConfigurationServerConfigurationJoinInfo using its canonical wire layout.
func (x *ServerConfigurationServerConfigurationJoinInfo) Marshal(io IO) {
	OptionalFunc(io, &x.Gathering, func(value *ServerConfigurationGatheringsConfigurationJoinInfo) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.ClientStoreEntryPoint, func(value *ServerConfigurationClientStoreEntryPointConfiguration) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Presence, func(value *ServerConfigurationPresenceConfiguration) {
		value.Marshal(io)
	})
}

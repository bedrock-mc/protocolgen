// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type Transfer struct {
	ServerAddress           string
	ServerPort              uint16
	ReloadWorld             bool
	GatheringsConfiguration Optional[ServerConfigurationGatheringsConfigurationJoinInfo]
}

// Marshal reads or writes Transfer using its canonical wire layout.
func (x *Transfer) Marshal(io IO) {
	io.String(&x.ServerAddress)
	io.Uint16(&x.ServerPort)
	io.Bool(&x.ReloadWorld)
	OptionalFunc(io, &x.GatheringsConfiguration, func(value *ServerConfigurationGatheringsConfigurationJoinInfo) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}

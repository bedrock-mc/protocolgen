// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type Transfer struct {
	ServerAddress           string
	ServerPort              uint16
	ReloadWorld             bool
	GatheringsConfiguration protocol.Optional[protocol.ServerConfigurationGatheringsConfigurationJoinInfo]
}

// Marshal reads or writes Transfer using its canonical wire layout.
func (x *Transfer) Marshal(io protocol.IO) {
	io.String(&x.ServerAddress)
	io.Uint16(&x.ServerPort)
	io.Bool(&x.ReloadWorld)
	protocol.OptionalFunc(io, &x.GatheringsConfiguration, func(value *protocol.ServerConfigurationGatheringsConfigurationJoinInfo) {
		value.Marshal(io)
	})
}

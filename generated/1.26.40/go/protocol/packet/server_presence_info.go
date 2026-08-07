// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ServerPresenceInfo struct {
	PresenceConfiguration protocol.Optional[protocol.ServerConfigurationPresenceConfiguration]
}

// Marshal reads or writes ServerPresenceInfo using its canonical wire layout.
func (x *ServerPresenceInfo) Marshal(io protocol.IO) {
	protocol.OptionalFunc(io, &x.PresenceConfiguration, func(value *protocol.ServerConfigurationPresenceConfiguration) {
		value.Marshal(io)
	})
}

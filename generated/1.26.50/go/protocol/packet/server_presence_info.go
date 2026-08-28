// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// ServerPresenceInfo is sent by the server to provide the client with presence info.
type ServerPresenceInfo struct {
	// PresenceConfiguration is the presence info to set, or nothing to fall back to the default.
	PresenceConfiguration protocol.Optional[protocol.ServerConfigurationPresenceConfiguration]
}

// Marshal reads or writes ServerPresenceInfo using its canonical wire layout.
func (x *ServerPresenceInfo) Marshal(io protocol.IO) {
	protocol.OptionalFunc(io, &x.PresenceConfiguration, func(value *protocol.ServerConfigurationPresenceConfiguration) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for ServerPresenceInfo.
func (*ServerPresenceInfo) ID() uint32 { return IDServerPresenceInfo }

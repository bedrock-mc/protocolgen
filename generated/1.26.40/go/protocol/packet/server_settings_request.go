// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// ServerSettingsRequest is sent by the client to request the settings specific to the server. These settings
// are shown in a separate tab client-side, and have the same structure as a custom form.
type ServerSettingsRequest struct {
}

// ID returns the protocol ID for ServerSettingsRequest.
func (*ServerSettingsRequest) ID() uint32 { return IDServerSettingsRequest }

// Marshal reads or writes ServerSettingsRequest using its canonical wire layout.
func (pk *ServerSettingsRequest) Marshal(io protocol.IO) {
}

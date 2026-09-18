// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// ServerStats is a packet sent from the server to the client to update the client on server statistics. It is
// purely used for telemetry.
type ServerStats struct {
	// ServerTime ...
	ServerTime float32
	// NetworkTime ...
	NetworkTime float32
}

// ID returns the protocol ID for ServerStats.
func (*ServerStats) ID() uint32 { return IDServerStats }

// Marshal reads or writes ServerStats using its canonical wire layout.
func (pk *ServerStats) Marshal(io protocol.IO) {
	io.Float32(&pk.ServerTime)
	io.Float32(&pk.NetworkTime)
}

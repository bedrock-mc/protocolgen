// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// TickingAreasLoadStatus is sent by the server to the client to notify the client of a ticking
// area's loading status.
type TickingAreasLoadStatus struct {
	// Preload is true if the server is waiting for the area's preload.
	WaitingForPreload bool
}

// Marshal reads or writes TickingAreasLoadStatus using its canonical wire layout.
func (x *TickingAreasLoadStatus) Marshal(io protocol.IO) {
	io.Bool(&x.WaitingForPreload)
}

// ID returns the protocol ID for TickingAreasLoadStatus.
func (*TickingAreasLoadStatus) ID() uint32 { return IDTickingAreasLoadStatus }

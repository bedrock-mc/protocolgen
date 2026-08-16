// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// LocatorBar is sent by the server to add, remove or update waypoints on the client's locator bar.
type LocatorBar struct {
	// Waypoints is a slice of waypoints to add, remove or update.
	Waypoints []protocol.LocatorBarWaypoint
}

// Marshal reads or writes LocatorBar using its canonical wire layout.
func (x *LocatorBar) Marshal(io protocol.IO) {
	protocol.SliceLimits(io, &x.Waypoints, 0, 40000)
}

// ID returns the protocol ID for LocatorBar.
func (*LocatorBar) ID() uint32 { return IDLocatorBar }

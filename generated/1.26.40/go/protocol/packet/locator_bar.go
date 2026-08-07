// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type LocatorBar struct {
	Waypoints []protocol.LocatorBarWaypoint
}

// Marshal reads or writes LocatorBar using its canonical wire layout.
func (x *LocatorBar) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.Waypoints)
}

// ID returns the protocol ID for LocatorBar.
func (*LocatorBar) ID() uint32 { return IDLocatorBar }

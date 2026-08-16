// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

// LocatorBarWaypoint represents a waypoint entry in the locator bar packet.
type LocatorBarWaypoint struct {
	// GroupHandle is the UUID handle for the waypoint group.
	GroupHandle           WaypointGroupWaypointHandle
	ServerWaypointPayload ServerWaypoint
	ActionFlag            ServerWaypointGroupAction
}

// Marshal reads or writes LocatorBarWaypoint using its canonical wire layout.
func (x *LocatorBarWaypoint) Marshal(io IO) {
	x.GroupHandle.Marshal(io)
	x.ServerWaypointPayload.Marshal(io)
	IntegerFunc(&x.ActionFlag, io.Uint8)
}

type WaypointGroupWaypointHandle struct {
	UUID uuid.UUID
}

// Marshal reads or writes WaypointGroupWaypointHandle using its canonical wire layout.
func (x *WaypointGroupWaypointHandle) Marshal(io IO) {
	io.UUID(&x.UUID)
}

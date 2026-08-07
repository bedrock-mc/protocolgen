// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LocatorBarWaypoint struct {
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

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type WaypointGroupWaypointHandle struct {
	UUID uuid.UUID
}

// Marshal reads or writes WaypointGroupWaypointHandle using its canonical wire layout.
func (x *WaypointGroupWaypointHandle) Marshal(io IO) {
	io.UUID(&x.UUID)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LocatorBar struct {
	Waypoints []LocatorBarWaypoint
}

// Marshal reads or writes LocatorBar using its canonical wire layout.
func (x *LocatorBar) Marshal(io IO) {
	FuncSlice(io, &x.Waypoints, io.Varuint32, func(value *LocatorBarWaypoint) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LocatorBar struct {
	Waypoints []LocatorBarWaypoint
}

// Marshal reads or writes LocatorBar using its canonical wire layout.
func (x *LocatorBar) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.Waypoints)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Waypoints), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.Waypoints))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.Waypoints = make([]LocatorBarWaypoint, int(count1))
	}
	for index2 := range x.Waypoints {
		x.Waypoints[index2].Marshal(io)
	}
}

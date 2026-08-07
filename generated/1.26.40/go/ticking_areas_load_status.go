// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type TickingAreasLoadStatus struct {
	WaitingForPreload bool
}

// Marshal reads or writes TickingAreasLoadStatus using its canonical wire layout.
func (x *TickingAreasLoadStatus) Marshal(io IO) {
	io.Bool(&x.WaitingForPreload)
}

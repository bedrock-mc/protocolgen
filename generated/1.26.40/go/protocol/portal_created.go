// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PortalCreated struct {
	DimensionID int32
}

func (*PortalCreated) isEventData() {}

// Marshal reads or writes PortalCreated using its canonical wire layout.
func (x *PortalCreated) Marshal(io IO) {
	io.Varint32(&x.DimensionID)
}

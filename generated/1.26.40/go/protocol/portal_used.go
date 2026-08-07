// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PortalUsed struct {
	SourceDimensionID int32
	TargetDimensionID int32
}

func (*PortalUsed) isEventData() {}

// Marshal reads or writes PortalUsed using its canonical wire layout.
func (x *PortalUsed) Marshal(io IO) {
	io.Varint32(&x.SourceDimensionID)
	io.Varint32(&x.TargetDimensionID)
}

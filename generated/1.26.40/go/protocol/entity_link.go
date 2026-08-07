// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type EntityLink struct {
	TargetA                int64
	TargetB                int64
	Type                   ActorLinkType
	Immediate              bool
	PassengerInitiated     bool
	VehicleAngularVelocity float32
}

// Marshal reads or writes EntityLink using its canonical wire layout.
func (x *EntityLink) Marshal(io IO) {
	io.ActorUniqueID(&x.TargetA)
	io.ActorUniqueID(&x.TargetB)
	IntegerFunc(&x.Type, io.Uint8)
	io.Bool(&x.Immediate)
	io.Bool(&x.PassengerInitiated)
	io.Float32(&x.VehicleAngularVelocity)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

// EntityLink is a link between two entities, typically being one entity riding another.
type EntityLink struct {
	TargetA int64
	TargetB int64
	// Type is one of the types above. It specifies the way the entity is linked to another entity.
	Type ActorLinkType
	// Immediate is set to immediately dismount an entity from another. This should be set when the
	// mount of an entity is killed.
	Immediate          bool
	PassengerInitiated bool
	// VehicleAngularVelocity is the angular velocity of the vehicle that the rider is riding.
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

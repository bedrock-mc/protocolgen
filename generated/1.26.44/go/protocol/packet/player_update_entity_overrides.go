// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// PlayerUpdateEntityOverrides is sent by the server to modify an entity's properties individually.
type PlayerUpdateEntityOverrides struct {
	TargetID int64
	// PropertyIndex is the index of the property to modify. The index is unique for each property of an
	// entity.
	PropertyIndex uint32
	Update        protocol.PlayerUpdateEntityOverridesData
}

// Marshal reads or writes PlayerUpdateEntityOverrides using its canonical wire layout.
func (x *PlayerUpdateEntityOverrides) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.TargetID)
	io.Varuint32(&x.PropertyIndex)
	protocol.Minimum(io, &x.PropertyIndex, 0)
	protocol.MarshalPlayerUpdateEntityOverridesData(io, &x.Update)
}

// ID returns the protocol ID for PlayerUpdateEntityOverrides.
func (*PlayerUpdateEntityOverrides) ID() uint32 { return IDPlayerUpdateEntityOverrides }

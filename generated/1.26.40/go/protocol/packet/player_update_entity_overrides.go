// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayerUpdateEntityOverrides struct {
	TargetID      int64
	PropertyIndex uint32
	Update        protocol.PlayerUpdateEntityOverridesUpdate
}

// Marshal reads or writes PlayerUpdateEntityOverrides using its canonical wire layout.
func (x *PlayerUpdateEntityOverrides) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.TargetID)
	io.Varuint32(&x.PropertyIndex)
	protocol.MarshalPlayerUpdateEntityOverridesUpdate(io, &x.Update)
}

// ID returns the protocol ID for PlayerUpdateEntityOverrides.
func (*PlayerUpdateEntityOverrides) ID() uint32 { return IDPlayerUpdateEntityOverrides }

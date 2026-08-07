// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type Animate struct {
	Action               protocol.AnimateAction
	TargetActorRuntimeID uint64
	Data                 float32
	SwingSource          protocol.Optional[string]
}

// Marshal reads or writes Animate using its canonical wire layout.
func (x *Animate) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.Action, io.Uint8)
	io.ActorRuntimeID(&x.TargetActorRuntimeID)
	io.Float32(&x.Data)
	protocol.OptionalFunc(io, &x.SwingSource, io.String)
}

// ID returns the protocol ID for Animate.
func (*Animate) ID() uint32 { return IDAnimate }

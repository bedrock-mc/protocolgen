// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// CameraInstruction gives a custom camera specific instructions to operate.
type CameraInstruction struct {
	Set              protocol.Optional[protocol.CameraInstructionSet]
	Clear            protocol.Optional[bool]
	Fade             protocol.Optional[protocol.CameraInstructionFade]
	Target           protocol.Optional[protocol.CameraInstructionTargetData]
	RemoveTarget     protocol.Optional[bool]
	FieldOfView      protocol.Optional[protocol.CameraInstructionFieldOfView]
	Spline           protocol.Optional[protocol.CameraSplineInstruction]
	AttachToEntity   protocol.Optional[protocol.CameraInstructionTarget]
	DetachFromEntity protocol.Optional[bool]
}

// ID ...
func (*CameraInstruction) ID() uint32 {
	return IDCameraInstruction
}

func (pk *CameraInstruction) Marshal(io protocol.IO) {
	protocol.OptionalMarshaler(io, &pk.Set)
	protocol.OptionalFunc(io, &pk.Clear, io.Bool)
	protocol.OptionalMarshaler(io, &pk.Fade)
	protocol.OptionalMarshaler(io, &pk.Target)
	protocol.OptionalFunc(io, &pk.RemoveTarget, io.Bool)
	protocol.OptionalMarshaler(io, &pk.FieldOfView)
	protocol.OptionalMarshaler(io, &pk.Spline)
	protocol.OptionalMarshaler(io, &pk.AttachToEntity)
	protocol.OptionalFunc(io, &pk.DetachFromEntity, io.Bool)
}

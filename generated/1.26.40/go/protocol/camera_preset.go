// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type CameraPreset struct {
	Name                    string
	InheritFrom             string
	PosX                    Optional[float32]
	PosY                    Optional[float32]
	PosZ                    Optional[float32]
	RotX                    Optional[float32]
	RotY                    Optional[float32]
	RotationSpeed           Optional[float32]
	SnapToTarget            Optional[bool]
	HorizontalRotationLimit Optional[mgl32.Vec2]
	VerticalRotationLimit   Optional[mgl32.Vec2]
	ContinueTargeting       Optional[bool]
	BlockListeningRadius    Optional[float32]
	ViewOffset              Optional[mgl32.Vec2]
	EntityOffset            Optional[mgl32.Vec3]
	Radius                  Optional[float32]
	YawLimitMin             Optional[float32]
	YawLimitMax             Optional[float32]
	Listener                Optional[CameraPresetAudioListener]
	PlayerEffects           Optional[bool]
	AimAssist               Optional[CameraAimAssistCommandPresetDefinition]
	ControlScheme           Optional[ControlScheme]
}

// Marshal reads or writes CameraPreset using its canonical wire layout.
func (x *CameraPreset) Marshal(io IO) {
	io.String(&x.Name)
	io.String(&x.InheritFrom)
	OptionalFunc(io, &x.PosX, io.Float32)
	OptionalFunc(io, &x.PosY, io.Float32)
	OptionalFunc(io, &x.PosZ, io.Float32)
	OptionalFunc(io, &x.RotX, io.Float32)
	OptionalFunc(io, &x.RotY, io.Float32)
	OptionalFunc(io, &x.RotationSpeed, io.Float32)
	OptionalFunc(io, &x.SnapToTarget, io.Bool)
	OptionalFunc(io, &x.HorizontalRotationLimit, io.Vec2)
	OptionalFunc(io, &x.VerticalRotationLimit, io.Vec2)
	OptionalFunc(io, &x.ContinueTargeting, io.Bool)
	OptionalFunc(io, &x.BlockListeningRadius, io.Float32)
	OptionalFunc(io, &x.ViewOffset, io.Vec2)
	OptionalFunc(io, &x.EntityOffset, io.Vec3)
	OptionalFunc(io, &x.Radius, io.Float32)
	OptionalFunc(io, &x.YawLimitMin, io.Float32)
	OptionalFunc(io, &x.YawLimitMax, io.Float32)
	OptionalFunc(io, &x.Listener, func(value *CameraPresetAudioListener) {
		IntegerFunc(value, io.Uint8)
	})
	OptionalFunc(io, &x.PlayerEffects, io.Bool)
	OptionalFunc(io, &x.AimAssist, func(value *CameraAimAssistCommandPresetDefinition) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.ControlScheme, func(value *ControlScheme) {
		IntegerFunc(value, io.Uint8)
	})
}

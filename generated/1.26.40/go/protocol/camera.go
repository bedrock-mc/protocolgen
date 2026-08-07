// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type CameraAimAssistAction uint8

const (
	CameraAimAssistActionSet   CameraAimAssistAction = 0
	CameraAimAssistActionClear CameraAimAssistAction = 1
)

// CameraAimAssistActorPriorityData represents priority data for aim assist actor targeting.
type CameraAimAssistActorPriorityData struct {
	// PresetIndex is the index of the aim assist preset.
	PresetIndex int32
	// CategoryIndex is the index of the aim assist category.
	CategoryIndex int32
	// ActorIndex is the index of the actor.
	ActorIndex int32
	// PriorityValue is the priority value for this actor.
	PriorityValue int32
}

// Marshal reads or writes CameraAimAssistActorPriorityData using its canonical wire layout.
func (x *CameraAimAssistActorPriorityData) Marshal(io IO) {
	io.Int32(&x.PresetIndex)
	io.Int32(&x.CategoryIndex)
	io.Int32(&x.ActorIndex)
	io.Int32(&x.PriorityValue)
}

type CameraAimAssistCategoryDefinition struct {
	Name       string
	Priorities CameraAimAssistCategoryPriorities
}

// Marshal reads or writes CameraAimAssistCategoryDefinition using its canonical wire layout.
func (x *CameraAimAssistCategoryDefinition) Marshal(io IO) {
	io.String(&x.Name)
	x.Priorities.Marshal(io)
}

type CameraAimAssistCategoryPriorities struct {
	Entities           []OrderedEntry[string, int32]
	Blocks             []OrderedEntry[string, int32]
	BlockTags          []OrderedEntry[string, int32]
	EntityTypeFamilies []OrderedEntry[string, int32]
	EntityDefault      Optional[int32]
	BlockDefault       Optional[int32]
}

// Marshal reads or writes CameraAimAssistCategoryPriorities using its canonical wire layout.
func (x *CameraAimAssistCategoryPriorities) Marshal(io IO) {
	OrderedMap(io, &x.Entities, io.Varuint32, io.String, io.Int32)
	OrderedMap(io, &x.Blocks, io.Varuint32, io.String, io.Int32)
	OrderedMap(io, &x.BlockTags, io.Varuint32, io.String, io.Int32)
	OrderedMap(io, &x.EntityTypeFamilies, io.Varuint32, io.String, io.Int32)
	OptionalFunc(io, &x.EntityDefault, io.Int32)
	OptionalFunc(io, &x.BlockDefault, io.Int32)
}

type CameraAimAssistCommandPresetDefinition struct {
	PresetID   Optional[string]
	TargetMode Optional[CameraAimAssistTargetMode]
	ViewAngle  Optional[mgl32.Vec2]
	Distance   Optional[float32]
}

// Marshal reads or writes CameraAimAssistCommandPresetDefinition using its canonical wire layout.
func (x *CameraAimAssistCommandPresetDefinition) Marshal(io IO) {
	OptionalFunc(io, &x.PresetID, io.String)
	OptionalFunc(io, &x.TargetMode, func(value *CameraAimAssistTargetMode) {
		IntegerFunc(value, io.Int32)
	})
	OptionalFunc(io, &x.ViewAngle, io.Vec2)
	OptionalFunc(io, &x.Distance, io.Float32)
}

type CameraAimAssistPresetDefinition struct {
	Identifier          string
	ExclusionSettings   CameraAimAssistPresetExclusionDefinition
	LiquidTargetingList []string
	ItemSettings        []OrderedEntry[string, string]
	DefaultItemSettings Optional[string]
	HandSettings        Optional[string]
}

// Marshal reads or writes CameraAimAssistPresetDefinition using its canonical wire layout.
func (x *CameraAimAssistPresetDefinition) Marshal(io IO) {
	io.String(&x.Identifier)
	x.ExclusionSettings.Marshal(io)
	FuncSlice(io, &x.LiquidTargetingList, io.Varuint32, io.String)
	OrderedMap(io, &x.ItemSettings, io.Varuint32, io.String, io.String)
	OptionalFunc(io, &x.DefaultItemSettings, io.String)
	OptionalFunc(io, &x.HandSettings, io.String)
}

type CameraAimAssistPresetExclusionDefinition struct {
	Blocks             []string
	Entities           []string
	BlockTags          []string
	EntityTypeFamilies []string
}

// Marshal reads or writes CameraAimAssistPresetExclusionDefinition using its canonical wire layout.
func (x *CameraAimAssistPresetExclusionDefinition) Marshal(io IO) {
	FuncSlice(io, &x.Blocks, io.Varuint32, io.String)
	FuncSlice(io, &x.Entities, io.Varuint32, io.String)
	FuncSlice(io, &x.BlockTags, io.Varuint32, io.String)
	FuncSlice(io, &x.EntityTypeFamilies, io.Varuint32, io.String)
}

type CameraAimAssistPresetOperation uint8

const (
	CameraAimAssistPresetOperationSet           CameraAimAssistPresetOperation = 0
	CameraAimAssistPresetOperationAddToExisting CameraAimAssistPresetOperation = 1
)

type CameraAimAssistTargetMode int32

const (
	CameraAimAssistTargetModeAngle    CameraAimAssistTargetMode = 0
	CameraAimAssistTargetModeDistance CameraAimAssistTargetMode = 1
)

// CameraEase represents an easing function that can be used by a CameraInstructionSet.
type CameraEase struct {
	// Type is the type of easing function used. This is one of the constants above.
	Type uint8
	// Time is the time in seconds that the easing function should take.
	Time float32
}

// Marshal reads or writes CameraEase using its canonical wire layout.
func (x *CameraEase) Marshal(io IO) {
	io.Uint8(&x.Type)
	io.Float32(&x.Time)
}

type CameraEntityOffset struct {
	EntityOffsetX float32
	EntityOffsetY float32
	EntityOffsetZ float32
}

// Marshal reads or writes CameraEntityOffset using its canonical wire layout.
func (x *CameraEntityOffset) Marshal(io IO) {
	io.Float32(&x.EntityOffsetX)
	io.Float32(&x.EntityOffsetY)
	io.Float32(&x.EntityOffsetZ)
}

type CameraFacing struct {
	Pos mgl32.Vec3
}

// Marshal reads or writes CameraFacing using its canonical wire layout.
func (x *CameraFacing) Marshal(io IO) {
	io.Vec3(&x.Pos)
}

type CameraFadeColor struct {
	Red   float32
	Green float32
	Blue  float32
}

// Marshal reads or writes CameraFadeColor using its canonical wire layout.
func (x *CameraFadeColor) Marshal(io IO) {
	io.Float32(&x.Red)
	io.Float32(&x.Green)
	io.Float32(&x.Blue)
}

// CameraFadeTimeData represents the time data for a CameraInstructionFade.
type CameraFadeTimeData struct {
	// FadeInTime is the time in seconds for the screen to fully fade in.
	FadeInTime float32
	// HoldTime is time in seconds to wait before fading out.
	HoldTime float32
	// FadeOutTime is the time in seconds for the screen to fully fade out.
	FadeOutTime float32
}

// Marshal reads or writes CameraFadeTimeData using its canonical wire layout.
func (x *CameraFadeTimeData) Marshal(io IO) {
	io.Float32(&x.FadeInTime)
	io.Float32(&x.HoldTime)
	io.Float32(&x.FadeOutTime)
}

type CameraInstructionData struct {
	Set              Optional[CameraInstructionSet]
	Clear            Optional[bool]
	Fade             Optional[CameraInstructionFade]
	Target           Optional[CameraInstructionTargetData]
	RemoveTarget     Optional[bool]
	FieldOfView      Optional[CameraInstructionFieldOfView]
	Spline           Optional[CameraSplineInstruction]
	AttachToEntity   Optional[CameraInstructionTarget]
	DetachFromEntity Optional[bool]
}

// Marshal reads or writes CameraInstructionData using its canonical wire layout.
func (x *CameraInstructionData) Marshal(io IO) {
	OptionalFunc(io, &x.Set, func(value *CameraInstructionSet) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Clear, io.Bool)
	OptionalFunc(io, &x.Fade, func(value *CameraInstructionFade) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Target, func(value *CameraInstructionTargetData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.RemoveTarget, io.Bool)
	OptionalFunc(io, &x.FieldOfView, func(value *CameraInstructionFieldOfView) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Spline, func(value *CameraSplineInstruction) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.AttachToEntity, func(value *CameraInstructionTarget) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.DetachFromEntity, io.Bool)
}

// CameraInstructionFade represents a camera instruction that fades the screen to a specified
// colour.
type CameraInstructionFade struct {
	// Time is the time data for the fade, which includes the fade in duration, wait duration and fade
	// out duration.
	Time Optional[CameraFadeTimeData]
	// Color is the colour of the screen to fade to. This only uses the red, green and blue components.
	Color Optional[CameraFadeColor]
}

// Marshal reads or writes CameraInstructionFade using its canonical wire layout.
func (x *CameraInstructionFade) Marshal(io IO) {
	OptionalFunc(io, &x.Time, func(value *CameraFadeTimeData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Color, func(value *CameraFadeColor) {
		value.Marshal(io)
	})
}

// CameraInstructionFieldOfView represents a camera instruction that updates the field of view.
type CameraInstructionFieldOfView struct {
	// FieldOfView is the field of view of the camera.
	FieldOfView      float32
	FOVEaseTime      float32
	FOVEaseType      string
	FieldOfViewClear bool
}

// Marshal reads or writes CameraInstructionFieldOfView using its canonical wire layout.
func (x *CameraInstructionFieldOfView) Marshal(io IO) {
	io.Float32(&x.FieldOfView)
	io.Float32(&x.FOVEaseTime)
	io.String(&x.FOVEaseType)
	io.Bool(&x.FieldOfViewClear)
}

// CameraInstructionSet represents a camera instruction that sets the camera to a specified preset
// and can be extended with easing functions and translations to the camera's position and rotation.
type CameraInstructionSet struct {
	// Preset is the index of the preset in the CameraPresets packet sent to the player.
	Preset uint32
	// Ease represents the easing function that is used by the instruction.
	Ease Optional[CameraEase]
	// Pos represents the position of the camera.
	Pos Optional[CameraPosition]
	// Rot represents the rotation of the camera.
	Rot Optional[CameraRotation]
	// Facing is a vector that the camera will always face towards during the duration of the
	// instruction.
	Facing Optional[CameraFacing]
	// ViewOffset is an offset based on a pivot point to the player, causing the camera to be shifted in
	// a certain direction.
	ViewOffset Optional[CameraViewOffset]
	// EntityOffset is an offset from the entity that the camera should be rendered at.
	EntityOffset Optional[CameraEntityOffset]
	// Default determines whether the camera is a default camera or not.
	Default Optional[bool]
	// RemoveIgnoreStartingValuesComponent behavior is currently unknown.
	RemoveIgnoreStartingValuesComponent bool
}

// Marshal reads or writes CameraInstructionSet using its canonical wire layout.
func (x *CameraInstructionSet) Marshal(io IO) {
	io.Uint32(&x.Preset)
	OptionalFunc(io, &x.Ease, func(value *CameraEase) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Pos, func(value *CameraPosition) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Rot, func(value *CameraRotation) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Facing, func(value *CameraFacing) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.ViewOffset, func(value *CameraViewOffset) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.EntityOffset, func(value *CameraEntityOffset) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Default, io.Bool)
	io.Bool(&x.RemoveIgnoreStartingValuesComponent)
}

// CameraInstructionTarget represents a camera instruction that targets a specific entity.
type CameraInstructionTarget struct {
	EntityActorID int64
}

// Marshal reads or writes CameraInstructionTarget using its canonical wire layout.
func (x *CameraInstructionTarget) Marshal(io IO) {
	io.Int64(&x.EntityActorID)
}

// CameraInstructionTarget represents a camera instruction that targets a specific entity.
type CameraInstructionTargetData struct {
	// TargetCenterOffset is the offset from the center of the entity that the camera should target.
	TargetCenterOffset Optional[mgl32.Vec3]
	// TargetActorID is the unique ID of the entity that the camera should target.
	TargetActorID int64
}

// Marshal reads or writes CameraInstructionTargetData using its canonical wire layout.
func (x *CameraInstructionTargetData) Marshal(io IO) {
	OptionalFunc(io, &x.TargetCenterOffset, io.Vec3)
	io.Int64(&x.TargetActorID)
}

type CameraPosition struct {
	Pos mgl32.Vec3
}

// Marshal reads or writes CameraPosition using its canonical wire layout.
func (x *CameraPosition) Marshal(io IO) {
	io.Vec3(&x.Pos)
}

// CameraPreset represents a basic preset that can be extended upon by more complex instructions.
type CameraPreset struct {
	// Name is the name of the preset. Each preset must have their own unique name.
	Name string
	// InheritFrom is the name of the preset that this preset extends upon. This can be left empty.
	InheritFrom string
	// PosX is the default X position of the camera.
	PosX Optional[float32]
	// PosY is the default Y position of the camera.
	PosY Optional[float32]
	// PosZ is the default Z position of the camera.
	PosZ Optional[float32]
	// RotX is the default pitch of the camera.
	RotX Optional[float32]
	// RotY is the default yaw of the camera.
	RotY Optional[float32]
	// RotationSpeed is the speed at which the camera should rotate.
	RotationSpeed Optional[float32]
	// SnapToTarget determines whether the camera should snap to the target entity or not.
	SnapToTarget Optional[bool]
	// HorizontalRotationLimit is the horizontal rotation limit of the camera.
	HorizontalRotationLimit Optional[mgl32.Vec2]
	// VerticalRotationLimit is the vertical rotation limit of the camera.
	VerticalRotationLimit Optional[mgl32.Vec2]
	// ContinueTargeting determines whether the camera should continue targeting when using aim assist.
	ContinueTargeting Optional[bool]
	// BlockListeningRadius is the radius around the camera that the aim assist should track targets.
	BlockListeningRadius Optional[float32]
	// ViewOffset is only used in a follow_orbit camera and controls an offset based on a pivot point to
	// the player, causing it to be shifted in a certain direction.
	ViewOffset Optional[mgl32.Vec2]
	// EntityOffset controls the offset from the entity that the camera should be rendered at.
	EntityOffset Optional[mgl32.Vec3]
	// Radius is only used in a follow_orbit camera and controls how far away from the player the camera
	// should be rendered.
	Radius Optional[float32]
	// YawLimitMin is the minimum yaw limit of the camera.
	YawLimitMin Optional[float32]
	// YawLimitMax is the maximum yaw limit of the camera.
	YawLimitMax Optional[float32]
	// Listener defines where the audio should be played from when using this preset. This is one of the
	// constants above.
	Listener Optional[CameraPresetAudioListener]
	// PlayerEffects is currently unknown.
	PlayerEffects Optional[bool]
	// AimAssist defines the aim assist to use when using this preset.
	AimAssist Optional[CameraAimAssistCommandPresetDefinition]
	// ControlScheme is the control scheme that the client should use in this camera. It is one of the
	// following: - ControlSchemeLockedPlayerRelativeStrafe is the default behaviour, this cannot be set
	// when the client is in a custom camera. - ControlSchemeCameraRelative makes movement relative to
	// the camera's transform, with the client's rotation being relative to the client's movement. -
	// ControlSchemeCameraRelativeStrafe makes movement relative to the camera's transform, with the
	// client's rotation being locked. - ControlSchemePlayerRelative makes movement relative to the
	// player's transform, meaning holding left/right will make the player turn in a circle. -
	// ControlSchemePlayerRelativeStrafe makes movement the same as the default behaviour, but can be
	// used in a custom camera.
	ControlScheme Optional[ControlScheme]
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

type CameraPresetAudioListener uint8

const (
	CameraPresetAudioListenerCamera CameraPresetAudioListener = 0
	CameraPresetAudioListenerPlayer CameraPresetAudioListener = 1
)

type CameraPresetList struct {
	Presets []CameraPreset
}

// Marshal reads or writes CameraPresetList using its canonical wire layout.
func (x *CameraPresetList) Marshal(io IO) {
	Slice(io, &x.Presets)
}

// CameraProgressOption represents a progress keyframe option for camera spline instructions.
type CameraProgressOption struct {
	KeyFrameValue      float32
	KeyFrameTime       float32
	KeyFrameEasingFunc string
}

// Marshal reads or writes CameraProgressOption using its canonical wire layout.
func (x *CameraProgressOption) Marshal(io IO) {
	io.Float32(&x.KeyFrameValue)
	io.Float32(&x.KeyFrameTime)
	io.String(&x.KeyFrameEasingFunc)
}

type CameraRotation struct {
	X float32
	Y float32
}

// Marshal reads or writes CameraRotation using its canonical wire layout.
func (x *CameraRotation) Marshal(io IO) {
	io.Float32(&x.X)
	io.Float32(&x.Y)
}

// CameraRotationOption represents a rotation option for camera spline instructions.
type CameraRotationOption struct {
	KeyFrameValue      mgl32.Vec3
	KeyFrameTime       float32
	KeyFrameEasingFunc string
}

// Marshal reads or writes CameraRotationOption using its canonical wire layout.
func (x *CameraRotationOption) Marshal(io IO) {
	io.Vec3(&x.KeyFrameValue)
	io.Float32(&x.KeyFrameTime)
	io.String(&x.KeyFrameEasingFunc)
}

type CameraShakeAction uint8

const (
	CameraShakeActionAdd  CameraShakeAction = 0
	CameraShakeActionStop CameraShakeAction = 1
)

type CameraShakeType uint8

const (
	CameraShakeTypePositional CameraShakeType = 0
	CameraShakeTypeRotational CameraShakeType = 1
)

type CameraSplineControlPoint struct {
	Position mgl32.Vec3
}

// Marshal reads or writes CameraSplineControlPoint using its canonical wire layout.
func (x *CameraSplineControlPoint) Marshal(io IO) {
	io.Vec3(&x.Position)
}

// CameraSplineDefinition represents a named camera spline definition.
type CameraSplineDefinition struct {
	// Name is the name of the spline definition.
	Name string
	// TotalTime is the total time for the spline animation.
	TotalTime float32
	// SplineType is the optional spline interpolation type.
	SplineType string
	// ControlPoints is a list of points that define the spline curve.
	ControlPoints []CameraSplineControlPoint
	// ProgressKeyFrames is a list of progress key frames for the spline.
	ProgressKeyFrames []CameraSplineProgressKeyFrame
	// RotationKeyFrames is a list of rotation key frames for the spline.
	RotationKeyFrames []CameraSplineRotationKeyFrame
}

// Marshal reads or writes CameraSplineDefinition using its canonical wire layout.
func (x *CameraSplineDefinition) Marshal(io IO) {
	io.String(&x.Name)
	io.Float32(&x.TotalTime)
	io.String(&x.SplineType)
	Slice(io, &x.ControlPoints)
	Slice(io, &x.ProgressKeyFrames)
	Slice(io, &x.RotationKeyFrames)
}

// CameraSplineInstruction represents a camera instruction that creates a spline path for the camera
// to follow.
type CameraSplineInstruction struct {
	// TotalTime is the total time for the spline animation.
	TotalTime float32
	Type      uint8
	// Curve is a list of points that define the spline curve.
	Curve []mgl32.Vec3
	// ProgressKeyFrames is a list of progress key frames for the spline.
	ProgressKeyFrames []CameraProgressOption
	RotationOption    []CameraRotationOption
	// SplineIdentifier is an optional identifier for referencing the spline by name.
	SplineIdentifier string
	// LoadFromJSON optionally determines whether the spline should be loaded from a JSON definition.
	LoadFromJSON bool
}

// Marshal reads or writes CameraSplineInstruction using its canonical wire layout.
func (x *CameraSplineInstruction) Marshal(io IO) {
	io.Float32(&x.TotalTime)
	io.Uint8(&x.Type)
	FuncSlice(io, &x.Curve, io.Varuint32, io.Vec3)
	Slice(io, &x.ProgressKeyFrames)
	Slice(io, &x.RotationOption)
	io.String(&x.SplineIdentifier)
	io.Bool(&x.LoadFromJSON)
}

type CameraSplineProgressKeyFrame struct {
	Progress float32
	Time     float32
	Easing   Optional[string]
}

// Marshal reads or writes CameraSplineProgressKeyFrame using its canonical wire layout.
func (x *CameraSplineProgressKeyFrame) Marshal(io IO) {
	io.Float32(&x.Progress)
	io.Float32(&x.Time)
	OptionalFunc(io, &x.Easing, io.String)
}

type CameraSplineRotationKeyFrame struct {
	Rotation mgl32.Vec3
	Time     float32
	Easing   Optional[string]
}

// Marshal reads or writes CameraSplineRotationKeyFrame using its canonical wire layout.
func (x *CameraSplineRotationKeyFrame) Marshal(io IO) {
	io.Vec3(&x.Rotation)
	io.Float32(&x.Time)
	OptionalFunc(io, &x.Easing, io.String)
}

type CameraViewOffset struct {
	X float32
	Y float32
}

// Marshal reads or writes CameraViewOffset using its canonical wire layout.
func (x *CameraViewOffset) Marshal(io IO) {
	io.Float32(&x.X)
	io.Float32(&x.Y)
}

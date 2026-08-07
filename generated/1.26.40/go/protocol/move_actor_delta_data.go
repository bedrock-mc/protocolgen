// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MoveActorDeltaData struct {
	ActorRuntimeID       uint64
	NewPositionX         Optional[float32]
	NewPositionY         Optional[float32]
	NewPositionZ         Optional[float32]
	RotationX            Optional[int8]
	RotationY            Optional[int8]
	RotationYHead        Optional[int8]
	IsOnGround           bool
	ForceMove            bool
	ForceMoveLocalEntity bool
	ForceCompletion      bool
}

// Marshal reads or writes MoveActorDeltaData using its canonical wire layout.
func (x *MoveActorDeltaData) Marshal(io IO) {
	io.ActorRuntimeID(&x.ActorRuntimeID)
	OptionalFunc(io, &x.NewPositionX, io.Float32)
	OptionalFunc(io, &x.NewPositionY, io.Float32)
	OptionalFunc(io, &x.NewPositionZ, io.Float32)
	OptionalFunc(io, &x.RotationX, io.Int8)
	OptionalFunc(io, &x.RotationY, io.Int8)
	OptionalFunc(io, &x.RotationYHead, io.Int8)
	io.Bool(&x.IsOnGround)
	io.Bool(&x.ForceMove)
	io.Bool(&x.ForceMoveLocalEntity)
	io.Bool(&x.ForceCompletion)
}

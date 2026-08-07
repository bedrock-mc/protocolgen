// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type GameTestRequest struct {
	MaxTestsPerBatch int32
	RepeatCount      int32
	Rotation         Rotation
	StopOnFailure    bool
	TestPos          BlockPos
	TestsPerRow      int32
	TestName         string
}

// Marshal reads or writes GameTestRequest using its canonical wire layout.
func (x *GameTestRequest) Marshal(io IO) {
	io.Varint32(&x.MaxTestsPerBatch)
	io.Varint32(&x.RepeatCount)
	enumValue1 := uint8(x.Rotation)
	io.Uint8(&enumValue1)
	x.Rotation = Rotation(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.Bool(&x.StopOnFailure)
	x.TestPos.Marshal(io)
	io.Varint32(&x.TestsPerRow)
	io.String(&x.TestName)
}

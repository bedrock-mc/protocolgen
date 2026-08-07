// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// GameTestRequest ...
type GameTestRequest struct {
	// MaxTestsPerBatch ...
	MaxTestsPerBatch int32
	RepeatCount      int32
	// Rotation represents the rotation of the test. It is one of the constants above.
	Rotation      protocol.Rotation
	StopOnFailure bool
	TestPos       protocol.BlockPos
	// TestsPerRow ...
	TestsPerRow int32
	TestName    string
}

// Marshal reads or writes GameTestRequest using its canonical wire layout.
func (x *GameTestRequest) Marshal(io protocol.IO) {
	io.Varint32(&x.MaxTestsPerBatch)
	io.Varint32(&x.RepeatCount)
	protocol.IntegerFunc(&x.Rotation, io.Uint8)
	io.Bool(&x.StopOnFailure)
	x.TestPos.Marshal(io)
	io.Varint32(&x.TestsPerRow)
	io.String(&x.TestName)
}

// ID returns the protocol ID for GameTestRequest.
func (*GameTestRequest) ID() uint32 { return IDGameTestRequest }

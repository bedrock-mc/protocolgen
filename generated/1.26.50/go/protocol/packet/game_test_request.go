// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

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

// ID returns the protocol ID for GameTestRequest.
func (*GameTestRequest) ID() uint32 { return IDGameTestRequest }

// Marshal reads or writes GameTestRequest using its canonical wire layout.
func (pk *GameTestRequest) Marshal(io protocol.IO) {
	io.Varint32(&pk.MaxTestsPerBatch)
	io.Varint32(&pk.RepeatCount)
	pk.Rotation.Marshal(io)
	io.Bool(&pk.StopOnFailure)
	pk.TestPos.Marshal(io)
	io.Varint32(&pk.TestsPerRow)
	io.String(&pk.TestName)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// GameTestResults is a packet sent in response to the GameTestRequest packet, with a boolean
// indicating whether the test was successful or not, and an error string if the test failed.
type GameTestResults struct {
	// Succeeded indicates whether the test succeeded or not.
	Succeeded bool
	// Error is the error that occurred. If Succeeded is true, this field is empty.
	Error    string
	TestName string
}

// Marshal reads or writes GameTestResults using its canonical wire layout.
func (x *GameTestResults) Marshal(io protocol.IO) {
	io.Bool(&x.Succeeded)
	io.String(&x.Error)
	io.String(&x.TestName)
}

// ID returns the protocol ID for GameTestResults.
func (*GameTestResults) ID() uint32 { return IDGameTestResults }

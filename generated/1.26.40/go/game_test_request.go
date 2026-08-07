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

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlaySound struct {
	Name              string
	Position          BlockPos
	Volume            float32
	Pitch             float32
	LoopCount         int32
	ServerSoundHandle *ServerSoundHandle
}

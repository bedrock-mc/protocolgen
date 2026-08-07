// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type RaidUpdate struct {
	CurrentWave int32
	TotalWaves  int32
	Success     bool
}

func (*RaidUpdate) isEventData() {}

// Marshal reads or writes RaidUpdate using its canonical wire layout.
func (x *RaidUpdate) Marshal(io IO) {
	io.Varint32(&x.CurrentWave)
	io.Varint32(&x.TotalWaves)
	io.Bool(&x.Success)
}

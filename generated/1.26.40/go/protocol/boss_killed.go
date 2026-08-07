// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BossKilled struct {
	BossActorID int64
	PartySize   int32
	BossType    int32
}

func (*BossKilled) isEventData() {}

// Marshal reads or writes BossKilled using its canonical wire layout.
func (x *BossKilled) Marshal(io IO) {
	io.Varint64(&x.BossActorID)
	io.Varint32(&x.PartySize)
	io.Varint32(&x.BossType)
}

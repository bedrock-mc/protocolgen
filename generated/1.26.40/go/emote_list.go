// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type EmoteList struct {
	RuntimeId     ActorRuntimeID
	EmotePieceIds []uuid.UUID
}

// Marshal reads or writes EmoteList using its canonical wire layout.
func (x *EmoteList) Marshal(io IO) {
	x.RuntimeId.Marshal(io)
	if !io.Reading() && uint64(len(x.EmotePieceIds)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.EmotePieceIds), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.EmotePieceIds))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.EmotePieceIds = make([]uuid.UUID, int(count1))
	}
	for index2 := range x.EmotePieceIds {
		io.UUID(&x.EmotePieceIds[index2])
	}
}

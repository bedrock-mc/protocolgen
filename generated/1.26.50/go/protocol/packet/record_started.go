// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type RecordStarted struct {
	BlockPosition     protocol.BlockPos
	ServerSoundHandle protocol.ServerSoundHandle
}

// ID ...
func (*RecordStarted) ID() uint32 {
	return IDRecordStarted
}

func (pk *RecordStarted) Marshal(io protocol.IO) {
	pk.BlockPosition.Marshal(io)
	pk.ServerSoundHandle.Marshal(io)
}

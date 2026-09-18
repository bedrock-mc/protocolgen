// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type RecordStarted struct {
	BlockPosition     protocol.BlockPos
	ServerSoundHandle protocol.ServerSoundHandle
}

// ID returns the protocol ID for RecordStarted.
func (*RecordStarted) ID() uint32 { return IDRecordStarted }

// Marshal reads or writes RecordStarted using its canonical wire layout.
func (pk *RecordStarted) Marshal(io protocol.IO) {
	pk.BlockPosition.Marshal(io)
	pk.ServerSoundHandle.Marshal(io)
}

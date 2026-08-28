// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

type RecordStarted struct {
	BlockPosition     protocol.BlockPos
	ServerSoundHandle protocol.ServerSoundHandle
}

// Marshal reads or writes RecordStarted using its canonical wire layout.
func (x *RecordStarted) Marshal(io protocol.IO) {
	x.BlockPosition.Marshal(io)
	x.ServerSoundHandle.Marshal(io)
}

// ID returns the protocol ID for RecordStarted.
func (*RecordStarted) ID() uint32 { return IDRecordStarted }

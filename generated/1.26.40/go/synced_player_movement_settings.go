// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SyncedPlayerMovementSettings struct {
	RewindHistorySize                int32
	ServerAuthoritativeBlockBreaking bool
}

// Marshal reads or writes SyncedPlayerMovementSettings using its canonical wire layout.
func (x *SyncedPlayerMovementSettings) Marshal(io IO) {
	io.Varint32(&x.RewindHistorySize)
	io.Bool(&x.ServerAuthoritativeBlockBreaking)
}

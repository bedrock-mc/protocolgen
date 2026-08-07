// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type CommandOriginData struct {
	Type      string
	UUID      uuid.UUID
	RequestID string
	PlayerID  int64
}

// Marshal reads or writes CommandOriginData using its canonical wire layout.
func (x *CommandOriginData) Marshal(io IO) {
	io.String(&x.Type)
	io.UUID(&x.UUID)
	io.String(&x.RequestID)
	io.Int64(&x.PlayerID)
}

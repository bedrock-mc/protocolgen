// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type PackIDVersionData struct {
	PackUUID    uuid.UUID
	PackVersion SemVersionData
}

// Marshal reads or writes PackIDVersionData using its canonical wire layout.
func (x *PackIDVersionData) Marshal(io IO) {
	io.UUID(&x.PackUUID)
	x.PackVersion.Marshal(io)
}

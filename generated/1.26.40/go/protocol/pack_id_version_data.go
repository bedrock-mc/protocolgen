// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type PackIdVersionData struct {
	PackUUID    uuid.UUID
	PackVersion SemVersionData
}

// Marshal reads or writes PackIdVersionData using its canonical wire layout.
func (x *PackIdVersionData) Marshal(io IO) {
	io.UUID(&x.PackUUID)
	x.PackVersion.Marshal(io)
}

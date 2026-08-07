// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type PackIDVersion struct {
	PackUUID    uuid.UUID
	PackVersion SemVersion
}

// Marshal reads or writes PackIDVersion using its canonical wire layout.
func (x *PackIDVersion) Marshal(io IO) {
	io.UUID(&x.PackUUID)
	x.PackVersion.Marshal(io)
}

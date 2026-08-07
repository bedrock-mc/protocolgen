// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type PackIdVersion struct {
	PackUUID    uuid.UUID
	PackVersion SemVersion
}

// Marshal reads or writes PackIdVersion using its canonical wire layout.
func (x *PackIdVersion) Marshal(io IO) {
	io.UUID(&x.PackUUID)
	x.PackVersion.Marshal(io)
}

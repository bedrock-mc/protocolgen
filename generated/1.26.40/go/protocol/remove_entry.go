// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type RemoveEntry struct {
	UUID uuid.UUID
}

func (*RemoveEntry) isPlayerListData() {}

// Marshal reads or writes RemoveEntry using its canonical wire layout.
func (x *RemoveEntry) Marshal(io IO) {
	io.UUID(&x.UUID)
}

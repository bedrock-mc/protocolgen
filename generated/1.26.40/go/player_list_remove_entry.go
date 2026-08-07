// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type PlayerListRemoveEntry struct {
	UUID uuid.UUID
}

func (PlayerListRemoveEntry) isPlayerListEntriesItem() {}

// Marshal reads or writes PlayerListRemoveEntry using its canonical wire layout.
func (x *PlayerListRemoveEntry) Marshal(io IO) {
	io.UUID(&x.UUID)
}

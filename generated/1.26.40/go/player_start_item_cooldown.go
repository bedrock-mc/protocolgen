// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerStartItemCooldown struct {
	ItemCategory  string
	DurationTicks int32
}

// Marshal reads or writes PlayerStartItemCooldown using its canonical wire layout.
func (x *PlayerStartItemCooldown) Marshal(io IO) {
	io.String(&x.ItemCategory)
	io.Varint32(&x.DurationTicks)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemUsed struct {
	ItemID    int16
	ItemAux   int32
	UseMethod int32
	Count     int32
}

func (*ItemUsed) isEventData() {}

// Marshal reads or writes ItemUsed using its canonical wire layout.
func (x *ItemUsed) Marshal(io IO) {
	io.Int16(&x.ItemID)
	io.Int32(&x.ItemAux)
	io.Int32(&x.UseMethod)
	io.Int32(&x.Count)
}

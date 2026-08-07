// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PotionMixDataEntry struct {
	FromPotionId   int32
	FromItemAux    int32
	ReagentItemId  int32
	ReagentItemAux int32
	ToPotionId     int32
	ToItemAux      int32
}

// Marshal reads or writes PotionMixDataEntry using its canonical wire layout.
func (x *PotionMixDataEntry) Marshal(io IO) {
	io.Varint32(&x.FromPotionId)
	io.Varint32(&x.FromItemAux)
	io.Varint32(&x.ReagentItemId)
	io.Varint32(&x.ReagentItemAux)
	io.Varint32(&x.ToPotionId)
	io.Varint32(&x.ToItemAux)
}

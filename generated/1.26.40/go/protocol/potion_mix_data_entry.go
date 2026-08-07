// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PotionMixDataEntry struct {
	FromPotionID   int32
	FromItemAux    int32
	ReagentItemID  int32
	ReagentItemAux int32
	ToPotionID     int32
	ToItemAux      int32
}

// Marshal reads or writes PotionMixDataEntry using its canonical wire layout.
func (x *PotionMixDataEntry) Marshal(io IO) {
	io.Varint32(&x.FromPotionID)
	io.Varint32(&x.FromItemAux)
	io.Varint32(&x.ReagentItemID)
	io.Varint32(&x.ReagentItemAux)
	io.Varint32(&x.ToPotionID)
	io.Varint32(&x.ToItemAux)
}

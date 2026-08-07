// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TrimPattern struct {
	ItemName  string
	PatternID string
}

// Marshal reads or writes TrimPattern using its canonical wire layout.
func (x *TrimPattern) Marshal(io IO) {
	io.String(&x.ItemName)
	io.String(&x.PatternID)
}

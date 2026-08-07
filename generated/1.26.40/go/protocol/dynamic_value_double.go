// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DynamicValueDouble struct {
	Value float64
}

func (*DynamicValueDouble) isDynamicValue() {}

// Marshal reads or writes DynamicValueDouble using its canonical wire layout.
func (x *DynamicValueDouble) Marshal(io IO) {
	io.Float64(&x.Value)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DynamicValueInt64 struct {
	Value int64
}

func (*DynamicValueInt64) isDynamicValue() {}

// Marshal reads or writes DynamicValueInt64 using its canonical wire layout.
func (x *DynamicValueInt64) Marshal(io IO) {
	io.Int64(&x.Value)
}

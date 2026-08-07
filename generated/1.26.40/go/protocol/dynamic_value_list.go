// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DynamicValueList struct {
	Value []DynamicValue
}

func (*DynamicValueList) isDynamicValue() {}

// Marshal reads or writes DynamicValueList using its canonical wire layout.
func (x *DynamicValueList) Marshal(io IO) {
	FuncSlice(io, &x.Value, io.Varuint32, func(value *DynamicValue) {
		MarshalDynamicValue(io, value)
	})
}

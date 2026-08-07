// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DynamicValueMap struct {
	Value []OrderedEntry[string, DynamicValue]
}

func (*DynamicValueMap) isDynamicValue() {}

// Marshal reads or writes DynamicValueMap using its canonical wire layout.
func (x *DynamicValueMap) Marshal(io IO) {
	OrderedMap(io, &x.Value, io.Varuint32, io.String, func(value *DynamicValue) {
		MarshalDynamicValue(io, value)
	})
}

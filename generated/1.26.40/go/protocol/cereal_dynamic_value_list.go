// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CerealDynamicValueList struct {
	Value []CerealDynamicValue
}

func (CerealDynamicValueList) isCerealDynamicValue() {}

// Marshal reads or writes CerealDynamicValueList using its canonical wire layout.
func (x *CerealDynamicValueList) Marshal(io IO) {
	FuncSlice(io, &x.Value, io.Varuint32, func(value *CerealDynamicValue) {
		MarshalCerealDynamicValue(io, value)
	})
}

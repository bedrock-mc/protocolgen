// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CerealDynamicValueInt64 struct {
	Value int64
}

func (CerealDynamicValueInt64) isCerealDynamicValue() {}

// Marshal reads or writes CerealDynamicValueInt64 using its canonical wire layout.
func (x *CerealDynamicValueInt64) Marshal(io IO) {
	io.Int64(&x.Value)
}

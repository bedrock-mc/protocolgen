// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CerealDynamicValueString struct {
	Value string
}

func (CerealDynamicValueString) isCerealDynamicValue() {}

// Marshal reads or writes CerealDynamicValueString using its canonical wire layout.
func (x *CerealDynamicValueString) Marshal(io IO) {
	io.String(&x.Value)
}

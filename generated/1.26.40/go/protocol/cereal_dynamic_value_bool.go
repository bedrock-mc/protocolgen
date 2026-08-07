// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CerealDynamicValueBool struct {
	Value bool
}

func (CerealDynamicValueBool) isCerealDynamicValue() {}

// Marshal reads or writes CerealDynamicValueBool using its canonical wire layout.
func (x *CerealDynamicValueBool) Marshal(io IO) {
	io.Bool(&x.Value)
}

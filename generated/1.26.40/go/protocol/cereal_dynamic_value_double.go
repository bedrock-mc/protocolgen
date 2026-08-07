// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CerealDynamicValueDouble struct {
	Value float64
}

func (CerealDynamicValueDouble) isCerealDynamicValue() {}

// Marshal reads or writes CerealDynamicValueDouble using its canonical wire layout.
func (x *CerealDynamicValueDouble) Marshal(io IO) {
	io.Float64(&x.Value)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DynamicValueBool struct {
	Value bool
}

func (*DynamicValueBool) isDynamicValue() {}

// Marshal reads or writes DynamicValueBool using its canonical wire layout.
func (x *DynamicValueBool) Marshal(io IO) {
	io.Bool(&x.Value)
}

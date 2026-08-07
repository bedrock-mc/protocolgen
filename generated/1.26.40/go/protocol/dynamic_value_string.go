// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DynamicValueString struct {
	Value string
}

func (*DynamicValueString) isDynamicValue() {}

// Marshal reads or writes DynamicValueString using its canonical wire layout.
func (x *DynamicValueString) Marshal(io IO) {
	io.String(&x.Value)
}

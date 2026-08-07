// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type InvalidItemDescriptor struct {
	DescriptorType ItemDescriptorType
}

func (*InvalidItemDescriptor) isItemDescriptor() {}

// Marshal reads or writes InvalidItemDescriptor using its canonical wire layout.
func (x *InvalidItemDescriptor) Marshal(io IO) {
	IntegerFunc(&x.DescriptorType, io.Uint8)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemTagItemDescriptor struct {
	DescriptorType ItemDescriptorType
	ItemTag        string
}

func (*ItemTagItemDescriptor) isItemDescriptor() {}

// Marshal reads or writes ItemTagItemDescriptor using its canonical wire layout.
func (x *ItemTagItemDescriptor) Marshal(io IO) {
	IntegerFunc(&x.DescriptorType, io.Uint8)
	io.String(&x.ItemTag)
}

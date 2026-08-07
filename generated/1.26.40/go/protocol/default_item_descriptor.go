// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DefaultItemDescriptor struct {
	DescriptorType ItemDescriptorType
	FullName       string
	AuxValue       int32
}

func (*DefaultItemDescriptor) isItemDescriptor() {}

// Marshal reads or writes DefaultItemDescriptor using its canonical wire layout.
func (x *DefaultItemDescriptor) Marshal(io IO) {
	IntegerFunc(&x.DescriptorType, io.Uint8)
	io.String(&x.FullName)
	io.Varint32(&x.AuxValue)
}

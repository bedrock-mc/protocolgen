// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MoLangItemDescriptor struct {
	DescriptorType ItemDescriptorType
	TagExpression  string
	MoLangVersion  MoLangVersion
}

func (*MoLangItemDescriptor) isItemDescriptor() {}

// Marshal reads or writes MoLangItemDescriptor using its canonical wire layout.
func (x *MoLangItemDescriptor) Marshal(io IO) {
	IntegerFunc(&x.DescriptorType, io.Uint8)
	io.String(&x.TagExpression)
	IntegerFunc(&x.MoLangVersion, io.Int16)
}

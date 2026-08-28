// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

// DefaultItemDescriptor represents an item descriptor for regular items. This is used for the
// significant majority of items.
type DefaultItemDescriptor struct {
	DescriptorType ItemDescriptorType
	FullName       string
	AuxValue       int32
}

func (*DefaultItemDescriptor) isItemDescriptor() {}

// Marshal reads or writes DefaultItemDescriptor using its canonical wire layout.
func (x *DefaultItemDescriptor) Marshal(io IO) {
	IntegerFunc(&x.DescriptorType, io.Uint8)
	io.StringLimits(&x.FullName, 1, 18446744073709551615)
	io.Varint32(&x.AuxValue)
	Minimum(io, &x.AuxValue, 0)
	Maximum(io, &x.AuxValue, 32767)
}

// InvalidItemDescriptor represents an invalid item descriptor. This is usually sent by the vanilla
// server for empty slots or ingredients.
type InvalidItemDescriptor struct {
	DescriptorType ItemDescriptorType
}

func (*InvalidItemDescriptor) isItemDescriptor() {}

// Marshal reads or writes InvalidItemDescriptor using its canonical wire layout.
func (x *InvalidItemDescriptor) Marshal(io IO) {
	IntegerFunc(&x.DescriptorType, io.Uint8)
}

// ItemDescriptor represents a type of item descriptor. This is one of the concrete types below. It
// is an alias of Marshaler.
type ItemDescriptorType uint8

const (
	ItemDescriptorTypeEmpty    ItemDescriptorType = 0
	ItemDescriptorTypeItemName ItemDescriptorType = 1
	ItemDescriptorTypeMoLang   ItemDescriptorType = 2
	ItemDescriptorTypeItemTag  ItemDescriptorType = 3
)

// ItemTagItemDescriptor represents an item descriptor that uses item tagging. This should be used
// to reduce duplicative entries for items that can be grouped under a single tag.
type ItemTagItemDescriptor struct {
	DescriptorType ItemDescriptorType
	ItemTag        string
}

func (*ItemTagItemDescriptor) isItemDescriptor() {}

// Marshal reads or writes ItemTagItemDescriptor using its canonical wire layout.
func (x *ItemTagItemDescriptor) Marshal(io IO) {
	IntegerFunc(&x.DescriptorType, io.Uint8)
	io.StringLimits(&x.ItemTag, 1, 18446744073709551615)
}

// MoLangItemDescriptor represents an item descriptor for items that use MoLang (e.g. behaviour
// packs).
type MoLangItemDescriptor struct {
	DescriptorType ItemDescriptorType
	TagExpression  string
	MoLangVersion  MoLangVersion
}

func (*MoLangItemDescriptor) isItemDescriptor() {}

// Marshal reads or writes MoLangItemDescriptor using its canonical wire layout.
func (x *MoLangItemDescriptor) Marshal(io IO) {
	IntegerFunc(&x.DescriptorType, io.Uint8)
	io.StringLimits(&x.TagExpression, 1, 18446744073709551615)
	IntegerFunc(&x.MoLangVersion, io.Int16)
}

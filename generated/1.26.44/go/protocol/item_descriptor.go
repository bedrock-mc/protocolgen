// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

// DefaultItemDescriptor represents an item descriptor for regular items. This is used for the significant
// majority of items.
type DefaultItemDescriptor struct {
	DescriptorType ItemDescriptorType
	// Name is the identifier of the item, such as minecraft:stone.
	FullName string
	// MetadataValue is the metadata value of the item. For some items, this is the damage value, whereas for
	// other items it is simply an identifier of a variant of the item.
	AuxValue int32
}

func (*DefaultItemDescriptor) tagItemDescriptor() uint32 { return 1 }

// Marshal reads or writes DefaultItemDescriptor using its canonical wire layout.
func (x *DefaultItemDescriptor) Marshal(io IO) {
	x.DescriptorType.Marshal(io)
	io.StringLimits(&x.FullName, 1, 18446744073709551615)
	io.Varint32(&x.AuxValue)
	Minimum(io, &x.AuxValue, 0)
	Maximum(io, &x.AuxValue, 32767)
}

// InvalidItemDescriptor represents an invalid item descriptor. This is usually sent by the vanilla server for
// empty slots or ingredients.
type InvalidItemDescriptor struct {
	DescriptorType ItemDescriptorType
}

func (*InvalidItemDescriptor) tagItemDescriptor() uint32 { return 0 }

// Marshal reads or writes InvalidItemDescriptor using its canonical wire layout.
func (x *InvalidItemDescriptor) Marshal(io IO) {
	x.DescriptorType.Marshal(io)
}

// ItemDescriptor represents a type of item descriptor. This is one of the concrete types below. It is an
// alias of Marshaler.
type ItemDescriptor interface {
	Marshaler
	tagItemDescriptor() uint32
}

// MarshalItemDescriptor reads or writes the ItemDescriptor union using its canonical wire layout.
func MarshalItemDescriptor(io IO, x *ItemDescriptor) {
	Union(io, x, io.Varuint32, ItemDescriptor.tagItemDescriptor, func(tag uint32) ItemDescriptor {
		switch tag {
		case 0:
			return new(InvalidItemDescriptor)
		case 1:
			return new(DefaultItemDescriptor)
		case 2:
			return new(MoLangItemDescriptor)
		case 3:
			return new(ItemTagItemDescriptor)
		}
		return nil
	})
}

// ItemDescriptor represents a type of item descriptor. This is one of the concrete types below. It is an
// alias of Marshaler.
type ItemDescriptorType uint8

const (
	ItemDescriptorTypeEmpty    ItemDescriptorType = 0
	ItemDescriptorTypeItemName ItemDescriptorType = 1
	ItemDescriptorTypeMoLang   ItemDescriptorType = 2
	ItemDescriptorTypeItemTag  ItemDescriptorType = 3
)

// Marshal reads or writes ItemDescriptorType through its uint8 wire encoding.
func (x *ItemDescriptorType) Marshal(io IO) { io.Uint8((*uint8)(x)) }

// ItemTagItemDescriptor represents an item descriptor that uses item tagging. This should be used to reduce
// duplicative entries for items that can be grouped under a single tag.
type ItemTagItemDescriptor struct {
	DescriptorType ItemDescriptorType
	// Tag represents the tag that the item is part of.
	ItemTag string
}

func (*ItemTagItemDescriptor) tagItemDescriptor() uint32 { return 3 }

// Marshal reads or writes ItemTagItemDescriptor using its canonical wire layout.
func (x *ItemTagItemDescriptor) Marshal(io IO) {
	x.DescriptorType.Marshal(io)
	io.StringLimits(&x.ItemTag, 1, 18446744073709551615)
}

// MoLangItemDescriptor represents an item descriptor for items that use MoLang (e.g. behaviour packs).
type MoLangItemDescriptor struct {
	DescriptorType ItemDescriptorType
	// Expression represents the MoLang expression used to identify the item/it's associated tag.
	TagExpression string
	// Version represents the version of MoLang to use.
	MoLangVersion MoLangVersion
}

func (*MoLangItemDescriptor) tagItemDescriptor() uint32 { return 2 }

// Marshal reads or writes MoLangItemDescriptor using its canonical wire layout.
func (x *MoLangItemDescriptor) Marshal(io IO) {
	x.DescriptorType.Marshal(io)
	io.StringLimits(&x.TagExpression, 1, 18446744073709551615)
	x.MoLangVersion.Marshal(io)
}

// ItemDescriptor represents a type of item descriptor. This is one of the concrete types below. It is an
// alias of Marshaler.
type StackRequestAction interface {
	Marshaler
	tagStackRequestAction() uint32
}

// MarshalStackRequestAction reads or writes the StackRequestAction union using its canonical wire layout.
func MarshalStackRequestAction(io IO, x *StackRequestAction) {
	Union(io, x, io.Varuint32, StackRequestAction.tagStackRequestAction, func(tag uint32) StackRequestAction {
		switch tag {
		case 0:
			return new(TakeStackRequestAction)
		case 1:
			return new(PlaceStackRequestAction)
		case 2:
			return new(SwapStackRequestAction)
		case 3:
			return new(DropStackRequestAction)
		case 4:
			return new(DestroyStackRequestAction)
		case 5:
			return new(ConsumeStackRequestAction)
		case 6:
			return new(CreateStackRequestAction)
		case 7:
			return new(LabTableCombineStackRequestAction)
		case 8:
			return new(BeaconPaymentStackRequestAction)
		case 9:
			return new(MineBlockStackRequestAction)
		case 10:
			return new(CraftRecipeStackRequestAction)
		case 11:
			return new(AutoCraftRecipeStackRequestAction)
		case 12:
			return new(CraftCreativeStackRequestAction)
		case 13:
			return new(CraftRecipeOptionalStackRequestAction)
		case 14:
			return new(CraftRepairAndDisenchantStackRequestAction)
		case 15:
			return new(CraftLoomStackRequestAction)
		case 16:
			return new(CraftNonImplementedStackRequestAction)
		case 17:
			return new(CraftResultsDeprecatedStackRequestAction)
		}
		return nil
	})
}

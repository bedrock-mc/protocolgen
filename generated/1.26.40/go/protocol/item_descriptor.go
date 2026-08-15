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
type ItemDescriptor interface {
	isItemDescriptor()
}

// MarshalItemDescriptor reads or writes the ItemDescriptor union using its canonical wire layout.
func MarshalItemDescriptor(io IO, x *ItemDescriptor) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(InvalidItemDescriptor)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(DefaultItemDescriptor)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(MoLangItemDescriptor)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(ItemTagItemDescriptor)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *InvalidItemDescriptor:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *DefaultItemDescriptor:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *MoLangItemDescriptor:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemTagItemDescriptor:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
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

// ItemDescriptor represents a type of item descriptor. This is one of the concrete types below. It
// is an alias of Marshaler.
type StackRequestAction interface {
	isStackRequestAction()
}

// MarshalStackRequestAction reads or writes the StackRequestAction union using its canonical wire layout.
func MarshalStackRequestAction(io IO, x *StackRequestAction) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(TakeStackRequestAction)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(PlaceStackRequestAction)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(SwapStackRequestAction)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(DropStackRequestAction)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(DestroyStackRequestAction)
				value.Marshal(io)
				*x = value
			case 5:
				value := new(ConsumeStackRequestAction)
				value.Marshal(io)
				*x = value
			case 6:
				value := new(CreateStackRequestAction)
				value.Marshal(io)
				*x = value
			case 7:
				value := new(LabTableCombineStackRequestAction)
				value.Marshal(io)
				*x = value
			case 8:
				value := new(BeaconPaymentStackRequestAction)
				value.Marshal(io)
				*x = value
			case 9:
				value := new(MineBlockStackRequestAction)
				value.Marshal(io)
				*x = value
			case 10:
				value := new(CraftRecipeStackRequestAction)
				value.Marshal(io)
				*x = value
			case 11:
				value := new(AutoCraftRecipeStackRequestAction)
				value.Marshal(io)
				*x = value
			case 12:
				value := new(CraftCreativeStackRequestAction)
				value.Marshal(io)
				*x = value
			case 13:
				value := new(CraftRecipeOptionalStackRequestAction)
				value.Marshal(io)
				*x = value
			case 14:
				value := new(CraftRepairAndDisenchantStackRequestAction)
				value.Marshal(io)
				*x = value
			case 15:
				value := new(CraftLoomStackRequestAction)
				value.Marshal(io)
				*x = value
			case 16:
				value := new(CraftNonImplementedStackRequestAction)
				value.Marshal(io)
				*x = value
			case 17:
				value := new(CraftResultsDeprecatedStackRequestAction)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *TakeStackRequestAction:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *PlaceStackRequestAction:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *SwapStackRequestAction:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *DropStackRequestAction:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *DestroyStackRequestAction:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ConsumeStackRequestAction:
				tag := uint32(5)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CreateStackRequestAction:
				tag := uint32(6)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *LabTableCombineStackRequestAction:
				tag := uint32(7)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *BeaconPaymentStackRequestAction:
				tag := uint32(8)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *MineBlockStackRequestAction:
				tag := uint32(9)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CraftRecipeStackRequestAction:
				tag := uint32(10)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *AutoCraftRecipeStackRequestAction:
				tag := uint32(11)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CraftCreativeStackRequestAction:
				tag := uint32(12)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CraftRecipeOptionalStackRequestAction:
				tag := uint32(13)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CraftRepairAndDisenchantStackRequestAction:
				tag := uint32(14)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CraftLoomStackRequestAction:
				tag := uint32(15)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CraftNonImplementedStackRequestAction:
				tag := uint32(16)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CraftResultsDeprecatedStackRequestAction:
				tag := uint32(17)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

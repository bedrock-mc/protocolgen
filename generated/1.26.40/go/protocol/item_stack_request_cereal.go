// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestCereal interface {
	isItemStackRequestCereal()
}

// MarshalItemStackRequestCereal reads or writes the ItemStackRequestCereal union using its canonical wire layout.
func MarshalItemStackRequestCereal(io IO, x *ItemStackRequestCereal) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(ItemStackRequestCerealTakeActionData)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(ItemStackRequestCerealPlaceActionData)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(ItemStackRequestCerealSwapActionData)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(ItemStackRequestCerealDropActionData)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(ItemStackRequestCerealDestroyActionData)
				value.Marshal(io)
				*x = value
			case 5:
				value := new(ItemStackRequestCerealConsumeActionData)
				value.Marshal(io)
				*x = value
			case 6:
				value := new(ItemStackRequestCerealCreateActionData)
				value.Marshal(io)
				*x = value
			case 7:
				value := new(ItemStackRequestCerealLabTableCombineActionData)
				value.Marshal(io)
				*x = value
			case 8:
				value := new(ItemStackRequestCerealBeaconPaymentActionData)
				value.Marshal(io)
				*x = value
			case 9:
				value := new(ItemStackRequestCerealMineBlockActionData)
				value.Marshal(io)
				*x = value
			case 10:
				value := new(ItemStackRequestCerealCraftRecipeActionData)
				value.Marshal(io)
				*x = value
			case 11:
				value := new(ItemStackRequestCerealCraftRecipeAutoActionData)
				value.Marshal(io)
				*x = value
			case 12:
				value := new(ItemStackRequestCerealCraftCreativeActionData)
				value.Marshal(io)
				*x = value
			case 13:
				value := new(ItemStackRequestCerealCraftRecipeOptionalActionData)
				value.Marshal(io)
				*x = value
			case 14:
				value := new(ItemStackRequestCerealCraftRepairAndDisenchantActionData)
				value.Marshal(io)
				*x = value
			case 15:
				value := new(ItemStackRequestCerealCraftLoomActionData)
				value.Marshal(io)
				*x = value
			case 16:
				value := new(ItemStackRequestCerealCraftNonImplementedActionData)
				value.Marshal(io)
				*x = value
			case 17:
				value := new(ItemStackRequestCerealCraftResultsActionData)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *ItemStackRequestCerealTakeActionData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealPlaceActionData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealSwapActionData:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealDropActionData:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealDestroyActionData:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealConsumeActionData:
				tag := uint32(5)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealCreateActionData:
				tag := uint32(6)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealLabTableCombineActionData:
				tag := uint32(7)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealBeaconPaymentActionData:
				tag := uint32(8)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealMineBlockActionData:
				tag := uint32(9)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealCraftRecipeActionData:
				tag := uint32(10)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealCraftRecipeAutoActionData:
				tag := uint32(11)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealCraftCreativeActionData:
				tag := uint32(12)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealCraftRecipeOptionalActionData:
				tag := uint32(13)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealCraftRepairAndDisenchantActionData:
				tag := uint32(14)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealCraftLoomActionData:
				tag := uint32(15)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealCraftNonImplementedActionData:
				tag := uint32(16)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemStackRequestCerealCraftResultsActionData:
				tag := uint32(17)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealRecipeIngredientDataItemDescriptor interface {
	isItemStackRequestCerealRecipeIngredientDataItemDescriptor()
}

func marshalItemStackRequestCerealRecipeIngredientDataItemDescriptor(io IO, x *ItemStackRequestCerealRecipeIngredientDataItemDescriptor) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value ItemStackRequestCerealEmptyItemDescriptorData
				value.Marshal(io)
				*x = value
			case 1:
				var value ItemStackRequestCerealItemNameDescriptorData
				value.Marshal(io)
				*x = value
			case 2:
				var value ItemStackRequestCerealMoLangItemDescriptorData
				value.Marshal(io)
				*x = value
			case 3:
				var value ItemStackRequestCerealItemTagDescriptorData
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case ItemStackRequestCerealEmptyItemDescriptorData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case ItemStackRequestCerealItemNameDescriptorData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case ItemStackRequestCerealMoLangItemDescriptorData:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case ItemStackRequestCerealItemTagDescriptorData:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

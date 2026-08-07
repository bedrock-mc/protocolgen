// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UnlockedRecipes struct {
	PacketType          UnlockedRecipesPacketType
	UnlockedRecipesList []string
}

// Marshal reads or writes UnlockedRecipes using its canonical wire layout.
func (x *UnlockedRecipes) Marshal(io IO) {
	enumValue1 := uint32(x.PacketType)
	io.Uint32(&enumValue1)
	x.PacketType = UnlockedRecipesPacketType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3, 4:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	if !io.Reading() && uint64(len(x.UnlockedRecipesList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.UnlockedRecipesList), "collection length overflows uint32")
		return
	}
	count2 := uint32(len(x.UnlockedRecipesList))
	io.Varuint32(&count2)
	if io.Reading() {
		if uint64(count2) > uint64(^uint(0)>>1) {
			io.InvalidValue(count2, "collection length overflows int")
			return
		}
		x.UnlockedRecipesList = make([]string, int(count2))
	}
	for index3 := range x.UnlockedRecipesList {
		io.String(&x.UnlockedRecipesList[index3])
	}
}

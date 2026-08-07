// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UnlockedRecipes struct {
	PacketType          UnlockedRecipesPacketType
	UnlockedRecipesList []string
}

// Marshal reads or writes UnlockedRecipes using its canonical wire layout.
func (x *UnlockedRecipes) Marshal(io IO) {
	IntegerFunc(&x.PacketType, io.Uint32)
	FuncSlice(io, &x.UnlockedRecipesList, io.Varuint32, io.String)
}

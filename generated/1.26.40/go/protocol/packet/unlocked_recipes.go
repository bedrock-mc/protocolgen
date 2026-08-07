// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// UnlockedRecipes gives the client a list of recipes that have been unlocked, restricting the
// recipes that appear in the recipe book.
type UnlockedRecipes struct {
	PacketType          protocol.PacketType
	UnlockedRecipesList []string
}

// Marshal reads or writes UnlockedRecipes using its canonical wire layout.
func (x *UnlockedRecipes) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.PacketType, io.Uint32)
	protocol.FuncSlice(io, &x.UnlockedRecipesList, io.Varuint32, io.String)
}

// ID returns the protocol ID for UnlockedRecipes.
func (*UnlockedRecipes) ID() uint32 { return IDUnlockedRecipes }

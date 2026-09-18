// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// UnlockedRecipes gives the client a list of recipes that have been unlocked, restricting the recipes that
// appear in the recipe book.
type UnlockedRecipes struct {
	PacketType          protocol.PacketType
	UnlockedRecipesList []string
}

// ID returns the protocol ID for UnlockedRecipes.
func (*UnlockedRecipes) ID() uint32 { return IDUnlockedRecipes }

// Marshal reads or writes UnlockedRecipes using its canonical wire layout.
func (pk *UnlockedRecipes) Marshal(io protocol.IO) {
	pk.PacketType.Marshal(io)
	protocol.FuncSlice(io, &pk.UnlockedRecipesList, io.Varuint32, io.String)
}

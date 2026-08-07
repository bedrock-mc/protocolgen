// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemEnchantOption struct {
	Cost         uint8
	Enchants     ItemEnchants
	EnchantName  string
	EnchantNetId TypedServerNetIdStructRecipeNetIdTag
}

// Marshal reads or writes ItemEnchantOption using its canonical wire layout.
func (x *ItemEnchantOption) Marshal(io IO) {
	io.Uint8(&x.Cost)
	x.Enchants.Marshal(io)
	io.String(&x.EnchantName)
	x.EnchantNetId.Marshal(io)
}

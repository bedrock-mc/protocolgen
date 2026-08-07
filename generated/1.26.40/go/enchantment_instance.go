// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type EnchantmentInstance struct {
	EnchantType  EnchantType
	EnchantLevel uint8
}

// Marshal reads or writes EnchantmentInstance using its canonical wire layout.
func (x *EnchantmentInstance) Marshal(io IO) {
	IntegerFunc(&x.EnchantType, io.Uint8)
	io.Uint8(&x.EnchantLevel)
}

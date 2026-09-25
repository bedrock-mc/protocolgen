// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type EnchantType uint8

const (
	EnchantTypeProtection           EnchantType = 0
	EnchantTypeFireprotection       EnchantType = 1
	EnchantTypeFeatherfalling       EnchantType = 2
	EnchantTypeBlastprotection      EnchantType = 3
	EnchantTypeProjectileprotection EnchantType = 4
	EnchantTypeThorns               EnchantType = 5
	EnchantTypeRespiration          EnchantType = 6
	EnchantTypeDepthstrider         EnchantType = 7
	EnchantTypeAquaaffinity         EnchantType = 8
	EnchantTypeSharpness            EnchantType = 9
	EnchantTypeSmite                EnchantType = 10
	EnchantTypeBaneofarthropods     EnchantType = 11
	EnchantTypeKnockback            EnchantType = 12
	EnchantTypeFireaspect           EnchantType = 13
	EnchantTypeLooting              EnchantType = 14
	EnchantTypeEfficiency           EnchantType = 15
	EnchantTypeSilktouch            EnchantType = 16
	EnchantTypeUnbreaking           EnchantType = 17
	EnchantTypeFortune              EnchantType = 18
	EnchantTypePower                EnchantType = 19
	EnchantTypePunch                EnchantType = 20
	EnchantTypeFlame                EnchantType = 21
	EnchantTypeInfinity             EnchantType = 22
	EnchantTypeLuckofthesea         EnchantType = 23
	EnchantTypeLure                 EnchantType = 24
	EnchantTypeFrostwalker          EnchantType = 25
	EnchantTypeMending              EnchantType = 26
	EnchantTypeCurseofbinding       EnchantType = 27
	EnchantTypeCurseofvanishing     EnchantType = 28
	EnchantTypeImpaling             EnchantType = 29
	EnchantTypeRiptide              EnchantType = 30
	EnchantTypeLoyalty              EnchantType = 31
	EnchantTypeChanneling           EnchantType = 32
	EnchantTypeMultishot            EnchantType = 33
	EnchantTypePiercing             EnchantType = 34
	EnchantTypeQuickcharge          EnchantType = 35
	EnchantTypeSoulspeed            EnchantType = 36
	EnchantTypeSwiftsneak           EnchantType = 37
	EnchantTypeWindburst            EnchantType = 38
	EnchantTypeDensity              EnchantType = 39
	EnchantTypeBreach               EnchantType = 40
	EnchantTypeLunge                EnchantType = 41
	EnchantTypeNumenchantments      EnchantType = 42
	EnchantTypeInvalidenchantment   EnchantType = 43
)

// Marshal reads or writes EnchantType through its uint8 wire encoding.
func (x *EnchantType) Marshal(io IO) { io.Uint8((*uint8)(x)) }

// EnchantmentInstance represents a single enchantment instance with the type of the enchantment and its
// level.
type EnchantmentInstance struct {
	EnchantType  EnchantType
	EnchantLevel uint8
}

// Marshal reads or writes EnchantmentInstance using its canonical wire layout.
func (x *EnchantmentInstance) Marshal(io IO) {
	x.EnchantType.Marshal(io)
	io.Uint8(&x.EnchantLevel)
}

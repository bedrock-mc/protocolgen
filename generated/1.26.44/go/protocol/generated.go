// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BedrockDDUIDataStoreUpdateData interface {
	isBedrockDDUIDataStoreUpdateData()
}

// MarshalBedrockDDUIDataStoreUpdateData reads or writes the BedrockDDUIDataStoreUpdateData union using its canonical wire layout.
func MarshalBedrockDDUIDataStoreUpdateData(io IO, x *BedrockDDUIDataStoreUpdateData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(BedrockDDUIDataStoreUpdateDataDouble)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(BedrockDDUIDataStoreUpdateDataBool)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(BedrockDDUIDataStoreUpdateDataString)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *BedrockDDUIDataStoreUpdateDataDouble:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *BedrockDDUIDataStoreUpdateDataBool:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *BedrockDDUIDataStoreUpdateDataString:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

type BedrockDDUIDataStoreUpdateDataBool struct {
	Value bool
}

func (*BedrockDDUIDataStoreUpdateDataBool) isBedrockDDUIDataStoreUpdateData() {}

// Marshal reads or writes BedrockDDUIDataStoreUpdateDataBool using its canonical wire layout.
func (x *BedrockDDUIDataStoreUpdateDataBool) Marshal(io IO) {
	io.Bool(&x.Value)
}

type BedrockDDUIDataStoreUpdateDataDouble struct {
	Value float64
}

func (*BedrockDDUIDataStoreUpdateDataDouble) isBedrockDDUIDataStoreUpdateData() {}

// Marshal reads or writes BedrockDDUIDataStoreUpdateDataDouble using its canonical wire layout.
func (x *BedrockDDUIDataStoreUpdateDataDouble) Marshal(io IO) {
	io.Float64(&x.Value)
}

type BedrockDDUIDataStoreUpdateDataString struct {
	Value string
}

func (*BedrockDDUIDataStoreUpdateDataString) isBedrockDDUIDataStoreUpdateData() {}

// Marshal reads or writes BedrockDDUIDataStoreUpdateDataString using its canonical wire layout.
func (x *BedrockDDUIDataStoreUpdateDataString) Marshal(io IO) {
	io.StringLimits(&x.Value, 0, 5000)
}

// Bitset131 stores the 131-bit value used by the wire bitset encoding.
type Bitset131 [3]uint64

const Bitset131Length = 131

// Set marks bit index i. It panics when i is outside [0, 131).
func (b *Bitset131) Set(i int) {
	if i < 0 || i >= Bitset131Length {
		panic("index out of bounds")
	}
	b[i/64] |= uint64(1) << uint(i%64)
}

// Unset clears bit index i. It panics when i is outside [0, 131).
func (b *Bitset131) Unset(i int) {
	if i < 0 || i >= Bitset131Length {
		panic("index out of bounds")
	}
	b[i/64] &^= uint64(1) << uint(i%64)
}

// Load reports whether bit index i is set. It panics when i is outside [0, 131).
func (b Bitset131) Load(i int) bool {
	if i < 0 || i >= Bitset131Length {
		panic("index out of bounds")
	}
	return b[i/64]&(uint64(1)<<uint(i%64)) != 0
}

// Len returns the number of bits in the bitset.
func (b Bitset131) Len() int { return Bitset131Length }

type DataItemEntryValue interface {
	isDataItemEntryValue()
}

// MarshalDataItemEntryValue reads or writes the DataItemEntryValue union using its canonical wire layout.
func MarshalDataItemEntryValue(io IO, x *DataItemEntryValue) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				value := new(DataItemByte)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(DataItemShort)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(DataItemInt)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(DataItemFloat)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(DataItemString)
				value.Marshal(io)
				*x = value
			case 5:
				value := new(DataItemCompoundTag)
				value.Marshal(io)
				*x = value
			case 6:
				value := new(DataItemPos)
				value.Marshal(io)
				*x = value
			case 7:
				value := new(DataItemInt64)
				value.Marshal(io)
				*x = value
			case 8:
				value := new(DataItemVec3)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *DataItemByte:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemShort:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemInt:
				tag := uint8(2)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemFloat:
				tag := uint8(3)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemString:
				tag := uint8(4)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemCompoundTag:
				tag := uint8(5)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemPos:
				tag := uint8(6)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemInt64:
				tag := uint8(7)
				io.Uint8(&tag)
				value.Marshal(io)
			case *DataItemVec3:
				tag := uint8(8)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

type DisconnectMessages interface {
	isDisconnectMessages()
}

// MarshalDisconnectMessages reads or writes the DisconnectMessages union using its canonical wire layout.
func MarshalDisconnectMessages(io IO, x *DisconnectMessages) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(DisconnectMessagesData)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(DisconnectMessagesEmpty)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *DisconnectMessagesData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *DisconnectMessagesEmpty:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

type DisconnectMessagesEmpty struct {
}

func (*DisconnectMessagesEmpty) isDisconnectMessages() {}

// Marshal reads or writes DisconnectMessagesEmpty using its canonical wire layout.
func (x *DisconnectMessagesEmpty) Marshal(io IO) {
}

type DynamicValueBool struct {
	Value bool
}

func (*DynamicValueBool) isDynamicValue() {}

// Marshal reads or writes DynamicValueBool using its canonical wire layout.
func (x *DynamicValueBool) Marshal(io IO) {
	io.Bool(&x.Value)
}

type DynamicValueDouble struct {
	Value float64
}

func (*DynamicValueDouble) isDynamicValue() {}

// Marshal reads or writes DynamicValueDouble using its canonical wire layout.
func (x *DynamicValueDouble) Marshal(io IO) {
	io.Float64(&x.Value)
}

type DynamicValueInt64 struct {
	Value int64
}

func (*DynamicValueInt64) isDynamicValue() {}

// Marshal reads or writes DynamicValueInt64 using its canonical wire layout.
func (x *DynamicValueInt64) Marshal(io IO) {
	io.Int64(&x.Value)
}

type DynamicValueList struct {
	Value []DynamicValue
}

func (*DynamicValueList) isDynamicValue() {}

// Marshal reads or writes DynamicValueList using its canonical wire layout.
func (x *DynamicValueList) Marshal(io IO) {
	FuncSlice(io, &x.Value, io.Varuint32, func(value *DynamicValue) {
		MarshalDynamicValue(io, value)
	})
}

type DynamicValueMap struct {
	Value []OrderedEntry[string, DynamicValue]
}

func (*DynamicValueMap) isDynamicValue() {}

// Marshal reads or writes DynamicValueMap using its canonical wire layout.
func (x *DynamicValueMap) Marshal(io IO) {
	OrderedMap(io, &x.Value, io.Varuint32, io.String, func(value *DynamicValue) {
		MarshalDynamicValue(io, value)
	})
}

type DynamicValueNone struct {
}

func (*DynamicValueNone) isDynamicValue() {}

// Marshal reads or writes DynamicValueNone using its canonical wire layout.
func (x *DynamicValueNone) Marshal(io IO) {
}

type DynamicValueString struct {
	Value string
}

func (*DynamicValueString) isDynamicValue() {}

// Marshal reads or writes DynamicValueString using its canonical wire layout.
func (x *DynamicValueString) Marshal(io IO) {
	io.String(&x.Value)
}

type GameRuleValue interface {
	isGameRuleValue()
}

// MarshalGameRuleValue reads or writes the GameRuleValue union using its canonical wire layout.
func MarshalGameRuleValue(io IO, x *GameRuleValue) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(GameRuleValueEmpty)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(GameRuleValueBool)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(GameRuleValueInt32)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(GameRuleValueFloat)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *GameRuleValueEmpty:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *GameRuleValueBool:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *GameRuleValueInt32:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *GameRuleValueFloat:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

type GameRuleValueBool struct {
	Value bool
}

func (*GameRuleValueBool) isGameRuleValue() {}

// Marshal reads or writes GameRuleValueBool using its canonical wire layout.
func (x *GameRuleValueBool) Marshal(io IO) {
	io.Bool(&x.Value)
}

type GameRuleValueEmpty struct {
}

func (*GameRuleValueEmpty) isGameRuleValue() {}

// Marshal reads or writes GameRuleValueEmpty using its canonical wire layout.
func (x *GameRuleValueEmpty) Marshal(io IO) {
}

type GameRuleValueFloat struct {
	Value float32
}

func (*GameRuleValueFloat) isGameRuleValue() {}

// Marshal reads or writes GameRuleValueFloat using its canonical wire layout.
func (x *GameRuleValueFloat) Marshal(io IO) {
	io.Float32(&x.Value)
}

type GameRuleValueInt32 struct {
	Value int32
}

func (*GameRuleValueInt32) isGameRuleValue() {}

// Marshal reads or writes GameRuleValueInt32 using its canonical wire layout.
func (x *GameRuleValueInt32) Marshal(io IO) {
	io.Int32(&x.Value)
}

type InventoryTransactionValue interface {
	isInventoryTransactionValue()
}

// MarshalInventoryTransactionValue reads or writes the InventoryTransactionValue union using its canonical wire layout.
func MarshalInventoryTransactionValue(io IO, x *InventoryTransactionValue) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(NormalTransactionData)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(InventoryMismatchData)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(ItemUseInventoryTransaction)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(ItemUseOnActorInventoryTransaction)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(ItemReleaseInventoryTransaction)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *NormalTransactionData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *InventoryMismatchData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemUseInventoryTransaction:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemUseOnActorInventoryTransaction:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ItemReleaseInventoryTransaction:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

type PrimitiveShapeExtraShapeData interface {
	isPrimitiveShapeExtraShapeData()
}

// MarshalPrimitiveShapeExtraShapeData reads or writes the PrimitiveShapeExtraShapeData union using its canonical wire layout.
func MarshalPrimitiveShapeExtraShapeData(io IO, x *PrimitiveShapeExtraShapeData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(PrimitiveShapeExtraShapeDataEmpty)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(ArrowData)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(TextShape)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(BoxData)
				value.Marshal(io)
				*x = value
			case 4:
				value := new(LineData)
				value.Marshal(io)
				*x = value
			case 5:
				value := new(SphereData)
				value.Marshal(io)
				*x = value
			case 6:
				value := new(CylinderData)
				value.Marshal(io)
				*x = value
			case 7:
				value := new(PyramidData)
				value.Marshal(io)
				*x = value
			case 8:
				value := new(EllipsoidData)
				value.Marshal(io)
				*x = value
			case 9:
				value := new(ConeData)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *PrimitiveShapeExtraShapeDataEmpty:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ArrowData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *TextShape:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *BoxData:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *LineData:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *SphereData:
				tag := uint32(5)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *CylinderData:
				tag := uint32(6)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *PyramidData:
				tag := uint32(7)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *EllipsoidData:
				tag := uint32(8)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ConeData:
				tag := uint32(9)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

type PrimitiveShapeExtraShapeDataEmpty struct {
}

func (*PrimitiveShapeExtraShapeDataEmpty) isPrimitiveShapeExtraShapeData() {}

// Marshal reads or writes PrimitiveShapeExtraShapeDataEmpty using its canonical wire layout.
func (x *PrimitiveShapeExtraShapeDataEmpty) Marshal(io IO) {
}

type ServerboundPackSettingChangePackSettingValue interface {
	isServerboundPackSettingChangePackSettingValue()
}

// MarshalServerboundPackSettingChangePackSettingValue reads or writes the ServerboundPackSettingChangePackSettingValue union using its canonical wire layout.
func MarshalServerboundPackSettingChangePackSettingValue(io IO, x *ServerboundPackSettingChangePackSettingValue) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(ServerboundPackSettingChangePackSettingValueFloat)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(ServerboundPackSettingChangePackSettingValueBool)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(ServerboundPackSettingChangePackSettingValueString)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *ServerboundPackSettingChangePackSettingValueFloat:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ServerboundPackSettingChangePackSettingValueBool:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *ServerboundPackSettingChangePackSettingValueString:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

type ServerboundPackSettingChangePackSettingValueBool struct {
	Value bool
}

func (*ServerboundPackSettingChangePackSettingValueBool) isServerboundPackSettingChangePackSettingValue() {
}

// Marshal reads or writes ServerboundPackSettingChangePackSettingValueBool using its canonical wire layout.
func (x *ServerboundPackSettingChangePackSettingValueBool) Marshal(io IO) {
	io.Bool(&x.Value)
}

type ServerboundPackSettingChangePackSettingValueFloat struct {
	Value float32
}

func (*ServerboundPackSettingChangePackSettingValueFloat) isServerboundPackSettingChangePackSettingValue() {
}

// Marshal reads or writes ServerboundPackSettingChangePackSettingValueFloat using its canonical wire layout.
func (x *ServerboundPackSettingChangePackSettingValueFloat) Marshal(io IO) {
	io.Float32(&x.Value)
}

type ServerboundPackSettingChangePackSettingValueString struct {
	Value string
}

func (*ServerboundPackSettingChangePackSettingValueString) isServerboundPackSettingChangePackSettingValue() {
}

// Marshal reads or writes ServerboundPackSettingChangePackSettingValueString using its canonical wire layout.
func (x *ServerboundPackSettingChangePackSettingValueString) Marshal(io IO) {
	io.String(&x.Value)
}

type SetScoreInfoItem interface {
	isSetScoreInfoItem()
}

// MarshalSetScoreInfoItem reads or writes the SetScoreInfoItem union using its canonical wire layout.
func MarshalSetScoreInfoItem(io IO, x *SetScoreInfoItem) {
	UnionFunc(io,
		func() {
			var tag uint8
			io.Uint8(&tag)
			switch int64(tag) {
			case 0:
				value := new(RemoveScore)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(ChangePlayerScore)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(ChangeEntityScore)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(ChangeFakePlayerScore)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *RemoveScore:
				tag := uint8(0)
				io.Uint8(&tag)
				value.Marshal(io)
			case *ChangePlayerScore:
				tag := uint8(1)
				io.Uint8(&tag)
				value.Marshal(io)
			case *ChangeEntityScore:
				tag := uint8(2)
				io.Uint8(&tag)
				value.Marshal(io)
			case *ChangeFakePlayerScore:
				tag := uint8(3)
				io.Uint8(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

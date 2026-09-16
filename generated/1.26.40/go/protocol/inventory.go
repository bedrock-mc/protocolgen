// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

// InventoryAction represents a single action that took place during an inventory transaction. On
// itself, this inventory action is always unbalanced: It must be combined with other actions in an
// inventory transaction to form a balanced transaction.
type InventoryAction struct {
	Source   InventorySource
	Slot     uint32
	FromItem NetworkItemStackDescriptorSerializedData
	ToItem   NetworkItemStackDescriptorSerializedData
}

// Marshal reads or writes InventoryAction using its canonical wire layout.
func (x *InventoryAction) Marshal(io IO) {
	x.Source.Marshal(io)
	io.Varuint32(&x.Slot)
	x.FromItem.Marshal(io)
	x.ToItem.Marshal(io)
}

type InventoryLayout int32

const (
	InventoryLayoutNone           InventoryLayout = 0
	InventoryLayoutInventoryOnly  InventoryLayout = 1
	InventoryLayoutDefault        InventoryLayout = 2
	InventoryLayoutRecipeBookOnly InventoryLayout = 3
)

// Marshal reads or writes InventoryLayout through its int32 wire encoding.
func (x *InventoryLayout) Marshal(io IO) { io.Varint32((*int32)(x)) }

type InventoryLeftTabIndex int32

const (
	InventoryLeftTabIndexNone               InventoryLeftTabIndex = 0
	InventoryLeftTabIndexRecipeConstruction InventoryLeftTabIndex = 1
	InventoryLeftTabIndexRecipeEquipment    InventoryLeftTabIndex = 2
	InventoryLeftTabIndexRecipeItems        InventoryLeftTabIndex = 3
	InventoryLeftTabIndexRecipeNature       InventoryLeftTabIndex = 4
	InventoryLeftTabIndexRecipeSearch       InventoryLeftTabIndex = 5
	InventoryLeftTabIndexSurvival           InventoryLeftTabIndex = 6
)

// Marshal reads or writes InventoryLeftTabIndex through its int32 wire encoding.
func (x *InventoryLeftTabIndex) Marshal(io IO) { io.Varint32((*int32)(x)) }

type InventoryMismatchData struct {
	Actions InventoryTransactionData
}

func (*InventoryMismatchData) tagInventoryTransactionValue() uint32 { return 1 }

// Marshal reads or writes InventoryMismatchData using its canonical wire layout.
func (x *InventoryMismatchData) Marshal(io IO) {
	x.Actions.Marshal(io)
}

type InventoryOptions struct {
	LeftInventoryTab  InventoryLeftTabIndex
	RightInventoryTab InventoryRightTabIndex
	Filtering         bool
	LayoutInv         InventoryLayout
	LayoutCraft       InventoryLayout
}

// Marshal reads or writes InventoryOptions using its canonical wire layout.
func (x *InventoryOptions) Marshal(io IO) {
	x.LeftInventoryTab.Marshal(io)
	x.RightInventoryTab.Marshal(io)
	io.Bool(&x.Filtering)
	x.LayoutInv.Marshal(io)
	x.LayoutCraft.Marshal(io)
}

type InventoryRightTabIndex int32

const (
	InventoryRightTabIndexNone       InventoryRightTabIndex = 0
	InventoryRightTabIndexFullScreen InventoryRightTabIndex = 1
	InventoryRightTabIndexCrafting   InventoryRightTabIndex = 2
	InventoryRightTabIndexArmor      InventoryRightTabIndex = 3
)

// Marshal reads or writes InventoryRightTabIndex through its int32 wire encoding.
func (x *InventoryRightTabIndex) Marshal(io IO) { io.Varint32((*int32)(x)) }

type InventorySource struct {
	SourceType  InventorySourceType
	ContainerID Optional[int8]
	BitFlags    Optional[InventorySourceInventorySourceFlags]
}

// Marshal reads or writes InventorySource using its canonical wire layout.
func (x *InventorySource) Marshal(io IO) {
	x.SourceType.Marshal(io)
	DoubleOptionalFunc(io, &x.ContainerID, io.Int8)
	DoubleOptionalFunc(io, &x.BitFlags, func(value *InventorySourceInventorySourceFlags) {
		value.Marshal(io)
	})
}

type InventorySourceInventorySourceFlags uint32

const (
	InventorySourceInventorySourceFlagsNoFlag                 InventorySourceInventorySourceFlags = 0
	InventorySourceInventorySourceFlagsWorldInteractionRandom InventorySourceInventorySourceFlags = 1
)

// Marshal reads or writes InventorySourceInventorySourceFlags through its uint32 wire encoding.
func (x *InventorySourceInventorySourceFlags) Marshal(io IO) { io.Varuint32((*uint32)(x)) }

type InventorySourceType uint32

const (
	InventorySourceTypeContainerInventory        InventorySourceType = 0
	InventorySourceTypeGlobalInventory           InventorySourceType = 1
	InventorySourceTypeWorldInteraction          InventorySourceType = 2
	InventorySourceTypeCreativeInventory         InventorySourceType = 3
	InventorySourceTypeNonImplementedFeatureTODO InventorySourceType = 99999
)

// Marshal reads or writes InventorySourceType through its uint32 wire encoding.
func (x *InventorySourceType) Marshal(io IO) { io.Varuint32((*uint32)(x)) }

// InventoryTransactionData represents an object that holds data specific to an inventory
// transaction type. The data it holds depends on the type.
type InventoryTransactionData struct {
	Actions Optional[[]InventoryAction]
}

// Marshal reads or writes InventoryTransactionData using its canonical wire layout.
func (x *InventoryTransactionData) Marshal(io IO) {
	OptionalFunc(io, &x.Actions, func(value *[]InventoryAction) {
		Slice(io, value)
	})
}

// NormalTransactionData represents an inventory transaction data object for normal transactions,
// such as crafting. It has no content.
type NormalTransactionData struct {
	Actions InventoryTransactionData
}

func (*NormalTransactionData) tagInventoryTransactionValue() uint32 { return 0 }

// Marshal reads or writes NormalTransactionData using its canonical wire layout.
func (x *NormalTransactionData) Marshal(io IO) {
	x.Actions.Marshal(io)
}

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

type InventoryMismatchData struct {
	Actions InventoryTransactionData
}

func (*InventoryMismatchData) isInventoryTransactionPacketData() {}

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
	IntegerFunc(&x.LeftInventoryTab, io.Varint32)
	IntegerFunc(&x.RightInventoryTab, io.Varint32)
	io.Bool(&x.Filtering)
	IntegerFunc(&x.LayoutInv, io.Varint32)
	IntegerFunc(&x.LayoutCraft, io.Varint32)
}

type InventoryRightTabIndex int32

const (
	InventoryRightTabIndexNone       InventoryRightTabIndex = 0
	InventoryRightTabIndexFullScreen InventoryRightTabIndex = 1
	InventoryRightTabIndexCrafting   InventoryRightTabIndex = 2
	InventoryRightTabIndexArmor      InventoryRightTabIndex = 3
)

type InventorySource struct {
	SourceType  InventorySourceType
	ContainerID Optional[int8]
	BitFlags    Optional[InventorySourceInventorySourceFlags]
}

// Marshal reads or writes InventorySource using its canonical wire layout.
func (x *InventorySource) Marshal(io IO) {
	IntegerFunc(&x.SourceType, io.Varuint32)
	OptionalFunc(io, &x.ContainerID, io.Int8)
	OptionalFunc(io, &x.BitFlags, func(value *InventorySourceInventorySourceFlags) {
		IntegerFunc(value, io.Varuint32)
	})
}

type InventorySourceInventorySourceFlags uint32

const (
	InventorySourceInventorySourceFlagsNoFlag                 InventorySourceInventorySourceFlags = 0
	InventorySourceInventorySourceFlagsWorldInteractionRandom InventorySourceInventorySourceFlags = 1
)

type InventorySourceType uint32

const (
	InventorySourceTypeContainerInventory        InventorySourceType = 0
	InventorySourceTypeGlobalInventory           InventorySourceType = 1
	InventorySourceTypeWorldInteraction          InventorySourceType = 2
	InventorySourceTypeCreativeInventory         InventorySourceType = 3
	InventorySourceTypeNonImplementedFeatureTODO InventorySourceType = 99999
)

// InventoryTransactionData represents an object that holds data specific to an inventory
// transaction type. The data it holds depends on the type.
type InventoryTransactionData struct {
	Actions []InventoryAction
}

// Marshal reads or writes InventoryTransactionData using its canonical wire layout.
func (x *InventoryTransactionData) Marshal(io IO) {
	Slice(io, &x.Actions)
}

// NormalTransactionData represents an inventory transaction data object for normal transactions,
// such as crafting. It has no content.
type NormalTransactionData struct {
	Actions InventoryTransactionData
}

func (*NormalTransactionData) isInventoryTransactionPacketData() {}

// Marshal reads or writes NormalTransactionData using its canonical wire layout.
func (x *NormalTransactionData) Marshal(io IO) {
	x.Actions.Marshal(io)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ContainerEnumName uint8

const (
	ContainerEnumNameAnvilinputcontainer                 ContainerEnumName = 0
	ContainerEnumNameAnvilmaterialcontainer              ContainerEnumName = 1
	ContainerEnumNameAnvilresultpreviewcontainer         ContainerEnumName = 2
	ContainerEnumNameSmithingtableinputcontainer         ContainerEnumName = 3
	ContainerEnumNameSmithingtablematerialcontainer      ContainerEnumName = 4
	ContainerEnumNameSmithingtableresultpreviewcontainer ContainerEnumName = 5
	ContainerEnumNameArmorcontainer                      ContainerEnumName = 6
	ContainerEnumNameLevelentitycontainer                ContainerEnumName = 7
	ContainerEnumNameBeaconpaymentcontainer              ContainerEnumName = 8
	ContainerEnumNameBrewingstandinputcontainer          ContainerEnumName = 9
	ContainerEnumNameBrewingstandresultcontainer         ContainerEnumName = 10
	ContainerEnumNameBrewingstandfuelcontainer           ContainerEnumName = 11
	ContainerEnumNameCombinedhotbarandinventorycontainer ContainerEnumName = 12
	ContainerEnumNameCraftinginputcontainer              ContainerEnumName = 13
	ContainerEnumNameCraftingoutputpreviewcontainer      ContainerEnumName = 14
	ContainerEnumNameRecipeconstructioncontainer         ContainerEnumName = 15
	ContainerEnumNameRecipenaturecontainer               ContainerEnumName = 16
	ContainerEnumNameRecipeitemscontainer                ContainerEnumName = 17
	ContainerEnumNameRecipesearchcontainer               ContainerEnumName = 18
	ContainerEnumNameRecipesearchbarcontainer            ContainerEnumName = 19
	ContainerEnumNameRecipeequipmentcontainer            ContainerEnumName = 20
	ContainerEnumNameRecipebookcontainer                 ContainerEnumName = 21
	ContainerEnumNameEnchantinginputcontainer            ContainerEnumName = 22
	ContainerEnumNameEnchantingmaterialcontainer         ContainerEnumName = 23
	ContainerEnumNameFurnacefuelcontainer                ContainerEnumName = 24
	ContainerEnumNameFurnaceingredientcontainer          ContainerEnumName = 25
	ContainerEnumNameFurnaceresultcontainer              ContainerEnumName = 26
	ContainerEnumNameHorseequipcontainer                 ContainerEnumName = 27
	ContainerEnumNameHotbarcontainer                     ContainerEnumName = 28
	ContainerEnumNameInventorycontainer                  ContainerEnumName = 29
	ContainerEnumNameShulkerboxcontainer                 ContainerEnumName = 30
	ContainerEnumNameTradeingredient1container           ContainerEnumName = 31
	ContainerEnumNameTradeingredient2container           ContainerEnumName = 32
	ContainerEnumNameTraderesultpreviewcontainer         ContainerEnumName = 33
	ContainerEnumNameOffhandcontainer                    ContainerEnumName = 34
	ContainerEnumNameCompoundcreatorinput                ContainerEnumName = 35
	ContainerEnumNameCompoundcreatoroutputpreview        ContainerEnumName = 36
	ContainerEnumNameElementconstructoroutputpreview     ContainerEnumName = 37
	ContainerEnumNameMaterialreducerinput                ContainerEnumName = 38
	ContainerEnumNameMaterialreduceroutput               ContainerEnumName = 39
	ContainerEnumNameLabtableinput                       ContainerEnumName = 40
	ContainerEnumNameLoominputcontainer                  ContainerEnumName = 41
	ContainerEnumNameLoomdyecontainer                    ContainerEnumName = 42
	ContainerEnumNameLoommaterialcontainer               ContainerEnumName = 43
	ContainerEnumNameLoomresultpreviewcontainer          ContainerEnumName = 44
	ContainerEnumNameBlastfurnaceingredientcontainer     ContainerEnumName = 45
	ContainerEnumNameSmokeringredientcontainer           ContainerEnumName = 46
	ContainerEnumNameTrade2ingredient1container          ContainerEnumName = 47
	ContainerEnumNameTrade2ingredient2container          ContainerEnumName = 48
	ContainerEnumNameTrade2resultpreviewcontainer        ContainerEnumName = 49
	ContainerEnumNameGrindstoneinputcontainer            ContainerEnumName = 50
	ContainerEnumNameGrindstoneadditionalcontainer       ContainerEnumName = 51
	ContainerEnumNameGrindstoneresultpreviewcontainer    ContainerEnumName = 52
	ContainerEnumNameStonecutterinputcontainer           ContainerEnumName = 53
	ContainerEnumNameStonecutterresultpreviewcontainer   ContainerEnumName = 54
	ContainerEnumNameCartographyinputcontainer           ContainerEnumName = 55
	ContainerEnumNameCartographyadditionalcontainer      ContainerEnumName = 56
	ContainerEnumNameCartographyresultpreviewcontainer   ContainerEnumName = 57
	ContainerEnumNameBarrelcontainer                     ContainerEnumName = 58
	ContainerEnumNameCursorcontainer                     ContainerEnumName = 59
	ContainerEnumNameCreatedoutputcontainer              ContainerEnumName = 60
	ContainerEnumNameSmithingtabletemplatecontainer      ContainerEnumName = 61
	ContainerEnumNameCrafterlevelentitycontainer         ContainerEnumName = 62
	ContainerEnumNameDynamiccontainer                    ContainerEnumName = 63
	ContainerEnumNameRecipefoodcontainer                 ContainerEnumName = 64
	ContainerEnumNameRecipeblockscontainer               ContainerEnumName = 65
	ContainerEnumNameRecipefurnaceitemscontainer         ContainerEnumName = 66
)

// Marshal reads or writes ContainerEnumName through its uint8 wire encoding.
func (x *ContainerEnumName) Marshal(io IO) { io.Uint8((*uint8)(x)) }

type ContainerMixDataEntry struct {
	FromItemID    int32
	ReagentItemID int32
	ToItemID      int32
}

// Marshal reads or writes ContainerMixDataEntry using its canonical wire layout.
func (x *ContainerMixDataEntry) Marshal(io IO) {
	io.Varint32(&x.FromItemID)
	io.Varint32(&x.ReagentItemID)
	io.Varint32(&x.ToItemID)
}

// FullContainerName contains information required to identify a container in a StackRequestSlotInfo.
type FullContainerName struct {
	// ContainerName is the ID of the container that the slot was in.
	ContainerName ContainerEnumName
	// DynamicID is the ID of the container if it is dynamic. If the container is not dynamic, this field should
	// be left empty. A non-optional value of 0 is assumed to be non-empty.
	DynamicID Optional[uint32]
}

// Marshal reads or writes FullContainerName using its canonical wire layout.
func (x *FullContainerName) Marshal(io IO) {
	x.ContainerName.Marshal(io)
	OptionalFunc(io, &x.DynamicID, io.Uint32)
}

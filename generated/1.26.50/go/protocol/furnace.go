// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type FurnaceLayout int32

const (
	FurnaceLayoutNone          FurnaceLayout = 0
	FurnaceLayoutInventoryOnly FurnaceLayout = 1
	FurnaceLayoutDefault       FurnaceLayout = 2
)

type FurnaceLeftTabIndex int32

const (
	FurnaceLeftTabIndexNone         FurnaceLeftTabIndex = 0
	FurnaceLeftTabIndexRecipeFood   FurnaceLeftTabIndex = 1
	FurnaceLeftTabIndexRecipeItems  FurnaceLeftTabIndex = 2
	FurnaceLeftTabIndexRecipeBlocks FurnaceLeftTabIndex = 3
	FurnaceLeftTabIndexRecipeSearch FurnaceLeftTabIndex = 4
	FurnaceLeftTabIndexInventory    FurnaceLeftTabIndex = 5
)

type FurnaceOptions struct {
	LeftFurnaceTab FurnaceLeftTabIndex
	Filtering      bool
	Layout         FurnaceLayout
}

// Marshal reads or writes FurnaceOptions using its canonical wire layout.
func (x *FurnaceOptions) Marshal(io IO) {
	IntegerFunc(&x.LeftFurnaceTab, io.Varint32)
	io.Bool(&x.Filtering)
	IntegerFunc(&x.Layout, io.Varint32)
}

type FurnaceType uint8

const (
	FurnaceTypeNone         FurnaceType = 0
	FurnaceTypeFurnace      FurnaceType = 1
	FurnaceTypeBlastFurnace FurnaceType = 2
	FurnaceTypeSmoker       FurnaceType = 3
)

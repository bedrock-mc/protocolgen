// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CraftingData struct {
	ShapedRecipes             []ShapedRecipe
	ShapelessRecipes          []ShapelessRecipe
	MultiRecipes              []MultiRecipe
	UserDataShapelessRecipes  []ShapelessRecipe
	ShapelessChemistryRecipes []ShapelessRecipe
	ShapedChemistryRecipes    []ShapedRecipe
	SmithingTransformRecipes  []SmithingTransformRecipe
	SmithingTrimRecipes       []SmithingTrimRecipe
	PotionMixes               []PotionMixDataEntry
	ContainerMixes            []ContainerMixDataEntry
	MaterialReducers          []MaterialReducerDataEntry
	ClearRecipes              bool
}

// Marshal reads or writes CraftingData using its canonical wire layout.
func (x *CraftingData) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.ShapedRecipes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ShapedRecipes), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.ShapedRecipes))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.ShapedRecipes = make([]ShapedRecipe, int(count1))
	}
	for index2 := range x.ShapedRecipes {
		x.ShapedRecipes[index2].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.ShapelessRecipes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ShapelessRecipes), "collection length overflows uint32")
		return
	}
	count3 := uint32(len(x.ShapelessRecipes))
	io.Varuint32(&count3)
	if io.Reading() {
		if uint64(count3) > uint64(^uint(0)>>1) {
			io.InvalidValue(count3, "collection length overflows int")
			return
		}
		x.ShapelessRecipes = make([]ShapelessRecipe, int(count3))
	}
	for index4 := range x.ShapelessRecipes {
		x.ShapelessRecipes[index4].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.MultiRecipes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.MultiRecipes), "collection length overflows uint32")
		return
	}
	count5 := uint32(len(x.MultiRecipes))
	io.Varuint32(&count5)
	if io.Reading() {
		if uint64(count5) > uint64(^uint(0)>>1) {
			io.InvalidValue(count5, "collection length overflows int")
			return
		}
		x.MultiRecipes = make([]MultiRecipe, int(count5))
	}
	for index6 := range x.MultiRecipes {
		x.MultiRecipes[index6].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.UserDataShapelessRecipes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.UserDataShapelessRecipes), "collection length overflows uint32")
		return
	}
	count7 := uint32(len(x.UserDataShapelessRecipes))
	io.Varuint32(&count7)
	if io.Reading() {
		if uint64(count7) > uint64(^uint(0)>>1) {
			io.InvalidValue(count7, "collection length overflows int")
			return
		}
		x.UserDataShapelessRecipes = make([]ShapelessRecipe, int(count7))
	}
	for index8 := range x.UserDataShapelessRecipes {
		x.UserDataShapelessRecipes[index8].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.ShapelessChemistryRecipes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ShapelessChemistryRecipes), "collection length overflows uint32")
		return
	}
	count9 := uint32(len(x.ShapelessChemistryRecipes))
	io.Varuint32(&count9)
	if io.Reading() {
		if uint64(count9) > uint64(^uint(0)>>1) {
			io.InvalidValue(count9, "collection length overflows int")
			return
		}
		x.ShapelessChemistryRecipes = make([]ShapelessRecipe, int(count9))
	}
	for index10 := range x.ShapelessChemistryRecipes {
		x.ShapelessChemistryRecipes[index10].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.ShapedChemistryRecipes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ShapedChemistryRecipes), "collection length overflows uint32")
		return
	}
	count11 := uint32(len(x.ShapedChemistryRecipes))
	io.Varuint32(&count11)
	if io.Reading() {
		if uint64(count11) > uint64(^uint(0)>>1) {
			io.InvalidValue(count11, "collection length overflows int")
			return
		}
		x.ShapedChemistryRecipes = make([]ShapedRecipe, int(count11))
	}
	for index12 := range x.ShapedChemistryRecipes {
		x.ShapedChemistryRecipes[index12].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.SmithingTransformRecipes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.SmithingTransformRecipes), "collection length overflows uint32")
		return
	}
	count13 := uint32(len(x.SmithingTransformRecipes))
	io.Varuint32(&count13)
	if io.Reading() {
		if uint64(count13) > uint64(^uint(0)>>1) {
			io.InvalidValue(count13, "collection length overflows int")
			return
		}
		x.SmithingTransformRecipes = make([]SmithingTransformRecipe, int(count13))
	}
	for index14 := range x.SmithingTransformRecipes {
		x.SmithingTransformRecipes[index14].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.SmithingTrimRecipes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.SmithingTrimRecipes), "collection length overflows uint32")
		return
	}
	count15 := uint32(len(x.SmithingTrimRecipes))
	io.Varuint32(&count15)
	if io.Reading() {
		if uint64(count15) > uint64(^uint(0)>>1) {
			io.InvalidValue(count15, "collection length overflows int")
			return
		}
		x.SmithingTrimRecipes = make([]SmithingTrimRecipe, int(count15))
	}
	for index16 := range x.SmithingTrimRecipes {
		x.SmithingTrimRecipes[index16].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.PotionMixes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.PotionMixes), "collection length overflows uint32")
		return
	}
	count17 := uint32(len(x.PotionMixes))
	io.Varuint32(&count17)
	if io.Reading() {
		if uint64(count17) > uint64(^uint(0)>>1) {
			io.InvalidValue(count17, "collection length overflows int")
			return
		}
		x.PotionMixes = make([]PotionMixDataEntry, int(count17))
	}
	for index18 := range x.PotionMixes {
		x.PotionMixes[index18].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.ContainerMixes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ContainerMixes), "collection length overflows uint32")
		return
	}
	count19 := uint32(len(x.ContainerMixes))
	io.Varuint32(&count19)
	if io.Reading() {
		if uint64(count19) > uint64(^uint(0)>>1) {
			io.InvalidValue(count19, "collection length overflows int")
			return
		}
		x.ContainerMixes = make([]ContainerMixDataEntry, int(count19))
	}
	for index20 := range x.ContainerMixes {
		x.ContainerMixes[index20].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.MaterialReducers)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.MaterialReducers), "collection length overflows uint32")
		return
	}
	count21 := uint32(len(x.MaterialReducers))
	io.Varuint32(&count21)
	if io.Reading() {
		if uint64(count21) > uint64(^uint(0)>>1) {
			io.InvalidValue(count21, "collection length overflows int")
			return
		}
		x.MaterialReducers = make([]MaterialReducerDataEntry, int(count21))
	}
	for index22 := range x.MaterialReducers {
		x.MaterialReducers[index22].Marshal(io)
	}
	io.Bool(&x.ClearRecipes)
}

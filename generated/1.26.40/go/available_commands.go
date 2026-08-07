// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type AvailableCommands struct {
	EnumValues              []string
	ChainedSubcommandValues []string
	PostFixes               []string
	EnumData                []AvailableCommandsEnumData
	ChainedSubcommandData   []AvailableCommandsChainedSubcommandData
	Commands                []AvailableCommandsPacketCommandData
	SoftEnums               []AvailableCommandsSoftEnumData
	Constraints             []AvailableCommandsConstrainedValueData
}

// Marshal reads or writes AvailableCommands using its canonical wire layout.
func (x *AvailableCommands) Marshal(io IO) {
	if !io.Reading() && uint64(len(x.EnumValues)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.EnumValues), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.EnumValues))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.EnumValues = make([]string, int(count1))
	}
	for index2 := range x.EnumValues {
		io.String(&x.EnumValues[index2])
	}
	if !io.Reading() && uint64(len(x.ChainedSubcommandValues)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ChainedSubcommandValues), "collection length overflows uint32")
		return
	}
	count3 := uint32(len(x.ChainedSubcommandValues))
	io.Varuint32(&count3)
	if io.Reading() {
		if uint64(count3) > uint64(^uint(0)>>1) {
			io.InvalidValue(count3, "collection length overflows int")
			return
		}
		x.ChainedSubcommandValues = make([]string, int(count3))
	}
	for index4 := range x.ChainedSubcommandValues {
		io.String(&x.ChainedSubcommandValues[index4])
	}
	if !io.Reading() && uint64(len(x.PostFixes)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.PostFixes), "collection length overflows uint32")
		return
	}
	count5 := uint32(len(x.PostFixes))
	io.Varuint32(&count5)
	if io.Reading() {
		if uint64(count5) > uint64(^uint(0)>>1) {
			io.InvalidValue(count5, "collection length overflows int")
			return
		}
		x.PostFixes = make([]string, int(count5))
	}
	for index6 := range x.PostFixes {
		io.String(&x.PostFixes[index6])
	}
	if !io.Reading() && uint64(len(x.EnumData)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.EnumData), "collection length overflows uint32")
		return
	}
	count7 := uint32(len(x.EnumData))
	io.Varuint32(&count7)
	if io.Reading() {
		if uint64(count7) > uint64(^uint(0)>>1) {
			io.InvalidValue(count7, "collection length overflows int")
			return
		}
		x.EnumData = make([]AvailableCommandsEnumData, int(count7))
	}
	for index8 := range x.EnumData {
		x.EnumData[index8].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.ChainedSubcommandData)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ChainedSubcommandData), "collection length overflows uint32")
		return
	}
	count9 := uint32(len(x.ChainedSubcommandData))
	io.Varuint32(&count9)
	if io.Reading() {
		if uint64(count9) > uint64(^uint(0)>>1) {
			io.InvalidValue(count9, "collection length overflows int")
			return
		}
		x.ChainedSubcommandData = make([]AvailableCommandsChainedSubcommandData, int(count9))
	}
	for index10 := range x.ChainedSubcommandData {
		x.ChainedSubcommandData[index10].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.Commands)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Commands), "collection length overflows uint32")
		return
	}
	count11 := uint32(len(x.Commands))
	io.Varuint32(&count11)
	if io.Reading() {
		if uint64(count11) > uint64(^uint(0)>>1) {
			io.InvalidValue(count11, "collection length overflows int")
			return
		}
		x.Commands = make([]AvailableCommandsPacketCommandData, int(count11))
	}
	for index12 := range x.Commands {
		x.Commands[index12].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.SoftEnums)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.SoftEnums), "collection length overflows uint32")
		return
	}
	count13 := uint32(len(x.SoftEnums))
	io.Varuint32(&count13)
	if io.Reading() {
		if uint64(count13) > uint64(^uint(0)>>1) {
			io.InvalidValue(count13, "collection length overflows int")
			return
		}
		x.SoftEnums = make([]AvailableCommandsSoftEnumData, int(count13))
	}
	for index14 := range x.SoftEnums {
		x.SoftEnums[index14].Marshal(io)
	}
	if !io.Reading() && uint64(len(x.Constraints)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Constraints), "collection length overflows uint32")
		return
	}
	count15 := uint32(len(x.Constraints))
	io.Varuint32(&count15)
	if io.Reading() {
		if uint64(count15) > uint64(^uint(0)>>1) {
			io.InvalidValue(count15, "collection length overflows int")
			return
		}
		x.Constraints = make([]AvailableCommandsConstrainedValueData, int(count15))
	}
	for index16 := range x.Constraints {
		x.Constraints[index16].Marshal(io)
	}
}

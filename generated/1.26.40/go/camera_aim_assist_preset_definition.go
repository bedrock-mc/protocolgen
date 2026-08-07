// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraAimAssistPresetDefinition struct {
	Identifier          string
	ExclusionSettings   CameraAimAssistPresetExclusionDefinition
	LiquidTargetingList []string
	ItemSettings        []OrderedEntry[string, string]
	DefaultItemSettings Optional[string]
	HandSettings        Optional[string]
}

// Marshal reads or writes CameraAimAssistPresetDefinition using its canonical wire layout.
func (x *CameraAimAssistPresetDefinition) Marshal(io IO) {
	io.String(&x.Identifier)
	x.ExclusionSettings.Marshal(io)
	FuncSlice(io, &x.LiquidTargetingList, io.Varuint32, io.String)
	OrderedMap(io, &x.ItemSettings, io.Varuint32, io.String, io.String)
	OptionalFunc(io, &x.DefaultItemSettings, io.String)
	OptionalFunc(io, &x.HandSettings, io.String)
}

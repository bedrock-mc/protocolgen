// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ServerboundPackSettingChangePackSettingValueString struct {
	Value string
}

func (*ServerboundPackSettingChangePackSettingValueString) isServerboundPackSettingChangePackSettingValue() {
}

// Marshal reads or writes ServerboundPackSettingChangePackSettingValueString using its canonical wire layout.
func (x *ServerboundPackSettingChangePackSettingValueString) Marshal(io IO) {
	io.String(&x.Value)
}

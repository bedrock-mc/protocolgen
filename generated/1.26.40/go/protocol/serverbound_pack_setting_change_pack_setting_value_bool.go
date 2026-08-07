// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ServerboundPackSettingChangePackSettingValueBool struct {
	Value bool
}

func (ServerboundPackSettingChangePackSettingValueBool) isServerboundPackSettingChangePackSettingValue() {
}

// Marshal reads or writes ServerboundPackSettingChangePackSettingValueBool using its canonical wire layout.
func (x *ServerboundPackSettingChangePackSettingValueBool) Marshal(io IO) {
	io.Bool(&x.Value)
}

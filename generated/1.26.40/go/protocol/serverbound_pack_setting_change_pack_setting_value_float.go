// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ServerboundPackSettingChangePackSettingValueFloat struct {
	Value float32
}

func (ServerboundPackSettingChangePackSettingValueFloat) isServerboundPackSettingChangePackSettingValue() {
}

// Marshal reads or writes ServerboundPackSettingChangePackSettingValueFloat using its canonical wire layout.
func (x *ServerboundPackSettingChangePackSettingValueFloat) Marshal(io IO) {
	io.Float32(&x.Value)
}

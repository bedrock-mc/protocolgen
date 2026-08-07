// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetDefaultGameType struct {
	DefaultGameType GameType
}

// Marshal reads or writes SetDefaultGameType using its canonical wire layout.
func (x *SetDefaultGameType) Marshal(io IO) {
	IntegerFunc(&x.DefaultGameType, io.Varint32)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetDifficulty struct {
	Difficulty uint32
}

// Marshal reads or writes SetDifficulty using its canonical wire layout.
func (x *SetDifficulty) Marshal(io IO) {
	io.Varuint32(&x.Difficulty)
}

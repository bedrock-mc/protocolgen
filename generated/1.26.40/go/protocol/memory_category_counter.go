// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MemoryCategoryCounter struct {
	Category     MemoryCategory
	CurrentBytes uint64
}

// Marshal reads or writes MemoryCategoryCounter using its canonical wire layout.
func (x *MemoryCategoryCounter) Marshal(io IO) {
	IntegerFunc(&x.Category, io.Uint8)
	io.Uint64(&x.CurrentBytes)
}

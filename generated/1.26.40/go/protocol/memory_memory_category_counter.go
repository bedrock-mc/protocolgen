// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MemoryMemoryCategoryCounter struct {
	Category     MemoryMemoryCategory
	CurrentBytes uint64
}

// Marshal reads or writes MemoryMemoryCategoryCounter using its canonical wire layout.
func (x *MemoryMemoryCategoryCounter) Marshal(io IO) {
	IntegerFunc(&x.Category, io.Uint8)
	io.Uint64(&x.CurrentBytes)
}

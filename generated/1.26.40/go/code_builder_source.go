// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CodeBuilderSource struct {
	Operation  CodeBuilderStorageQueryOptionsOperation
	Category   CodeBuilderStorageQueryOptionsCategory
	CodeStatus CodeBuilderExecutionStateCodeStatus
}

// Marshal reads or writes CodeBuilderSource using its canonical wire layout.
func (x *CodeBuilderSource) Marshal(io IO) {
	IntegerFunc(&x.Operation, io.Uint8)
	IntegerFunc(&x.Category, io.Uint8)
	IntegerFunc(&x.CodeStatus, io.Uint8)
}

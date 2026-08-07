// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CodeBuilderSource struct {
	Operation  CodeBuilderStorageQueryOptionsOperation
	Category   CodeBuilderStorageQueryOptionsCategory
	CodeStatus CodeBuilderExecutionStateCodeStatus
}

// Marshal reads or writes CodeBuilderSource using its canonical wire layout.
func (x *CodeBuilderSource) Marshal(io IO) {
	enumValue1 := uint8(x.Operation)
	io.Uint8(&enumValue1)
	x.Operation = CodeBuilderStorageQueryOptionsOperation(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	enumValue2 := uint8(x.Category)
	io.Uint8(&enumValue2)
	x.Category = CodeBuilderStorageQueryOptionsCategory(enumValue2)
	switch int64(enumValue2) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue2, "unknown enum value")
	}
	enumValue3 := uint8(x.CodeStatus)
	io.Uint8(&enumValue3)
	x.CodeStatus = CodeBuilderExecutionStateCodeStatus(enumValue3)
	switch int64(enumValue3) {
	case 0, 1, 2, 3, 4, 5:
	default:
		io.InvalidValue(enumValue3, "unknown enum value")
	}
}

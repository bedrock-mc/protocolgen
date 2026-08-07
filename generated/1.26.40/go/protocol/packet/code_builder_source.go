// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CodeBuilderSource struct {
	Operation  protocol.CodeBuilderStorageQueryOptionsOperation
	Category   protocol.CodeBuilderStorageQueryOptionsCategory
	CodeStatus protocol.CodeBuilderExecutionStateCodeStatus
}

// Marshal reads or writes CodeBuilderSource using its canonical wire layout.
func (x *CodeBuilderSource) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.Operation, io.Uint8)
	protocol.IntegerFunc(&x.Category, io.Uint8)
	protocol.IntegerFunc(&x.CodeStatus, io.Uint8)
}

// ID returns the protocol ID for CodeBuilderSource.
func (*CodeBuilderSource) ID() uint32 { return IDCodeBuilderSource }

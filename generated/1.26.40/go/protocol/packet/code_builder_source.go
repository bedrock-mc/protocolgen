// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// CodeBuilderSource is an Education Edition packet sent by the client to the server to run an
// operation with a code builder.
type CodeBuilderSource struct {
	// Operation is used to distinguish the operation performed. It is always one of the constants
	// listed above.
	Operation protocol.CodeBuilderStorageQueryOptionsOperation
	// Category is used to distinguish the category of the operation performed. It is always one of the
	// constants listed above.
	Category protocol.CodeBuilderStorageQueryOptionsCategory
	// CodeStatus is the status of the code builder. It is always one of the constants listed above.
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

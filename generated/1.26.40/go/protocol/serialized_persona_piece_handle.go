// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type SerializedPersonaPieceHandle struct {
	PieceID        string
	PieceType      PersonaPieceType
	PackID         uuid.UUID
	IsDefaultPiece bool
	ProductID      string
}

// Marshal reads or writes SerializedPersonaPieceHandle using its canonical wire layout.
func (x *SerializedPersonaPieceHandle) Marshal(io IO) {
	io.String(&x.PieceID)
	IntegerFunc(&x.PieceType, io.Uint32)
	io.UUID(&x.PackID)
	io.Bool(&x.IsDefaultPiece)
	io.String(&x.ProductID)
}

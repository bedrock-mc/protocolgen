// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type SerializedPersonaPieceHandle struct {
	PieceId        string
	PieceType      PersonaPieceType
	PackId         uuid.UUID
	IsDefaultPiece bool
	ProductId      string
}

// Marshal reads or writes SerializedPersonaPieceHandle using its canonical wire layout.
func (x *SerializedPersonaPieceHandle) Marshal(io IO) {
	io.String(&x.PieceId)
	IntegerFunc(&x.PieceType, io.Uint32)
	io.UUID(&x.PackId)
	io.Bool(&x.IsDefaultPiece)
	io.String(&x.ProductId)
}

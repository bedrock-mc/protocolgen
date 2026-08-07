// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"image/color"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// IO is the minimal symmetric wire interface used by generated Marshal methods.
// Reading reports whether calls populate values. InvalidValue must stop the
// current codec operation, typically by panicking or recording a terminal error.
// String and Bytes use a varuint32 byte-length prefix. UUID uses Bedrock's
// little-endian 64-bit halves, NBT consumes exactly one little-endian tag, and
// Bitset uses seven payload bits per continuation byte.
type IO interface {
	Reading() bool
	InvalidValue(value any, context string)

	Bool(*bool)
	Int8(*int8)
	Uint8(*uint8)
	Int16(*int16)
	Uint16(*uint16)
	BEInt16(*int16)
	BEUint16(*uint16)
	Int32(*int32)
	Uint32(*uint32)
	BEInt32(*int32)
	BEUint32(*uint32)
	Int64(*int64)
	Uint64(*uint64)
	BEInt64(*int64)
	BEUint64(*uint64)
	Float32(*float32)
	Float64(*float64)
	BEFloat32(*float32)
	BEFloat64(*float64)
	Varint32(*int32)
	Varuint32(*uint32)
	Varint64(*int64)
	Varuint64(*uint64)
	SignedVarint32(*int32)
	SignedVarint64(*int64)

	String(*string)
	Bytes(*[]byte)
	NBT(*[]byte)
	UUID(*uuid.UUID)
	Vec2(*mgl32.Vec2)
	Vec3(*mgl32.Vec3)
	RGBA(*color.RGBA)
	Bitset(words []uint64, bits uint64)
}

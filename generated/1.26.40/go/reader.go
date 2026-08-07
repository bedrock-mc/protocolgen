// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"encoding/binary"
	"fmt"
	"image/color"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// Reader decodes one generated protocol value from a byte slice. Decode
// failures are retained by Err instead of panicking.
type Reader struct {
	data           []byte
	pos            int
	err            error
	maxSliceLength uint64
}

// NewReader creates a Reader over data.
func NewReader(data []byte) *Reader {
	return NewReaderWithLimit(data, maxSliceLength)
}

// NewReaderWithLimit creates a Reader with an explicit maximum decoded
// collection length. A zero limit rejects every non-empty collection.
func NewReaderWithLimit(data []byte, max uint64) *Reader {
	return &Reader{data: data, maxSliceLength: max}
}

// Reading reports that this IO implementation decodes values.
func (*Reader) Reading() bool { return true }

// Err returns the first decode error, if any.
func (r *Reader) Err() error { return r.err }

// Remaining reports the number of unread bytes.
func (r *Reader) Remaining() int {
	if r.pos >= len(r.data) {
		return 0
	}
	return len(r.data) - r.pos
}

func (r *Reader) fail(err error) {
	if r.err == nil {
		r.err = err
	}
}

// InvalidValue records a malformed wire value.
func (r *Reader) InvalidValue(value any, context string) {
	r.fail(fmt.Errorf("invalid value %v for %s", value, context))
}

func (r *Reader) readByte() byte {
	if r.err != nil {
		return 0
	}
	if r.pos >= len(r.data) {
		r.fail(fmt.Errorf("unexpected end of input"))
		return 0
	}
	v := r.data[r.pos]
	r.pos++
	return v
}

func (r *Reader) readN(n int) []byte {
	if r.err != nil {
		return nil
	}
	if n < 0 || n > len(r.data)-r.pos {
		r.fail(fmt.Errorf("unexpected end of input reading %d bytes", n))
		return nil
	}
	v := r.data[r.pos : r.pos+n]
	r.pos += n
	return v
}

func (r *Reader) readUvarint() uint64 {
	var value uint64
	for shift := uint(0); shift < 64; shift += 7 {
		b := r.readByte()
		if r.err != nil {
			return 0
		}
		if shift == 63 && b > 1 {
			r.fail(fmt.Errorf("varint overflows uint64"))
			return 0
		}
		value |= uint64(b&0x7f) << shift
		if b&0x80 == 0 {
			return value
		}
	}
	r.fail(fmt.Errorf("varint exceeds ten bytes"))
	return 0
}

func (r *Reader) Bool(x *bool) {
	value := r.readByte()
	if r.err != nil {
		return
	}
	if value > 1 {
		r.InvalidValue(value, "boolean must be 0 or 1")
		return
	}
	*x = value == 1
}
func (r *Reader) Int8(x *int8)   { *x = int8(r.readByte()) }
func (r *Reader) Uint8(x *uint8) { *x = r.readByte() }

func (r *Reader) Int16(x *int16) {
	data := r.readN(2)
	if data != nil {
		*x = int16(binary.LittleEndian.Uint16(data))
	}
}
func (r *Reader) Uint16(x *uint16) {
	data := r.readN(2)
	if data != nil {
		*x = binary.LittleEndian.Uint16(data)
	}
}
func (r *Reader) BEInt16(x *int16) {
	data := r.readN(2)
	if data != nil {
		*x = int16(binary.BigEndian.Uint16(data))
	}
}
func (r *Reader) BEUint16(x *uint16) {
	data := r.readN(2)
	if data != nil {
		*x = binary.BigEndian.Uint16(data)
	}
}

func (r *Reader) Int32(x *int32) {
	data := r.readN(4)
	if data != nil {
		*x = int32(binary.LittleEndian.Uint32(data))
	}
}
func (r *Reader) Uint32(x *uint32) {
	data := r.readN(4)
	if data != nil {
		*x = binary.LittleEndian.Uint32(data)
	}
}
func (r *Reader) BEInt32(x *int32) {
	data := r.readN(4)
	if data != nil {
		*x = int32(binary.BigEndian.Uint32(data))
	}
}
func (r *Reader) BEUint32(x *uint32) {
	data := r.readN(4)
	if data != nil {
		*x = binary.BigEndian.Uint32(data)
	}
}

func (r *Reader) Int64(x *int64) {
	data := r.readN(8)
	if data != nil {
		*x = int64(binary.LittleEndian.Uint64(data))
	}
}
func (r *Reader) Uint64(x *uint64) {
	data := r.readN(8)
	if data != nil {
		*x = binary.LittleEndian.Uint64(data)
	}
}
func (r *Reader) BEInt64(x *int64) {
	data := r.readN(8)
	if data != nil {
		*x = int64(binary.BigEndian.Uint64(data))
	}
}
func (r *Reader) BEUint64(x *uint64) {
	data := r.readN(8)
	if data != nil {
		*x = binary.BigEndian.Uint64(data)
	}
}

func (r *Reader) Float32(x *float32) {
	data := r.readN(4)
	if data != nil {
		*x = math.Float32frombits(binary.LittleEndian.Uint32(data))
	}
}
func (r *Reader) Float64(x *float64) {
	data := r.readN(8)
	if data != nil {
		*x = math.Float64frombits(binary.LittleEndian.Uint64(data))
	}
}
func (r *Reader) BEFloat32(x *float32) {
	data := r.readN(4)
	if data != nil {
		*x = math.Float32frombits(binary.BigEndian.Uint32(data))
	}
}
func (r *Reader) BEFloat64(x *float64) {
	data := r.readN(8)
	if data != nil {
		*x = math.Float64frombits(binary.BigEndian.Uint64(data))
	}
}

func (r *Reader) Varint32(x *int32) {
	value := r.readUvarint()
	if r.err != nil {
		return
	}
	if value > uint64(^uint32(0)) {
		r.InvalidValue(value, "zigzag int32 overflows uint32")
		return
	}
	value32 := uint32(value)
	*x = int32(value32 >> 1)
	if value32&1 != 0 {
		*x = ^*x
	}
}

func (r *Reader) Varuint32(x *uint32) {
	value := r.readUvarint()
	if r.err != nil {
		return
	}
	if value > uint64(^uint32(0)) {
		r.InvalidValue(value, "varuint32 overflows uint32")
		return
	}
	*x = uint32(value)
}

func (r *Reader) Varint64(x *int64) {
	value := r.readUvarint()
	if r.err != nil {
		return
	}
	*x = int64(value >> 1)
	if value&1 != 0 {
		*x = ^*x
	}
}

func (r *Reader) Varuint64(x *uint64) { *x = r.readUvarint() }
func (r *Reader) SignedVarint32(x *int32) {
	value := r.readUvarint()
	if r.err != nil {
		return
	}
	if value > uint64(^uint32(0)) {
		r.InvalidValue(value, "signed varint32 overflows uint32")
		return
	}
	*x = int32(value)
}
func (r *Reader) SignedVarint64(x *int64) { *x = int64(r.readUvarint()) }

func (r *Reader) ActorRuntimeID(x *uint64)          { r.Varuint64(x) }
func (r *Reader) ActorRuntimeIDVarint64(x *int64)   { r.Varint64(x) }
func (r *Reader) ActorRuntimeIDVaruint32(x *uint32) { r.Varuint32(x) }
func (r *Reader) ActorUniqueID(x *int64)            { r.Varint64(x) }
func (r *Reader) ActorUniqueIDInt64(x *int64)       { r.Int64(x) }
func (r *Reader) ActorUniqueIDUint64(x *uint64)     { r.Uint64(x) }
func (r *Reader) ActorUniqueIDVaruint64(x *uint64)  { r.Varuint64(x) }
func (r *Reader) PlayerInputTick(x *uint64)         { r.Varuint64(x) }

func (r *Reader) String(x *string) {
	length, ok := r.readLength("string")
	if !ok {
		return
	}
	*x = string(r.readN(length))
}

func (r *Reader) Bytes(x *[]byte) {
	length, ok := r.readLength("byte slice")
	if !ok {
		return
	}
	data := r.readN(length)
	if data == nil {
		return
	}
	*x = append((*x)[:0], data...)
}

func (r *Reader) readLength(context string) (int, bool) {
	var length uint32
	r.Varuint32(&length)
	if r.err != nil {
		return 0, false
	}
	if uint64(length) > uint64(^uint(0)>>1) {
		r.InvalidValue(length, context+" length overflows int")
		return 0, false
	}
	return int(length), true
}

func (r *Reader) NBT(x *[]byte) {
	if r.err != nil {
		return
	}
	start := r.pos
	if err := scanNBT(r.data, &r.pos, true); err != nil {
		r.fail(err)
		return
	}
	*x = append((*x)[:0], r.data[start:r.pos]...)
}

func (r *Reader) UUID(x *uuid.UUID) {
	data := r.readN(16)
	if data == nil {
		return
	}
	copy((*x)[:8], data[:8])
	copy((*x)[8:], data[8:])
	reverseBytes((*x)[:8])
	reverseBytes((*x)[8:])
}

func (r *Reader) Vec2(x *mgl32.Vec2) {
	r.Float32(&(*x)[0])
	r.Float32(&(*x)[1])
}

func (r *Reader) Vec3(x *mgl32.Vec3) {
	r.Float32(&(*x)[0])
	r.Float32(&(*x)[1])
	r.Float32(&(*x)[2])
}

func (r *Reader) RGBA(x *color.RGBA) {
	var value uint32
	r.Uint32(&value)
	*x = color.RGBA{R: byte(value), G: byte(value >> 8), B: byte(value >> 16), A: byte(value >> 24)}
}

func (r *Reader) Bitset(words []uint64, bits uint64) {
	wordCount := bits / 64
	if bits%64 != 0 {
		wordCount++
	}
	if uint64(len(words)) != wordCount {
		r.InvalidValue(len(words), "bitset word count does not match declared width")
		return
	}
	for i := range words {
		words[i] = 0
	}
	for offset := uint64(0); offset < bits; offset += 7 {
		value := r.readByte()
		if r.err != nil {
			return
		}
		width := bits - offset
		if width > 7 {
			width = 7
		}
		if width < 7 && uint64(value&0x7f) >= (uint64(1)<<width) {
			r.InvalidValue(value, "bitset contains bits outside its declared width")
			return
		}
		for bit := uint64(0); bit < width; bit++ {
			if value&(1<<bit) != 0 {
				index := offset + bit
				words[index/64] |= uint64(1) << (index % 64)
			}
		}
		if value&0x80 == 0 {
			return
		}
		if offset+7 >= bits {
			r.InvalidValue(value, "bitset exceeds its declared width")
			return
		}
	}
}

func (r *Reader) SliceLength(value uint64, max uint64) bool {
	if r.maxSliceLength < max {
		max = r.maxSliceLength
	}
	if value > max {
		r.InvalidValue(value, "collection length exceeds decoder limit")
		return false
	}
	return true
}

func reverseBytes(value []byte) {
	for i, j := 0, len(value)-1; i < j; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]
	}
}

func scanNBT(data []byte, pos *int, named bool) error {
	if *pos >= len(data) {
		return fmt.Errorf("NBT ended before its tag")
	}
	tag := data[*pos]
	*pos += 1
	if tag == 0 {
		return nil
	}
	if named {
		if _, err := nbtReadBytes(data, pos, 2); err != nil {
			return err
		}
		nameLength := int(binary.LittleEndian.Uint16(data[*pos-2 : *pos]))
		if _, err := nbtReadBytes(data, pos, nameLength); err != nil {
			return err
		}
	}
	return scanNBTPayload(data, pos, tag)
}

func scanNBTPayload(data []byte, pos *int, tag byte) error {
	switch tag {
	case 1:
		_, err := nbtReadBytes(data, pos, 1)
		return err
	case 2:
		_, err := nbtReadBytes(data, pos, 2)
		return err
	case 3, 5:
		_, err := nbtReadBytes(data, pos, 4)
		return err
	case 4, 6:
		_, err := nbtReadBytes(data, pos, 8)
		return err
	case 7:
		length, err := nbtReadInt32(data, pos)
		if err != nil || length < 0 {
			return fmt.Errorf("invalid NBT byte-array length %d", length)
		}
		_, err = nbtReadBytes(data, pos, int(length))
		return err
	case 8:
		length, err := nbtReadUint16(data, pos)
		if err != nil {
			return err
		}
		_, err = nbtReadBytes(data, pos, int(length))
		return err
	case 9:
		element, err := nbtReadByte(data, pos)
		if err != nil {
			return err
		}
		length, err := nbtReadInt32(data, pos)
		if err != nil || length < 0 {
			return fmt.Errorf("invalid NBT list length %d", length)
		}
		for i := int32(0); i < length; i++ {
			if err := scanNBTPayload(data, pos, element); err != nil {
				return err
			}
		}
		return nil
	case 10:
		for {
			if *pos >= len(data) {
				return fmt.Errorf("NBT compound has no end tag")
			}
			if data[*pos] == 0 {
				*pos += 1
				return nil
			}
			if err := scanNBT(data, pos, true); err != nil {
				return err
			}
		}
	case 11, 12:
		length, err := nbtReadInt32(data, pos)
		if err != nil || length < 0 {
			return fmt.Errorf("invalid NBT array length %d", length)
		}
		width := 4
		if tag == 12 {
			width = 8
		}
		_, err = nbtReadBytes(data, pos, int(length)*width)
		return err
	default:
		return fmt.Errorf("unknown NBT tag %d", tag)
	}
}

func nbtReadByte(data []byte, pos *int) (byte, error) {
	if *pos >= len(data) {
		return 0, fmt.Errorf("NBT ended unexpectedly")
	}
	value := data[*pos]
	*pos += 1
	return value, nil
}

func nbtReadBytes(data []byte, pos *int, length int) ([]byte, error) {
	if length < 0 || length > len(data)-*pos {
		return nil, fmt.Errorf("NBT ended unexpectedly")
	}
	value := data[*pos : *pos+length]
	*pos += length
	return value, nil
}

func nbtReadUint16(data []byte, pos *int) (uint16, error) {
	value, err := nbtReadBytes(data, pos, 2)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(value), nil
}

func nbtReadInt32(data []byte, pos *int) (int32, error) {
	value, err := nbtReadBytes(data, pos, 4)
	if err != nil {
		return 0, err
	}
	return int32(binary.LittleEndian.Uint32(value)), nil
}

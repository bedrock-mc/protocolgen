package protocol

import (
	"bytes"
	"math"
	"testing"

	"github.com/google/uuid"
)

func TestVarintsRoundTrip(t *testing.T) {
	writer := NewWriter()
	u32 := uint32(math.MaxUint32)
	i32 := int32(math.MinInt32)
	u64 := uint64(math.MaxUint64)
	i64 := int64(math.MinInt64)
	writer.Varuint32(&u32)
	writer.Varint32(&i32)
	writer.Varuint64(&u64)
	writer.Varint64(&i64)
	if err := writer.Err(); err != nil {
		t.Fatal(err)
	}

	reader := NewReader(writer.Data())
	var gotU32 uint32
	var gotI32 int32
	var gotU64 uint64
	var gotI64 int64
	reader.Varuint32(&gotU32)
	reader.Varint32(&gotI32)
	reader.Varuint64(&gotU64)
	reader.Varint64(&gotI64)
	if err := reader.Err(); err != nil {
		t.Fatal(err)
	}
	if gotU32 != u32 || gotI32 != i32 || gotU64 != u64 || gotI64 != i64 {
		t.Fatalf("round trip = %d, %d, %d, %d", gotU32, gotI32, gotU64, gotI64)
	}
}

func TestReaderRejectsSliceOverLimit(t *testing.T) {
	reader := NewReader(nil)
	if reader.SliceLength(maxSliceLength+1, maxSliceLength) {
		t.Fatal("SliceLength accepted a collection over the default limit")
	}
	if reader.Err() == nil {
		t.Fatal("SliceLength did not record an error")
	}
}

func TestUUIDRoundTrip(t *testing.T) {
	want := uuid.MustParse("00112233-4455-6677-8899-aabbccddeeff")
	writer := NewWriter()
	writer.UUID(&want)
	if err := writer.Err(); err != nil {
		t.Fatal(err)
	}

	reader := NewReader(writer.Data())
	var got uuid.UUID
	reader.UUID(&got)
	if err := reader.Err(); err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("UUID = %s, want %s", got, want)
	}
	if !bytes.Equal(writer.Data(), []byte{0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11, 0x00, 0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88}) {
		t.Fatalf("wire UUID = %x", writer.Data())
	}
}

func TestReaderRejectsTruncatedString(t *testing.T) {
	reader := NewReader([]byte{3, 'a'})
	var value string
	reader.String(&value)
	if reader.Err() == nil {
		t.Fatal("truncated string was accepted")
	}
}

package roundtrip

import (
	"bytes"
	"encoding/hex"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	protocol "protocolgen/generated/1.26.40/go/protocol"
	packet "protocolgen/generated/1.26.40/go/protocol/packet"
)

func TestPacketCorpus(t *testing.T) {
	for _, data := range corpus(t) {
		checkPackets(t, data)
	}
}

func FuzzPackets(f *testing.F) {
	for _, data := range corpus(f) {
		f.Add(data)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		checkPackets(t, data)
	})
}

func TestRuntimeVarintsRoundTrip(t *testing.T) {
	for _, value := range []int32{0, 1, -1, 63, -64, 1 << 20, -(1 << 20), -1 << 31, 1<<31 - 1} {
		var encoded int32
		writer := protocol.NewWriter()
		writer.Varint32(&value)
		if err := writer.Err(); err != nil {
			t.Fatalf("write zigzag int32 %d: %v", value, err)
		}
		reader := protocol.NewReader(writer.Data())
		reader.Varint32(&encoded)
		if err := reader.Err(); err != nil || reader.Remaining() != 0 || encoded != value {
			t.Fatalf("zigzag int32 %d -> %x -> %d, err=%v remaining=%d", value, writer.Data(), encoded, err, reader.Remaining())
		}
	}

	for _, value := range []int64{0, 1, -1, 63, -64, 1 << 40, -(1 << 40), -1 << 63, 1<<63 - 1} {
		var encoded int64
		writer := protocol.NewWriter()
		writer.Varint64(&value)
		if err := writer.Err(); err != nil {
			t.Fatalf("write zigzag int64 %d: %v", value, err)
		}
		reader := protocol.NewReader(writer.Data())
		reader.Varint64(&encoded)
		if err := reader.Err(); err != nil || reader.Remaining() != 0 || encoded != value {
			t.Fatalf("zigzag int64 %d -> %x -> %d, err=%v remaining=%d", value, writer.Data(), encoded, err, reader.Remaining())
		}
	}

	for _, value := range []uint64{0, 1, 127, 128, 1 << 32, ^uint64(0)} {
		var encoded uint64
		writer := protocol.NewWriter()
		writer.Varuint64(&value)
		if err := writer.Err(); err != nil {
			t.Fatalf("write varuint64 %d: %v", value, err)
		}
		reader := protocol.NewReader(writer.Data())
		reader.Varuint64(&encoded)
		if err := reader.Err(); err != nil || reader.Remaining() != 0 || encoded != value {
			t.Fatalf("varuint64 %d -> %x -> %d, err=%v remaining=%d", value, writer.Data(), encoded, err, reader.Remaining())
		}
	}
}

func TestRuntimeSliceLimit(t *testing.T) {
	reader := protocol.NewReaderWithLimit(nil, 2)
	if reader.SliceLength(3, 4096) {
		t.Fatal("slice length above the reader limit was accepted")
	}
	if reader.Err() == nil {
		t.Fatal("slice length rejection did not record an error")
	}

	reader = protocol.NewReaderWithLimit(nil, 4)
	if !reader.SliceLength(1, 2) {
		t.Fatal("slice length below the helper limit was rejected")
	}
}

func TestRuntimeStringBounds(t *testing.T) {
	for _, data := range [][]byte{{3, 'a'}, {0xff, 0xff, 0xff, 0xff, 0x0f}} {
		reader := protocol.NewReader(data)
		var value string
		reader.String(&value)
		if reader.Err() == nil {
			t.Fatalf("malformed string %x was accepted", data)
		}
	}

	writer := protocol.NewWriter()
	value := "bounded"
	writer.String(&value)
	if err := writer.Err(); err != nil {
		t.Fatalf("write string: %v", err)
	}
	reader := protocol.NewReader(writer.Data())
	reader.String(&value)
	if err := reader.Err(); err != nil || reader.Remaining() != 0 || value != "bounded" {
		t.Fatalf("string round trip value=%q err=%v remaining=%d", value, err, reader.Remaining())
	}
}

func corpus(t testing.TB) [][]byte {
	t.Helper()
	paths, err := filepath.Glob(filepath.Join("testdata", "corpus", "*.hex"))
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(paths)
	data := make([][]byte, 0, len(paths))
	for _, path := range paths {
		contents, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		value, err := hex.DecodeString(strings.TrimSpace(string(contents)))
		if err != nil {
			t.Fatalf("decode %s: %v", path, err)
		}
		data = append(data, value)
	}
	return data
}

func checkPackets(t *testing.T, data []byte) {
	t.Helper()
	pool := packet.NewPool()
	ids := make([]uint32, 0, len(pool))
	for id := range pool {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	for _, id := range ids {
		factory := pool[id]
		value := factory()
		reader := protocol.NewReader(data)
		if panicValue := callMarshal(value, reader); panicValue != nil {
			t.Fatalf("packet %d decode panicked for %x: %v", id, data, panicValue)
		}
		if err := reader.Err(); err != nil || reader.Remaining() != 0 {
			continue
		}

		writer := protocol.NewWriter()
		if panicValue := callMarshal(value, writer); panicValue != nil {
			t.Fatalf("packet %d encode panicked after decoding %x: %v", id, data, panicValue)
		}
		if err := writer.Err(); err != nil {
			t.Fatalf("packet %d encode failed after decoding %x: %v", id, data, err)
		}
		if !bytes.Equal(writer.Data(), data) {
			t.Fatalf("packet %d changed bytes: input=%x output=%x", id, data, writer.Data())
		}
	}
}

func callMarshal(value packet.Packet, io protocol.IO) (panicValue any) {
	defer func() {
		panicValue = recover()
	}()
	value.Marshal(io)
	return nil
}

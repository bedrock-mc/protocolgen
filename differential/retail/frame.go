package retail

import (
	"errors"
	"fmt"
)

const batchHeader = 0xfe

// Entry is one batch entry: its header varint and the body after it.
type Entry struct {
	Header  uint32
	Payload []byte // header varint and body, exactly as framed
	Body    []byte
}

// NewEntry frames body under a canonical header varint.
func NewEntry(header uint32, body []byte) Entry {
	payload := appendVaruint32(nil, header)
	headerLen := len(payload)
	payload = append(payload, body...)
	return Entry{Header: header, Payload: payload, Body: payload[headerLen:]}
}

// SplitBatch splits an uncompressed batch into entries. On a framing error it
// returns the entries before the malformed one.
func SplitBatch(batch []byte) ([]Entry, error) {
	if len(batch) == 0 || batch[0] != batchHeader {
		return nil, errors.New("batch does not start with 0xfe")
	}
	rest := batch[1:]
	var entries []Entry
	for len(rest) != 0 {
		length, n, err := readVaruint32(rest)
		if err != nil {
			return entries, fmt.Errorf("entry %d length: %w", len(entries), err)
		}
		rest = rest[n:]
		if uint64(length) > uint64(len(rest)) {
			return entries, fmt.Errorf("entry %d declares %d bytes, %d remain", len(entries), length, len(rest))
		}
		payload := rest[:length:length]
		rest = rest[length:]
		header, n, err := readVaruint32(payload)
		if err != nil {
			return entries, fmt.Errorf("entry %d header: %w", len(entries), err)
		}
		entries = append(entries, Entry{Header: header, Payload: payload, Body: payload[n:]})
	}
	return entries, nil
}

// JoinBatch frames entries into an uncompressed batch.
func JoinBatch(entries []Entry) []byte {
	batch := []byte{batchHeader}
	for _, entry := range entries {
		batch = appendVaruint32(batch, uint32(len(entry.Payload)))
		batch = append(batch, entry.Payload...)
	}
	return batch
}

func readVaruint32(data []byte) (uint32, int, error) {
	var value uint64
	for i := 0; i < 5; i++ {
		if i == len(data) {
			return 0, 0, errors.New("varuint32 is truncated")
		}
		value |= uint64(data[i]&0x7f) << (7 * i)
		if data[i]&0x80 == 0 {
			if value > 0xffffffff {
				return 0, 0, errors.New("varuint32 overflows 32 bits")
			}
			return uint32(value), i + 1, nil
		}
	}
	return 0, 0, errors.New("varuint32 does not end within 5 bytes")
}

func appendVaruint32(dst []byte, value uint32) []byte {
	for value >= 0x80 {
		dst = append(dst, byte(value)|0x80)
		value >>= 7
	}
	return append(dst, byte(value))
}

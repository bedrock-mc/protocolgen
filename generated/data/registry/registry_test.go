package registry

import (
	"bytes"
	"testing"
)

// TestPayloadsReturnPrivateCopies prevents a consumer from corrupting another caller's data.
func TestPayloadsReturnPrivateCopies(t *testing.T) {
	for name, read := range map[string]func() []byte{
		"blocks": BlockStatesNBT, "data-driven blocks": DataDrivenBlocksNBT, "items": ItemsNBT,
	} {
		t.Run(name, func(t *testing.T) {
			first, original := read(), read()
			if len(first) == 0 {
				t.Fatal("payload is empty")
			}
			first[0] ^= 0xff
			if !bytes.Equal(read(), original) {
				t.Fatal("caller changed the shared payload")
			}
		})
	}
}

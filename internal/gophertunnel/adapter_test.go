package gophertunnel

import (
	"path/filepath"
	"testing"
)

func TestExtractPreservesASTWireOrderAsNarrowAdapter(t *testing.T) {
	packets, err := Extract(filepath.Join("..", "driftcheck", "testdata", "packets"))
	if err != nil {
		t.Fatalf("Extract: %v", err)
	}
	if len(packets) != 1 || packets[0].TypeName != "Alpha" || len(packets[0].Ops) != 2 {
		t.Fatalf("packets = %#v, want Alpha with two operations", packets)
	}
	if packets[0].Ops[0].Method != "Uint8" || packets[0].Ops[1].Method != "Varuint32" {
		t.Fatalf("wire order = %#v", packets[0].Ops)
	}
}

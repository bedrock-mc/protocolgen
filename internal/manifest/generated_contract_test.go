package manifest

import (
	"path/filepath"
	"testing"
)

func TestCheckedInResourcePackResponseUsesVarUint32Selector(t *testing.T) {
	value, err := Load(filepath.Join("..", "..", "generated", "1.26.40", "manifest.json"))
	if err != nil {
		t.Fatalf("Load checked-in manifest: %v", err)
	}
	for _, packet := range value.Packets {
		if packet.ID != 8 {
			continue
		}
		if packet.Name != "ResourcePackClientResponsePacket" || len(packet.Fields) != 1 {
			t.Fatalf("packet 8 shape = %s with %d fields", packet.Name, len(packet.Fields))
		}
		control := packet.Fields[0].Encode.Control
		if control == nil || control.Primitive == nil {
			t.Fatalf("packet 8 selector is not primitive: %#v", control)
		}
		if got := control.Primitive.Code; got != "var_u32" {
			t.Fatalf("packet 8 selector primitive = %q, want var_u32", got)
		}
		return
	}
	t.Fatal("packet 8 is missing from the checked-in manifest")
}

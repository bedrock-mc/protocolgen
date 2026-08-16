package manifest

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestCheckedIn12644IsSameIDHotfixWithOnlyNestedScoreRemovalOptional(t *testing.T) {
	base, err := Load(filepath.Join("..", "..", "generated", "1.26.40", "manifest.json"))
	if err != nil {
		t.Fatalf("Load 1.26.40 manifest: %v", err)
	}
	hotfix, err := Load(filepath.Join("..", "..", "generated", "1.26.44", "manifest.json"))
	if err != nil {
		t.Fatalf("Load 1.26.44 manifest: %v", err)
	}
	if hotfix.Target.MinecraftVersion != "1.26.44" || hotfix.Target.ProtocolVersion != 2168 {
		t.Fatalf("1.26.44 target = %#v, want Minecraft 1.26.44 / protocol 2168", hotfix.Target)
	}

	base.Target = hotfix.Target
	base.Sources = hotfix.Sources
	base.Overrides = hotfix.Overrides
	base.Derivation = hotfix.Derivation
	for packetIndex := range base.Packets {
		if base.Packets[packetIndex].ID != 108 {
			continue
		}
		remove := &base.Packets[packetIndex].Fields[0].Encode.Element.Variants[0].Encode.Fields[2].Encode
		if remove.Kind != KindOptional {
			t.Fatalf("1.26.40 remove objective = %s, want optional", remove.Kind)
		}
		*remove = Optional(*remove)
	}
	if !reflect.DeepEqual(base, hotfix) {
		t.Fatal("1.26.44 manifest has wire changes beyond the nested RemoveScore objective optional")
	}
}

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

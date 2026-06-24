package main

import "testing"

func TestGoFieldName(t *testing.T) {
	cases := map[string]string{
		"Action":                     "Action",
		"Swing Source":               "SwingSource",
		"Target Actor Runtime ID":    "TargetActorRuntimeID",
		"Sender's XUID":              "SendersXUID",
		"Is Trial":                   "IsTrial",
		"BlockNetworkIds Are Hashes": "BlockNetworkIdsAreHashes",
		"Container Id":               "ContainerId",
	}
	for in, want := range cases {
		if got := goFieldName(in); got != want {
			t.Errorf("goFieldName(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestGoTypeName(t *testing.T) {
	cases := map[string]string{
		"AnimatePacket":   "Animate",
		"StartGamePacket": "StartGame",
		"TextPacket":      "Text",
		"NoSuffix":        "NoSuffix",
	}
	for in, want := range cases {
		if got := goTypeName(in); got != want {
			t.Errorf("goTypeName(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestCleanCompositeTitle(t *testing.T) {
	cases := map[string]string{
		"mce::UUID":                    "UUID",
		"Json::Value":                  "Value",
		"MapItemTrackedActor UniqueId": "MapItemTrackedActorUniqueId",
		"(anonymous namespace)::MapInfoRequestPacketAnon::ClientPixelsProxy": "ClientPixelsProxy",
		"FeatureRegistry::FeatureBinaryJsonFormat":                           "FeatureBinaryJsonFormat",
		"server_config":              "ServerConfig",
		"NetworkItemStackDescriptor": "NetworkItemStackDescriptor",
	}
	for in, want := range cases {
		if got := cleanCompositeTitle(in); got != want {
			t.Errorf("cleanCompositeTitle(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestFileName(t *testing.T) {
	cases := map[string]string{
		"Animate":                "animate.go",
		"StartGame":              "start_game.go",
		"ClientboundMapItemData": "clientbound_map_item_data.go",
		"SubChunk":               "sub_chunk.go",
	}
	for in, want := range cases {
		if got := fileName(in); got != want {
			t.Errorf("fileName(%q) = %q, want %q", in, got, want)
		}
	}
}

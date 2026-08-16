package vanilladata

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func testSource() SourceConfig {
	return SourceConfig{
		SchemaVersion:    1,
		MinecraftVersion: "1.26.44",
		ProtocolVersion:  2168,
		BDS: BDSSource{Version: "1.26.44.3", Linux: LinuxSource{
			URL: "https://example.invalid/bedrock-server-1.26.44.3.zip", ArchiveSHA256: strings.Repeat("a", 64),
		}},
		Gophertunnel: GophertunnelSource{
			Repository: "https://github.com/HashimTheArab/gophertunnel", ModulePath: "github.com/HashimTheArab/gophertunnel", Revision: strings.Repeat("b", 40), ModuleVersion: "v0.0.0-20260815100934-bbbbbbbbbbbb",
		},
		Endstone: &EndstoneSource{
			Repository:       "https://github.com/EndstoneMC/endstone",
			Revision:         strings.Repeat("d", 40),
			BDSVersion:       "1.26.44.3",
			HeadlessPatchSHA: "sha256:" + strings.Repeat("e", 64),
		},
		ServerProperties: map[string]string{"online-mode": "false", "server-port": "19132", "level-name": "protocolgen-vanilla-data", "level-seed": "1"},
	}
}

func TestValidateContractRequiresExactManifestSourceAndCodec(t *testing.T) {
	source := testSource()
	m := Manifest{
		Target:  Target{MinecraftVersion: "1.26.44", ProtocolVersion: 2168},
		Sources: []SourcePin{{Kind: "gophertunnel-exact-codec", Revision: source.Gophertunnel.Revision}},
		Packets: []Packet{{ID: 162, Name: "ItemRegistryPacket", Direction: "clientbound"}},
	}
	specs := []PacketSpec{{ID: 162, Name: "ItemRegistryPacket", File: "item_registry.dat"}}
	if err := ValidateContract(m, source, "1.26.44", 2168, specs); err != nil {
		t.Fatalf("ValidateContract exact match: %v", err)
	}
	source.ProtocolVersion++
	if err := ValidateContract(m, source, "1.26.44", 2168, specs); err == nil || !strings.Contains(err.Error(), "vanilla source") {
		t.Fatalf("ValidateContract source mismatch = %v", err)
	}
	source = testSource()
	m.Packets[0].Direction = "serverbound"
	if err := ValidateContract(m, source, "1.26.44", 2168, specs); err == nil || !strings.Contains(err.Error(), "clientbound") {
		t.Fatalf("ValidateContract direction mismatch = %v", err)
	}
}

func TestRecorderCopiesFirstPayloadAndDoesNotRequireOptionalPackets(t *testing.T) {
	specs := []PacketSpec{
		{ID: 119, Name: "AvailableActorIdentifiersPacket", File: "available_actor_identifiers.dat"},
		{ID: 162, Name: "ItemRegistryPacket", File: "item_registry.dat"},
		{ID: 180, Name: "DimensionDataPacket", File: "dimension_data.dat", Optional: true},
	}
	recorder := NewRecorder(specs)
	payload := []byte{1, 2, 3}
	recorder.Observe(119, payload)
	payload[0] = 9
	recorder.Observe(119, []byte{4})
	if got, want := recorder.Missing(), []string{"ItemRegistryPacket"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Recorder.Missing = %v, want %v", got, want)
	}
	recorder.Observe(162, []byte{7, 8})
	if !recorder.Complete() {
		t.Fatalf("Recorder.Complete = false; missing %v", recorder.Missing())
	}
	if got := recorder.Payloads()["AvailableActorIdentifiersPacket"]; !reflect.DeepEqual(got, []byte{1, 2, 3}) {
		t.Fatalf("recorded payload = %v, want copied first payload", got)
	}
}

func TestWriteArtifactsRecordsOptionalAbsenceAndReplacesStaleOutput(t *testing.T) {
	out := filepath.Join(t.TempDir(), "vanilla-data")
	if err := os.MkdirAll(out, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(out, "stale.dat"), []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}
	source := testSource()
	input := ArtifactInput{
		Target: Target{MinecraftVersion: "1.26.44", ProtocolVersion: 2168, Build: "hotfix4"},
		BDS: BDSProvenance{
			Version: source.BDS.Version, ArchiveURL: source.BDS.Linux.URL, ArchiveSHA256: source.BDS.Linux.ArchiveSHA256,
			BinarySHA256: strings.Repeat("c", 64), ServerProperties: source.ServerProperties,
		},
		Gophertunnel: GophertunnelProvenance(source.Gophertunnel),
		Specs: []PacketSpec{
			{ID: 162, Name: "ItemRegistryPacket", File: "item_registry.dat"},
			{ID: 180, Name: "DimensionDataPacket", File: "dimension_data.dat", Optional: true},
		},
		Payloads: map[string][]byte{"ItemRegistryPacket": {3, 1, 4}},
		DerivedFiles: map[string][]byte{
			"required_item_list.json": []byte("{}\n"),
		},
		Warning: "optional settling stopped early",
	}
	if err := WriteArtifacts(out, input); err != nil {
		t.Fatalf("WriteArtifacts: %v", err)
	}
	if _, err := os.Stat(filepath.Join(out, "stale.dat")); !os.IsNotExist(err) {
		t.Fatalf("stale file survived atomic replacement: %v", err)
	}
	if got, err := os.ReadFile(filepath.Join(out, "item_registry.dat")); err != nil || !reflect.DeepEqual(got, []byte{3, 1, 4}) {
		t.Fatalf("item_registry.dat = %v, %v", got, err)
	}
	metadataData, err := os.ReadFile(filepath.Join(out, "capture.json"))
	if err != nil {
		t.Fatal(err)
	}
	var metadata CaptureMetadata
	if err := json.Unmarshal(metadataData, &metadata); err != nil {
		t.Fatalf("decode capture.json: %v", err)
	}
	if metadata.Packets[1].Status != "absent" || !metadata.Packets[1].Optional || metadata.Packets[1].File != "" {
		t.Fatalf("optional packet status = %#v", metadata.Packets[1])
	}
	if metadata.Warning != input.Warning {
		t.Fatalf("capture warning = %q, want %q", metadata.Warning, input.Warning)
	}
	if got, want := metadata.Files[0].SHA256, "sha256:9ceed0d818acc42d0318974c548e2251f888f8512bc5c4c70378aa969a883cad"; got != want {
		t.Fatalf("packet SHA-256 = %q, want %q", got, want)
	}
	if got, err := os.ReadFile(filepath.Join(out, "required_item_list.json")); err != nil || string(got) != "{}\n" {
		t.Fatalf("canonical compatibility file = %q, %v", got, err)
	}
	if got := metadata.Files[1]; got.Kind != "derived" || got.File != "required_item_list.json" {
		t.Fatalf("derived capture metadata = %#v", got)
	}
}

func TestWriteArtifactsRejectsMissingRequiredPacketAndUnsafeName(t *testing.T) {
	source := testSource()
	base := ArtifactInput{
		Target:       Target{MinecraftVersion: "1.26.44", ProtocolVersion: 2168},
		BDS:          BDSProvenance{Version: source.BDS.Version, ArchiveURL: source.BDS.Linux.URL, ArchiveSHA256: source.BDS.Linux.ArchiveSHA256, BinarySHA256: strings.Repeat("c", 64)},
		Gophertunnel: GophertunnelProvenance(source.Gophertunnel),
		Specs:        []PacketSpec{{ID: 1, Name: "Required", File: "required.dat"}},
		Payloads:     map[string][]byte{},
	}
	if err := WriteArtifacts(filepath.Join(t.TempDir(), "out"), base); err == nil || !strings.Contains(err.Error(), "required packet") {
		t.Fatalf("missing packet error = %v", err)
	}
	base.Specs[0].Optional = true
	base.Specs[0].File = ".."
	if err := WriteArtifacts(filepath.Join(t.TempDir(), "out"), base); err == nil || !strings.Contains(err.Error(), "invalid capture file") {
		t.Fatalf("unsafe name error = %v", err)
	}
}

func TestWriteArtifactsRecordsInternalDataProvenanceAndKind(t *testing.T) {
	source := testSource()
	palette := endstonePalette(t)
	canonical, err := canonicalBlockStates(palette)
	if err != nil {
		t.Fatal(err)
	}
	paletteDigest := sha256.Sum256(palette)
	internal := InternalDataManifest{
		SchemaVersion: 1,
		Target:        Target{MinecraftVersion: "1.26.44", ProtocolVersion: 2168},
		BDSVersion:    source.BDS.Version,
		Endstone:      *source.Endstone,
		Files:         []InternalDataFile{{File: "block_palette.nbt", Bytes: len(palette), SHA256: "sha256:" + hex.EncodeToString(paletteDigest[:])}},
	}
	out := filepath.Join(t.TempDir(), "out")
	if err := WriteArtifacts(out, ArtifactInput{
		Target: Target{MinecraftVersion: "1.26.44", ProtocolVersion: 2168},
		BDS: BDSProvenance{
			Version: source.BDS.Version, ArchiveURL: source.BDS.Linux.URL, ArchiveSHA256: source.BDS.Linux.ArchiveSHA256, BinarySHA256: strings.Repeat("c", 64),
		},
		Gophertunnel: GophertunnelProvenance(source.Gophertunnel),
		Specs:        []PacketSpec{{ID: 1, Name: "Required", File: "required.dat"}},
		Payloads:     map[string][]byte{"Required": {1}},
		InternalData: &internal,
		InternalFiles: map[string][]byte{
			"block_palette.nbt":          palette,
			"canonical_block_states.nbt": canonical,
		},
	}); err != nil {
		t.Fatalf("WriteArtifacts: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(out, "capture.json"))
	if err != nil {
		t.Fatal(err)
	}
	var metadata CaptureMetadata
	if err := json.Unmarshal(data, &metadata); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(metadata.InternalData, &internal) {
		t.Fatalf("internal provenance = %#v, want %#v", metadata.InternalData, &internal)
	}
	seen := map[string]string{}
	for _, file := range metadata.Files {
		seen[file.File] = file.Kind
	}
	if seen["block_palette.nbt"] != "internal_data" || seen["canonical_block_states.nbt"] != "internal_data" {
		t.Fatalf("internal file kinds = %#v", seen)
	}
}

func TestLoadSourceConfigValidatesPins(t *testing.T) {
	source := testSource()
	data, err := json.Marshal(source)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "source.json")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}
	if got, err := LoadSourceConfig(path); err != nil || !reflect.DeepEqual(got, source) {
		t.Fatalf("LoadSourceConfig = %#v, %v", got, err)
	}
	source.ServerProperties["online-mode"] = "true"
	data, err = json.Marshal(source)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadSourceConfig(path); err == nil || !strings.Contains(err.Error(), "online-mode=false") {
		t.Fatalf("LoadSourceConfig online mode error = %v", err)
	}
	source = testSource()
	source.Endstone.BDSVersion = "1.26.40.0"
	data, err = json.Marshal(source)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadSourceConfig(path); err == nil || !strings.Contains(err.Error(), "Endstone provenance") {
		t.Fatalf("LoadSourceConfig Endstone version error = %v", err)
	}
}

func TestLoadSourceConfigAcceptsReleaseVersionWithPinnedArchiveBuild(t *testing.T) {
	source := testSource()
	source.MinecraftVersion = "1.26.40"
	source.BDS = BDSSource{
		Version:        "1.26.40",
		ArchiveVersion: "1.26.40.8",
		Linux: LinuxSource{
			URL:           "https://www.minecraft.net/bedrockdedicatedserver/bin-linux/bedrock-server-1.26.40.8.zip",
			ArchiveSHA256: strings.Repeat("a", 64),
		},
	}
	source.Endstone.BDSVersion = "1.26.40"
	data, err := json.Marshal(source)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "source.json")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadSourceConfig(path); err != nil {
		t.Fatalf("LoadSourceConfig release/archive build: %v", err)
	}
}

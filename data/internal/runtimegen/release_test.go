package runtimegen

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"go/parser"
	"go/token"
	"testing"

	"github.com/bedrock-mc/protocolgen/data/source"
)

func TestReleaseIdentifiesCompleteCanonicalLock(t *testing.T) {
	lock := source.Lock{
		SchemaVersion: 1, MinecraftVersion: "1.26.50", ProtocolVersion: 2193,
		Inputs: map[string]source.Input{
			"second": {Revision: "two", SHA256: "bb"},
			"first":  {Revision: "one", SHA256: "aa"},
		},
		Semantic: source.SemanticInputs{BDSVersion: "1.26.32.2"},
	}
	files, err := Release(lock)
	if err != nil {
		t.Fatal(err)
	}
	var manifest struct {
		MinecraftVersion string      `json:"minecraft_version"`
		ProtocolVersion  int         `json:"protocol_version"`
		SourceLockSHA256 string      `json:"source_lock_sha256"`
		SourceLock       source.Lock `json:"source_lock"`
	}
	if err := json.Unmarshal(files["release.json"], &manifest); err != nil {
		t.Fatal(err)
	}
	canonical, err := json.Marshal(lock)
	if err != nil {
		t.Fatal(err)
	}
	wantHash := fmt.Sprintf("%x", sha256.Sum256(canonical))
	if manifest.SourceLockSHA256 != wantHash || manifest.MinecraftVersion != lock.MinecraftVersion || manifest.ProtocolVersion != lock.ProtocolVersion {
		t.Fatalf("wrong release identity: %+v", manifest)
	}
	recorded, err := json.Marshal(manifest.SourceLock)
	if err != nil || !bytes.Equal(canonical, recorded) {
		t.Fatalf("release lost source pins: %s, %v", recorded, err)
	}
	if _, err := parser.ParseFile(token.NewFileSet(), "release_generated.go", files["release_generated.go"], 0); err != nil {
		t.Fatal(err)
	}
	for _, declaration := range []string{
		`const MinecraftVersion = "1.26.50"`,
		`const ProtocolVersion = 2193`,
		fmt.Sprintf("const SourceLockSHA256 = %q", wantHash),
	} {
		if !bytes.Contains(files["release_generated.go"], []byte(declaration)) {
			t.Errorf("missing declaration %s", declaration)
		}
	}
	// JSON key ordering, not insertion ordering, defines the lock identity.
	lock.Inputs = map[string]source.Input{
		"first": {Revision: "one", SHA256: "aa"}, "second": {Revision: "two", SHA256: "bb"},
	}
	again, err := Release(lock)
	if err != nil {
		t.Fatal(err)
	}
	for name, content := range files {
		if !bytes.Equal(content, again[name]) {
			t.Errorf("%s changed with map insertion order", name)
		}
	}
	lock.Inputs["first"] = source.Input{Revision: "one", SHA256: "changed"}
	changed, err := Release(lock)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(files["release_generated.go"], changed["release_generated.go"]) {
		t.Fatal("changed source digest did not change release identity")
	}
}

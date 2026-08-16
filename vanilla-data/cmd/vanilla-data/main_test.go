package main

import (
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunRequiresVersionedCaptureInputs(t *testing.T) {
	err := run(nil)
	if err == nil {
		t.Fatal("run(nil) succeeded, want missing input error")
	}
	for _, required := range []string{"-manifest", "-source", "-out", "-bds-binary"} {
		if !strings.Contains(err.Error(), required) {
			t.Fatalf("run(nil) error = %q, want %s", err, required)
		}
	}
}

func TestRunValidateOnlyChecksGeneratedProtocolContract(t *testing.T) {
	root := filepath.Join("..", "..", "..")
	if err := run([]string{
		"-manifest", filepath.Join(root, "generated", "1.26.44", "manifest.json"),
		"-source", filepath.Join(root, "generated", "1.26.44", "vanilla-source.json"),
		"-validate-only",
	}); err != nil {
		t.Fatalf("run validate-only: %v", err)
	}
}

func TestRunRejectsInternalDataWithoutVersionLockedEndstoneSource(t *testing.T) {
	root := filepath.Join("..", "..", "..")
	err := run([]string{
		"-manifest", filepath.Join(root, "generated", "1.26.44", "manifest.json"),
		"-source", filepath.Join(root, "generated", "1.26.44", "vanilla-source.json"),
		"-internal-data", t.TempDir(),
		"-validate-only",
	})
	if err == nil || !strings.Contains(err.Error(), "Endstone source lock") {
		t.Fatalf("run internal data without source lock = %v", err)
	}
}

func TestVerifyServerProperties(t *testing.T) {
	path := filepath.Join(t.TempDir(), "server.properties")
	if err := os.WriteFile(path, []byte("# BDS\nonline-mode=false\nallow-list=false\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := verifyServerProperties(path, map[string]string{"online-mode": "false", "allow-list": "false"}); err != nil {
		t.Fatalf("verifyServerProperties exact match: %v", err)
	}
	if err := verifyServerProperties(path, map[string]string{"online-mode": "true"}); err == nil || !strings.Contains(err.Error(), "want") {
		t.Fatalf("verifyServerProperties mismatch = %v", err)
	}
}

func TestRequireLoopback(t *testing.T) {
	for _, address := range []string{"127.0.0.1:19132", "[::1]:19132", "localhost:19132"} {
		if err := requireLoopback(address); err != nil {
			t.Fatalf("requireLoopback(%q): %v", address, err)
		}
	}
	if err := requireLoopback("example.com:19132"); err == nil {
		t.Fatal("requireLoopback accepted remote host")
	}
}

func TestValidateAddressPort(t *testing.T) {
	if err := validateAddressPort("127.0.0.1:19132", "19132"); err != nil {
		t.Fatalf("validateAddressPort exact match: %v", err)
	}
	if err := validateAddressPort("127.0.0.1:19133", "19132"); err == nil {
		t.Fatal("validateAddressPort accepted mismatched port")
	}
}

func TestPacketFromServer(t *testing.T) {
	expected := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 19132}
	if !packetFromServer(&net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 19132}, expected) {
		t.Fatal("packetFromServer rejected server endpoint")
	}
	if packetFromServer(&net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 50000}, expected) {
		t.Fatal("packetFromServer accepted client endpoint")
	}
}

func TestModuleVersionMatchesRevision(t *testing.T) {
	revision := "4794743244247a9395360bced1f294d4882c4df0"
	if !moduleVersionMatchesRevision("v1.25.3-0.20260815100934-479474324424", revision) {
		t.Fatal("pseudo-version did not match its full revision")
	}
	if moduleVersionMatchesRevision("v1.58.1", revision) {
		t.Fatal("unrelated tag matched revision")
	}
}

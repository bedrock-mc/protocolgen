package gophertunneloracle

import (
	"testing"

	"protocolgen/internal/manifest"
)

// TestReviewedDivergenceDoesNotAcceptNewShapes ensures a packet ID alone cannot hide a regression.
func TestReviewedDivergenceDoesNotAcceptNewShapes(t *testing.T) {
	m := fixtureManifest(manifest.Primitive("u8"))
	source := extraction{Packets: []sourcePacket{{ID: 1, Operations: []sourceOperation{{Kind: "primitive", Code: "u16le", Site: "/first/checkout/packet.go:1"}}}}}
	lock := fixtureLock()
	report := Compare(m, source, lock, emptyAccepted(), "manifest.json")
	baseline := emptyAccepted()
	baseline.Divergences = []AcceptedDivergence{{ID: 1, Fingerprint: report.Packets[0].Fingerprint}}
	if len(Compare(m, source, lock, baseline, "manifest.json").Accepted) != 1 {
		t.Fatal("exact reviewed divergence was not accepted")
	}
	source.Packets[0].Operations[0].Site = "/different/checkout/packet.go:1"
	if len(Compare(m, source, lock, baseline, "manifest.json").Accepted) != 1 {
		t.Fatal("checkout location changed the comparison fingerprint")
	}
	source.Packets[0].Operations[0].Code = "u32le"
	changed := Compare(m, source, lock, baseline, "manifest.json")
	if len(changed.Unaccepted) != 1 || len(changed.Accepted) != 0 || len(changed.ResolvedAccepted) != 0 {
		t.Fatalf("changed oracle shape escaped review: %+v", changed)
	}
	source.Packets[0].Operations[0].Code = "u16le"
	m.Packets[0].Fields[0].Encode = manifest.Primitive("i32le")
	if len(Compare(m, source, lock, baseline, "manifest.json").Unaccepted) != 1 {
		t.Fatal("changed manifest shape escaped review")
	}
	m.Packets[0].Fields[0].Encode = manifest.Primitive("u16le")
	if len(Compare(m, source, lock, baseline, "manifest.json").ResolvedAccepted) != 1 {
		t.Fatal("resolved divergence did not require baseline cleanup")
	}
}

// TestDivergenceFingerprintPreservesLargeIntegers prevents JSON number rounding from masking a changed input.
func TestDivergenceFingerprintPreservesLargeIntegers(t *testing.T) {
	m := fixtureManifest(manifest.FixedArray(1<<54, manifest.Primitive("u8")))
	first, err := divergenceFingerprint(m.Packets[0], sourcePacket{}, nil, fixtureLock())
	if err != nil {
		t.Fatal(err)
	}
	m.Packets[0].Fields[0].Encode.Length++
	second, err := divergenceFingerprint(m.Packets[0], sourcePacket{}, nil, fixtureLock())
	if err != nil {
		t.Fatal(err)
	}
	if first == second {
		t.Fatal("distinct 64-bit lengths have identical fingerprints")
	}
}

// TestDivergenceFingerprintIncludesComparedPaths requires review when improved analysis exposes a new difference.
func TestDivergenceFingerprintIncludesComparedPaths(t *testing.T) {
	packet := fixtureManifest(manifest.Primitive("u8")).Packets[0]
	paths := []PathResult{{Classification: "DIVERGENCE", ManifestSequence: []string{"u8"}, GophertunnelSequence: []string{"u16le"}, GophertunnelSite: "/first/source.go:1"}}
	first, err := divergenceFingerprint(packet, sourcePacket{}, paths, fixtureLock())
	if err != nil {
		t.Fatal(err)
	}
	paths[0].GophertunnelSite = "/second/source.go:1"
	relocated, err := divergenceFingerprint(packet, sourcePacket{}, paths, fixtureLock())
	if err != nil {
		t.Fatal(err)
	}
	if first != relocated {
		t.Fatal("diagnostic path changed comparison identity")
	}
	paths[0].GophertunnelSequence = []string{"u32le"}
	changed, err := divergenceFingerprint(packet, sourcePacket{}, paths, fixtureLock())
	if err != nil {
		t.Fatal(err)
	}
	if first == changed {
		t.Fatal("new comparison result reused the accepted fingerprint")
	}
}

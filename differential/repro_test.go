package differential

import (
	"encoding/hex"
	"testing"

	gophertunnelpacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	generatedpacket "protocolgen/generated/1.26.40/go/protocol/packet"
)

// TestReproCase diagnoses a single divergence input against both codecs.
func TestReproCase(t *testing.T) {
	cases := []struct {
		id   uint32
		hexs string
	}{
		{122, "0130"},
	}
	for _, c := range cases {
		data, _ := hex.DecodeString(c.hexs)
		gen := decodeGenerated(generatedpacket.NewServerPool()[c.id], data)
		ora := decodeOracle(gophertunnelpacket.NewServerPool()[c.id], data, false)
		t.Logf("packet %d input %s\n  generated: ok=%v full=%v err=%v enc=%x\n  oracle:    ok=%v full=%v err=%v enc=%x",
			c.id, c.hexs, gen.ok, gen.full, gen.err, gen.encoded, ora.ok, ora.full, ora.err, ora.encoded)
	}
}

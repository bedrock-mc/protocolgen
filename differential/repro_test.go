package differential

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	gophertunnelpacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// TestReproCases replays recorded divergence inputs through both codecs so a
// reviewer can inspect exact byte behavior per case.
func TestReproCases(t *testing.T) {
	path := os.Getenv("REPRO_CASES")
	if path == "" {
		t.Skip("set REPRO_CASES to a JSON case file")
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var cases []struct {
		Packet    uint32 `json:"packet"`
		Direction string `json:"direction"`
		InputHex  string `json:"input_hex"`
	}
	if err := json.Unmarshal(raw, &cases); err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		data, err := hex.DecodeString(c.InputHex)
		if err != nil {
			t.Fatal(err)
		}
		dir := direction(c.Direction)
		genFactory, ok1 := generatedFactory(dir, c.Packet)
		oraFactory, ok2 := oracleFactory(dir, c.Packet)
		if !ok1 || !ok2 {
			t.Logf("packet %d %s: missing from a pool (gen=%v oracle=%v)", c.Packet, c.Direction, ok1, ok2)
			continue
		}
		gen := decodeGenerated(genFactory, data)
		ora := decodeOracle(oraFactory, data, dir == directionServer)
		var oraName string
		if pk := oraFactory(); pk != nil {
			oraName = pkName(pk)
		}
		t.Logf("packet %d (%s) %s input=%s\n  generated: ok=%v err=%v enc=%x\n  oracle:    ok=%v err=%v enc=%x",
			c.Packet, oraName, c.Direction, c.InputHex, gen.ok, gen.err, gen.encoded, ora.ok, ora.err, ora.encoded)
	}
}

func pkName(pk gophertunnelpacket.Packet) string {
	return fmt.Sprintf("%T", pk)
}

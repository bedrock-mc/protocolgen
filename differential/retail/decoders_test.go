package retail

import (
	"testing"

	genpacket "protocolgen/generated/1.26.51/go/protocol/packet"
)

func generatedInput(t *testing.T, caseName, sender string, header uint32, pk genpacket.Packet) Input {
	t.Helper()
	body, err := genpacket.Encode(pk)
	if err != nil {
		t.Fatal(err)
	}
	return Input{Fixture: "f.bin", Case: caseName, Sender: sender, Batch: JoinBatch([]Entry{NewEntry(header, body)})}
}

func TestGeneratedDecoderRoundTripsItsOwnEncoding(t *testing.T) {
	input := generatedInput(t, CaseRetail, SenderServer, genpacket.IDSetTime|1<<10, &genpacket.SetTime{Time: -5})
	verdicts := DecodeGenerated(input)
	if len(verdicts) != 2 || verdicts[1].Case != ProbeElementLimit {
		t.Fatalf("verdicts %+v", verdicts)
	}
	v := verdicts[0]
	if v.Outcome != OutcomeOK || v.Encoded != v.Payload || v.RoundTrip != RoundTripSame || v.Header != "10:1:0" {
		t.Fatalf("verdict %+v", v)
	}
}

func TestGeneratedDecoderRejectsTheWrongSender(t *testing.T) {
	v := DecodeGenerated(generatedInput(t, CaseWrongSender, SenderClient, genpacket.IDSetTime, &genpacket.SetTime{}))[0]
	if v.Outcome != OutcomeError || v.Class != ClassDirection {
		t.Fatalf("verdict %+v", v)
	}
}

func TestUnknownIDsAreForwardedUnchanged(t *testing.T) {
	unknown, err := UnusedPacketID(KnownPacket)
	if err != nil {
		t.Fatal(err)
	}
	input := Input{Fixture: "f.bin", Case: CaseUnknownID, Sender: SenderServer, Batch: JoinBatch([]Entry{NewEntry(unknown, []byte{1, 2, 3})})}
	for name, v := range map[string]Verdict{"go": DecodeGenerated(input)[0], "gophertunnel": DecodeGophertunnel(input)[0]} {
		if v.Outcome != OutcomeUnknown || v.Encoded != v.Payload {
			t.Errorf("%s verdict %+v", name, v)
		}
	}
}

func TestTrailingBytesFailBothGoDecoders(t *testing.T) {
	entry := NewEntry(genpacket.IDSetTime, []byte{0x0a, 0x00})
	input := Input{Fixture: "f.bin", Case: CaseTrailing, Sender: SenderServer, Batch: JoinBatch([]Entry{entry})}
	for name, v := range map[string]Verdict{"go": DecodeGenerated(input)[0], "gophertunnel": DecodeGophertunnel(input)[0]} {
		if v.Outcome != OutcomeError || v.Class != ClassTrailing {
			t.Errorf("%s verdict %+v", name, v)
		}
	}
}

func TestCraftingVariantsRoundTripThroughTheirEncoder(t *testing.T) {
	inputs, err := CraftingInputs()
	if err != nil {
		t.Fatal(err)
	}
	if len(inputs) != len(craftingVariants()) {
		t.Fatalf("%d inputs for %d variants", len(inputs), len(craftingVariants()))
	}
	for _, input := range inputs {
		verdicts := DecodeGophertunnel(input)
		if len(verdicts) != 1 {
			t.Fatalf("%s: %d entries", input.Fixture, len(verdicts))
		}
		v := verdicts[0]
		if v.Header != "52:0:0" || v.Outcome != OutcomeOK || v.Encoded != v.Payload || v.RoundTrip != RoundTripSame {
			t.Errorf("%s: verdict %+v", input.Fixture, v)
		}
	}
}

func TestLargeInputRoundTripsThroughTheGeneratedCodec(t *testing.T) {
	inputs, err := LargeInputs()
	if err != nil {
		t.Fatal(err)
	}
	for _, input := range inputs {
		v := DecodeGenerated(input)[0]
		if v.Outcome != OutcomeOK || v.Encoded != v.Payload || v.RoundTrip != RoundTripSame {
			t.Errorf("%s: verdict %+v", input.Fixture, v)
		}
	}
}

package retail

import (
	"bytes"
	"testing"
)

func TestBatchesSplitAndJoinExactly(t *testing.T) {
	entries := []Entry{NewEntry(9|1<<10|2<<12, []byte{1, 2, 3}), NewEntry(300, nil)}
	batch := JoinBatch(entries)
	got, err := SplitBatch(batch)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].Header != entries[0].Header || !bytes.Equal(got[0].Body, []byte{1, 2, 3}) || got[1].Header != 300 || len(got[1].Body) != 0 {
		t.Fatalf("split %+v", got)
	}
	if !bytes.Equal(JoinBatch(got), batch) {
		t.Fatal("rejoined batch differs")
	}
}

func TestSplitBatchKeepsEntriesBeforeAnOverrun(t *testing.T) {
	batch := append(JoinBatch([]Entry{NewEntry(9, []byte{1})}), 5, 9)
	got, err := SplitBatch(batch)
	if err == nil || len(got) != 1 {
		t.Fatalf("split %d entries, error %v", len(got), err)
	}
	if _, err := SplitBatch([]byte{0xfe, 0x80, 0x80, 0x80, 0x80, 0x80}); err == nil {
		t.Fatal("accepted a six-byte length")
	}
	if _, err := SplitBatch([]byte{0x01}); err == nil {
		t.Fatal("accepted a batch without 0xfe")
	}
}

func TestBuildInputsDerivesEveryCase(t *testing.T) {
	corpus := Corpus{Fixtures: []Fixture{
		{File: "one_way.bin", Direction: "clientbound", Sender: SenderServer},
		{File: "both_ways.bin", Direction: "bidirectional", Sender: SenderClient},
		{File: "empty_body.bin", Direction: "bidirectional", Sender: SenderClient},
	}}
	batches := map[string][]byte{
		"one_way.bin":    JoinBatch([]Entry{NewEntry(11|3<<12, []byte{7, 8})}),
		"both_ways.bin":  JoinBatch([]Entry{NewEntry(9, []byte{1}), NewEntry(9, []byte{2, 3})}),
		"empty_body.bin": JoinBatch([]Entry{NewEntry(10, nil)}),
	}
	inputs, err := BuildInputs(corpus, batches, 1000)
	if err != nil {
		t.Fatal(err)
	}
	byKey := map[string]Input{}
	for _, input := range inputs {
		byKey[input.Case+"/"+input.Fixture] = input
	}
	if len(byKey) != len(inputs) {
		t.Fatal("inputs repeat a case and fixture")
	}
	if _, ok := byKey[CaseWrongSender+"/both_ways.bin"]; ok {
		t.Fatal("derived a wrong sender for a bidirectional packet")
	}
	if input := byKey[CaseWrongSender+"/one_way.bin"]; input.Sender != SenderClient || !bytes.Equal(input.Batch, batches["one_way.bin"]) {
		t.Fatalf("wrong-sender input %+v", input)
	}
	if _, ok := byKey[CaseTruncated+"/empty_body.bin"]; ok {
		t.Fatal("truncated an empty body")
	}
	entries := func(key string) []Entry {
		got, err := SplitBatch(byKey[key].Batch)
		if err != nil {
			t.Fatalf("%s: %v", key, err)
		}
		return got
	}
	if got := entries(CaseUnknownID + "/one_way.bin"); got[0].Header != 1000|3<<12 || !bytes.Equal(got[0].Body, []byte{7, 8}) {
		t.Fatalf("unknown-id entry %+v", got[0])
	}
	if got := entries(CaseTruncated + "/both_ways.bin"); len(got[0].Body) != 0 || !bytes.Equal(got[1].Body, []byte{2}) {
		t.Fatalf("truncated entries %+v", got)
	}
	if got := entries(CaseTrailing + "/both_ways.bin"); !bytes.Equal(got[1].Body, []byte{2, 3, 0}) {
		t.Fatalf("trailing entries %+v", got)
	}
	if !bytes.Equal(entries(CaseRetail + "/both_ways.bin")[1].Body, []byte{2, 3}) {
		t.Fatal("deriving cases changed the retail batch")
	}
}

func TestUnusedPacketIDIsTheHighestUndefined(t *testing.T) {
	id, err := UnusedPacketID(func(id uint32) bool { return id > 1000 })
	if err != nil || id != 1000 {
		t.Fatalf("got %d, %v", id, err)
	}
	if _, err := UnusedPacketID(func(uint32) bool { return true }); err == nil {
		t.Fatal("found an id when every id is defined")
	}
}

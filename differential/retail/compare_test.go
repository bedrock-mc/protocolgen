package retail

import (
	"os"
	"path/filepath"
	"testing"
)

func testExpected() []Expected {
	var expected []Expected
	for _, caseName := range []string{CaseRetail, CaseCrafting, CaseUnknownID, CaseWrongSender, CaseTruncated, CaseTrailing, CaseLarge} {
		for _, fixture := range []string{"a.bin", "b.bin"} {
			for entry := range 2 {
				expected = append(expected, Expected{fixture, caseName, entry, "9:0:0", caseName + fixture + string(rune('0'+entry))})
			}
		}
	}
	return expected
}

// agreeing returns verdicts that meet every expectation.
func agreeing(expected []Expected) map[string][]Verdict {
	verdicts := map[string][]Verdict{}
	for caseName, r := range rules {
		for _, decoder := range r.decoders {
			for _, want := range expected {
				if want.Case != r.source {
					continue
				}
				v := Verdict{Fixture: want.Fixture, Case: caseName, Entry: want.Entry, Header: want.Header, Payload: want.Payload}
				switch r.expect {
				case expectRoundTrip:
					v.Outcome, v.Encoded, v.RoundTrip = OutcomeOK, want.Payload, RoundTripSame
				case expectUnknown:
					v.Outcome, v.Encoded = OutcomeUnknown, want.Payload
				case expectReject:
					v.Outcome, v.Class = OutcomeError, ClassDecode
				case expectDirection:
					v.Outcome, v.Class = OutcomeError, ClassDirection
				case expectAlloc:
					v.Outcome, v.Class = OutcomeError, ClassAlloc
				default:
					v.Outcome = OutcomeOK
				}
				verdicts[decoder] = append(verdicts[decoder], v)
			}
		}
	}
	return verdicts
}

func edit(verdicts map[string][]Verdict, decoder, caseName, fixture string, entry int, change func(*Verdict)) {
	for i := range verdicts[decoder] {
		v := &verdicts[decoder][i]
		if v.Case == caseName && v.Fixture == fixture && v.Entry == entry {
			change(v)
		}
	}
}

func remove(verdicts map[string][]Verdict, decoder, caseName, fixture string, entry int) {
	kept := verdicts[decoder][:0]
	for _, v := range verdicts[decoder] {
		if v.Case != caseName || v.Fixture != fixture || v.Entry != entry {
			kept = append(kept, v)
		}
	}
	verdicts[decoder] = kept
}

func kindsOf(found []Disagreement) []string {
	var out []string
	for _, d := range found {
		out = append(out, d.Decoder+" "+d.Case+" "+d.Fixture+" "+d.Kind)
	}
	return out
}

func wantKinds(t *testing.T, found []Disagreement, want ...string) {
	t.Helper()
	got := kindsOf(found)
	if len(got) != len(want) {
		t.Fatalf("disagreements %q, want %q", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("disagreements %q, want %q", got, want)
		}
	}
}

func TestAgreeingVerdictsHaveNoDisagreements(t *testing.T) {
	expected := testExpected()
	wantKinds(t, Compare(expected, agreeing(expected)))
}

func TestReencodedBytesMustMatchTheEntry(t *testing.T) {
	expected := testExpected()
	verdicts := agreeing(expected)
	edit(verdicts, DecoderGo, CaseRetail, "a.bin", 1, func(v *Verdict) { v.Encoded = "other" })
	wantKinds(t, Compare(expected, verdicts), "go retail a.bin reencode")
}

// An accepted byte mismatch must not hide a round-trip failure on the same
// entry.
func TestReencodeExceptionDoesNotHideRoundTrip(t *testing.T) {
	expected := testExpected()
	verdicts := agreeing(expected)
	edit(verdicts, DecoderGophertunnel, CaseCrafting, "a.bin", 0, func(v *Verdict) { v.Encoded, v.RoundTrip = "other", RoundTripDiffers })
	found := Compare(expected, verdicts)
	wantKinds(t, found, "gophertunnel crafting a.bin reencode", "gophertunnel crafting a.bin roundtrip")
	exception := Exception{Decoder: DecoderGophertunnel, Case: CaseCrafting, Fixtures: []string{"a.bin"}, Kind: KindReencode}
	report := Apply(found, []Exception{exception})
	if len(report.Unexpected) != 1 || report.Unexpected[0].Kind != KindRoundTrip || len(report.Stale) != 0 {
		t.Fatalf("report %+v", report)
	}
}

func TestUnknownIDMustStayUnknown(t *testing.T) {
	expected := testExpected()
	verdicts := agreeing(expected)
	edit(verdicts, DecoderValentine, CaseUnknownID, "b.bin", 0, func(v *Verdict) { v.Outcome, v.Class = OutcomeError, "unknown-id" })
	wantKinds(t, Compare(expected, verdicts), "valentine unknown-id b.bin decode")
}

func TestWrongSenderMustFailOnDirection(t *testing.T) {
	expected := testExpected()
	verdicts := agreeing(expected)
	edit(verdicts, DecoderGophertunnel, CaseWrongSender, "a.bin", 0, func(v *Verdict) { v.Outcome, v.Class = OutcomeUnknown, "" })
	edit(verdicts, DecoderRust, CaseWrongSender, "a.bin", 0, func(v *Verdict) { v.Class = "eof" })
	wantKinds(t, Compare(expected, verdicts), "gophertunnel wrong-sender a.bin accepted", "rust wrong-sender a.bin class")
}

func TestFrameErrorSuppressesLaterEntries(t *testing.T) {
	expected := testExpected()
	verdicts := agreeing(expected)
	edit(verdicts, DecoderGophertunnel, CaseTrailing, "a.bin", 0, func(v *Verdict) { v.Outcome, v.Header, v.Payload = OutcomeFrameError, "", "" })
	remove(verdicts, DecoderGophertunnel, CaseTrailing, "a.bin", 1)
	wantKinds(t, Compare(expected, verdicts), "gophertunnel trailing a.bin frame")
}

func TestMissingAndExtraVerdictsAreReported(t *testing.T) {
	expected := testExpected()
	verdicts := agreeing(expected)
	remove(verdicts, DecoderRust, CaseCrafting, "b.bin", 1)
	verdicts[DecoderGo] = append(verdicts[DecoderGo], Verdict{Fixture: "c.bin", Case: CaseRetail, Outcome: OutcomeOK})
	wantKinds(t, Compare(expected, verdicts), "go retail c.bin extra", "rust crafting b.bin missing")
}

func TestProbesSkipEntriesWhoseBaseCaseFailed(t *testing.T) {
	expected := testExpected()
	verdicts := agreeing(expected)
	edit(verdicts, DecoderRust, CaseRetail, "a.bin", 0, func(v *Verdict) { v.Outcome, v.Class = OutcomeError, "eof" })
	remove(verdicts, DecoderRust, ProbeElementLimit, "a.bin", 0)
	edit(verdicts, DecoderRust, CaseLarge, "b.bin", 1, func(v *Verdict) { v.Outcome, v.Class = OutcomeError, "eof" })
	edit(verdicts, DecoderRust, ProbeAllocFailure, "b.bin", 1, func(v *Verdict) { v.Class = "eof" })
	wantKinds(t, Compare(expected, verdicts), "rust large b.bin decode", "rust retail a.bin decode")
}

func TestElementLimitOutcomesMustAgree(t *testing.T) {
	expected := testExpected()
	verdicts := agreeing(expected)
	edit(verdicts, DecoderRust, ProbeElementLimit, "b.bin", 1, func(v *Verdict) { v.Outcome, v.Class = OutcomeError, ClassLimit })
	wantKinds(t, Compare(expected, verdicts), "rust element-limit b.bin limit-mismatch")
}

func TestAllocationProbeMustFailOnTheRefusedAllocation(t *testing.T) {
	expected := testExpected()
	verdicts := agreeing(expected)
	edit(verdicts, DecoderRust, ProbeAllocFailure, "a.bin", 0, func(v *Verdict) { v.Outcome, v.Class = OutcomeOK, "" })
	edit(verdicts, DecoderRust, ProbeAllocFailure, "b.bin", 0, func(v *Verdict) { v.Class = "eof" })
	wantKinds(t, Compare(expected, verdicts), "rust alloc-failure a.bin unexercised", "rust alloc-failure b.bin class")
}

func TestExceptionsAcceptMatchesAndStaleOnesFail(t *testing.T) {
	found := []Disagreement{
		{DecoderValentine, CaseUnknownID, "a.bin", 0, KindDecode, ""},
		{DecoderValentine, CaseUnknownID, "b.bin", 0, KindDecode, ""},
		{DecoderGo, CaseRetail, "a.bin", 0, KindReencode, ""},
		{DecoderGophertunnel, CaseTruncated, "a.bin", 0, KindAccepted, ""},
	}
	exceptions := []Exception{
		{Decoder: DecoderValentine, Case: CaseUnknownID, Fixtures: []string{AnyFixture}, Kind: KindDecode},
		{Decoder: DecoderGophertunnel, Case: CaseTruncated, Fixtures: []string{"a.bin", "b.bin"}, Kind: KindAccepted},
		{Decoder: DecoderRust, Case: CaseRetail, Fixtures: []string{"a.bin"}, Kind: KindDecode},
	}
	report := Apply(found, exceptions)
	if len(report.Accepted) != 3 || len(report.Unexpected) != 1 || len(report.Stale) != 2 || !report.Failed() {
		t.Fatalf("report %+v", report)
	}
	if stale := report.Stale[0]; stale.Exception.Decoder != DecoderGophertunnel || stale.Fixture != "b.bin" {
		t.Fatalf("stale fixture %+v", stale)
	}
	if Apply(found[:2], exceptions[:1]).Failed() {
		t.Fatal("fully accepted comparison failed")
	}
}

func TestPanicsAreNeverAcceptedAsRejections(t *testing.T) {
	expected := testExpected()
	verdicts := agreeing(expected)
	edit(verdicts, DecoderRust, CaseTruncated, "a.bin", 0, func(v *Verdict) { v.Outcome, v.Class = OutcomePanic, "" })
	edit(verdicts, DecoderRust, CaseRetail, "b.bin", 0, func(v *Verdict) { v.Outcome = OutcomeError })
	edit(verdicts, DecoderRust, ProbeElementLimit, "b.bin", 0, func(v *Verdict) { v.Outcome = OutcomePanic })
	wantKinds(t, Compare(expected, verdicts), "rust element-limit b.bin panic", "rust retail b.bin decode", "rust truncated a.bin panic")
}

func TestFrameErrorsPastTheLastEntryAreReported(t *testing.T) {
	expected := testExpected()
	verdicts := agreeing(expected)
	verdicts[DecoderValentine] = append(verdicts[DecoderValentine], Verdict{Fixture: "a.bin", Case: CaseRetail, Entry: 2, Outcome: OutcomeFrameError})
	wantKinds(t, Compare(expected, verdicts), "valentine retail a.bin frame")
}

func TestLoadExceptionsRejectsUnreviewedEntries(t *testing.T) {
	corpus := Corpus{MinecraftVersion: "1.26.51", ProtocolVersion: 2193}
	const oracle = "github.com/fork/gophertunnel v1"
	entry := func(decoder, fixtures, kind, extra string) string {
		return `{"decoder":"` + decoder + `","case":"retail","fixtures":` + fixtures + `,"kind":"` + kind + `","reason":"why","evidence":["where"]` + extra + `}`
	}
	valid := entry("go", `["a.bin"]`, "decode", "")
	for name, body := range map[string]string{
		"no evidence":      `{"decoder":"go","case":"retail","fixtures":["a.bin"],"kind":"decode","reason":"why","evidence":[]}`,
		"no reason":        `{"decoder":"go","case":"retail","fixtures":["a.bin"],"kind":"decode","reason":" ","evidence":["where"]}`,
		"unknown decoder":  entry("java", `["a.bin"]`, "decode", ""),
		"unknown kind":     entry("go", `["a.bin"]`, "slow", ""),
		"excused panic":    entry("go", `["a.bin"]`, "panic", ""),
		"no fixtures":      entry("go", `[]`, "decode", ""),
		"wildcard mixed":   entry("go", `["*","a.bin"]`, "decode", ""),
		"unpinned oracle":  entry("gophertunnel", `["a.bin"]`, "decode", ""),
		"other oracle":     entry("gophertunnel", `["a.bin"]`, "decode", `,"oracle":"github.com/fork/gophertunnel v2"`),
		"repeated fixture": valid + "," + entry("go", `["b.bin","a.bin"]`, "decode", ""),
	} {
		path := filepath.Join(t.TempDir(), "exceptions.json")
		data := `{"schema_version":1,"minecraft_version":"1.26.51","protocol_version":2193,"exceptions":[` + body + `]}`
		if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
			t.Fatal(err)
		}
		if _, err := LoadExceptions(path, corpus, oracle); err == nil {
			t.Errorf("%s: loaded", name)
		}
	}
}

func TestVerdictLinesRoundTrip(t *testing.T) {
	v := Verdict{Fixture: "a.bin", Case: CaseRetail, Entry: 3, Header: "9:1:2", Outcome: OutcomeError, Class: ClassTrailing, Payload: "00ff", Detail: "two\tlines\n"}
	got, err := ParseVerdict(v.Line())
	if err != nil {
		t.Fatal(err)
	}
	v.Detail = "two lines"
	if got != v {
		t.Fatalf("parsed %+v, want %+v", got, v)
	}
	if _, err := ParseVerdict("a\tb"); err == nil {
		t.Fatal("parsed a short line")
	}
}

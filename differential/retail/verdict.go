// Package retail runs Cinnabar's pinned retail packet fixtures through every
// decoder and compares their per-entry verdicts.
package retail

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Decoder names, which are also the verdict file names.
const (
	DecoderValentine    = "valentine"
	DecoderRust         = "rust"
	DecoderGo           = "go"
	DecoderGophertunnel = "gophertunnel"
)

// Decoders lists every decoder the comparison requires.
var Decoders = []string{DecoderValentine, DecoderRust, DecoderGo, DecoderGophertunnel}

// Verdict outcomes.
const (
	OutcomeOK         = "ok"
	OutcomeUnknown    = "unknown"
	OutcomeError      = "error"
	OutcomeFrameError = "frame-error"
	OutcomePanic      = "panic" // the decoder panicked; never an acceptable rejection
)

// Round-trip results: whether decoding the re-encoded entry reproduces the
// decoder's own field rendering.
const (
	RoundTripSame    = "same"
	RoundTripDiffers = "differs"
)

// Error classes the comparator checks; decoders may report others.
const (
	ClassDirection = "direction"
	ClassLimit     = "limit"
	ClassAlloc     = "alloc"
	ClassTrailing  = "trailing"
	ClassZeroCopy  = "zero-copy"
	ClassDecode    = "decode"
)

const (
	none         = "-"
	verdictCols  = 11
	maxDetailLen = 240
	idMask       = 0x3ff
)

// Verdict is one decoder's result for one batch entry. Digests are the first
// eight bytes of SHA-256 in hex.
type Verdict struct {
	Fixture   string
	Case      string
	Entry     int
	Header    string // id:sender:target subclients, or "-" when framing failed
	Outcome   string
	Class     string
	Payload   string // digest of the entry's header varint and body as received
	Fields    string // digest of the decoder's own rendering of the decoded value
	RoundTrip string
	Encoded   string // digest of the re-encoded header varint and body
	Detail    string
}

// Line formats v as one tab-separated verdict line.
func (v Verdict) Line() string {
	return strings.Join([]string{
		v.Fixture, v.Case, strconv.Itoa(v.Entry), orNone(v.Header), v.Outcome, orNone(v.Class),
		orNone(v.Payload), orNone(v.Fields), orNone(v.RoundTrip), orNone(v.Encoded), sanitize(v.Detail),
	}, "\t")
}

// ParseVerdict parses a line written by Line or by the Rust runner.
func ParseVerdict(line string) (Verdict, error) {
	cols := strings.Split(line, "\t")
	if len(cols) != verdictCols {
		return Verdict{}, fmt.Errorf("verdict has %d columns, want %d", len(cols), verdictCols)
	}
	entry, err := strconv.Atoi(cols[2])
	if err != nil || entry < 0 {
		return Verdict{}, fmt.Errorf("verdict entry %q is not an index", cols[2])
	}
	switch cols[4] {
	case OutcomeOK, OutcomeUnknown, OutcomeError, OutcomeFrameError, OutcomePanic:
	default:
		return Verdict{}, fmt.Errorf("verdict outcome %q is not recognised", cols[4])
	}
	field := func(i int) string {
		if cols[i] == none {
			return ""
		}
		return cols[i]
	}
	return Verdict{
		Fixture: cols[0], Case: cols[1], Entry: entry, Header: field(3), Outcome: cols[4], Class: field(5),
		Payload: field(6), Fields: field(7), RoundTrip: field(8), Encoded: field(9), Detail: field(10),
	}, nil
}

// fail records err, as a panic outcome when it came from a recovered panic.
func (v *Verdict) fail(class string, err error) {
	var panicked *panicError
	if errors.As(err, &panicked) {
		v.Outcome, v.Class, v.Detail = OutcomePanic, "", err.Error()
		return
	}
	v.Outcome, v.Class, v.Detail = OutcomeError, class, err.Error()
}

// encodeFailed records a re-encoding error, keeping the decode outcome unless
// the encoder panicked.
func (v *Verdict) encodeFailed(err error) {
	var panicked *panicError
	if errors.As(err, &panicked) {
		v.fail("", err)
		return
	}
	v.Detail = "encode: " + err.Error()
}

// Digest returns the verdict digest of data.
func Digest(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:8])
}

// HeaderText formats the id and subclient bits of a raw header varint.
func HeaderText(raw uint32) string {
	return fmt.Sprintf("%d:%d:%d", raw&idMask, (raw>>10)&3, (raw>>12)&3)
}

func orNone(value string) string {
	if value == "" {
		return none
	}
	return value
}

func sanitize(detail string) string {
	detail = strings.Map(func(r rune) rune {
		if r < 0x20 || r == 0x7f {
			return ' '
		}
		return r
	}, detail)
	if runes := []rune(detail); len(runes) > maxDetailLen {
		detail = string(runes[:maxDetailLen])
	}
	return orNone(strings.TrimSpace(detail))
}

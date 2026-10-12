package retail

import (
	"errors"
	"fmt"

	genpacket "protocolgen/generated/1.26.51/go/protocol/packet"
)

// BuildInputs returns the retail batches and the cases derived from them.
// unknownID must be an id that no decoder defines.
func BuildInputs(corpus Corpus, batches map[string][]byte, unknownID uint32) ([]Input, error) {
	var inputs []Input
	for _, fixture := range corpus.Fixtures {
		batch, ok := batches[fixture.File]
		if !ok {
			return nil, fmt.Errorf("fixture %s was not read", fixture.File)
		}
		entries, err := SplitBatch(batch)
		if err != nil {
			return nil, fmt.Errorf("fixture %s: %w", fixture.File, err)
		}
		add := func(caseName, sender string, batch []byte) {
			inputs = append(inputs, Input{Fixture: fixture.File, Case: caseName, Sender: sender, Batch: batch})
		}
		add(CaseRetail, fixture.Sender, batch)
		add(CaseUnknownID, fixture.Sender, JoinBatch(mapEntries(entries, func(e Entry) Entry {
			return NewEntry(e.Header&^idMask|unknownID, e.Body)
		})))
		if fixture.Direction != "bidirectional" {
			add(CaseWrongSender, otherSender(fixture.Sender), batch)
		}
		if allBodies(entries) {
			add(CaseTruncated, fixture.Sender, JoinBatch(mapEntries(entries, func(e Entry) Entry {
				return NewEntry(e.Header, e.Body[:len(e.Body)-1])
			})))
		}
		add(CaseTrailing, fixture.Sender, JoinBatch(mapEntries(entries, func(e Entry) Entry {
			return NewEntry(e.Header, append(append([]byte(nil), e.Body...), 0))
		})))
	}
	return inputs, nil
}

// LargeInputs encodes, with the generated Go codec, a server packet whose
// collection needs a reservation well above the allocation probe's ceiling:
// 3000 strings of at least 32 bytes each in the contract's Rust model.
func LargeInputs() ([]Input, error) {
	body, err := genpacket.Encode(&genpacket.UpdateSoftEnum{EnumName: "harness", Values: make([]string, 3000)})
	if err != nil {
		return nil, err
	}
	return []Input{{
		Fixture: "update_soft_enum_3000_values.bin",
		Case:    CaseLarge,
		Sender:  SenderServer,
		Batch:   JoinBatch([]Entry{NewEntry(genpacket.IDUpdateSoftEnum, body)}),
	}}, nil
}

// UnusedPacketID returns the highest ten-bit packet id for which known is
// false.
func UnusedPacketID(known func(uint32) bool) (uint32, error) {
	for id := uint32(idMask); id > 0; id-- {
		if !known(id) {
			return id, nil
		}
	}
	return 0, errors.New("every packet id is defined")
}

func mapEntries(entries []Entry, f func(Entry) Entry) []Entry {
	out := make([]Entry, len(entries))
	for i, entry := range entries {
		out[i] = f(entry)
	}
	return out
}

func allBodies(entries []Entry) bool {
	for _, entry := range entries {
		if len(entry.Body) == 0 {
			return false
		}
	}
	return len(entries) != 0
}

func otherSender(sender string) string {
	if sender == SenderServer {
		return SenderClient
	}
	return SenderServer
}

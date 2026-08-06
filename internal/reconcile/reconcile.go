// Package reconcile turns independent source claims into the canonical v2
// manifest. It has no precedence rule: equal claims merge source pins and
// different claims require a fingerprinted, evidenced adjudication.
package reconcile

import (
	"fmt"
	"reflect"
	"sort"

	"protocolgen/internal/claims"
	"protocolgen/internal/manifest"
)

type fieldKey struct {
	PacketID uint32
	Ordinal  int
}

func Reconcile(target manifest.Target, results []claims.Result, adjudications []manifest.Adjudication) (manifest.Manifest, error) {
	if len(results) == 0 {
		return manifest.Manifest{}, fmt.Errorf("reconcile has no source results")
	}
	allPins := make([]manifest.SourcePin, 0, len(results))
	pins := map[string]manifest.SourcePin{}
	allOverrides := map[string]manifest.OverrideProof{}
	groups := map[fieldKey][]claims.Claim{}
	for resultIndex, result := range results {
		if result.Pin.ID == "" {
			return manifest.Manifest{}, fmt.Errorf("source result %d has no source pin", resultIndex)
		}
		if result.Target != target {
			return manifest.Manifest{}, fmt.Errorf("source %q target does not match; mixing protocol snapshots is forbidden", result.Pin.ID)
		}
		if old, ok := pins[result.Pin.ID]; ok && !reflect.DeepEqual(old, result.Pin) {
			return manifest.Manifest{}, fmt.Errorf("source pin %q is not stable", result.Pin.ID)
		}
		if _, ok := pins[result.Pin.ID]; !ok {
			pins[result.Pin.ID] = result.Pin
			allPins = append(allPins, result.Pin)
		}
		for _, claim := range result.Claims {
			if claim.SourceID != result.Pin.ID {
				return manifest.Manifest{}, fmt.Errorf("claim %s has source %q but result is %q", claim.FieldPath, claim.SourceID, result.Pin.ID)
			}
			groups[fieldKey{PacketID: claim.PacketID, Ordinal: claim.Ordinal}] = append(groups[fieldKey{PacketID: claim.PacketID, Ordinal: claim.Ordinal}], claim)
		}
		for _, proof := range result.Overrides {
			if old, ok := allOverrides[proof.ID]; ok && !reflect.DeepEqual(old, proof) {
				return manifest.Manifest{}, fmt.Errorf("override proof %q is not stable", proof.ID)
			}
			allOverrides[proof.ID] = proof
		}
	}
	if len(groups) == 0 {
		return manifest.Manifest{}, fmt.Errorf("reconcile has no packet field claims")
	}

	keys := make([]fieldKey, 0, len(groups))
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].PacketID != keys[j].PacketID {
			return keys[i].PacketID < keys[j].PacketID
		}
		return keys[i].Ordinal < keys[j].Ordinal
	})

	packets := map[uint32]*manifest.Packet{}
	usedAdjudications := map[string]bool{}
	for _, key := range keys {
		group := groups[key]
		sort.SliceStable(group, func(i, j int) bool { return group[i].SourceID < group[j].SourceID })
		selected, pinsForField, evidence, used, err := selectClaim(target, group, adjudications)
		if err != nil {
			return manifest.Manifest{}, err
		}
		if used != "" {
			usedAdjudications[used] = true
		}
		packet := packets[selected.PacketID]
		if packet == nil {
			packet = &manifest.Packet{ID: selected.PacketID, Name: selected.PacketName, Direction: selected.Direction}
			packets[selected.PacketID] = packet
		}
		if packet.Name != selected.PacketName || packet.Direction != selected.Direction {
			return manifest.Manifest{}, fmt.Errorf("packet %d metadata disagrees without a field adjudication", selected.PacketID)
		}
		field := manifest.Field{
			Ordinal: selected.Ordinal, Name: selected.Name, Semantic: selected.Semantic, TypeID: selected.TypeID,
			Encode: selected.Encode, Decode: selected.Decode, Symmetry: selected.Symmetry,
			Reserved: selected.Reserved, Ignored: selected.Ignored, Compatibility: append([]string(nil), selected.Compatibility...),
			Provenance: manifest.Provenance{Pins: pinsForField, Evidence: evidence},
		}
		packet.Fields = append(packet.Fields, field)
	}

	packetsOut := make([]manifest.Packet, 0, len(packets))
	for _, packet := range packets {
		sort.Slice(packet.Fields, func(i, j int) bool { return packet.Fields[i].Ordinal < packet.Fields[j].Ordinal })
		packetsOut = append(packetsOut, *packet)
	}
	sort.Slice(packetsOut, func(i, j int) bool { return packetsOut[i].ID < packetsOut[j].ID })
	sort.Slice(allPins, func(i, j int) bool { return allPins[i].ID < allPins[j].ID })

	used := make([]manifest.Adjudication, 0, len(usedAdjudications))
	for _, adjudication := range adjudications {
		if usedAdjudications[adjudication.ID] {
			used = append(used, adjudication)
		}
	}
	proofs := make([]manifest.OverrideProof, 0, len(allOverrides))
	for _, proof := range allOverrides {
		proofs = append(proofs, proof)
	}
	sort.Slice(proofs, func(i, j int) bool { return proofs[i].ID < proofs[j].ID })
	result := manifest.Manifest{SchemaVersion: manifest.SchemaVersion, Target: target, Sources: allPins, Packets: packetsOut, Adjudications: used, Overrides: proofs}
	if err := manifest.Validate(result); err != nil {
		return manifest.Manifest{}, err
	}
	return result, nil
}

func selectClaim(target manifest.Target, group []claims.Claim, adjudications []manifest.Adjudication) (claims.Claim, []string, []manifest.Evidence, string, error) {
	if len(group) == 0 {
		return claims.Claim{}, nil, nil, "", fmt.Errorf("empty claim group")
	}
	first := comparable(group[0])
	same := true
	for _, claim := range group[1:] {
		if !reflect.DeepEqual(first, comparable(claim)) {
			same = false
			break
		}
	}
	if same {
		pinIDs := make([]string, 0, len(group))
		for _, claim := range group {
			pinIDs = append(pinIDs, claim.SourceID)
		}
		sort.Strings(pinIDs)
		return group[0], pinIDs, nil, "", nil
	}

	context, err := claims.ContextFingerprint(target, group)
	if err != nil {
		return claims.Claim{}, nil, nil, "", err
	}
	for _, adjudication := range adjudications {
		if adjudication.Target != group[0].FieldPath {
			continue
		}
		if adjudication.PrePatchContextSHA256 != context {
			return claims.Claim{}, nil, nil, "", fmt.Errorf("stale adjudication %q for %s: pre-patch context fingerprint changed", adjudication.ID, adjudication.Target)
		}
		if err := matchClaimFingerprints(adjudication, group); err != nil {
			return claims.Claim{}, nil, nil, "", fmt.Errorf("stale adjudication %q for %s: %w", adjudication.ID, adjudication.Target, err)
		}
		var selected *claims.Claim
		for i := range group {
			if group[i].SourceID == adjudication.SelectedSource {
				selected = &group[i]
				break
			}
		}
		if selected == nil {
			return claims.Claim{}, nil, nil, "", fmt.Errorf("adjudication %q selects a claim not present in disagreement", adjudication.ID)
		}
		return *selected, []string{selected.SourceID}, append([]manifest.Evidence(nil), adjudication.Evidence...), adjudication.ID, nil
	}
	return claims.Claim{}, nil, nil, "", fmt.Errorf("source claims for %s disagree; an evidenced fingerprinted adjudication is required", group[0].FieldPath)
}

func comparable(claim claims.Claim) claims.Claim {
	claim.SourceID = ""
	claim.Locator = ""
	return claim
}

func matchClaimFingerprints(adjudication manifest.Adjudication, group []claims.Claim) error {
	want := map[string]string{}
	for _, claim := range group {
		digest, err := claims.Fingerprint(claim)
		if err != nil {
			return err
		}
		want[claim.SourceID] = digest
	}
	got := map[string]string{}
	for _, claim := range adjudication.Claims {
		got[claim.SourceID] = claim.Digest
	}
	if !reflect.DeepEqual(want, got) {
		return fmt.Errorf("claim fingerprints no longer match")
	}
	return nil
}

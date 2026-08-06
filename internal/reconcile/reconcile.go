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

type packetMetadata struct {
	Name      string
	Direction manifest.Direction
}

func Reconcile(target manifest.Target, results []claims.Result, adjudications []manifest.Adjudication) (manifest.Manifest, error) {
	if len(results) == 0 {
		return manifest.Manifest{}, fmt.Errorf("reconcile has no source results")
	}
	allPins := make([]manifest.SourcePin, 0, len(results))
	pins := map[string]manifest.SourcePin{}
	allOverrides := map[string]manifest.OverrideProof{}
	groups := map[fieldKey][]claims.Claim{}
	packetMetadataByID := map[uint32]packetMetadata{}
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
		for _, packet := range result.Packets {
			if packet.SourceID != result.Pin.ID {
				return manifest.Manifest{}, fmt.Errorf("packet %d has source %q but result is %q", packet.ID, packet.SourceID, result.Pin.ID)
			}
			metadata := packetMetadata{Name: packet.Name, Direction: packet.Direction}
			if old, ok := packetMetadataByID[packet.ID]; ok {
				if old.Name != metadata.Name {
					return manifest.Manifest{}, fmt.Errorf("packet %d name disagrees: %q vs %q", packet.ID, old.Name, metadata.Name)
				}
				direction, ok := mergeDirection(old.Direction, metadata.Direction)
				if !ok {
					return manifest.Manifest{}, fmt.Errorf("packet %d direction disagrees: %q vs %q", packet.ID, old.Direction, metadata.Direction)
				}
				metadata.Direction = direction
			}
			packetMetadataByID[packet.ID] = metadata
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
	if len(groups) == 0 && len(packetMetadataByID) == 0 {
		return manifest.Manifest{}, fmt.Errorf("reconcile has no packet claims")
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
	for id, metadata := range packetMetadataByID {
		packets[id] = &manifest.Packet{ID: id, Name: metadata.Name, Direction: metadata.Direction}
	}
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
		direction, directionsAgree := mergeDirection(packet.Direction, selected.Direction)
		if packet.Name != selected.PacketName || !directionsAgree {
			return manifest.Manifest{}, fmt.Errorf("packet %d metadata disagrees without a field adjudication", selected.PacketID)
		}
		packet.Direction = direction
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

func mergeDirection(left, right manifest.Direction) (manifest.Direction, bool) {
	if left == manifest.DirectionUnknown {
		return right, true
	}
	if right == manifest.DirectionUnknown || left == right {
		return left, true
	}
	return "", false
}

func selectClaim(target manifest.Target, group []claims.Claim, adjudications []manifest.Adjudication) (claims.Claim, []string, []manifest.Evidence, string, error) {
	if len(group) == 0 {
		return claims.Claim{}, nil, nil, "", fmt.Errorf("empty claim group")
	}
	if merged, pinIDs, ok := mergeClaimGroup(group); ok {
		return merged, pinIDs, nil, "", nil
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

func mergeClaimGroup(group []claims.Claim) (claims.Claim, []string, bool) {
	selected := group[0]
	for _, claim := range group[1:] {
		if semanticScore(claim) > semanticScore(selected) {
			selected = claim
		}
	}
	merged := selected
	for _, claim := range group {
		if claim.PacketID != merged.PacketID || claim.Ordinal != merged.Ordinal || claim.Direction != merged.Direction || claim.Symmetry != merged.Symmetry || claim.Reserved != merged.Reserved || claim.Ignored != merged.Ignored || !reflect.DeepEqual(claim.Compatibility, merged.Compatibility) {
			return claims.Claim{}, nil, false
		}
		encode, ok := mergeNode(merged.Encode, claim.Encode)
		if !ok {
			return claims.Claim{}, nil, false
		}
		merged.Encode = encode
		leftDecode, rightDecode := merged.Encode, claim.Encode
		if merged.Decode != nil {
			leftDecode = *merged.Decode
		}
		if claim.Decode != nil {
			rightDecode = *claim.Decode
		}
		decode, ok := mergeNode(leftDecode, rightDecode)
		if !ok {
			return claims.Claim{}, nil, false
		}
		if reflect.DeepEqual(wireNode(merged.Encode), wireNode(decode)) {
			merged.Decode = nil
		} else {
			merged.Decode = &decode
		}
	}
	var pinIDs []string
	for _, claim := range group {
		if hasConcreteEvidence(claim.Encode) || claim.Decode != nil && hasConcreteEvidence(*claim.Decode) {
			pinIDs = append(pinIDs, claim.SourceID)
		}
	}
	if len(pinIDs) == 0 {
		for _, claim := range group {
			pinIDs = append(pinIDs, claim.SourceID)
		}
	}
	sort.Strings(pinIDs)
	return merged, pinIDs, true
}

func hasConcreteEvidence(node manifest.Node) bool {
	if node.Kind == manifest.KindUnresolved || node.Kind == manifest.KindOpaque {
		return false
	}
	if node.Kind != manifest.KindStruct && node.Kind != manifest.KindSequence && node.Kind != manifest.KindUnion && node.Kind != manifest.KindConditional {
		return true
	}
	for _, child := range node.Elements {
		if hasConcreteEvidence(child) {
			return true
		}
	}
	for _, field := range node.Fields {
		if hasConcreteEvidence(field.Encode) {
			return true
		}
	}
	for _, variant := range node.Variants {
		if hasConcreteEvidence(variant.Encode) {
			return true
		}
	}
	for _, oneCase := range node.Cases {
		for _, child := range oneCase.Encode {
			if hasConcreteEvidence(child) {
				return true
			}
		}
	}
	return false
}

func mergeNode(left, right manifest.Node) (manifest.Node, bool) {
	if hasEvidenceGap(left) && (left.Kind == manifest.KindUnresolved || left.Kind == manifest.KindOpaque) {
		return right, true
	}
	if hasEvidenceGap(right) && (right.Kind == manifest.KindUnresolved || right.Kind == manifest.KindOpaque) {
		return left, true
	}
	if left.Kind == manifest.KindEnum && right.Kind == manifest.KindPrimitive && reflect.DeepEqual(left.Primitive, right.Primitive) {
		return left, true
	}
	if right.Kind == manifest.KindEnum && left.Kind == manifest.KindPrimitive && reflect.DeepEqual(left.Primitive, right.Primitive) {
		return right, true
	}
	if left.Kind == manifest.KindOptional && left.Value != nil && hasEvidenceGap(*left.Value) && right.Kind != manifest.KindOptional {
		result := left
		value := right
		result.Value = &value
		return result, true
	}
	if right.Kind == manifest.KindOptional && right.Value != nil && hasEvidenceGap(*right.Value) && left.Kind != manifest.KindOptional {
		result := right
		value := left
		result.Value = &value
		return result, true
	}
	if reflect.DeepEqual(wireNode(left), wireNode(right)) {
		if nodeSemanticScore(right) > nodeSemanticScore(left) {
			return right, true
		}
		return left, true
	}
	if left.Kind != right.Kind {
		return manifest.Node{}, false
	}
	result := left
	if nodeSemanticScore(right) > nodeSemanticScore(left) {
		result.Semantic, result.TypeID = right.Semantic, right.TypeID
	}
	var ok bool
	switch left.Kind {
	case manifest.KindStruct:
		if len(left.Fields) != len(right.Fields) {
			return manifest.Node{}, false
		}
		result.Fields = append([]manifest.Field(nil), left.Fields...)
		for index := range left.Fields {
			if left.Fields[index].Ordinal != right.Fields[index].Ordinal {
				return manifest.Node{}, false
			}
			result.Fields[index], ok = mergeField(left.Fields[index], right.Fields[index])
			if !ok {
				return manifest.Node{}, false
			}
		}
	case manifest.KindArray, manifest.KindFixedArray:
		if left.Length != right.Length {
			return manifest.Node{}, false
		}
		result.Prefix, ok = mergeNodePointer(left.Prefix, right.Prefix)
		if !ok {
			return manifest.Node{}, false
		}
		result.Element, ok = mergeNodePointer(left.Element, right.Element)
		if !ok {
			return manifest.Node{}, false
		}
	case manifest.KindOptional, manifest.KindReserved, manifest.KindIgnored:
		result.Value, ok = mergeNodePointer(left.Value, right.Value)
		if left.Kind == manifest.KindReserved || left.Kind == manifest.KindIgnored {
			result.Element, ok = mergeNodePointer(left.Element, right.Element)
		}
		if !ok {
			return manifest.Node{}, false
		}
	case manifest.KindMap:
		result.Prefix, ok = mergeNodePointer(left.Prefix, right.Prefix)
		if !ok {
			return manifest.Node{}, false
		}
		result.Key, ok = mergeNodePointer(left.Key, right.Key)
		if !ok {
			return manifest.Node{}, false
		}
		result.Value, ok = mergeNodePointer(left.Value, right.Value)
		if !ok {
			return manifest.Node{}, false
		}
	case manifest.KindSequence:
		if len(left.Elements) != len(right.Elements) {
			return manifest.Node{}, false
		}
		result.Elements = append([]manifest.Node(nil), left.Elements...)
		for index := range left.Elements {
			result.Elements[index], ok = mergeNode(left.Elements[index], right.Elements[index])
			if !ok {
				return manifest.Node{}, false
			}
		}
	case manifest.KindUnion:
		if len(left.Variants) != len(right.Variants) {
			return manifest.Node{}, false
		}
		result.Control, ok = mergeNodePointer(left.Control, right.Control)
		if !ok {
			return manifest.Node{}, false
		}
		result.Variants = append([]manifest.Variant(nil), left.Variants...)
		for index := range left.Variants {
			if left.Variants[index].Value != right.Variants[index].Value {
				return manifest.Node{}, false
			}
			result.Variants[index].Encode, ok = mergeNode(left.Variants[index].Encode, right.Variants[index].Encode)
			if !ok {
				return manifest.Node{}, false
			}
		}
	case manifest.KindString, manifest.KindBytes:
		if left.Encoding != right.Encoding || left.Representation != right.Representation {
			return manifest.Node{}, false
		}
		result.Prefix, ok = mergeNodePointer(left.Prefix, right.Prefix)
		if !ok {
			return manifest.Node{}, false
		}
	case manifest.KindEnum:
		if !reflect.DeepEqual(left.Primitive, right.Primitive) {
			return manifest.Node{}, false
		}
		if len(right.Variants) > len(left.Variants) {
			return right, true
		}
		return left, true
	default:
		return manifest.Node{}, false
	}
	return result, true
}

func mergeField(left, right manifest.Field) (manifest.Field, bool) {
	result := left
	if semanticScoreForField(right) > semanticScoreForField(left) {
		result.Name, result.Semantic, result.TypeID = right.Name, right.Semantic, right.TypeID
	}
	leftGap, rightGap := hasEvidenceGap(left.Encode), hasEvidenceGap(right.Encode)
	encode, ok := mergeNode(left.Encode, right.Encode)
	if !ok {
		return manifest.Field{}, false
	}
	result.Encode = encode
	if !leftGap && !rightGap {
		result.Provenance.Pins = mergeStrings(left.Provenance.Pins, right.Provenance.Pins)
	} else if leftGap && !rightGap {
		result.Provenance = right.Provenance
	}
	return result, true
}

func mergeNodePointer(left, right *manifest.Node) (*manifest.Node, bool) {
	if left == nil || right == nil {
		return nil, left == nil && right == nil
	}
	merged, ok := mergeNode(*left, *right)
	return &merged, ok
}

func semanticScoreForField(field manifest.Field) int {
	score := nodeSemanticScore(field.Encode)
	if field.Semantic != "" {
		score++
	}
	if field.TypeID != "" {
		score += 2
	}
	return score
}

func mergeStrings(left, right []string) []string {
	seen := map[string]bool{}
	for _, value := range append(append([]string(nil), left...), right...) {
		seen[value] = true
	}
	result := make([]string, 0, len(seen))
	for value := range seen {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func hasEvidenceGap(node manifest.Node) bool {
	if node.Kind == manifest.KindUnresolved || node.Kind == manifest.KindOpaque {
		return true
	}
	for _, child := range []*manifest.Node{node.Prefix, node.Element, node.Value, node.Key, node.Control, node.Default} {
		if child != nil && hasEvidenceGap(*child) {
			return true
		}
	}
	for _, child := range node.Elements {
		if hasEvidenceGap(child) {
			return true
		}
	}
	for _, field := range node.Fields {
		if hasEvidenceGap(field.Encode) || field.Decode != nil && hasEvidenceGap(*field.Decode) {
			return true
		}
	}
	for _, variant := range node.Variants {
		if hasEvidenceGap(variant.Encode) || variant.Decode != nil && hasEvidenceGap(*variant.Decode) {
			return true
		}
	}
	for _, oneCase := range node.Cases {
		for _, child := range oneCase.Encode {
			if hasEvidenceGap(child) {
				return true
			}
		}
		for _, child := range oneCase.Decode {
			if hasEvidenceGap(child) {
				return true
			}
		}
	}
	return false
}

func comparable(claim claims.Claim) claims.Claim {
	claim.SourceID = ""
	claim.Locator = ""
	return claim
}

func wireComparable(claim claims.Claim) claims.Claim {
	claim = comparable(claim)
	claim.PacketName = ""
	claim.FieldPath = ""
	claim.Name = ""
	claim.Semantic = ""
	claim.TypeID = ""
	claim.Encode = wireNode(claim.Encode)
	if claim.Decode != nil {
		decode := wireNode(*claim.Decode)
		claim.Decode = &decode
	}
	return claim
}

func wireNode(node manifest.Node) manifest.Node {
	node.Elements = append([]manifest.Node(nil), node.Elements...)
	node.Fields = append([]manifest.Field(nil), node.Fields...)
	node.Variants = append([]manifest.Variant(nil), node.Variants...)
	node.Cases = append([]manifest.Case(nil), node.Cases...)
	for index := range node.Cases {
		node.Cases[index].Encode = append([]manifest.Node(nil), node.Cases[index].Encode...)
		node.Cases[index].Decode = append([]manifest.Node(nil), node.Cases[index].Decode...)
	}
	node.Semantic = ""
	node.TypeID = ""
	if node.Kind == manifest.KindEnum {
		return manifest.Node{Kind: manifest.KindPrimitive, Primitive: node.Primitive}
	}
	node.Prefix = wireNodePointer(node.Prefix)
	node.Element = wireNodePointer(node.Element)
	node.Value = wireNodePointer(node.Value)
	node.Key = wireNodePointer(node.Key)
	node.Control = wireNodePointer(node.Control)
	node.Default = wireNodePointer(node.Default)
	for index := range node.Elements {
		node.Elements[index] = wireNode(node.Elements[index])
	}
	for index := range node.Fields {
		field := &node.Fields[index]
		field.Name = ""
		field.Semantic = ""
		field.TypeID = ""
		field.Provenance = manifest.Provenance{}
		field.Encode = wireNode(field.Encode)
		if field.Decode != nil {
			decode := wireNode(*field.Decode)
			field.Decode = &decode
		}
	}
	for index := range node.Variants {
		node.Variants[index].Name = ""
		node.Variants[index].Encode = wireNode(node.Variants[index].Encode)
		if node.Variants[index].Decode != nil {
			decode := wireNode(*node.Variants[index].Decode)
			node.Variants[index].Decode = &decode
		}
	}
	for caseIndex := range node.Cases {
		for nodeIndex := range node.Cases[caseIndex].Encode {
			node.Cases[caseIndex].Encode[nodeIndex] = wireNode(node.Cases[caseIndex].Encode[nodeIndex])
		}
		for nodeIndex := range node.Cases[caseIndex].Decode {
			node.Cases[caseIndex].Decode[nodeIndex] = wireNode(node.Cases[caseIndex].Decode[nodeIndex])
		}
	}
	return node
}

func wireNodePointer(node *manifest.Node) *manifest.Node {
	if node == nil {
		return nil
	}
	result := wireNode(*node)
	return &result
}

func semanticScore(claim claims.Claim) int {
	score := 0
	if claim.Semantic != "" {
		score++
	}
	if claim.TypeID != "" {
		score += 2
	}
	return score + nodeSemanticScore(claim.Encode)
}

func nodeSemanticScore(node manifest.Node) int {
	score := 0
	if node.Semantic != "" {
		score++
	}
	if node.TypeID != "" {
		score += 2
	}
	if node.Kind == manifest.KindEnum {
		score += 10 + len(node.Variants)
	}
	for _, field := range node.Fields {
		score += nodeSemanticScore(field.Encode)
	}
	return score
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

package gophertunneloracle

import "strings"

const missingVariantMetadataReason = "gophertunnel: byte operations match but union selector binding is unproved because conditional paths lack variant metadata"

// missingVariantMetadataPaths distinguishes missing extractor annotations from
// proved wire disagreement. Every path on both sides must match; uncertainty in
// one branch must not hide a payload, control width, or known selector mismatch.
func missingVariantMetadataPaths(want, got []wirePath) ([]PathResult, bool) {
	if len(want) == 0 || len(got) == 0 {
		return nil, false
	}
	matched := make([]bool, len(got))
	var results []PathResult
	uncertain := false
	for _, expected := range want {
		if len(expected.Reasons) != 0 {
			return nil, false
		}
		found := false
		for index, actual := range got {
			if len(actual.Reasons) != 0 {
				return nil, false
			}
			missing, equivalent := matchMissingVariantMetadata(expected.Atoms, actual.Atoms)
			if !equivalent {
				continue
			}
			matched[index] = true
			found = true
			path := PathResult{Classification: "AGREEMENT", ManifestConstraint: expected.Constraint, GophertunnelConstraint: actual.Constraint}
			if missing {
				uncertain = true
				path.Classification = "UNRESOLVED"
				path.Reasons = []string{missingVariantMetadataReason}
				path.ManifestSequence = atomDisplays(expected.Atoms)
				path.GophertunnelSequence = atomDisplays(actual.Atoms)
				path.GophertunnelSite = firstAtomSite(actual.Atoms)
			}
			results = append(results, path)
		}
		if !found {
			return nil, false
		}
	}
	for _, found := range matched {
		if !found {
			return nil, false
		}
	}
	return results, uncertain
}

// matchMissingVariantMetadata permits only omitted canonical variant markers.
// Existing oracle markers remain binding: a known different selector or any
// different byte operation prevents even an unresolved equivalence claim.
func matchMissingVariantMetadata(want, got []atom) (missing, equivalent bool) {
	want = normalizeFixedArrayGrouping(want)
	got = normalizeFixedArrayGrouping(got)
	left, right := 0, 0
	for left < len(want) {
		if right < len(got) && want[left].Token == got[right].Token {
			left++
			right++
			continue
		}
		if strings.HasPrefix(want[left].Token, "VARIANT:") && (right == len(got) || !strings.HasPrefix(got[right].Token, "VARIANT:")) {
			missing = true
			left++
			continue
		}
		return false, false
	}
	return missing, right == len(got)
}

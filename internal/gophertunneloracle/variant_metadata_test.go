package gophertunneloracle

import "testing"

// variantTestPath builds a small path while keeping wire tokens visible in tests.
func variantTestPath(tokens ...string) wirePath {
	path := wirePath{}
	for _, token := range tokens {
		path.Atoms = append(path.Atoms, atom{Token: token, Display: token})
	}
	return path
}

func TestMissingVariantMetadataIsUnresolved(t *testing.T) {
	want := []wirePath{
		variantTestPath("P:u8", "VARIANT:0"),
		variantTestPath("P:u8", "VARIANT:1", "P:i32le"),
	}
	got := []wirePath{
		variantTestPath("P:u8"),
		variantTestPath("P:u8", "P:i32le"),
	}
	paths, reasons, divergent := comparePaths(want, got)
	if divergent || len(reasons) != 1 || reasons[0] != missingVariantMetadataReason {
		t.Fatalf("paths=%#v reasons=%v divergent=%v", paths, reasons, divergent)
	}
	for _, path := range paths {
		if path.Classification != "UNRESOLVED" {
			t.Fatalf("missing selector binding was asserted as %s", path.Classification)
		}
	}
}

func TestMissingVariantMetadataPreservesDisagreements(t *testing.T) {
	for name, got := range map[string][]wirePath{
		"known-selector": {variantTestPath("P:u8", "VARIANT:2", "P:i32le")},
		"control-width":  {variantTestPath("P:var_u32", "P:i32le")},
		"payload":        {variantTestPath("P:u8", "P:i32be")},
		"extra-path":     {variantTestPath("P:u8", "P:i32le"), variantTestPath("P:u8", "P:f32le")},
	} {
		t.Run(name, func(t *testing.T) {
			want := []wirePath{variantTestPath("P:u8", "VARIANT:1", "P:i32le")}
			_, reasons, divergent := comparePaths(want, got)
			if !divergent || len(reasons) != 0 {
				t.Fatalf("real disagreement was hidden: reasons=%v divergent=%v", reasons, divergent)
			}
		})
	}
}

func TestMissingVariantMetadataDoesNotHideAnotherBranch(t *testing.T) {
	want := []wirePath{
		variantTestPath("P:u8", "VARIANT:0"),
		variantTestPath("P:u8", "VARIANT:1", "P:i32le"),
	}
	got := []wirePath{
		variantTestPath("P:u8"),
		variantTestPath("P:u8", "P:i32be"),
	}
	_, reasons, divergent := comparePaths(want, got)
	if !divergent || len(reasons) != 0 {
		t.Fatalf("one unknown selector hid a byte conflict: reasons=%v divergent=%v", reasons, divergent)
	}
}

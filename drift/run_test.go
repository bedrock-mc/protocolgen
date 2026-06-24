package main

import "testing"

func TestLoadSourcePacketsResolvesNumericIDs(t *testing.T) {
	pks, err := loadSourcePackets("testdata/packets")
	if err != nil {
		t.Fatalf("loadSourcePackets: %v", err)
	}
	var alpha *SourcePacket
	for i := range pks {
		if pks[i].TypeName == "Alpha" {
			alpha = &pks[i]
		}
	}
	if alpha == nil {
		t.Fatalf("Alpha packet not found in %+v", pks)
	}
	if alpha.ID != 1 {
		t.Errorf("Alpha.ID = %d, want 1 (resolved from IDAlpha)", alpha.ID)
	}
	if len(alpha.Ops) != 2 || alpha.Ops[0].Method != "Uint8" || alpha.Ops[1].Method != "Varuint32" {
		t.Errorf("Alpha.Ops = %+v, want [Uint8, Varuint32]", alpha.Ops)
	}
}

func TestLoadSourcePacketsDropsUnresolvedIDs(t *testing.T) {
	pks, err := loadSourcePackets("testdata/packets")
	if err != nil {
		t.Fatalf("loadSourcePackets: %v", err)
	}
	for _, p := range pks {
		if p.TypeName == "Dynamic" {
			t.Errorf("Dynamic packet (runtime id) should be dropped, got %+v", p)
		}
		if p.ID == 0 {
			t.Errorf("packet %q has unresolved id 0 and should have been dropped", p.TypeName)
		}
	}
}

func TestLoadDocPacketsReadsDirectory(t *testing.T) {
	docs, err := loadDocPackets("testdata/docs")
	if err != nil {
		t.Fatalf("loadDocPackets: %v", err)
	}
	var found bool
	for _, d := range docs {
		if d.Name == "AlphaPacket" && d.ID == 1 {
			found = true
		}
	}
	if !found {
		t.Errorf("AlphaPacket (id 1) not found in %+v", docs)
	}
}

func TestLoadDocPacketsSkipsLegacyCombinedFile(t *testing.T) {
	// Mojang's main branch still ships json/__protocoldoc.json, a JSON array
	// rather than a per-packet object. It must be skipped, not error the run.
	docs, err := loadDocPackets("testdata/docs")
	if err != nil {
		t.Fatalf("loadDocPackets must skip __protocoldoc.json, got error: %v", err)
	}
	for _, d := range docs {
		if d.Name == "" {
			t.Errorf("a malformed/legacy doc leaked into results: %+v", d)
		}
	}
}

func TestEndToEndCleanFixturesHaveNoDrift(t *testing.T) {
	docs, err := loadDocPackets("testdata/docs")
	if err != nil {
		t.Fatalf("loadDocPackets: %v", err)
	}
	srcs, err := loadSourcePackets("testdata/packets")
	if err != nil {
		t.Fatalf("loadSourcePackets: %v", err)
	}
	rep := Compare(docs, srcs)
	if rep.HasDrift() {
		t.Errorf("HasDrift() = true for matching fixtures; report = %+v", rep)
	}
}

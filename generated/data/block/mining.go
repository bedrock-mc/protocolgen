package block

import "sort"

// ToolFamilies is a set of tools that can mine a block or qualify for its drops.
type ToolFamilies uint8

const (
	ToolPickaxe ToolFamilies = 1 << iota
	ToolAxe
	ToolShovel
	ToolHoe
	ToolSword
	ToolShears
)

// Has reports whether every family in tools is in the set.
func (families ToolFamilies) Has(tools ToolFamilies) bool {
	return families&tools == tools
}

// MiningEvidenceStatus describes how strongly a mining classification is known.
type MiningEvidenceStatus uint8

const (
	// MiningUnknown means the pinned source provided no usable classification.
	MiningUnknown MiningEvidenceStatus = iota
	// MiningProvisional means the classification comes from an older Bedrock release.
	MiningProvisional
)

// MiningFacts are provisional tool classifications for one block identifier.
// EffectiveTools includes tools inferred from material and tools that qualify
// for drops. HarvestTools only contains tools that qualify for drops. A
// HarvestLevel of -1 means no minimum tier is recorded. The facts are from
// an older Bedrock release and are not verified for this catalog's target.
type MiningFacts struct {
	EffectiveTools ToolFamilies
	HarvestTools   ToolFamilies
	HarvestLevel   int8
	Evidence       MiningEvidenceStatus
}

// MiningByName finds provisional mining facts for a block identifier. A false
// result means the pinned evidence has no usable tool classification; callers
// must not interpret it as confirmation that the hand can harvest the block.
func MiningByName(name string) (MiningFacts, bool) {
	index := sort.Search(len(miningRows), func(i int) bool { return miningRows[i].Name >= name })
	if index == len(miningRows) || miningRows[index].Name != name {
		return MiningFacts{}, false
	}
	facts := miningRows[index].Facts
	facts.Evidence = MiningProvisional
	return facts, true
}

type miningRow struct {
	Name  string
	Facts MiningFacts
}

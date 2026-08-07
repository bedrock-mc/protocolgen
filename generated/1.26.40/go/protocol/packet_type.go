// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PacketType uint32

const (
	PacketTypeEmpty                    PacketType = 0
	PacketTypeInitiallyUnlockedRecipes PacketType = 1
	PacketTypeNewlyUnlockedRecipes     PacketType = 2
	PacketTypeRemoveUnlockedRecipes    PacketType = 3
	PacketTypeRemoveAllUnlockedRecipes PacketType = 4
)

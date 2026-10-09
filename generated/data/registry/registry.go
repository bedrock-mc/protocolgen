// Package registry provides authenticated vanilla registry payloads in their consumer encodings.
package registry

import "bytes"

// BlockStatesNBT returns a private copy of concatenated NetworkLittleEndian NBT
// compounds with name, states, and version fields, in source palette order.
func BlockStatesNBT() []byte { return bytes.Clone(blockStates) }

// DataDrivenBlocksNBT returns a private copy of the LittleEndian NBT compound
// containing a blocks list of name and complete components compounds.
func DataDrivenBlocksNBT() []byte { return bytes.Clone(dataDrivenBlocks) }

// ItemsNBT returns a private copy of the NetworkLittleEndian NBT name dictionary.
// Each entry contains runtime_id, component_based, version, and complete data.
func ItemsNBT() []byte { return bytes.Clone(items) }

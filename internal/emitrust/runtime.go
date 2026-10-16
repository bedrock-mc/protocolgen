package emitrust

import (
	_ "embed"
	"strings"
)

// The wire runtime is hand-authored Rust emitted verbatim. It carries no
// protocol shape: everything version-specific is generated against it.
var (
	//go:embed runtime/wire.rs
	runtimeWire string
	//go:embed runtime/collections.rs
	runtimeCollections string
	//go:embed runtime/bytes.rs
	runtimeBytes string
	//go:embed runtime/nbt_scan.rs
	runtimeNbtScan string
	//go:embed runtime/nbt.rs
	runtimeNbt string
	//go:embed runtime/uuid.rs
	runtimeUUID string
	//go:embed runtime/glam.rs
	runtimeGlam string
	//go:embed runtime/pattern.rs
	runtimePattern string
	//go:embed runtime/tests.rs
	runtimeTests string
	//go:embed runtime/recursive.rs
	runtimeRecursive string
)

// The codec-contract runtime replaces every fragment whose signatures the
// contract changes; the NBT scanner is shared.
var (
	//go:embed runtime/contract/wire.rs
	contractWire string
	//go:embed runtime/contract/frame.rs
	contractFrame string
	//go:embed runtime/contract/nbt.rs
	contractNbt string
	//go:embed runtime/contract/uuid.rs
	contractUUID string
	//go:embed runtime/contract/glam.rs
	contractGlam string
	//go:embed runtime/contract/pattern.rs
	contractPattern string
	//go:embed runtime/contract/tests.rs
	contractTests string
)

// emitWire assembles the runtime, including only the fragments whose supporting
// dependency this manifest actually pulls in.
func emitWire(g *generator) string {
	if g.contract > 0 {
		return emitContractWire(g)
	}
	var b strings.Builder
	b.WriteString(runtimeWire)
	b.WriteString(runtimeRecursive)
	b.WriteString(runtimeCollections)
	// bytes is an unconditional runtime dependency: Reader::from_shared and
	// schema-bounded byte buffers use it even when this manifest has no public
	// byte-buffer field.
	b.WriteString(runtimeBytes)
	if g.usesNbt {
		b.WriteString(runtimeNbtScan)
		b.WriteString(runtimeNbt)
	}
	if g.usesUUID {
		b.WriteString(runtimeUUID)
	}
	if g.usesGlam {
		b.WriteString(runtimeGlam)
	}
	if g.usesPattern {
		b.WriteString(runtimePattern)
	}
	b.WriteString(runtimeTests)
	return strings.TrimSpace(b.String()) + "\n"
}

func emitContractWire(g *generator) string {
	var b strings.Builder
	b.WriteString(contractWire)
	b.WriteString(runtimeRecursive)
	b.WriteString(contractFrame)
	if g.usesNbt {
		b.WriteString(runtimeNbtScan)
		b.WriteString(contractNbt)
	}
	if g.usesUUID {
		b.WriteString(contractUUID)
	}
	if g.usesGlam {
		b.WriteString(contractGlam)
	}
	if g.usesPattern {
		b.WriteString(contractPattern)
	}
	b.WriteString(contractTests)
	return strings.TrimSpace(b.String()) + "\n"
}

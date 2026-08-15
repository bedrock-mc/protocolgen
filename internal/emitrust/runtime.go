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
)

// emitWire assembles the runtime, including only the fragments whose supporting
// dependency this manifest actually pulls in.
func emitWire(g *generator) string {
	var b strings.Builder
	b.WriteString(runtimeWire)
	b.WriteString(runtimeCollections)
	// bytes is an unconditional runtime dependency: Reader::from_shared and
	// schema-bounded byte buffers use it even when this manifest has no public
	// byte-buffer field.
	b.WriteString(runtimeBytes)
	if g.usesNbt {
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

package main

// Format is a code-generation strategy.
//
// The "gophertunnel" format maps doc fields onto gophertunnel's existing
// protocol.* types and io methods and applies gophertunnel conventions (it is
// meant to be diffed against / merged into the hand-written packets). The "raw"
// format generates a self-contained package: faithful structs for every packet
// AND every composite definition, marshalling against a generated IO interface,
// with no dependency on gophertunnel.
type Format interface {
	// Name returns the format identifier ("gophertunnel" or "raw").
	Name() string
	// EmitPacket renders a single packet to Go source, with any TODO notes.
	EmitPacket(pk genPacket, idConst string) (src string, todos []string, err error)
	// EmitExtra renders shared files keyed by filename. The gophertunnel format
	// returns nothing (it references existing protocol.* types); the raw format
	// returns the composite type definitions and the IO interface. It receives the
	// parsed packets too so it can stub any referenced-but-undefined type.
	EmitExtra(comps []composite, packets []genPacket) (files map[string]string, err error)
}

// gophertunnelFormat emits gophertunnel-flavoured packets (all the special
// mapping/preservation logic lives in resolve.go / emit.go / merge.go).
type gophertunnelFormat struct{}

func (gophertunnelFormat) Name() string { return "gophertunnel" }

func (gophertunnelFormat) EmitPacket(pk genPacket, idConst string) (string, []string, error) {
	return emitPacket(pk, idConst)
}

func (gophertunnelFormat) EmitExtra([]composite, []genPacket) (map[string]string, error) {
	return nil, nil
}

// selectFormat returns the Format for the given -format value.
func selectFormat(name, rawPkg string) (Format, bool) {
	switch name {
	case "gophertunnel":
		return gophertunnelFormat{}, true
	case "raw":
		return rawFormat{pkg: rawPkg}, true
	default:
		return nil, false
	}
}

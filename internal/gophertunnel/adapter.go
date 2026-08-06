// Package gophertunnel is a narrow adapter around the preserved go/ast
// extractor. It intentionally does not reverse-lower generated code into the
// canonical manifest.
package gophertunnel

import "protocolgen/internal/driftcheck"

type Packet struct {
	TypeName string
	ID       int
	Ops      []driftcheck.WireOp
}

func Extract(directory string) ([]Packet, error) {
	source, err := driftcheck.ExtractSourcePackets(directory)
	if err != nil {
		return nil, err
	}
	packets := make([]Packet, 0, len(source))
	for _, packet := range source {
		packets = append(packets, Packet{TypeName: packet.TypeName, ID: packet.ID, Ops: append([]driftcheck.WireOp(nil), packet.Ops...)})
	}
	return packets, nil
}

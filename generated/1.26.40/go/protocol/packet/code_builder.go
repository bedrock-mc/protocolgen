// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CodeBuilder struct {
	URL                   string
	ShouldOpenCodeBuilder bool
}

// Marshal reads or writes CodeBuilder using its canonical wire layout.
func (x *CodeBuilder) Marshal(io protocol.IO) {
	io.String(&x.URL)
	io.Bool(&x.ShouldOpenCodeBuilder)
}

// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// Packet is the common runtime contract for every generated Bedrock packet.
// Marshal reads from or writes to the supplied protocol IO implementation.
type Packet interface {
	ID() uint32
	Marshal(protocol.IO)
}

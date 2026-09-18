// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// CommandOutput is sent by the server to the client to send text as output of a command. Most servers do not
// use this packet and instead simply send Text packets, but there is reason to send it. If the origin of a
// CommandRequest packet is not the player itself, but, for example, a websocket server, sending a Text packet
// will not do what is expected: The message should go to the websocket server, not to the client's chat. The
// CommandOutput packet will make sure the messages are relayed to the correct origin of the command request.
type CommandOutput struct {
	// CommandOrigin is the data specifying the origin of the command. In other words, the source that the command
	// request was from, such as the player itself or a websocket server. The client forwards the messages in this
	// packet to the right origin, depending on what is sent here.
	OriginData protocol.CommandOriginData
	Output     protocol.CommandOutputData
}

// ID ...
func (*CommandOutput) ID() uint32 {
	return IDCommandOutput
}

func (pk *CommandOutput) Marshal(io protocol.IO) {
	pk.OriginData.Marshal(io)
	pk.Output.Marshal(io)
}

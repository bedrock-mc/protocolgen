// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetDisplayObjective struct {
	DisplaySlotName      string
	ObjectiveName        string
	ObjectiveDisplayName string
	CriteriaName         string
	SortOrder            int32
}

// Marshal reads or writes SetDisplayObjective using its canonical wire layout.
func (x *SetDisplayObjective) Marshal(io protocol.IO) {
	io.String(&x.DisplaySlotName)
	io.String(&x.ObjectiveName)
	io.String(&x.ObjectiveDisplayName)
	io.String(&x.CriteriaName)
	io.Varint32(&x.SortOrder)
}

// ID returns the protocol ID for SetDisplayObjective.
func (*SetDisplayObjective) ID() uint32 { return IDSetDisplayObjective }

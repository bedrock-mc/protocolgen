// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ActorDefinition struct {
	EventName string
}

func (*ActorDefinition) isEventData() {}

// Marshal reads or writes ActorDefinition using its canonical wire layout.
func (x *ActorDefinition) Marshal(io IO) {
	io.String(&x.EventName)
}

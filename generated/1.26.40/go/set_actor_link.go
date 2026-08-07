// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetActorLink struct {
	Link ActorLink
}

// Marshal reads or writes SetActorLink using its canonical wire layout.
func (x *SetActorLink) Marshal(io IO) {
	x.Link.Marshal(io)
}

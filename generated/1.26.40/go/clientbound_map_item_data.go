// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundMapItemData struct {
	MapID           ActorUniqueID
	Dimension       uint8
	IsLocked        bool
	MapOrigin       BlockPos
	CreationMapIDs  Optional[[]ActorUniqueID]
	Scale           Optional[int8]
	TrackedActorIDs Optional[[]MapItemTrackedActorUniqueId]
	Decorations     Optional[[]MapDecoration]
	Width           Optional[int32]
	Height          Optional[int32]
	StartX          Optional[int32]
	StartY          Optional[int32]
	Pixels          Optional[[]uint32]
}

// Marshal reads or writes ClientboundMapItemData using its canonical wire layout.
func (x *ClientboundMapItemData) Marshal(io IO) {
	x.MapID.Marshal(io)
	io.Uint8(&x.Dimension)
	io.Bool(&x.IsLocked)
	x.MapOrigin.Marshal(io)
	io.Bool(&x.CreationMapIDs.set)
	if x.CreationMapIDs.set {
		if !io.Reading() && uint64(len(x.CreationMapIDs.val)) > uint64(^uint32(0)) {
			io.InvalidValue(len(x.CreationMapIDs.val), "collection length overflows uint32")
			return
		}
		count1 := uint32(len(x.CreationMapIDs.val))
		io.Varuint32(&count1)
		if io.Reading() {
			if uint64(count1) > uint64(^uint(0)>>1) {
				io.InvalidValue(count1, "collection length overflows int")
				return
			}
			x.CreationMapIDs.val = make([]ActorUniqueID, int(count1))
		}
		for index2 := range x.CreationMapIDs.val {
			x.CreationMapIDs.val[index2].Marshal(io)
		}
	} else if io.Reading() {
		var zero []ActorUniqueID
		x.CreationMapIDs.val = zero
	}
	io.Bool(&x.Scale.set)
	if x.Scale.set {
		io.Int8(&x.Scale.val)
	} else if io.Reading() {
		var zero int8
		x.Scale.val = zero
	}
	io.Bool(&x.TrackedActorIDs.set)
	if x.TrackedActorIDs.set {
		if !io.Reading() && uint64(len(x.TrackedActorIDs.val)) > uint64(^uint32(0)) {
			io.InvalidValue(len(x.TrackedActorIDs.val), "collection length overflows uint32")
			return
		}
		count3 := uint32(len(x.TrackedActorIDs.val))
		io.Varuint32(&count3)
		if io.Reading() {
			if uint64(count3) > uint64(^uint(0)>>1) {
				io.InvalidValue(count3, "collection length overflows int")
				return
			}
			x.TrackedActorIDs.val = make([]MapItemTrackedActorUniqueId, int(count3))
		}
		for index4 := range x.TrackedActorIDs.val {
			x.TrackedActorIDs.val[index4].Marshal(io)
		}
	} else if io.Reading() {
		var zero []MapItemTrackedActorUniqueId
		x.TrackedActorIDs.val = zero
	}
	io.Bool(&x.Decorations.set)
	if x.Decorations.set {
		if !io.Reading() && uint64(len(x.Decorations.val)) > uint64(^uint32(0)) {
			io.InvalidValue(len(x.Decorations.val), "collection length overflows uint32")
			return
		}
		count5 := uint32(len(x.Decorations.val))
		io.Varuint32(&count5)
		if io.Reading() {
			if uint64(count5) > uint64(^uint(0)>>1) {
				io.InvalidValue(count5, "collection length overflows int")
				return
			}
			x.Decorations.val = make([]MapDecoration, int(count5))
		}
		for index6 := range x.Decorations.val {
			x.Decorations.val[index6].Marshal(io)
		}
	} else if io.Reading() {
		var zero []MapDecoration
		x.Decorations.val = zero
	}
	io.Bool(&x.Width.set)
	if x.Width.set {
		io.Varint32(&x.Width.val)
	} else if io.Reading() {
		var zero int32
		x.Width.val = zero
	}
	io.Bool(&x.Height.set)
	if x.Height.set {
		io.Varint32(&x.Height.val)
	} else if io.Reading() {
		var zero int32
		x.Height.val = zero
	}
	io.Bool(&x.StartX.set)
	if x.StartX.set {
		io.Varint32(&x.StartX.val)
	} else if io.Reading() {
		var zero int32
		x.StartX.val = zero
	}
	io.Bool(&x.StartY.set)
	if x.StartY.set {
		io.Varint32(&x.StartY.val)
	} else if io.Reading() {
		var zero int32
		x.StartY.val = zero
	}
	io.Bool(&x.Pixels.set)
	if x.Pixels.set {
		if !io.Reading() && uint64(len(x.Pixels.val)) > uint64(^uint32(0)) {
			io.InvalidValue(len(x.Pixels.val), "collection length overflows uint32")
			return
		}
		count7 := uint32(len(x.Pixels.val))
		io.Varuint32(&count7)
		if io.Reading() {
			if uint64(count7) > uint64(^uint(0)>>1) {
				io.InvalidValue(count7, "collection length overflows int")
				return
			}
			x.Pixels.val = make([]uint32, int(count7))
		}
		for index8 := range x.Pixels.val {
			io.Uint32(&x.Pixels.val[index8])
		}
	} else if io.Reading() {
		var zero []uint32
		x.Pixels.val = zero
	}
}

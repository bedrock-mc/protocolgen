// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SoundDataEvent interface {
	isSoundDataEvent()
}

func marshalSoundDataEvent(io IO, x *SoundDataEvent) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value SoundDataEventStop
				value.Marshal(io)
				*x = value
			case 1:
				var value SoundDataEventSetVolume
				value.Marshal(io)
				*x = value
			case 2:
				var value SoundDataEventSetPitch
				value.Marshal(io)
				*x = value
			case 3:
				var value SoundDataEventFade
				value.Marshal(io)
				*x = value
			case 4:
				var value SoundDataEventSeekTo
				value.Marshal(io)
				*x = value
			case 5:
				var value SoundDataEventPause
				value.Marshal(io)
				*x = value
			case 6:
				var value SoundDataEventResume
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case SoundDataEventStop:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case SoundDataEventSetVolume:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case SoundDataEventSetPitch:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case SoundDataEventFade:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case SoundDataEventSeekTo:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			case SoundDataEventPause:
				tag := uint32(5)
				io.Varuint32(&tag)
				value.Marshal(io)
			case SoundDataEventResume:
				tag := uint32(6)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

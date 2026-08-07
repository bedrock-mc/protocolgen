// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BookEditAction interface {
	isBookEditAction()
}

func marshalBookEditAction(io IO, x *BookEditAction) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				var value BookEditActionReplacePage
				value.Marshal(io)
				*x = value
			case 1:
				var value BookEditActionAddPage
				value.Marshal(io)
				*x = value
			case 2:
				var value BookEditActionDeletePage
				value.Marshal(io)
				*x = value
			case 3:
				var value BookEditActionSwapPages
				value.Marshal(io)
				*x = value
			case 4:
				var value BookEditActionFinalize
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case BookEditActionReplacePage:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case BookEditActionAddPage:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case BookEditActionDeletePage:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case BookEditActionSwapPages:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			case BookEditActionFinalize:
				tag := uint32(4)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

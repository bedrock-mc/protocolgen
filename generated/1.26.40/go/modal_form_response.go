// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ModalFormResponse struct {
	FormID           uint32
	JSONResponse     Optional[string]
	FormCancelReason Optional[ModalFormCancelReason]
}

// Marshal reads or writes ModalFormResponse using its canonical wire layout.
func (x *ModalFormResponse) Marshal(io IO) {
	io.Varuint32(&x.FormID)
	io.Bool(&x.JSONResponse.set)
	if x.JSONResponse.set {
		io.String(&x.JSONResponse.val)
	} else if io.Reading() {
		var zero string
		x.JSONResponse.val = zero
	}
	io.Bool(&x.FormCancelReason.set)
	if x.FormCancelReason.set {
		enumValue1 := uint8(x.FormCancelReason.val)
		io.Uint8(&enumValue1)
		x.FormCancelReason.val = ModalFormCancelReason(enumValue1)
		switch int64(enumValue1) {
		case 0, 1:
		default:
			io.InvalidValue(enumValue1, "unknown enum value")
		}
	} else if io.Reading() {
		var zero ModalFormCancelReason
		x.FormCancelReason.val = zero
	}
}

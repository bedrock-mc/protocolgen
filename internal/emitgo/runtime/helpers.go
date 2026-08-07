// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

// Marshaler is implemented by every generated struct and union payload.
type Marshaler interface {
	Marshal(IO)
}

// PtrMarshaler constrains a value whose pointer implements Marshaler.
type PtrMarshaler[T any] interface {
	Marshaler
	*T
}

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// IntegerFunc marshals a value through a wire integer type while preserving
// the public type used by the generated model.
func IntegerFunc[S, W integer](x *S, f func(*W)) {
	w := W(*x)
	f(&w)
	*x = S(w)
}

// OptionalFunc marshals a bool-prefixed optional value.
func OptionalFunc[T any](io IO, x *Optional[T], f func(*T)) {
	io.Bool(&x.set)
	if x.set {
		f(&x.val)
	}
}

// DoubleOptionalFunc marshals an optional nested inside an always-present
// outer optional, as used by several Cereal fields.
func DoubleOptionalFunc[T any](io IO, x *Optional[T], f func(*T)) {
	outer := true
	io.Bool(&outer)
	if outer {
		OptionalFunc(io, x, f)
	} else {
		*x = Optional[T]{}
	}
}

// OptionalMarshaler marshals an Optional whose value has a Marshal method.
func OptionalMarshaler[T any, A PtrMarshaler[T]](io IO, x *Optional[T]) {
	io.Bool(&x.set)
	if x.set {
		A(&x.val).Marshal(io)
	}
}

// UnionFunc dispatches a generated union codec without exposing the IO
// direction branch in every packet definition.
func UnionFunc(io IO, read, write func()) {
	if io.Reading() {
		read()
	} else {
		write()
	}
}

// sliceReader is implemented by decoders that allocate slices from a wire
// count. Writers intentionally do not implement it, so the same helpers work
// in both directions without exposing direction checks in packet methods.
type sliceReader interface {
	SliceLength(value uint64, max uint64) bool
}

const maxSliceLength = 4096

// FuncSlice marshals a length-prefixed slice using a count encoder and element
// callback. The count encoder determines the exact wire prefix type.
func FuncSlice[T any, C integer](io IO, x *[]T, count func(*C), f func(*T)) {
	if io.Reading() {
		reader, ok := io.(sliceReader)
		if !ok {
			io.InvalidValue(io, "reader does not implement SliceLength")
			return
		}
		var n C
		count(&n)
		length, valid := sliceLength(io, n)
		if !valid {
			return
		}
		if !reader.SliceLength(uint64(length), maxSliceLength) {
			return
		}
		*x = make([]T, length)
	} else {
		n, valid := sliceCount[C](io, len(*x))
		if !valid {
			return
		}
		count(&n)
	}
	for i := range *x {
		f(&(*x)[i])
	}
}

// OrderedMap marshals an ordered map representation while preserving duplicate
// keys and source order.
func OrderedMap[K, V any, C integer](io IO, x *[]OrderedEntry[K, V], count func(*C), key func(*K), value func(*V)) {
	if io.Reading() {
		reader, ok := io.(sliceReader)
		if !ok {
			io.InvalidValue(io, "reader does not implement SliceLength")
			return
		}
		var n C
		count(&n)
		length, valid := sliceLength(io, n)
		if !valid {
			return
		}
		if !reader.SliceLength(uint64(length), maxSliceLength) {
			return
		}
		*x = make([]OrderedEntry[K, V], length)
	} else {
		n, valid := sliceCount[C](io, len(*x))
		if !valid {
			return
		}
		count(&n)
	}
	for i := range *x {
		key(&(*x)[i].Key)
		value(&(*x)[i].Value)
	}
}

func sliceCount[C integer](io IO, length int) (C, bool) {
	var zero C
	max := ^uint64(0)
	switch zero := any(zero).(type) {
	case int:
		max = uint64(^uint(0) >> 1)
	case int8:
		max = uint64(^uint8(0) >> 1)
	case int16:
		max = uint64(^uint16(0) >> 1)
	case int32:
		max = uint64(^uint32(0) >> 1)
	case int64:
		max = uint64(^uint64(0) >> 1)
	case uint:
		max = uint64(^uint(0))
	case uint8:
		max = uint64(^uint8(0))
	case uint16:
		max = uint64(^uint16(0))
	case uint32:
		max = uint64(^uint32(0))
	case uint64:
		max = ^uint64(0)
	case uintptr:
		max = uint64(^uintptr(0))
	default:
		_ = zero
	}
	if uint64(length) > max {
		io.InvalidValue(length, "collection length exceeds its wire count")
		return zero, false
	}
	return C(length), true
}

func sliceLength[C integer](io IO, value C) (int, bool) {
	var n int64
	var unsigned uint64
	switch value := any(value).(type) {
	case int:
		n = int64(value)
	case int8:
		n = int64(value)
	case int16:
		n = int64(value)
	case int32:
		n = int64(value)
	case int64:
		n = value
	case uint:
		unsigned = uint64(value)
	case uint8:
		unsigned = uint64(value)
	case uint16:
		unsigned = uint64(value)
	case uint32:
		unsigned = uint64(value)
	case uint64:
		unsigned = value
	case uintptr:
		unsigned = uint64(value)
	}
	if n < 0 {
		io.InvalidValue(value, "negative collection length")
		return 0, false
	}
	if n > 0 {
		if uint64(n) > uint64(^uint(0)>>1) {
			io.InvalidValue(value, "collection length overflows int")
			return 0, false
		}
		return int(n), true
	}
	if unsigned > uint64(^uint(0)>>1) {
		io.InvalidValue(value, "collection length overflows int")
		return 0, false
	}
	return int(unsigned), true
}

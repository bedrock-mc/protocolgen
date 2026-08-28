// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import (
	"regexp"
	"sync"
)

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

type number interface {
	integer | ~float32 | ~float64
}

var schemaPatterns sync.Map

// Pattern validates a schema-published regular expression. Manifests validate
// patterns before generation; the cache avoids recompiling them per packet.
func Pattern(io IO, x *string, pattern string) {
	compiledValue, ok := schemaPatterns.Load(pattern)
	if !ok {
		compiled, err := regexp.Compile(pattern)
		if err != nil {
			io.InvalidValue(pattern, "invalid schema pattern")
			return
		}
		compiledValue, _ = schemaPatterns.LoadOrStore(pattern, compiled)
	}
	if !compiledValue.(*regexp.Regexp).MatchString(*x) {
		io.InvalidValue(*x, "string does not match schema pattern")
	}
}

// Minimum and Maximum validate schema-published numeric bounds after the wire
// operation, so the same generated call checks both encoded and decoded data.
func Minimum[T number](io IO, x *T, minimum T) {
	if *x < minimum {
		io.InvalidValue(*x, "value is below schema minimum")
	}
}

func Maximum[T number](io IO, x *T, maximum T) {
	if *x > maximum {
		io.InvalidValue(*x, "value exceeds schema maximum")
	}
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
	FuncSliceLimits(io, x, count, 0, ^uint64(0), f)
}

// FuncSliceLimits is FuncSlice with schema-published element-count bounds.
func FuncSliceLimits[T any, C integer](io IO, x *[]T, count func(*C), min, max uint64, f func(*T)) {
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
		if !schemaLength(io, uint64(length), min, max) {
			return
		}
		if !reader.SliceLength(uint64(length), maxSliceLength) {
			return
		}
		*x = make([]T, length)
	} else {
		if !schemaLength(io, uint64(len(*x)), min, max) {
			return
		}
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

// Slice marshals a varuint32-prefixed slice whose element pointer marshals itself.
func Slice[T any, A PtrMarshaler[T]](io IO, x *[]T) {
	SliceLimits[T, A](io, x, 0, ^uint64(0))
}

// SliceLimits is Slice with schema-published element-count bounds.
func SliceLimits[T any, A PtrMarshaler[T]](io IO, x *[]T, min, max uint64) {
	if io.Reading() {
		reader, ok := io.(sliceReader)
		if !ok {
			io.InvalidValue(io, "reader does not implement SliceLength")
			return
		}
		var n uint32
		io.Varuint32(&n)
		length, valid := sliceLength(io, n)
		if !valid || !schemaLength(io, uint64(length), min, max) || !reader.SliceLength(uint64(length), maxSliceLength) {
			return
		}
		*x = make([]T, length)
	} else {
		if !schemaLength(io, uint64(len(*x)), min, max) {
			return
		}
		n, valid := sliceCount[uint32](io, len(*x))
		if !valid {
			return
		}
		io.Varuint32(&n)
	}
	for i := range *x {
		A(&(*x)[i]).Marshal(io)
	}
}

// OrderedMap marshals an ordered map representation while preserving duplicate
// keys and source order.
func OrderedMap[K, V any, C integer](io IO, x *[]OrderedEntry[K, V], count func(*C), key func(*K), value func(*V)) {
	OrderedMapLimits(io, x, count, 0, ^uint64(0), key, value)
}

// OrderedMapLimits is OrderedMap with schema-published entry-count bounds.
func OrderedMapLimits[K, V any, C integer](io IO, x *[]OrderedEntry[K, V], count func(*C), min, max uint64, key func(*K), value func(*V)) {
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
		if !schemaLength(io, uint64(length), min, max) {
			return
		}
		if !reader.SliceLength(uint64(length), maxSliceLength) {
			return
		}
		*x = make([]OrderedEntry[K, V], length)
	} else {
		if !schemaLength(io, uint64(len(*x)), min, max) {
			return
		}
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

func schemaLength(io IO, value, min, max uint64) bool {
	if value < min || value > max {
		io.InvalidValue(value, "collection length outside schema limits")
		return false
	}
	return true
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

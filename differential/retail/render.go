package retail

import (
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

const maxRenderDepth = 256

// renderDigest digests a deterministic rendering of a decoded Go value:
// pointers followed, map keys sorted, floats by bit pattern and nil
// collections equal to empty ones.
func renderDigest(value any) string {
	var b strings.Builder
	render(&b, reflect.ValueOf(value), 0)
	return Digest([]byte(b.String()))
}

func render(b *strings.Builder, v reflect.Value, depth int) {
	if depth > maxRenderDepth {
		b.WriteString("<depth>")
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		b.WriteString("nil")
	case reflect.Pointer, reflect.Interface:
		if v.IsNil() {
			b.WriteString("nil")
			return
		}
		if v.Kind() == reflect.Interface {
			b.WriteString(v.Elem().Type().String())
		}
		render(b, v.Elem(), depth+1)
	case reflect.Struct:
		b.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			b.WriteString(v.Type().Field(i).Name)
			b.WriteByte(':')
			render(b, v.Field(i), depth+1)
			b.WriteByte(',')
		}
		b.WriteByte('}')
	case reflect.Slice, reflect.Array:
		b.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			render(b, v.Index(i), depth+1)
			b.WriteByte(',')
		}
		b.WriteByte(']')
	case reflect.Map:
		entries := make([]string, 0, v.Len())
		iter := v.MapRange()
		for iter.Next() {
			var entry strings.Builder
			render(&entry, iter.Key(), depth+1)
			entry.WriteByte('=')
			render(&entry, iter.Value(), depth+1)
			entries = append(entries, entry.String())
		}
		sort.Strings(entries)
		b.WriteString("map[" + strings.Join(entries, ",") + "]")
	case reflect.String:
		b.WriteString(strconv.Quote(v.String()))
	case reflect.Bool:
		b.WriteString(strconv.FormatBool(v.Bool()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		b.WriteString(strconv.FormatInt(v.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		b.WriteString(strconv.FormatUint(v.Uint(), 10))
	case reflect.Float32, reflect.Float64:
		b.WriteString("f" + strconv.FormatUint(math.Float64bits(v.Float()), 16))
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		b.WriteString("c" + strconv.FormatUint(math.Float64bits(real(c)), 16) + "," + strconv.FormatUint(math.Float64bits(imag(c)), 16))
	default:
		b.WriteString(v.Type().String())
	}
}

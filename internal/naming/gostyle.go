// Package-level Go casing rules shared by the Go emitter and overlay tooling.
package naming

import (
	"strings"
	"unicode"
)

// GoExportName converts a wire field or type token to an exported Go
// identifier with Go initialism casing applied.
func GoExportName(value string) string {
	var b strings.Builder
	upperNext := true
	for _, r := range value {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if upperNext {
				r = unicode.ToUpper(r)
				upperNext = false
			}
			b.WriteRune(r)
		} else {
			upperNext = true
		}
	}
	result := b.String()
	if result == "" {
		return "Generated"
	}
	if unicode.IsDigit([]rune(result)[0]) {
		return "Generated" + NormalizeGoInitialisms(result)
	}
	return NormalizeGoInitialisms(result)
}

var goInitialisms = map[string]string{
	"acl": "ACL", "api": "API", "argb": "ARGB", "ascii": "ASCII", "cpu": "CPU", "css": "CSS",
	"dns": "DNS", "eof": "EOF", "guid": "GUID", "gpu": "GPU", "html": "HTML", "http": "HTTP",
	"https": "HTTPS", "id": "ID", "ip": "IP", "json": "JSON", "nbt": "NBT", "osx": "OSX",
	"qps": "QPS", "ram": "RAM", "rgba": "RGBA", "rgb": "RGB", "rpc": "RPC", "sql": "SQL",
	"ssh": "SSH", "tcp": "TCP", "tls": "TLS", "tnt": "TNT", "ttl": "TTL", "udp": "UDP",
	"ui": "UI", "uid": "UID", "uint": "UINT", "uri": "URI", "url": "URL", "uuid": "UUID",
	"utf8": "UTF8", "uwp": "UWP", "vm": "VM", "xml": "XML", "xz": "XZ", "yaml": "YAML", "zip": "ZIP",
	"molang": "MoLang",
}

func NormalizeGoInitialisms(value string) string {
	words := goCamelWords(value)
	if len(words) == 0 {
		return value
	}
	var b strings.Builder
	for _, word := range words {
		if replacement, ok := goInitialisms[strings.ToLower(word)]; ok {
			b.WriteString(replacement)
		} else {
			b.WriteString(word)
		}
	}
	return b.String()
}

func goCamelWords(value string) []string {
	runes := []rune(value)
	if len(runes) == 0 {
		return nil
	}
	start := 0
	words := make([]string, 0, 4)
	for index := 1; index < len(runes); index++ {
		previous, current := runes[index-1], runes[index]
		nextLower := index+1 < len(runes) && unicode.IsLower(runes[index+1])
		boundary := unicode.IsUpper(current) && (unicode.IsLower(previous) || unicode.IsDigit(previous) || unicode.IsUpper(previous) && nextLower)
		if boundary {
			words = append(words, string(runes[start:index]))
			start = index
		}
	}
	return append(words, string(runes[start:]))
}

// IsExportedGoIdentifier reports whether value is a valid exported Go identifier.
func IsExportedGoIdentifier(value string) bool {
	if value == "" {
		return false
	}
	for index, r := range value {
		switch {
		case index == 0 && !unicode.IsUpper(r):
			return false
		case r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r):
		default:
			return false
		}
	}
	return true
}

// EnumVariantName is the Go constant suffix for a wire variant name:
// all-caps names are title-cased by token and initialisms kept.
func EnumVariantName(value string) string {
	if value == "" {
		return "Unknown"
	}
	allUpper := true
	for _, r := range value {
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			allUpper = false
			break
		}
	}
	if !allUpper {
		return normalizeEnumInitialisms(GoExportName(value))
	}
	var b strings.Builder
	for _, token := range strings.FieldsFunc(value, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) }) {
		if token == "" {
			continue
		}
		if initialism, ok := enumInitialisms[token]; ok {
			b.WriteString(initialism)
			continue
		}
		lower := strings.ToLower(token)
		b.WriteString(GoExportName(lower))
	}
	if b.Len() == 0 {
		return "Unknown"
	}
	return NormalizeGoInitialisms(b.String())
}

func normalizeEnumInitialisms(value string) string {
	for _, replacement := range []struct{ from, to string }{
		{from: "Tntcart", to: "TNTCart"},
		{from: "Fishpos", to: "FishPosition"},
		{from: "Hooktime", to: "HookTime"},
		{from: "Tnt", to: "TNT"},
		{from: "Nbt", to: "NBT"},
		{from: "Uuid", to: "UUID"},
		{from: "Argb", to: "ARGB"},
		{from: "Rgba", to: "RGBA"},
		{from: "Rgb", to: "RGB"},
		{from: "Uwp", to: "UWP"},
		{from: "Osx", to: "OSX"},
	} {
		value = strings.ReplaceAll(value, replacement.from, replacement.to)
	}
	return value
}

var enumInitialisms = map[string]string{
	"ANIM": "Animation", "FISHPOS": "FishPosition", "HOOKTIME": "HookTime", "ID": "ID", "NBT": "NBT", "OSX": "OSX", "RGBA": "RGBA", "RGB": "RGB", "TNT": "TNT", "TNTCART": "TNTCart", "UWP": "UWP", "URL": "URL", "URI": "URI", "UUID": "UUID", "X": "X", "Y": "Y", "Z": "Z",
}

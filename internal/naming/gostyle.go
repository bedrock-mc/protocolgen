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

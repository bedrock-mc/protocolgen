package main

import (
	"strings"
	"unicode"
)

// goFieldName converts a documented field name into an exported Go identifier,
// e.g. "Swing Source" -> "SwingSource", "Target Actor Runtime ID" -> "TargetActorRuntimeID".
func goFieldName(docName string) string {
	cleaned := strings.ReplaceAll(docName, "'", "")
	tokens := strings.FieldsFunc(cleaned, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	var b strings.Builder
	for _, tok := range tokens {
		runes := []rune(tok)
		runes[0] = unicode.ToUpper(runes[0])
		b.WriteString(string(runes))
	}
	out := b.String()
	if out == "" {
		return "Field"
	}
	if unicode.IsDigit([]rune(out)[0]) {
		return "Field" + out
	}
	return out
}

// goTypeName converts a packet title into its Go type name, stripping the
// trailing "Packet" suffix, e.g. "AnimatePacket" -> "Animate".
func goTypeName(title string) string {
	return strings.TrimSuffix(title, "Packet")
}

// fileName converts a Go type name into gophertunnel's snake_case file name,
// e.g. "StartGame" -> "start_game.go".
func fileName(typeName string) string {
	return toSnake(typeName) + ".go"
}

// cleanCompositeTitle turns a Mojang composite definition title into a valid,
// exported Go identifier: it drops C++ namespace qualifiers (taking the segment
// after the last "::") and strips spaces/punctuation, e.g. "mce::UUID" -> "UUID",
// "MapItemTrackedActor UniqueId" -> "MapItemTrackedActorUniqueId".
func cleanCompositeTitle(title string) string {
	if i := strings.LastIndex(title, "::"); i >= 0 {
		title = title[i+2:]
	}
	return goFieldName(title)
}

// toSnake converts a PascalCase/camelCase identifier into snake_case.
func toSnake(s string) string {
	var b strings.Builder
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			// Insert an underscore at a lower->upper boundary, or before the last
			// letter of an acronym run that starts a new word (e.g. "IDFoo").
			prevLower := i > 0 && unicode.IsLower(runes[i-1])
			nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
			if i > 0 && (prevLower || nextLower) {
				b.WriteByte('_')
			}
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

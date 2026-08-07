package gophertunneloracle

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"protocolgen/internal/manifest"
)

func CompareFile(options Options) (Report, error) {
	canonical, err := manifest.Load(options.ManifestPath)
	if err != nil {
		return Report{}, err
	}
	lock, err := LoadLock(options.LockPath)
	if err != nil {
		return Report{}, err
	}
	accepted, err := LoadAccepted(options.AcceptedPath)
	if err != nil {
		return Report{}, err
	}
	if lock.ProtocolVersion != canonical.Target.ProtocolVersion || lock.MinecraftVersion != canonical.Target.MinecraftVersion {
		return Report{}, fmt.Errorf("gophertunnel lock targets Minecraft %s/protocol %d, manifest targets %s/%d", lock.MinecraftVersion, lock.ProtocolVersion, canonical.Target.MinecraftVersion, canonical.Target.ProtocolVersion)
	}
	if accepted.ProtocolVersion != canonical.Target.ProtocolVersion || accepted.MinecraftVersion != canonical.Target.MinecraftVersion {
		return Report{}, fmt.Errorf("accepted-divergences targets Minecraft %s/protocol %d, manifest targets %s/%d", accepted.MinecraftVersion, accepted.ProtocolVersion, canonical.Target.MinecraftVersion, canonical.Target.ProtocolVersion)
	}
	checkout, err := resolveCheckout(lock, options.GophertunnelPath, options.CacheDir)
	if err != nil {
		return Report{}, err
	}
	extracted, err := Extract(checkout)
	if err != nil {
		return Report{}, err
	}
	report := Compare(canonical, extracted, lock, accepted, options.ManifestPath)
	if options.ReportPath != "" {
		if err := WriteReport(options.ReportPath, report); err != nil {
			return report, err
		}
	}
	if options.FailOnUnaccepted && len(report.Unaccepted) > 0 {
		return report, fmt.Errorf("%d unaccepted gophertunnel wire divergence(s): %s", len(report.Unaccepted), joinIDs(report.Unaccepted))
	}
	return report, nil
}

// Compare is the pure comparison layer. Extraction and checkout are kept
// outside it so normalization rules can be tested with tiny synthetic trees.
func Compare(canonical manifest.Manifest, source extraction, lock Lock, accepted AcceptedFile, manifestPath string) Report {
	report := Report{
		SchemaVersion:         ReportSchemaVersion,
		MinecraftVersion:      canonical.Target.MinecraftVersion,
		ProtocolVersion:       canonical.Target.ProtocolVersion,
		Manifest:              manifestPath,
		Gophertunnel:          OracleSource{Repo: lock.Gophertunnel.Repo, Commit: lock.Gophertunnel.Commit},
		Normalization:         defaultNormalization(),
		UnresolvedDiagnostics: source.Diagnostics,
	}
	acceptedByID := make(map[uint32]AcceptedDivergence, len(accepted.Divergences))
	for _, entry := range accepted.Divergences {
		acceptedByID[entry.ID] = entry
	}
	sourceByID := make(map[uint32]sourcePacket, len(source.Packets))
	for _, packet := range source.Packets {
		sourceByID[packet.ID] = packet
	}
	canonicalIDs := make(map[uint32]bool, len(canonical.Packets))
	for _, packet := range canonical.Packets {
		canonicalIDs[packet.ID] = true
		result := PacketResult{ID: packet.ID, Name: packet.Name}
		oracle, ok := sourceByID[packet.ID]
		if !ok {
			result.Classification = "NO_ORACLE_PACKET"
			report.Counts.NoOraclePacket++
			result.Reasons = []string{"gophertunnel has no packet with this ID"}
			report.Packets = append(report.Packets, result)
			continue
		}
		result.GophertunnelName = oracle.Name
		want, wantReasons := canonicalAtoms(packet)
		got, gotReasons := sourceAtoms(oracle)
		reasons := append([]string{}, wantReasons...)
		reasons = append(reasons, gotReasons...)
		if len(reasons) > 0 {
			result.Classification = "UNRESOLVED"
			result.Reasons = uniqueStrings(reasons)
			report.Counts.Unresolved++
			report.Packets = append(report.Packets, result)
			continue
		}
		if atomsEqual(want, got) {
			result.Classification = "AGREEMENT"
			result.OperationCount = len(want)
			report.Counts.Agreement++
			report.Packets = append(report.Packets, result)
			continue
		}
		result.Classification = "DIVERGENCE"
		result.ManifestSequence = atomDisplays(want)
		result.GophertunnelSequence = atomDisplays(got)
		result.Differences = differences(want, got)
		report.Counts.Divergence++
		if _, accepted := acceptedByID[packet.ID]; accepted {
			report.Accepted = append(report.Accepted, packet.ID)
		} else {
			report.Unaccepted = append(report.Unaccepted, packet.ID)
		}
		report.Packets = append(report.Packets, result)
	}
	for _, packet := range source.Packets {
		if !canonicalIDs[packet.ID] {
			report.OracleOnly = append(report.OracleOnly, PacketIdentity{ID: packet.ID, Name: packet.Name})
		}
	}
	for _, entry := range accepted.Divergences {
		if !containsID(report.Accepted, entry.ID) {
			report.ResolvedAccepted = append(report.ResolvedAccepted, entry.ID)
		}
	}
	sort.Slice(report.Packets, func(i, j int) bool { return report.Packets[i].ID < report.Packets[j].ID })
	sort.Slice(report.OracleOnly, func(i, j int) bool { return report.OracleOnly[i].ID < report.OracleOnly[j].ID })
	sortIDs(report.Accepted)
	sortIDs(report.Unaccepted)
	sortIDs(report.ResolvedAccepted)
	return report
}

func defaultNormalization() Normalization {
	return Normalization{
		FixedWidth:    "Signed and unsigned fixed-width integers with identical width and endianness are equivalent.",
		Strings:       "A length-prefixed UTF-8 string and a length-prefixed byte slice with the same prefix are wire-equivalent.",
		ByteArrays:    "A prefixed array of single u8 elements is equivalent to a byte slice with the same prefix.",
		UUID:          "UUID is compared as 16 bytes at its wire position; gophertunnel's internal UUID byte ordering is intentionally not validated.",
		PreencodedNBT: "RawBytes named SerialisedOffers, SerialisedInventoryData, SerialisedEntityIdentifiers, or SerialisedEventData normalize to nbt_le.",
		Preserved: []string{
			"integer width",
			"endianness",
			"fixed-width versus varint",
			"varint versus zigzag",
			"float versus integer",
			"option presence",
			"array prefix type",
			"fixed-array length",
			"union control and variant discriminants",
		},
	}
}

func canonicalAtoms(packet manifest.Packet) ([]atom, []string) {
	var result []atom
	var reasons []string
	fields := append([]manifest.Field(nil), packet.Fields...)
	sort.SliceStable(fields, func(i, j int) bool { return fields[i].Ordinal < fields[j].Ordinal })
	for _, field := range fields {
		atoms, unresolved := manifestNodeAtoms(field.Name, field.Encode)
		result = append(result, atoms...)
		reasons = append(reasons, unresolved...)
	}
	return result, uniqueStrings(reasons)
}

func manifestNodeAtoms(path string, node manifest.Node) ([]atom, []string) {
	switch node.Kind {
	case manifest.KindVoid:
		return nil, nil
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return nil, []string{"manifest: primitive at " + path + " has no shape"}
		}
		if node.Primitive.Code == "uuid" {
			return []atom{{Token: "UUID16", Field: path, Display: "uuid(16 bytes)"}}, nil
		}
		return []atom{{Token: "P:" + canonicalPrimitive(node.Primitive.Code), Field: path, Display: node.Primitive.Code}}, nil
	case manifest.KindEnum:
		if node.Primitive == nil {
			return nil, []string{"manifest: enum at " + path + " has no underlying shape"}
		}
		return []atom{{Token: "P:" + canonicalPrimitive(node.Primitive.Code), Field: path, Display: "enum(" + node.Primitive.Code + ")"}}, nil
	case manifest.KindString, manifest.KindBytes:
		prefix, err := manifestPrefix(node.Prefix)
		if err != nil {
			return nil, []string{"manifest: " + path + ": " + err.Error()}
		}
		kind := "string"
		if node.Kind == manifest.KindBytes {
			kind = "bytes"
		}
		return []atom{{Token: "LEN:" + canonicalPrimitive(prefix), Field: path, Display: kind + "(prefix=" + prefix + ")"}}, nil
	case manifest.KindArray:
		prefix, err := manifestPrefix(node.Prefix)
		if err != nil || node.Element == nil {
			if err == nil {
				err = fmt.Errorf("array has no element")
			}
			return nil, []string{"manifest: " + path + ": " + err.Error()}
		}
		if isManifestU8(*node.Element) {
			return []atom{{Token: "LEN:" + canonicalPrimitive(prefix), Field: path, Display: "byte-array(prefix=" + prefix + ")"}}, nil
		}
		children, reasons := manifestNodeAtoms(path+"[]", *node.Element)
		result := []atom{{Token: "ARRAY:" + canonicalPrimitive(prefix), Field: path, Display: "array(prefix=" + prefix + ")"}}
		result = append(result, children...)
		result = append(result, atom{Token: "/ARRAY", Field: path, Display: "/array"})
		return result, reasons
	case manifest.KindFixedArray:
		if node.Element == nil || node.Length == 0 {
			return nil, []string{"manifest: fixed array at " + path + " is incomplete"}
		}
		if node.Length == 16 && isManifestU8(*node.Element) {
			return []atom{{Token: "UUID16", Field: path, Display: "uuid(16 bytes)"}}, nil
		}
		children, reasons := manifestNodeAtoms(path+"[]", *node.Element)
		result := []atom{{Token: fmt.Sprintf("FIXED:%d", node.Length), Field: path, Display: fmt.Sprintf("fixed-array(length=%d)", node.Length)}}
		result = append(result, children...)
		result = append(result, atom{Token: "/FIXED", Field: path, Display: "/fixed-array"})
		return result, reasons
	case manifest.KindSequence:
		var result []atom
		var reasons []string
		for index, child := range node.Elements {
			atoms, childReasons := manifestNodeAtoms(fmt.Sprintf("%s[%d]", path, index), child)
			result = append(result, atoms...)
			reasons = append(reasons, childReasons...)
		}
		return result, reasons
	case manifest.KindOptional:
		if node.Value == nil {
			return nil, []string{"manifest: optional at " + path + " has no value"}
		}
		children, reasons := manifestNodeAtoms(path, *node.Value)
		result := []atom{{Token: "OPTION:bool", Field: path, Display: "option(presence=bool)"}}
		result = append(result, children...)
		result = append(result, atom{Token: "/OPTION", Field: path, Display: "/option"})
		return result, reasons
	case manifest.KindStruct:
		fields := append([]manifest.Field(nil), node.Fields...)
		sort.SliceStable(fields, func(i, j int) bool { return fields[i].Ordinal < fields[j].Ordinal })
		var result []atom
		var reasons []string
		for _, field := range fields {
			atoms, childReasons := manifestNodeAtoms(path+"."+field.Name, field.Encode)
			result = append(result, atoms...)
			reasons = append(reasons, childReasons...)
		}
		return result, reasons
	case manifest.KindMap:
		if node.Prefix == nil || node.Key == nil || node.Value == nil {
			return nil, []string{"manifest: map at " + path + " is incomplete"}
		}
		prefix, err := manifestPrefix(node.Prefix)
		if err != nil {
			return nil, []string{"manifest: map at " + path + ": " + err.Error()}
		}
		keyAtoms, keyReasons := manifestNodeAtoms(path+".<key>", *node.Key)
		valueAtoms, valueReasons := manifestNodeAtoms(path+".<value>", *node.Value)
		result := []atom{{Token: "ARRAY:" + canonicalPrimitive(prefix), Field: path, Display: "map(prefix=" + prefix + ")"}}
		result = append(result, keyAtoms...)
		result = append(result, valueAtoms...)
		result = append(result, atom{Token: "/ARRAY", Field: path, Display: "/map"})
		return result, append(keyReasons, valueReasons...)
	case manifest.KindUnion:
		control, err := manifestPrimitive(node.Control)
		if err != nil {
			return nil, []string{"manifest: union at " + path + ": " + err.Error()}
		}
		variants := append([]manifest.Variant(nil), node.Variants...)
		sort.SliceStable(variants, func(i, j int) bool { return variants[i].Value < variants[j].Value })
		result := []atom{{Token: "UNION:" + canonicalPrimitive(control), Field: path, Display: "union(control=" + control + ")"}}
		var reasons []string
		for _, variant := range variants {
			children, childReasons := manifestNodeAtoms(path+".variant", variant.Encode)
			result = append(result, atom{Token: fmt.Sprintf("VARIANT:%d", variant.Value), Field: path, Display: fmt.Sprintf("variant(%d)", variant.Value)})
			result = append(result, children...)
			result = append(result, atom{Token: "/VARIANT", Field: path, Display: "/variant"})
			reasons = append(reasons, childReasons...)
		}
		result = append(result, atom{Token: "/UNION", Field: path, Display: "/union"})
		return result, reasons
	case manifest.KindBitset:
		if node.Length == 0 {
			return nil, []string{"manifest: bitset at " + path + " has no length"}
		}
		return []atom{{Token: fmt.Sprintf("BITSET:%d", node.Length), Field: path, Display: fmt.Sprintf("bitset(length=%d)", node.Length)}}, nil
	case manifest.KindReserved, manifest.KindIgnored:
		if node.Element == nil {
			return nil, []string{"manifest: compatibility node at " + path + " has no element"}
		}
		return manifestNodeAtoms(path, *node.Element)
	case manifest.KindConditional:
		return nil, []string{"manifest: conditional at " + path + " requires runtime branch evidence"}
	case manifest.KindRecursive:
		return nil, []string{"manifest: recursive node at " + path + " is not statically finite"}
	case manifest.KindOpaque, manifest.KindUnresolved:
		return nil, []string{"manifest: " + string(node.Kind) + " at " + path + ": " + node.Reason}
	default:
		return nil, []string{"manifest: unsupported node " + string(node.Kind) + " at " + path}
	}
}

func sourceAtoms(packet sourcePacket) ([]atom, []string) {
	var result []atom
	var reasons []string
	for _, operation := range packet.Operations {
		atoms, unresolved := sourceOperationAtoms(operation)
		result = append(result, atoms...)
		reasons = append(reasons, unresolved...)
	}
	return result, uniqueStrings(reasons)
}

func sourceOperationAtoms(operation sourceOperation) ([]atom, []string) {
	path := operation.Field
	switch operation.Kind {
	case "primitive":
		code := operation.Code
		if code == "raw_bytes" {
			if isPreencodedNBTField(path) {
				code = "nbt_le"
			}
		}
		return []atom{{Token: "P:" + canonicalPrimitive(code), Field: path, Display: code}}, nil
	case "string", "bytes":
		return []atom{{Token: "LEN:" + canonicalPrimitive(operation.Prefix), Field: path, Display: operation.Kind + "(prefix=" + operation.Prefix + ")"}}, nil
	case "uuid":
		return []atom{{Token: "UUID16", Field: path, Display: "uuid(16 bytes)"}}, nil
	case "bitset":
		if operation.Length == 0 {
			return nil, []string{"gophertunnel: bitset at " + path + " has no static length"}
		}
		return []atom{{Token: fmt.Sprintf("BITSET:%d", operation.Length), Field: path, Display: fmt.Sprintf("bitset(length=%d)", operation.Length)}}, nil
	case "array":
		if isSourceU8(operation.Element) {
			return []atom{{Token: "LEN:" + canonicalPrimitive(operation.Prefix), Field: path, Display: "byte-array(prefix=" + operation.Prefix + ")"}}, nil
		}
		var result []atom
		var reasons []string
		result = append(result, atom{Token: "ARRAY:" + canonicalPrimitive(operation.Prefix), Field: path, Display: "array(prefix=" + operation.Prefix + ")"})
		for _, child := range operation.Element {
			atoms, childReasons := sourceOperationAtoms(child)
			result = append(result, atoms...)
			reasons = append(reasons, childReasons...)
		}
		result = append(result, atom{Token: "/ARRAY", Field: path, Display: "/array"})
		return result, reasons
	case "fixed_array":
		if operation.Length == 0 {
			return nil, []string{"gophertunnel: fixed array at " + path + " has no length"}
		}
		if operation.Length == 16 && isSourceU8(operation.Element) {
			return []atom{{Token: "UUID16", Field: path, Display: "uuid(16 bytes)"}}, nil
		}
		var result []atom
		var reasons []string
		result = append(result, atom{Token: fmt.Sprintf("FIXED:%d", operation.Length), Field: path, Display: fmt.Sprintf("fixed-array(length=%d)", operation.Length)})
		for _, child := range operation.Element {
			atoms, childReasons := sourceOperationAtoms(child)
			result = append(result, atoms...)
			reasons = append(reasons, childReasons...)
		}
		result = append(result, atom{Token: "/FIXED", Field: path, Display: "/fixed-array"})
		return result, reasons
	case "optional":
		var result []atom
		var reasons []string
		result = append(result, atom{Token: "OPTION:" + canonicalPrimitive(operation.Presence), Field: path, Display: "option(presence=" + operation.Presence + ")"})
		for _, child := range operation.Value {
			atoms, childReasons := sourceOperationAtoms(child)
			result = append(result, atoms...)
			reasons = append(reasons, childReasons...)
		}
		result = append(result, atom{Token: "/OPTION", Field: path, Display: "/option"})
		return result, reasons
	case "union":
		result := []atom{{Token: "UNION:" + canonicalPrimitive(operation.Control), Field: path, Display: "union(control=" + operation.Control + ")"}}
		variants := append([]sourceVariant(nil), operation.Variants...)
		sort.SliceStable(variants, func(i, j int) bool { return variants[i].Value < variants[j].Value })
		var reasons []string
		for _, variant := range variants {
			result = append(result, atom{Token: fmt.Sprintf("VARIANT:%d", variant.Value), Field: path, Display: fmt.Sprintf("variant(%d)", variant.Value)})
			for _, child := range variant.Ops {
				atoms, childReasons := sourceOperationAtoms(child)
				result = append(result, atoms...)
				reasons = append(reasons, childReasons...)
			}
			result = append(result, atom{Token: "/VARIANT", Field: path, Display: "/variant"})
		}
		result = append(result, atom{Token: "/UNION", Field: path, Display: "/union"})
		return result, reasons
	case "unresolved", "recursive":
		reason := operation.Reason
		if reason == "" {
			reason = "operation is not statically resolvable"
		}
		where := path
		if operation.Site != "" {
			where += " (" + operation.Site + ")"
		}
		return nil, []string{"gophertunnel: " + reason + " at " + where}
	default:
		return nil, []string{"gophertunnel: unsupported operation " + operation.Kind + " at " + path}
	}
}

func manifestPrefix(node *manifest.Node) (string, error) {
	if node == nil || node.Kind != manifest.KindPrimitive || node.Primitive == nil {
		return "", fmt.Errorf("length prefix is not an explicit primitive")
	}
	return node.Primitive.Code, nil
}

func manifestPrimitive(node *manifest.Node) (string, error) {
	if node == nil || node.Kind != manifest.KindPrimitive || node.Primitive == nil {
		return "", fmt.Errorf("union control is not an explicit primitive")
	}
	return node.Primitive.Code, nil
}

func canonicalPrimitive(code string) string {
	switch code {
	case "i8", "u8":
		return "FIXED8"
	case "i16le", "u16le":
		return "FIXED16LE"
	case "i16be", "u16be":
		return "FIXED16BE"
	case "i32le", "u32le":
		return "FIXED32LE"
	case "i32be", "u32be":
		return "FIXED32BE"
	case "i64le", "u64le":
		return "FIXED64LE"
	case "i64be", "u64be":
		return "FIXED64BE"
	default:
		return code
	}
}

func isManifestU8(node manifest.Node) bool {
	return node.Kind == manifest.KindPrimitive && node.Primitive != nil && node.Primitive.Code == "u8"
}

func isSourceU8(operations []sourceOperation) bool {
	return len(operations) == 1 && operations[0].Kind == "primitive" && operations[0].Code == "u8"
}

func isPreencodedNBTField(field string) bool {
	for _, suffix := range []string{"SerialisedOffers", "SerialisedInventoryData", "SerialisedEntityIdentifiers", "SerialisedEventData"} {
		if strings.HasSuffix(field, suffix) {
			return true
		}
	}
	return false
}

func atomsEqual(left, right []atom) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index].Token != right[index].Token {
			return false
		}
	}
	return true
}

func atomDisplays(atoms []atom) []string {
	result := make([]string, len(atoms))
	for index, current := range atoms {
		result[index] = current.Display
	}
	return result
}

func differences(left, right []atom) []Difference {
	n, m := len(left), len(right)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := n; i >= 0; i-- {
		for j := m; j >= 0; j-- {
			switch {
			case i == n:
				dp[i][j] = m - j
			case j == m:
				dp[i][j] = n - i
			case left[i].Token == right[j].Token:
				dp[i][j] = dp[i+1][j+1]
			default:
				dp[i][j] = minInt(1+dp[i+1][j+1], 1+dp[i+1][j], 1+dp[i][j+1])
			}
		}
	}
	var result []Difference
	i, j, position := 0, 0, 0
	for i < n || j < m {
		if i < n && j < m && left[i].Token == right[j].Token {
			i++
			j++
			position++
			continue
		}
		if i < n && j < m && dp[i][j] == 1+dp[i+1][j+1] {
			result = append(result, Difference{Position: position, Manifest: left[i].Display, ManifestField: left[i].Field, Gophertunnel: right[j].Display, GophertunnelField: right[j].Field})
			i++
			j++
			position++
			continue
		}
		if i < n && dp[i][j] == 1+dp[i+1][j] {
			result = append(result, Difference{Position: position, Manifest: left[i].Display, ManifestField: left[i].Field, Gophertunnel: "missing"})
			i++
			position++
			continue
		}
		result = append(result, Difference{Position: position, Manifest: "missing", Gophertunnel: right[j].Display, GophertunnelField: right[j].Field})
		j++
		position++
	}
	return result
}

func minInt(values ...int) int {
	result := values[0]
	for _, value := range values[1:] {
		if value < result {
			result = value
		}
	}
	return result
}

func uniqueStrings(input []string) []string {
	seen := map[string]bool{}
	result := make([]string, 0, len(input))
	for _, value := range input {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	return result
}

func containsID(ids []uint32, wanted uint32) bool {
	for _, id := range ids {
		if id == wanted {
			return true
		}
	}
	return false
}

func joinIDs(ids []uint32) string {
	parts := make([]string, len(ids))
	for index, id := range ids {
		parts[index] = fmt.Sprintf("%d", id)
	}
	return strings.Join(parts, ", ")
}

func WriteReport(path string, report Report) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal gophertunnel report: %w", err)
	}
	data = append(data, '\n')
	if parent := filepath.Dir(path); parent != "." && parent != "" {
		if err := os.MkdirAll(parent, 0o755); err != nil {
			return fmt.Errorf("create gophertunnel report directory: %w", err)
		}
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write gophertunnel report: %w", err)
	}
	return nil
}

func Summary(report Report, reportPath string) string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "gophertunnel oracle: %s @ %s\n", report.Gophertunnel.Repo, report.Gophertunnel.Commit)
	fmt.Fprintf(&builder, "Compared %d manifest packets for Minecraft %s / protocol %d\n", len(report.Packets), report.MinecraftVersion, report.ProtocolVersion)
	fmt.Fprintf(&builder, "AGREEMENT: %d  DIVERGENCE: %d  UNRESOLVED: %d  NO_ORACLE_PACKET: %d\n", report.Counts.Agreement, report.Counts.Divergence, report.Counts.Unresolved, report.Counts.NoOraclePacket)
	fmt.Fprintf(&builder, "accepted divergences: %d  unaccepted divergences: %d  resolved accepted entries: %d\n", len(report.Accepted), len(report.Unaccepted), len(report.ResolvedAccepted))
	if len(report.Unaccepted) > 0 {
		fmt.Fprintf(&builder, "unaccepted packet IDs: %s\n", joinIDs(report.Unaccepted))
	}
	if len(report.ResolvedAccepted) > 0 {
		fmt.Fprintf(&builder, "review baseline entries no longer diverging: %s\n", joinIDs(report.ResolvedAccepted))
	}
	if len(report.OracleOnly) > 0 {
		fmt.Fprintf(&builder, "oracle-only packets: %d\n", len(report.OracleOnly))
	}
	if reportPath != "" {
		fmt.Fprintf(&builder, "JSON report: %s\n", reportPath)
	}
	return builder.String()
}

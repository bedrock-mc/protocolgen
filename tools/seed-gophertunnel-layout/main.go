// Command seed-gophertunnel-layout derives a layout overlay from a gophertunnel
// checkout and writes a gap report: which fork types, fields, and constants
// the generated tree already covers, and which it does not.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"protocolgen/internal/docs"
	"protocolgen/internal/flatten"
	"protocolgen/internal/layout"
	"protocolgen/internal/manifest"
	"protocolgen/internal/naming"
)

func main() {
	manifestPath := flag.String("manifest", "", "canonical manifest")
	namingPath := flag.String("naming", "", "reviewed naming overlay")
	gopherPath := flag.String("gophertunnel", "", "gophertunnel checkout")
	outPath := flag.String("out", "", "layout overlay output")
	reportPath := flag.String("report", "", "gap report output (Markdown)")
	docsPath := flag.String("docs", "", "reviewed docs overlay to extend with fork comments for matched types and fields (optional)")
	flag.Parse()
	if *manifestPath == "" || *namingPath == "" || *gopherPath == "" || *outPath == "" || *reportPath == "" {
		fail("-manifest, -naming, -gophertunnel, -out and -report are required")
	}
	m, err := manifest.Load(*manifestPath)
	if err != nil {
		fail("%v", err)
	}
	overlay, err := naming.LoadOverlay(*namingPath, m)
	if err != nil {
		fail("%v", err)
	}
	index, err := buildIndex(m, overlay)
	if err != nil {
		fail("%v", err)
	}
	fork, err := parseFork(*gopherPath)
	if err != nil {
		fail("%v", err)
	}
	var docOverlay docs.Overlay
	if *docsPath != "" {
		docOverlay, err = docs.LoadOverlay(*docsPath, m)
		if err != nil {
			fail("%v", err)
		}
	}
	document, report := seed(m, index, fork, docOverlay)
	document.Target = m.Target
	if *docsPath != "" {
		if err := writeJSON(*docsPath, docs.Document{SchemaVersion: 1, Target: m.Target, Entries: docs.SortedEntries(docOverlay)}); err != nil {
			fail("%v", err)
		}
		fmt.Printf("docs: %d fork comments ported -> %s\n", report.portedDocs, *docsPath)
	}
	if err := layout.ValidateOverlay(m, document); err != nil {
		fail("seeded overlay is invalid: %v", err)
	}
	if err := writeJSON(*outPath, layout.SortedDocument(document)); err != nil {
		fail("%v", err)
	}
	if err := os.WriteFile(*reportPath, []byte(report.render(m, fork)), 0o644); err != nil {
		fail("%v", err)
	}
	fmt.Printf("layout: %d constant placements (%d variants named), %d field names -> %s\n", len(document.Constants), report.namedVariants, len(document.Fields), *outPath)
	fmt.Printf("gaps: %d fork types unmatched, %d generated types unmatched, %d const groups unmatched -> %s\n", len(report.forkOnlyTypes), len(report.generatedOnlyTypes), len(report.unmatchedGroups), *reportPath)
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

// --- manifest side ---------------------------------------------------------

type enumInfo struct {
	TypeID   string
	Name     string
	Variants []manifest.Variant
}

type ownerInfo struct {
	TypeID  string // type ID that owns the fields: the inlined payload struct for a flattened packet
	FileKey string // key for file placement: the packet name for packets, else TypeID
	Name    string // Go name
	Fields  []manifest.Field
	Packet  bool
}

type index struct {
	enums  []enumInfo
	owners []ownerInfo
}

func buildIndex(m manifest.Manifest, overlay naming.Overlay) (index, error) {
	resolver := naming.NewResolver(overlay)
	for _, packet := range m.Packets {
		if err := resolver.Reserve(packet.Name, naming.PacketTypeName(packet.Name), naming.GoExportName); err != nil {
			return index{}, err
		}
	}
	var result index
	seenEnum := map[string]int{}
	seenOwner := map[string]bool{}
	// A payload struct inlined into its packet is indexed as that packet, not
	// as a shared type.
	usage := flatten.Count(m)
	native := func(node manifest.Node) bool { return nativeTypeIDs[node.TypeID] }
	for _, packet := range m.Packets {
		for _, item := range usage.PacketFields(packet, native) {
			if item.Owner != packet.Name {
				seenOwner[item.Owner] = true
			}
		}
	}
	var walk func(node manifest.Node) error
	walk = func(node manifest.Node) error {
		typeID := node.TypeID
		if typeID == "" {
			typeID = naming.InferredTypeName(node)
		}
		if typeID != "" {
			switch node.Kind {
			case manifest.KindEnum:
				position, seen := seenEnum[typeID]
				if !seen {
					name, err := resolver.Resolve(node, "", naming.GoExportName)
					if err != nil {
						return err
					}
					position = len(result.enums)
					seenEnum[typeID] = position
					result.enums = append(result.enums, enumInfo{TypeID: typeID, Name: name})
				}
				// Union the variants of every site so partial claims still map.
				known := map[string]bool{}
				for _, variant := range result.enums[position].Variants {
					known[variant.Name] = true
				}
				for _, variant := range node.Variants {
					if !known[variant.Name] {
						result.enums[position].Variants = append(result.enums[position].Variants, variant)
					}
				}
			case manifest.KindStruct:
				if !seenOwner[typeID] && len(node.Fields) > 0 {
					name, err := resolver.Resolve(node, "", naming.GoExportName)
					if err != nil {
						return err
					}
					seenOwner[typeID] = true
					result.owners = append(result.owners, ownerInfo{TypeID: typeID, FileKey: typeID, Name: name, Fields: node.Fields})
				}
			}
		}
		for _, field := range node.Fields {
			if err := walk(field.Encode); err != nil {
				return err
			}
		}
		for _, variant := range node.Variants {
			if err := walk(variant.Encode); err != nil {
				return err
			}
		}
		for _, child := range []*manifest.Node{node.Prefix, node.Element, node.Value, node.Key, node.Control, node.Default} {
			if child != nil {
				if err := walk(*child); err != nil {
					return err
				}
			}
		}
		for _, child := range node.Elements {
			if err := walk(child); err != nil {
				return err
			}
		}
		for _, oneCase := range node.Cases {
			for _, child := range append(oneCase.Encode, oneCase.Decode...) {
				if err := walk(child); err != nil {
					return err
				}
			}
		}
		return nil
	}
	for _, packet := range m.Packets {
		// The emitter inlines a sole single-use payload struct, so the packet's
		// effective fields and their owning type ID must match what it keys
		// docs and layout names by.
		effective := usage.PacketFields(packet, native)
		owner := packet.Name
		fields := make([]manifest.Field, 0, len(effective))
		for _, item := range effective {
			owner = item.Owner
			fields = append(fields, item.Field)
		}
		result.owners = append(result.owners, ownerInfo{TypeID: owner, FileKey: packet.Name, Name: naming.GoExportName(naming.PacketTypeName(packet.Name)), Fields: fields, Packet: true})
		for _, field := range packet.Fields {
			if err := walk(field.Encode); err != nil {
				return index{}, err
			}
		}
	}
	return result, nil
}

// --- fork side --------------------------------------------------------------

type forkField struct {
	Name string
	Type string
	Doc  string
}

type forkType struct {
	Name     string
	Package  string
	File     string
	Doc      string
	Fields   []forkField
	Struct   bool
	Marshals bool // has a Marshal method, so it is a wire type rather than runtime plumbing
}

type forkConst struct {
	Name  string
	Value int64
}

type forkGroup struct {
	Package string
	File    string
	Consts  []forkConst
	prefix  string
}

type forkIndex struct {
	commit string
	types  map[string]forkType // key: package + "." + name
	groups []forkGroup
}

func parseFork(root string) (forkIndex, error) {
	result := forkIndex{types: map[string]forkType{}}
	marshals := map[string]bool{}
	if out, err := exec.Command("git", "-C", root, "rev-parse", "HEAD").Output(); err == nil {
		result.commit = strings.TrimSpace(string(out))
	}
	for _, pkg := range []string{"protocol", "packet"} {
		directory := filepath.Join(root, "minecraft", "protocol")
		if pkg == "packet" {
			directory = filepath.Join(directory, "packet")
		}
		paths, err := filepath.Glob(filepath.Join(directory, "*.go"))
		if err != nil {
			return forkIndex{}, err
		}
		sort.Strings(paths)
		fset := token.NewFileSet()
		var files []*ast.File
		stems := map[*ast.File]string{}
		for _, path := range paths {
			if strings.HasSuffix(path, "_test.go") {
				continue
			}
			file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				return forkIndex{}, fmt.Errorf("parse %s: %w", path, err)
			}
			if file.Name.Name != pkg {
				continue
			}
			files = append(files, file)
			stems[file] = strings.TrimSuffix(filepath.Base(path), ".go")
		}
		evaluator := newConstEvaluator(files)
		for _, file := range files {
			for _, declaration := range file.Decls {
				if fn, ok := declaration.(*ast.FuncDecl); ok && fn.Name.Name == "Marshal" && fn.Recv != nil && len(fn.Recv.List) == 1 {
					receiver := fn.Recv.List[0].Type
					if star, ok := receiver.(*ast.StarExpr); ok {
						receiver = star.X
					}
					if ident, ok := receiver.(*ast.Ident); ok {
						marshals[pkg+"."+ident.Name] = true
					}
					continue
				}
				gen, ok := declaration.(*ast.GenDecl)
				if !ok {
					continue
				}
				switch gen.Tok {
				case token.TYPE:
					for _, spec := range gen.Specs {
						typeSpec := spec.(*ast.TypeSpec)
						item := forkType{Name: typeSpec.Name.Name, Package: pkg, File: stems[file], Doc: commentText(typeSpec.Doc)}
						if item.Doc == "" {
							item.Doc = commentText(gen.Doc)
						}
						if structure, ok := typeSpec.Type.(*ast.StructType); ok {
							item.Struct = true
							for _, field := range structure.Fields.List {
								var typ strings.Builder
								_ = format.Node(&typ, fset, field.Type)
								doc := commentText(field.Doc)
								if doc == "" {
									doc = commentText(field.Comment)
								}
								for _, name := range field.Names {
									item.Fields = append(item.Fields, forkField{Name: name.Name, Type: typ.String(), Doc: doc})
								}
							}
						}
						key := pkg + "." + item.Name
						if _, exists := result.types[key]; !exists {
							result.types[key] = item
						}
					}
				case token.CONST:
					group := forkGroup{Package: pkg, File: stems[file]}
					for _, c := range evaluator.group(gen) {
						group.Consts = append(group.Consts, c)
					}
					if len(group.Consts) > 0 {
						group.prefix = commonPrefix(group.Consts)
						result.groups = append(result.groups, group)
					}
				}
			}
		}
	}
	for key, item := range result.types {
		item.Marshals = marshals[key]
		result.types[key] = item
	}
	return result, nil
}

// constEvaluator resolves integer constant blocks, including iota, shifts,
// and references to other constants in the package.
type constEvaluator struct {
	values map[string]int64
	exprs  map[string]ast.Expr
	iotas  map[string]int
}

func newConstEvaluator(files []*ast.File) *constEvaluator {
	e := &constEvaluator{values: map[string]int64{}, exprs: map[string]ast.Expr{}, iotas: map[string]int{}}
	for _, file := range files {
		for _, declaration := range file.Decls {
			gen, ok := declaration.(*ast.GenDecl)
			if !ok || gen.Tok != token.CONST {
				continue
			}
			var last []ast.Expr
			for iota, spec := range gen.Specs {
				valueSpec := spec.(*ast.ValueSpec)
				if len(valueSpec.Values) > 0 {
					last = valueSpec.Values
				}
				for i, name := range valueSpec.Names {
					if name.Name == "_" || i >= len(last) {
						continue
					}
					e.exprs[name.Name] = last[i]
					e.iotas[name.Name] = iota
				}
			}
		}
	}
	return e
}

func (e *constEvaluator) group(gen *ast.GenDecl) []forkConst {
	var result []forkConst
	for _, spec := range gen.Specs {
		for _, name := range spec.(*ast.ValueSpec).Names {
			if name.Name == "_" {
				continue
			}
			if value, ok := e.resolve(name.Name, 0); ok {
				result = append(result, forkConst{Name: name.Name, Value: value})
			}
		}
	}
	return result
}

func (e *constEvaluator) resolve(name string, depth int) (int64, bool) {
	if value, ok := e.values[name]; ok {
		return value, true
	}
	expr, ok := e.exprs[name]
	if !ok || depth > 8 {
		return 0, false
	}
	value, ok := e.eval(expr, e.iotas[name], depth)
	if ok {
		e.values[name] = value
	}
	return value, ok
}

func (e *constEvaluator) eval(expr ast.Expr, iota int, depth int) (int64, bool) {
	switch value := expr.(type) {
	case *ast.BasicLit:
		if value.Kind != token.INT {
			return 0, false
		}
		n, err := strconv.ParseInt(value.Value, 0, 64)
		return n, err == nil
	case *ast.Ident:
		if value.Name == "iota" {
			return int64(iota), true
		}
		return e.resolve(value.Name, depth+1)
	case *ast.ParenExpr:
		return e.eval(value.X, iota, depth)
	case *ast.CallExpr:
		if len(value.Args) == 1 {
			return e.eval(value.Args[0], iota, depth)
		}
	case *ast.UnaryExpr:
		operand, ok := e.eval(value.X, iota, depth)
		if !ok {
			return 0, false
		}
		switch value.Op {
		case token.SUB:
			return -operand, true
		case token.ADD:
			return operand, true
		case token.XOR:
			return ^operand, true
		}
	case *ast.BinaryExpr:
		left, ok := e.eval(value.X, iota, depth)
		if !ok {
			return 0, false
		}
		right, ok := e.eval(value.Y, iota, depth)
		if !ok {
			return 0, false
		}
		switch value.Op {
		case token.ADD:
			return left + right, true
		case token.SUB:
			return left - right, true
		case token.MUL:
			return left * right, true
		case token.SHL:
			return left << uint(right), true
		case token.SHR:
			return left >> uint(right), true
		case token.OR:
			return left | right, true
		case token.AND:
			return left & right, true
		}
	}
	return 0, false
}

func commentText(group *ast.CommentGroup) string {
	if group == nil {
		return ""
	}
	return strings.TrimSpace(group.Text())
}

func commonPrefix(consts []forkConst) string {
	if len(consts) == 0 {
		return ""
	}
	prefix := consts[0].Name
	for _, c := range consts[1:] {
		for !strings.HasPrefix(c.Name, prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	// Cut back to a word boundary so "GameTypeS" becomes "GameType" while
	// "IDLogin"/"IDPlayStatus" keeps "ID".
	for k := len(prefix); k > 0; k-- {
		boundary := true
		for _, c := range consts {
			if len(c.Name) > k && !unicode.IsUpper(rune(c.Name[k])) && !unicode.IsDigit(rune(c.Name[k])) {
				boundary = false
				break
			}
		}
		if boundary {
			return prefix[:k]
		}
	}
	return ""
}

// --- matching -----------------------------------------------------------------

type gapReport struct {
	namedVariants      int
	portedDocs         int
	placements         []placementNote
	unmatchedEnums     []enumInfo
	unmatchedGroups    []forkGroup
	forkOnlyTypes      []forkType
	generatedOnlyTypes []ownerInfo
	fieldGaps          []fieldGap
}

type placementNote struct {
	Enum     enumInfo
	Group    forkGroup
	Unnamed  []string
	Foreign  []forkConst // fork constants with no manifest variant
	Mismatch []string    // same name, different value
	Score    int
}

type fieldGap struct {
	Owner        ownerInfo
	Fork         forkType
	Positional   bool
	OnlyManifest []string
	OnlyFork     []string
	TypeMismatch []string
	Renamed      int
}

func seed(m manifest.Manifest, idx index, fork forkIndex, docOverlay docs.Overlay) (layout.Document, *gapReport) {
	document := layout.Document{SchemaVersion: 1}
	report := &gapReport{}
	usedGroups := map[int]bool{}
	for i, group := range fork.groups {
		// Packet IDs are emitted by the generator itself; never relocate them.
		if group.prefix == "ID" {
			usedGroups[i] = true
		}
	}
	for _, enum := range idx.enums {
		best, bestScore := -1, 0
		for i, group := range fork.groups {
			if usedGroups[i] {
				continue
			}
			if score := matchScore(enum, group); score > bestScore {
				best, bestScore = i, score
			}
		}
		if best < 0 || bestScore < 2 || bestScore*2 < len(enum.Variants) {
			report.unmatchedEnums = append(report.unmatchedEnums, enum)
			continue
		}
		usedGroups[best] = true
		group := fork.groups[best]
		if group.Package == "protocol" {
			document.Files = append(document.Files, layout.FileEntry{TypeID: enum.TypeID, Package: "protocol", File: group.File, Rationale: fmt.Sprintf("gophertunnel keeps the %s constants in protocol/%s.go.", group.prefix, group.File)})
		}
		names, note := variantNames(enum, group)
		note.Score = bestScore
		report.placements = append(report.placements, note)
		report.namedVariants += len(names)
		document.Constants = append(document.Constants, layout.ConstantEntry{
			TypeID:    enum.TypeID,
			Package:   group.Package,
			File:      group.File,
			Names:     names,
			Rationale: fmt.Sprintf("Matches the %s const block in gophertunnel %s/%s.go (%d of %d variants named by value and name).", group.prefix, group.Package, group.File, len(names), len(enum.Variants)),
		})
	}
	for i, group := range fork.groups {
		if !usedGroups[i] {
			report.unmatchedGroups = append(report.unmatchedGroups, group)
		}
	}

	matchedFork := map[string]bool{}
	var unmatchedOwners []ownerInfo
	pair := func(owner ownerInfo, source forkType) {
		matchedFork[source.Package+"."+source.Name] = true
		document.Files = append(document.Files, layout.FileEntry{TypeID: owner.FileKey, Package: source.Package, File: source.File, Rationale: fmt.Sprintf("gophertunnel keeps %s in %s/%s.go.", source.Name, source.Package, source.File)})
		if docOverlay.Types != nil && source.Doc != "" && docOverlay.Types[owner.FileKey] == "" {
			docOverlay.Types[owner.FileKey] = docs.LeadWith(source.Doc, source.Name, owner.Name)
			report.portedDocs++
		}
		gap := fieldGap{Owner: owner, Fork: source}
		entries := matchFields(owner, source, &gap, docOverlay, report)
		document.Fields = append(document.Fields, entries...)
		gap.Renamed = len(entries)
		if len(gap.OnlyManifest)+len(gap.OnlyFork)+len(gap.TypeMismatch) > 0 || gap.Renamed > 0 {
			report.fieldGaps = append(report.fieldGaps, gap)
		}
	}
	for _, owner := range idx.owners {
		if nativeTypeIDs[owner.TypeID] {
			continue
		}
		if source, ok := forkTypeFor(owner.Name, ownerPackage(owner), fork); ok {
			pair(owner, source)
		} else {
			unmatchedOwners = append(unmatchedOwners, owner)
		}
	}
	// Second pass: types whose names differ but whose field names mostly agree.
	for _, owner := range unmatchedOwners {
		source, ok := forkTypeByFields(owner, fork, matchedFork)
		if ok {
			pair(owner, source)
		} else {
			report.generatedOnlyTypes = append(report.generatedOnlyTypes, owner)
			document.Fields = append(document.Fields, conventionEntries(owner, nil)...)
		}
	}
	for key, item := range fork.types {
		if item.Struct && item.Marshals && !matchedFork[key] {
			report.forkOnlyTypes = append(report.forkOnlyTypes, item)
		}
	}
	sort.Slice(report.forkOnlyTypes, func(i, j int) bool {
		return report.forkOnlyTypes[i].Package+report.forkOnlyTypes[i].Name < report.forkOnlyTypes[j].Package+report.forkOnlyTypes[j].Name
	})
	return document, report
}

// matchScore counts variants a const group explains: two points when a
// constant shares the value and the name, one when only the value matches an
// otherwise plausible group.
func matchScore(enum enumInfo, group forkGroup) int {
	byValue := map[int64][]forkConst{}
	for _, c := range group.Consts {
		byValue[c.Value] = append(byValue[c.Value], c)
	}
	score, named := 0, 0
	for _, variant := range enum.Variants {
		candidates := byValue[variant.Value]
		if len(candidates) == 0 {
			continue
		}
		if _, ok := nameMatch(variant.Name, group.prefix, candidates); ok {
			score += 2
			named++
		} else if len(candidates) == 1 {
			score++
		}
	}
	if named == 0 && !similar(enum.Name, group.prefix) {
		return 0
	}
	return score
}

func nameMatch(variant, prefix string, candidates []forkConst) (forkConst, bool) {
	want := normalize(variant)
	for _, c := range candidates {
		suffix := normalize(strings.TrimPrefix(c.Name, prefix))
		if suffix == want || strings.HasSuffix(normalize(c.Name), want) || (want != "" && strings.HasSuffix(want, suffix) && suffix != "") {
			return c, true
		}
	}
	return forkConst{}, false
}

func variantNames(enum enumInfo, group forkGroup) (map[string]string, placementNote) {
	note := placementNote{Enum: enum, Group: group}
	byValue := map[int64][]forkConst{}
	for _, c := range group.Consts {
		byValue[c.Value] = append(byValue[c.Value], c)
	}
	names := map[string]string{}
	taken := map[string]bool{}
	for _, variant := range enum.Variants {
		candidates := byValue[variant.Value]
		var chosen forkConst
		ok := false
		if c, matched := nameMatch(variant.Name, group.prefix, candidates); matched {
			chosen, ok = c, true
		} else if len(candidates) == 1 {
			chosen, ok = candidates[0], true
		}
		if !ok || taken[chosen.Name] {
			note.Unnamed = append(note.Unnamed, fmt.Sprintf("%s=%d", variant.Name, variant.Value))
			continue
		}
		taken[chosen.Name] = true
		names[variant.Name] = chosen.Name
	}
	for _, c := range group.Consts {
		if !taken[c.Name] {
			note.Foreign = append(note.Foreign, c)
		}
	}
	return names, note
}

func ownerPackage(owner ownerInfo) string {
	if owner.Packet {
		return "packet"
	}
	return "protocol"
}

// forkTypeByFields matches a type to an unclaimed fork struct in the same
// package when at least two thirds of its field names agree.
func forkTypeByFields(owner ownerInfo, fork forkIndex, claimed map[string]bool) (forkType, bool) {
	if len(owner.Fields) < 2 {
		return forkType{}, false
	}
	want := map[string]bool{}
	for _, field := range owner.Fields {
		want[normalize(field.Name)] = true
	}
	var best forkType
	bestShared := 0
	for key, item := range fork.types {
		if !item.Struct || claimed[key] || item.Package != ownerPackage(owner) || len(item.Fields) < 2 {
			continue
		}
		shared := 0
		for _, field := range item.Fields {
			if want[normalize(field.Name)] {
				shared++
			}
		}
		if shared*3 >= len(owner.Fields)*2 && shared*3 >= len(item.Fields)*2 && shared > bestShared {
			best, bestShared = item, shared
		}
	}
	return best, bestShared > 0
}

func forkTypeFor(name, pkg string, fork forkIndex) (forkType, bool) {
	if item, ok := fork.types[pkg+"."+name]; ok && item.Struct {
		return item, true
	}
	want := normalize(name)
	var candidates []forkType
	for _, item := range fork.types {
		if !item.Struct || item.Package != pkg {
			continue
		}
		have := normalize(item.Name)
		if want == have {
			return item, true
		}
		if want == have+"data" || want == have+"info" || want == have+"type" || have == want+"data" {
			candidates = append(candidates, item)
		}
	}
	if len(candidates) == 1 {
		return candidates[0], true
	}
	return forkType{}, false
}

// nativeTypeIDs are emitted as mgl32, uuid, or plain integers, so they never
// need a fork struct.
var nativeTypeIDs = map[string]bool{"Vec2": true, "Vec3": true, "BlockPos": true, "ActorUniqueID": true, "ActorRuntimeID": true, "PlayerInputTick": true, "mce::Color": true, "mce::UUID": true}

// forkSpelling applies the fork's naming conventions to a generated field
// name: entities rather than actors, and British spellings.
func forkSpelling(name string) string {
	for _, pair := range [][2]string{{"Actor", "Entity"}, {"Armor", "Armour"}, {"Color", "Colour"}, {"Behavior", "Behaviour"}} {
		name = strings.ReplaceAll(name, pair[0], pair[1])
	}
	return name
}

// conventionEntries renames fields the fork does not name so they still read
// like the fork's; matched is the set of wire names that already have a fork
// name.
func conventionEntries(owner ownerInfo, matched map[string]bool) []layout.FieldEntry {
	var entries []layout.FieldEntry
	for _, field := range owner.Fields {
		if matched[field.Name] {
			continue
		}
		generated := naming.GoExportName(field.Name)
		if styled := forkSpelling(generated); styled != generated {
			entries = append(entries, layout.FieldEntry{TypeID: owner.TypeID, Field: field.Name, Name: styled, Rationale: "gophertunnel spelling convention for a field it does not name."})
		}
	}
	return entries
}

func matchFields(owner ownerInfo, source forkType, gap *fieldGap, docOverlay docs.Overlay, report *gapReport) []layout.FieldEntry {
	var entries []layout.FieldEntry
	usedFork := map[int]bool{}
	matched := map[int]int{} // manifest index -> fork index
	for i, field := range owner.Fields {
		for j, forkField := range source.Fields {
			if usedFork[j] {
				continue
			}
			if normalize(field.Name) == normalize(forkField.Name) || normalize(naming.GoExportName(field.Name)) == normalize(forkField.Name) {
				matched[i], usedFork[j] = j, true
				break
			}
		}
	}
	if len(matched) < len(owner.Fields) && len(owner.Fields) == len(source.Fields) {
		positional := true
		for i := range owner.Fields {
			if _, ok := matched[i]; ok {
				continue
			}
			if usedFork[i] || typeCategory(owner.Fields[i].Encode) != forkCategory(source.Fields[i].Type) {
				positional = false
				break
			}
		}
		if positional {
			for i := range owner.Fields {
				if _, ok := matched[i]; !ok {
					matched[i], usedFork[i] = i, true
				}
			}
			gap.Positional = true
		}
	}
	named := map[string]bool{}
	for i, field := range owner.Fields {
		j, ok := matched[i]
		if !ok {
			gap.OnlyManifest = append(gap.OnlyManifest, field.Name)
			continue
		}
		forkField := source.Fields[j]
		if typeCategory(field.Encode) != forkCategory(forkField.Type) {
			gap.TypeMismatch = append(gap.TypeMismatch, fmt.Sprintf("%s (%s vs %s)", forkField.Name, describe(field.Encode), forkField.Type))
		}
		goName := naming.GoExportName(field.Name)
		if goName != forkField.Name && naming.IsExportedGoIdentifier(forkField.Name) {
			entries = append(entries, layout.FieldEntry{TypeID: owner.TypeID, Field: field.Name, Name: forkField.Name, Rationale: fmt.Sprintf("gophertunnel %s.%s field name.", source.Name, forkField.Name)})
			goName = forkField.Name
		}
		named[field.Name] = true
		if docOverlay.Fields != nil && forkField.Doc != "" {
			key := docs.FieldKey(owner.TypeID, field.Name)
			if docOverlay.Fields[key] == "" {
				docOverlay.Fields[key] = docs.LeadWith(forkField.Doc, forkField.Name, goName)
				report.portedDocs++
			}
		}
	}
	entries = append(entries, conventionEntries(owner, named)...)
	for j, forkField := range source.Fields {
		if !usedFork[j] {
			gap.OnlyFork = append(gap.OnlyFork, forkField.Name+" "+forkField.Type)
		}
	}
	return entries
}

func typeCategory(node manifest.Node) string {
	switch node.Kind {
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return "unknown"
		}
		switch node.Primitive.Code {
		case "bool":
			return "bool"
		case "f32le", "f32be":
			return "float32"
		case "f64le", "f64be":
			return "float64"
		case "uuid":
			return "named"
		default:
			return "integer"
		}
	case manifest.KindEnum:
		// gophertunnel keeps enums as plain integers.
		return "integer"
	case manifest.KindString:
		return "string"
	case manifest.KindBytes, manifest.KindArray:
		return "slice"
	case manifest.KindFixedArray:
		return "array"
	case manifest.KindOptional:
		return "optional"
	case manifest.KindMap:
		return "map"
	default:
		return "named"
	}
}

func forkCategory(value string) string {
	value = strings.TrimSpace(value)
	switch {
	case strings.HasPrefix(value, "[]"):
		return "slice"
	case strings.HasPrefix(value, "["):
		return "array"
	case strings.HasPrefix(value, "map["):
		return "map"
	case strings.Contains(value, "Optional["):
		return "optional"
	}
	switch value {
	case "bool":
		return "bool"
	case "float32":
		return "float32"
	case "float64":
		return "float64"
	case "int8", "int16", "int32", "int64", "int", "uint8", "uint16", "uint32", "uint64", "uint", "byte":
		return "integer"
	case "string":
		return "string"
	default:
		return "named"
	}
}

func describe(node manifest.Node) string {
	switch node.Kind {
	case manifest.KindPrimitive:
		if node.Primitive != nil {
			return node.Primitive.Code
		}
	case manifest.KindEnum:
		if node.Primitive != nil {
			return "enum " + node.Primitive.Code
		}
	case manifest.KindOptional:
		if node.Value != nil {
			return "optional " + describe(*node.Value)
		}
	case manifest.KindArray:
		if node.Element != nil {
			return "[]" + describe(*node.Element)
		}
	}
	if node.TypeID != "" {
		return node.TypeID
	}
	return string(node.Kind)
}

func similar(a, b string) bool {
	x, y := normalize(a), normalize(b)
	return x != "" && y != "" && (strings.Contains(x, y) || strings.Contains(y, x))
}

// normalize folds case, punctuation, and the British spellings gophertunnel
// uses so names compare by meaning.
func normalize(value string) string {
	var b strings.Builder
	for _, r := range value {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	result := b.String()
	for _, pair := range [][2]string{{"armour", "armor"}, {"colour", "color"}, {"behaviour", "behavior"}, {"centre", "center"}, {"entity", "actor"}} {
		result = strings.ReplaceAll(result, pair[0], pair[1])
	}
	return result
}

func writeJSON(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}

// --- report -------------------------------------------------------------------

func (r *gapReport) render(m manifest.Manifest, fork forkIndex) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# gophertunnel layout gaps for Minecraft %s / protocol %d\n\n", m.Target.MinecraftVersion, m.Target.ProtocolVersion)
	fmt.Fprintf(&b, "Fork checkout `%s`. Generated by `tools/seed-gophertunnel-layout`; the numbers describe what the layout overlay could map automatically, not wire correctness.\n\n", fork.commit)
	fmt.Fprintf(&b, "| | count |\n|---|---:|\n")
	fmt.Fprintf(&b, "| enums placed onto a fork const block | %d |\n", len(r.placements))
	fmt.Fprintf(&b, "| enums with no fork const block | %d |\n", len(r.unmatchedEnums))
	fmt.Fprintf(&b, "| fork const blocks with no enum | %d |\n", len(r.unmatchedGroups))
	fmt.Fprintf(&b, "| fork struct types with no generated type | %d |\n", len(r.forkOnlyTypes))
	fmt.Fprintf(&b, "| generated types with no fork type | %d |\n", len(r.generatedOnlyTypes))
	fmt.Fprintf(&b, "| matched types with field differences | %d |\n\n", len(r.fieldGaps))

	b.WriteString("## Fork struct types with no generated counterpart\n\nHand-written decoding or a name the overlay does not map yet.\n\n")
	for _, item := range r.forkOnlyTypes {
		fmt.Fprintf(&b, "- `%s.%s` (%s.go, %d fields)\n", item.Package, item.Name, item.File, len(item.Fields))
	}
	b.WriteString("\n## Generated types with no fork counterpart\n\n")
	for _, owner := range r.generatedOnlyTypes {
		fmt.Fprintf(&b, "- `%s` (%s, %d fields)\n", owner.Name, owner.TypeID, len(owner.Fields))
	}
	b.WriteString("\n## Field differences on matched types\n\n")
	for _, gap := range r.fieldGaps {
		fmt.Fprintf(&b, "### `%s` vs fork `%s.%s`", gap.Owner.Name, gap.Fork.Package, gap.Fork.Name)
		if gap.Positional {
			b.WriteString(" (some fields matched by position)")
		}
		fmt.Fprintf(&b, "\n\n- renamed to fork names: %d\n", gap.Renamed)
		if len(gap.OnlyManifest) > 0 {
			fmt.Fprintf(&b, "- only in manifest: %s\n", strings.Join(quote(gap.OnlyManifest), ", "))
		}
		if len(gap.OnlyFork) > 0 {
			fmt.Fprintf(&b, "- only in fork: %s\n", strings.Join(quote(gap.OnlyFork), ", "))
		}
		if len(gap.TypeMismatch) > 0 {
			fmt.Fprintf(&b, "- type category differs: %s\n", strings.Join(quote(gap.TypeMismatch), ", "))
		}
		b.WriteString("\n")
	}
	b.WriteString("## Constant placements needing review\n\nVariants the fork block does not name keep their generated name; fork constants with no variant are listed as foreign.\n\n")
	for _, note := range r.placements {
		if len(note.Unnamed) == 0 && len(note.Foreign) == 0 {
			continue
		}
		fmt.Fprintf(&b, "- `%s` -> `%s.%s` block `%s*` (score %d)", note.Enum.Name, note.Group.Package, note.Group.File, note.Group.prefix, note.Score)
		if len(note.Unnamed) > 0 {
			fmt.Fprintf(&b, "; unnamed: %s", strings.Join(note.Unnamed, ", "))
		}
		if len(note.Foreign) > 0 {
			names := make([]string, 0, len(note.Foreign))
			for _, c := range note.Foreign {
				names = append(names, fmt.Sprintf("%s=%d", c.Name, c.Value))
			}
			fmt.Fprintf(&b, "; foreign: %s", strings.Join(names, ", "))
		}
		b.WriteString("\n")
	}
	b.WriteString("\n## Enums with no fork const block\n\n")
	for _, enum := range r.unmatchedEnums {
		fmt.Fprintf(&b, "- `%s` (%s, %d variants)\n", enum.Name, enum.TypeID, len(enum.Variants))
	}
	b.WriteString("\n## Fork const blocks with no enum\n\n")
	for _, group := range r.unmatchedGroups {
		fmt.Fprintf(&b, "- `%s.%s` `%s*` (%d constants)\n", group.Package, group.File, group.prefix, len(group.Consts))
	}
	return b.String()
}

func quote(values []string) []string {
	result := make([]string, len(values))
	for i, value := range values {
		result[i] = "`" + value + "`"
	}
	return result
}

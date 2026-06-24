package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"sort"
	"strings"
)

// mergeIntoExisting rewrites an existing packet source in place: it replaces the
// packet struct's field list and the Marshal method body with generated content
// while preserving the package/type doc comments, const blocks, and helper
// functions.
//
// When preserveMatching is true, a generated field whose wire encoding matches
// the existing field at the same wire position (read from the Marshal op
// sequence, not struct order) is kept verbatim — name, type, comment, and op —
// so gophertunnel's clean API names (and []byte/Optional/variant choices) are
// not churned by Mojang's inconsistent doc names. Only genuinely new/changed
// fields use generated content.
func mergeIntoExisting(existingSrc []byte, pk genPacket, preserveMatching bool) (merged string, todos []string, err error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", existingSrc, parser.ParseComments)
	if err != nil {
		return "", nil, fmt.Errorf("parse existing: %w", err)
	}

	var marshalFn *ast.FuncDecl
	var importDecl *ast.GenDecl
	for _, d := range f.Decls {
		switch decl := d.(type) {
		case *ast.GenDecl:
			if decl.Tok == token.IMPORT && importDecl == nil {
				importDecl = decl
			}
		case *ast.FuncDecl:
			if decl.Name.Name == "Marshal" && decl.Recv != nil && len(decl.Recv.List) == 1 {
				marshalFn = decl
			}
		}
	}
	if marshalFn == nil || marshalFn.Body == nil {
		return "", nil, fmt.Errorf("no Marshal method found")
	}
	recv := recvType(marshalFn.Recv.List[0].Type)

	var structType *ast.StructType
	for _, d := range f.Decls {
		gen, ok := d.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}
		for _, spec := range gen.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || ts.Name.Name != recv {
				continue
			}
			if st, ok := ts.Type.(*ast.StructType); ok {
				structType = st
			}
		}
	}
	if structType == nil || structType.Fields == nil {
		return "", nil, fmt.Errorf("no struct %q found", recv)
	}

	// Existing field doc comments and types: by position (for fallback carry-over)
	// and by name (for wire-matched preservation, which is order-independent).
	existingDocs := make([]string, 0, len(structType.Fields.List))
	existingTypes := make([]string, 0, len(structType.Fields.List))
	docByName := map[string]string{}
	typeByName := map[string]string{}
	for _, fld := range structType.Fields.List {
		doc, typ := commentText(fld.Doc), types.ExprString(fld.Type)
		n := len(fld.Names)
		if n == 0 {
			n = 1
		}
		for range n {
			existingDocs = append(existingDocs, doc)
			existingTypes = append(existingTypes, typ)
		}
		for _, nm := range fld.Names {
			docByName[nm.Name] = doc
			typeByName[nm.Name] = typ
		}
	}

	// Existing Marshal op sequence (wire order). Preservation only applies when the
	// Marshal is fully linear (no conditionals/helpers we can't position), which is
	// exactly the rename-noise case; union/conditional packets fall back to generated.
	ops, linear := extractMarshalOps(marshalFn, existingSrc, fset, firstParamName(marshalFn))

	needImports := map[string]bool{}
	var fieldBlock, marshalBody strings.Builder
	for i, gf := range pk.Fields {
		shape := resolve(gf)

		// Wire-matched preservation: keep gopher's field verbatim when its wire
		// encoding equals the generated field's at the same wire position.
		if preserveMatching && linear && i < len(ops) && ops[i].field != "" {
			genM, genI := parseGenOp(shape.Marshal)
			exType, known := typeByName[ops[i].field]
			if known {
				gs := wireSig(genM, genI, shape.GoType)
				es := wireSig(ops[i].method, ops[i].inner, exType)
				if gs != "" && gs == es {
					if d := docByName[ops[i].field]; d != "" {
						fieldBlock.WriteString(d)
					}
					fmt.Fprintf(&fieldBlock, "\t%s %s\n", ops[i].field, exType)
					marshalBody.WriteString("\t" + ops[i].src + "\n")
					continue
				}
			}
		}

		// Preserve gophertunnel's []byte where the docs would generate a plain
		// string (io.String and io.ByteSlice are wire-identical).
		if shape.GoType == "string" && shape.Todo == "" && i < len(existingTypes) && existingTypes[i] == "[]byte" {
			shape.GoType = "[]byte"
			shape.Marshal = fmt.Sprintf("io.ByteSlice(&pk.%s)", gf.GoName)
		}
		for _, imp := range shape.Imports {
			needImports[imp] = true
		}
		if shape.Todo != "" {
			todos = append(todos, fmt.Sprintf("%s.%s: %s", pk.TypeName, gf.GoName, shape.Todo))
		}
		if i < len(existingDocs) && existingDocs[i] != "" {
			fieldBlock.WriteString(existingDocs[i])
		}
		fmt.Fprintf(&fieldBlock, "\t%s %s\n", gf.GoName, shape.GoType)
		fmt.Fprintf(&marshalBody, "\t%s\n", shape.Marshal)
	}

	type edit struct {
		start, end int
		text       string
	}
	off := func(p token.Pos) int { return fset.Position(p).Offset }
	edits := []edit{
		{off(structType.Fields.Opening) + 1, off(structType.Fields.Closing), "\n" + fieldBlock.String()},
		{off(marshalFn.Body.Lbrace) + 1, off(marshalFn.Body.Rbrace), "\n" + marshalBody.String()},
	}

	if importDecl != nil {
		all := map[string]bool{protocolImport: true}
		for _, spec := range importDecl.Specs {
			if is, ok := spec.(*ast.ImportSpec); ok {
				all[strings.Trim(is.Path.Value, `"`)] = true
			}
		}
		for p := range needImports {
			all[p] = true
		}
		edits = append(edits, edit{off(importDecl.Pos()), off(importDecl.End()), renderImports(all)})
	}

	sort.Slice(edits, func(i, j int) bool { return edits[i].start > edits[j].start })
	out := append([]byte(nil), existingSrc...)
	for _, e := range edits {
		out = append(out[:e.start], append([]byte(e.text), out[e.end:]...)...)
	}

	formatted, ferr := format.Source(out)
	if ferr != nil {
		return string(out), todos, fmt.Errorf("gofmt: %w", ferr)
	}
	return string(formatted), todos, nil
}

// existingOp is one io/protocol call in a packet's Marshal method, in wire order.
type existingOp struct {
	field  string // the pk.<field> referenced
	method string // io method or protocol helper
	inner  string // nested io method, e.g. the String in FuncSlice(io, &x, io.String)
	src    string // verbatim source of the statement
}

// extractMarshalOps returns the Marshal's op sequence and whether it is fully
// linear (every statement is a single io/protocol call). Non-linear bodies
// (conditionals, helpers) disable preservation for that packet.
func extractMarshalOps(fn *ast.FuncDecl, src []byte, fset *token.FileSet, ioName string) (ops []existingOp, linear bool) {
	for _, stmt := range fn.Body.List {
		es, ok := stmt.(*ast.ExprStmt)
		if !ok {
			return nil, false
		}
		call, ok := es.X.(*ast.CallExpr)
		if !ok {
			return nil, false
		}
		op, ok := classifyExistingCall(call, ioName)
		if !ok {
			return nil, false
		}
		op.src = strings.TrimSpace(string(src[fset.Position(stmt.Pos()).Offset:fset.Position(stmt.End()).Offset]))
		ops = append(ops, op)
	}
	return ops, true
}

func classifyExistingCall(call *ast.CallExpr, ioName string) (existingOp, bool) {
	fun := call.Fun
	switch e := fun.(type) {
	case *ast.IndexExpr:
		fun = e.X
	case *ast.IndexListExpr:
		fun = e.X
	}
	sel, ok := fun.(*ast.SelectorExpr)
	if !ok {
		return existingOp{}, false
	}
	base, ok := sel.X.(*ast.Ident)
	if !ok || (base.Name != ioName && base.Name != "protocol") {
		return existingOp{}, false
	}
	return existingOp{
		method: sel.Sel.Name,
		field:  fieldRefFromArgs(call.Args),
		inner:  innerIOFromArgs(call.Args, ioName),
	}, true
}

func fieldRefFromArgs(args []ast.Expr) string {
	for _, a := range args {
		if u, ok := a.(*ast.UnaryExpr); ok && u.Op == token.AND {
			if s, ok := u.X.(*ast.SelectorExpr); ok {
				return s.Sel.Name
			}
		}
	}
	return ""
}

func innerIOFromArgs(args []ast.Expr, ioName string) string {
	for _, a := range args {
		if s, ok := a.(*ast.SelectorExpr); ok {
			if b, ok := s.X.(*ast.Ident); ok && b.Name == ioName {
				return s.Sel.Name
			}
		}
	}
	return ""
}

// bytesOps are the io methods that encode a varuint32-length-prefixed byte run;
// they are wire-identical, so string and []byte fields are interchangeable.
var bytesOps = map[string]bool{"String": true, "StringUTF": true, "ByteSlice": true, "Bytes": true}

// wireSig reduces an op + Go type to a wire-equivalence signature: two fields
// with the same signature encode identically (so a name difference is a pure
// rename). Field names are deliberately excluded.
func wireSig(method, inner, goType string) string {
	switch {
	case method == "":
		return ""
	case bytesOps[method]:
		return "bytes"
	case method == "Single":
		return "single:" + goType
	case strings.HasPrefix(method, "Slice"):
		return "slice:" + strings.TrimPrefix(goType, "[]")
	case strings.HasPrefix(method, "FuncSlice"):
		if bytesOps[inner] {
			return "slice:bytes"
		}
		return "slice:" + inner
	default:
		return method
	}
}

// parseGenOp extracts the op method (and any nested io method) from a generated
// Marshal statement string.
func parseGenOp(marshal string) (method, inner string) {
	if i := strings.Index(marshal, "//"); i >= 0 {
		marshal = marshal[:i]
	}
	m := strings.TrimSpace(marshal)
	switch {
	case strings.HasPrefix(m, "io."):
		method = leadingIdent(m[len("io."):])
	case strings.HasPrefix(m, "protocol."):
		method = leadingIdent(m[len("protocol."):])
		if j := strings.LastIndex(m, "io."); j >= 0 {
			inner = leadingIdent(m[j+len("io."):])
		}
	}
	return method, inner
}

func leadingIdent(s string) string {
	i := 0
	for i < len(s) && (s[i] == '_' || s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9') {
		i++
	}
	return s[:i]
}

// firstParamName returns the name of a function's first parameter (the io var).
func firstParamName(fn *ast.FuncDecl) string {
	if fn.Type.Params == nil || len(fn.Type.Params.List) == 0 || len(fn.Type.Params.List[0].Names) == 0 {
		return "io"
	}
	return fn.Type.Params.List[0].Names[0].Name
}

// commentText renders a comment group as raw source lines (with trailing newline).
func commentText(cg *ast.CommentGroup) string {
	if cg == nil {
		return ""
	}
	var b strings.Builder
	for _, c := range cg.List {
		b.WriteString(c.Text)
		b.WriteByte('\n')
	}
	return b.String()
}

// renderImports renders a sorted import block.
func renderImports(paths map[string]bool) string {
	keys := make([]string, 0, len(paths))
	for p := range paths {
		keys = append(keys, p)
	}
	sort.Strings(keys)
	var b strings.Builder
	b.WriteString("import (\n")
	for _, p := range keys {
		fmt.Fprintf(&b, "\t%q\n", p)
	}
	b.WriteString(")")
	return b.String()
}

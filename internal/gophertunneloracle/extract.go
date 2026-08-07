package gophertunneloracle

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const extractionDepthLimit = 40

type typeRef struct {
	Kind string // named, primitive, slice, array, optional, pointer, unknown
	Name string
	Elem *typeRef
	Len  int
}

type typeInfo struct {
	Key      string
	Name     string
	Pkg      string
	Fields   map[string]typeRef
	Embedded []typeRef
}

type sourceFile struct {
	Path    string
	Rel     string
	Pkg     string
	AST     *ast.File
	Imports map[string]string
}

type marshalInfo struct {
	Key  string
	Recv string
	IO   string
	Decl *ast.FuncDecl
	File *sourceFile
}

type extractor struct {
	root        string
	fset        *token.FileSet
	files       []*sourceFile
	types       map[string]*typeInfo
	marshals    map[string]*marshalInfo
	functions   map[string]*marshalInfo
	consts      map[string]int
	diagnostics []diagnostic
	packet      string
}

var sourcePrimitive = map[string]string{
	"Uint8":                   "u8",
	"Int8":                    "i8",
	"Uint16":                  "u16le",
	"Int16":                   "i16le",
	"Uint32":                  "u32le",
	"Int32":                   "i32le",
	"BEInt32":                 "i32be",
	"Uint64":                  "u64le",
	"Int64":                   "i64le",
	"Float32":                 "f32le",
	"Float64":                 "f64le",
	"Bool":                    "bool",
	"Varint32":                "zigzag_i32",
	"Varuint32":               "var_u32",
	"Varint64":                "zigzag_i64",
	"Varuint64":               "var_u64",
	"ActorRuntimeID":          "var_u64",
	"ActorRuntimeIDVarint64":  "zigzag_i64",
	"ActorRuntimeIDVaruint32": "var_u32",
	"ActorUniqueID":           "zigzag_i64",
	"ActorUniqueIDInt64":      "i64le",
	"ActorUniqueIDUint64":     "u64le",
	"ActorUniqueIDVaruint64":  "var_u64",
	"PlayerInputTick":         "var_u64",
}

var sourceBuiltinPrimitive = map[string]string{
	"byte": "u8", "uint8": "u8", "int8": "i8",
	"uint16": "u16le", "int16": "i16le",
	"uint32": "u32le", "int32": "i32le",
	"uint64": "u64le", "int64": "i64le",
	"float32": "f32le", "float64": "f64le", "bool": "bool",
}

var sourceArrayPrefixes = map[string]string{
	"Slice":                   "var_u32",
	"FuncSlice":               "var_u32",
	"FuncIOSlice":             "var_u32",
	"SliceOfFunc":             "var_u32",
	"SliceUint8Length":        "u8",
	"SliceUint16Length":       "u16le",
	"SliceUint32Length":       "u32le",
	"SliceVarint32Length":     "zigzag_i32",
	"FuncSliceUint8Length":    "u8",
	"FuncSliceUint16Length":   "u16le",
	"FuncSliceUint32Length":   "u32le",
	"FuncSliceVarint32Length": "zigzag_i32",
}

func Extract(root string) (extraction, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return extraction{}, fmt.Errorf("resolve gophertunnel root: %w", err)
	}
	e := &extractor{
		root:      abs,
		fset:      token.NewFileSet(),
		types:     map[string]*typeInfo{},
		marshals:  map[string]*marshalInfo{},
		functions: map[string]*marshalInfo{},
		consts:    map[string]int{},
	}
	if err := e.load(); err != nil {
		return extraction{}, err
	}
	ids, packetTypes := e.packetIDs()
	keys := make([]string, 0, len(packetTypes))
	for key := range packetTypes {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	packets := make([]sourcePacket, 0, len(keys))
	for _, key := range keys {
		idName := packetTypes[key]
		id, ok := ids[idName]
		if !ok {
			e.packet = shortType(key)
			e.addDiagnostic(nil, key, "", "packet ID constant could not be evaluated", idName)
			continue
		}
		e.packet = shortType(key)
		ops := e.expandType(key, e.packet, 0, map[string]bool{})
		packets = append(packets, sourcePacket{ID: uint32(id), Name: e.packet, Operations: ops})
	}
	sort.Slice(packets, func(i, j int) bool {
		if packets[i].ID != packets[j].ID {
			return packets[i].ID < packets[j].ID
		}
		return packets[i].Name < packets[j].Name
	})
	return extraction{Packets: packets, Diagnostics: e.diagnostics}, nil
}

func (e *extractor) load() error {
	base := filepath.Join(e.root, "minecraft", "protocol")
	if info, err := os.Stat(base); err != nil || !info.IsDir() {
		return fmt.Errorf("gophertunnel root has no minecraft/protocol directory: %s", e.root)
	}
	err := filepath.WalkDir(base, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			return nil
		}
		file, err := parser.ParseFile(e.fset, path, nil, parser.SkipObjectResolution)
		if err != nil {
			return fmt.Errorf("parse %s: %w", path, err)
		}
		rel, err := filepath.Rel(base, path)
		if err != nil {
			return err
		}
		sf := &sourceFile{Path: path, Rel: filepath.ToSlash(rel), Pkg: file.Name.Name, AST: file, Imports: map[string]string{}}
		for _, imp := range file.Imports {
			importPath, err := strconv.Unquote(imp.Path.Value)
			if err != nil {
				continue
			}
			alias := filepath.Base(importPath)
			if imp.Name != nil {
				alias = imp.Name.Name
			}
			sf.Imports[alias] = filepath.Base(importPath)
		}
		e.files = append(e.files, sf)
		return nil
	})
	if err != nil {
		return fmt.Errorf("walk gophertunnel protocol source: %w", err)
	}
	sort.Slice(e.files, func(i, j int) bool { return e.files[i].Rel < e.files[j].Rel })
	for _, sf := range e.files {
		e.collectConstants(sf)
	}
	for _, sf := range e.files {
		e.collectTypes(sf)
	}
	for _, sf := range e.files {
		e.collectMethods(sf)
	}
	return nil
}

func (e *extractor) collectConstants(sf *sourceFile) {
	for _, decl := range sf.AST.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.CONST {
			continue
		}
		var previous ast.Expr
		for i, spec := range gen.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			if len(valueSpec.Values) > 0 {
				previous = valueSpec.Values[0]
			}
			value, ok := e.evalConst(previous, i, sf.Pkg)
			if !ok {
				continue
			}
			for _, name := range valueSpec.Names {
				if name.Name != "_" {
					e.consts[sf.Pkg+"."+name.Name] = value
				}
			}
		}
	}
}

func (e *extractor) collectTypes(sf *sourceFile) {
	for _, decl := range sf.AST.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}
		for _, spec := range gen.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}
			key := sf.Pkg + "." + ts.Name.Name
			info := &typeInfo{Key: key, Name: ts.Name.Name, Pkg: sf.Pkg, Fields: map[string]typeRef{}}
			for _, field := range st.Fields.List {
				ref := e.parseType(field.Type, sf)
				if len(field.Names) == 0 {
					info.Embedded = append(info.Embedded, ref)
					info.Fields[shortType(ref.Name)] = ref
					continue
				}
				for _, name := range field.Names {
					info.Fields[name.Name] = ref
				}
			}
			e.types[key] = info
		}
	}
}

func (e *extractor) collectMethods(sf *sourceFile) {
	for _, decl := range sf.AST.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		if fn.Recv == nil {
			e.functions[sf.Pkg+"."+fn.Name.Name] = &marshalInfo{Key: sf.Pkg + "." + fn.Name.Name, Decl: fn, File: sf}
			continue
		}
		if fn.Name.Name != "Marshal" {
			continue
		}
		recvType, recvVar := receiver(fn)
		if recvType == "" {
			continue
		}
		ioName := ""
		if fn.Type.Params != nil && len(fn.Type.Params.List) > 0 && len(fn.Type.Params.List[0].Names) > 0 {
			ioName = fn.Type.Params.List[0].Names[0].Name
		}
		key := sf.Pkg + "." + recvType
		e.marshals[key] = &marshalInfo{Key: key, Recv: recvVar, IO: ioName, Decl: fn, File: sf}
	}
}

func (e *extractor) packetIDs() (map[string]int, map[string]string) {
	ids := map[string]int{}
	for name, value := range e.consts {
		if strings.HasPrefix(name, "packet.ID") {
			ids[strings.TrimPrefix(name, "packet.")] = value
		}
	}
	packetTypes := map[string]string{}
	for _, sf := range e.files {
		if sf.Pkg != "packet" {
			continue
		}
		for _, decl := range sf.AST.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Name.Name != "ID" || fn.Recv == nil || fn.Body == nil {
				continue
			}
			recv, _ := receiver(fn)
			idName := idConstFromReturn(fn)
			if idName != "" {
				packetTypes["packet."+recv] = idName
			}
		}
	}
	return ids, packetTypes
}

func receiver(fn *ast.FuncDecl) (string, string) {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		return "", ""
	}
	field := fn.Recv.List[0]
	name := ""
	if len(field.Names) > 0 {
		name = field.Names[0].Name
	}
	switch typ := field.Type.(type) {
	case *ast.Ident:
		return typ.Name, name
	case *ast.StarExpr:
		if ident, ok := typ.X.(*ast.Ident); ok {
			return ident.Name, name
		}
	}
	return "", name
}

func idConstFromReturn(fn *ast.FuncDecl) string {
	for _, stmt := range fn.Body.List {
		ret, ok := stmt.(*ast.ReturnStmt)
		if !ok || len(ret.Results) != 1 {
			continue
		}
		if ident, ok := ret.Results[0].(*ast.Ident); ok {
			return ident.Name
		}
	}
	return ""
}

func (e *extractor) evalConst(expr ast.Expr, iota int, pkg string) (int, bool) {
	switch value := expr.(type) {
	case nil:
		return iota, true
	case *ast.BasicLit:
		if value.Kind != token.INT {
			return 0, false
		}
		n, err := strconv.ParseInt(value.Value, 0, 64)
		return int(n), err == nil
	case *ast.Ident:
		if value.Name == "iota" {
			return iota, true
		}
		n, ok := e.consts[pkg+"."+value.Name]
		return n, ok
	case *ast.SelectorExpr:
		base, ok := value.X.(*ast.Ident)
		if !ok {
			return 0, false
		}
		n, ok := e.consts[base.Name+"."+value.Sel.Name]
		return n, ok
	case *ast.ParenExpr:
		return e.evalConst(value.X, iota, pkg)
	case *ast.BinaryExpr:
		left, ok := e.evalConst(value.X, iota, pkg)
		if !ok {
			return 0, false
		}
		right, ok := e.evalConst(value.Y, iota, pkg)
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
			return left << right, true
		case token.OR:
			return left | right, true
		}
	}
	return 0, false
}

func (e *extractor) expandType(key, field string, depth int, stack map[string]bool) []sourceOperation {
	if depth > extractionDepthLimit {
		return []sourceOperation{e.unresolved(nil, key, field, "recursion depth limit exceeded", key)}
	}
	if stack[key] {
		return []sourceOperation{{Kind: "recursive", Field: field, TypeName: key, Reason: "recursive reference"}}
	}
	method := e.marshals[key]
	if method == nil {
		return []sourceOperation{e.unresolved(nil, key, field, "target type has no Marshal method", key)}
	}
	next := cloneBoolMap(stack)
	next[key] = true
	env := map[string]typeRef{}
	if method.Recv != "" {
		env[method.Recv] = typeRef{Kind: "named", Name: key}
	}
	return e.extractBlock(method.Decl.Body.List, method, env, field, depth, next)
}

func cloneBoolMap(input map[string]bool) map[string]bool {
	result := make(map[string]bool, len(input))
	for key, value := range input {
		result[key] = value
	}
	return result
}

func (e *extractor) extractBlock(stmts []ast.Stmt, method *marshalInfo, env map[string]typeRef, base string, depth int, stack map[string]bool) []sourceOperation {
	var result []sourceOperation
	for _, stmt := range stmts {
		switch current := stmt.(type) {
		case *ast.ExprStmt:
			if call, ok := current.X.(*ast.CallExpr); ok {
				result = append(result, e.extractCall(call, method, env, base, depth, stack)...)
			}
		case *ast.AssignStmt:
			for _, rhs := range current.Rhs {
				if call, ok := rhs.(*ast.CallExpr); ok && e.isWireCall(call, method.IO) {
					result = append(result, e.extractCall(call, method, env, base, depth, stack)...)
				}
			}
		case *ast.DeclStmt:
			if gen, ok := current.Decl.(*ast.GenDecl); ok {
				for _, spec := range gen.Specs {
					valueSpec, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}
					for _, value := range valueSpec.Values {
						if call, ok := value.(*ast.CallExpr); ok && e.isWireCall(call, method.IO) {
							result = append(result, e.extractCall(call, method, env, base, depth, stack)...)
						}
					}
				}
			}
		case *ast.IfStmt:
			if e.statementHasWire(current, method.IO) {
				result = append(result, e.unresolved(current, method.Key, base, "runtime conditional branch cannot be represented as one wire sequence", e.nodeString(current.Cond)))
			}
		case *ast.SwitchStmt:
			if e.statementHasWire(current, method.IO) {
				result = append(result, e.unresolved(current, method.Key, base, "runtime switch/union requires a discriminator value", e.nodeString(current.Tag)))
			}
		case *ast.TypeSwitchStmt:
			if e.statementHasWire(current, method.IO) {
				result = append(result, e.unresolved(current, method.Key, base, "runtime type switch/union cannot be linearized", e.nodeString(current.Assign)))
			}
		case *ast.ForStmt, *ast.RangeStmt:
			if e.statementHasWire(current, method.IO) {
				result = append(result, e.unresolved(current, method.Key, base, "runtime loop outside a recognized slice helper", e.nodeString(current)))
			}
		case *ast.DeferStmt:
			if e.statementHasWire(current, method.IO) {
				result = append(result, e.unresolved(current, method.Key, base, "deferred wire call in Marshal", e.nodeString(current.Call)))
			}
		case *ast.BlockStmt:
			result = append(result, e.extractBlock(current.List, method, env, base, depth, stack)...)
		}
	}
	return result
}

func (e *extractor) statementHasWire(node ast.Node, ioName string) bool {
	found := false
	ast.Inspect(node, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if ok && e.isWireCall(call, ioName) {
			found = true
			return false
		}
		return !found
	})
	return found
}

func (e *extractor) isWireCall(call *ast.CallExpr, ioName string) bool {
	fun := unwrapIndex(call.Fun)
	switch current := fun.(type) {
	case *ast.SelectorExpr:
		if base, ok := current.X.(*ast.Ident); ok {
			if base.Name == ioName {
				return current.Sel.Name != "InvalidValue" && current.Sel.Name != "UnknownEnumOption" && current.Sel.Name != "ShieldID"
			}
			if base.Name == "protocol" {
				return true
			}
		}
		return current.Sel.Name == "Marshal" || referencesIO(call, ioName)
	case *ast.Ident:
		return referencesIO(call, ioName)
	default:
		return false
	}
}

func (e *extractor) extractCall(call *ast.CallExpr, method *marshalInfo, env map[string]typeRef, base string, depth int, stack map[string]bool) []sourceOperation {
	fun := unwrapIndex(call.Fun)
	selector, isSelector := fun.(*ast.SelectorExpr)
	if !isSelector {
		if ident, ok := fun.(*ast.Ident); ok && method.File.Pkg == "protocol" && isProtocolHelper(ident.Name) {
			return e.expandProtocol(call, ident.Name, method, env, base, depth, stack)
		}
		if referencesIO(call, method.IO) {
			return []sourceOperation{e.unresolved(call, method.Key, base, "unrecognized function call in Marshal", e.nodeString(call))}
		}
		return nil
	}
	name := selector.Sel.Name
	if ident, ok := selector.X.(*ast.Ident); ok && ident.Name == method.IO {
		if name == "InvalidValue" || name == "UnknownEnumOption" || name == "ShieldID" {
			return nil
		}
		field := base
		if len(call.Args) > 0 {
			field = e.fieldPath(call.Args[0], method, env, base)
		}
		if code, ok := sourcePrimitive[name]; ok {
			return []sourceOperation{{Kind: "primitive", Field: field, Code: code}}
		}
		return e.expandIOHelper(call, name, field, method)
	}
	if ident, ok := selector.X.(*ast.Ident); ok && ident.Name == "protocol" {
		return e.expandProtocol(call, name, method, env, base, depth, stack)
	}
	if name == "Marshal" {
		ref := e.resolveExprType(selector.X, method, env)
		for ref.Kind == "pointer" && ref.Elem != nil {
			ref = *ref.Elem
		}
		field := e.fieldPath(selector.X, method, env, base)
		if ref.Kind != "named" {
			return []sourceOperation{e.unresolved(call, method.Key, field, "Marshal receiver type could not be resolved", e.nodeString(selector.X))}
		}
		return e.expandType(ref.Name, field, depth+1, stack)
	}
	if referencesIO(call, method.IO) {
		return []sourceOperation{e.unresolved(call, method.Key, base, "opaque method/helper call in Marshal", e.nodeString(call))}
	}
	return nil
}

func isProtocolHelper(name string) bool {
	if _, ok := sourceArrayPrefixes[name]; ok {
		return true
	}
	switch name {
	case "SliceOfLen", "FuncSliceOfLen", "FuncIOSliceOfLen", "Optional", "OptionalFunc", "DoubleOptionalFunc", "OptionalMarshaler", "Single", "IntegerFunc":
		return true
	default:
		return false
	}
}

func (e *extractor) expandIOHelper(call *ast.CallExpr, name, field string, method *marshalInfo) []sourceOperation {
	primN := func(code string, count int) []sourceOperation {
		result := make([]sourceOperation, count)
		axis := []string{"X", "Y", "Z"}
		for i := range result {
			path := field
			if count > 1 {
				path += "." + axis[i]
			}
			result[i] = sourceOperation{Kind: "primitive", Field: path, Code: code}
		}
		return result
	}
	switch name {
	case "Vec3":
		return primN("f32le", 3)
	case "Vec2":
		return primN("f32le", 2)
	case "BlockPos":
		return primN("zigzag_i32", 3)
	case "ChunkPos":
		return primN("zigzag_i32", 2)
	case "SubChunkPos":
		return primN("i32le", 3)
	case "SoundPos":
		return primN("zigzag_i32", 3)
	case "ByteFloat":
		return primN("u8", 1)
	case "UUID":
		return []sourceOperation{{Kind: "uuid", Field: field, Length: 16}}
	case "String":
		return []sourceOperation{{Kind: "string", Field: field, Prefix: "var_u32", Encoding: "utf8"}}
	case "StringUTF":
		return []sourceOperation{{Kind: "string", Field: field, Prefix: "i16le", Encoding: "utf8"}}
	case "ByteSlice":
		return []sourceOperation{{Kind: "bytes", Field: field, Prefix: "var_u32"}}
	case "NBT", "NBTList":
		return []sourceOperation{{Kind: "primitive", Field: field, Code: "nbt_le"}}
	case "Bytes":
		return []sourceOperation{{Kind: "primitive", Field: field, Code: "raw_bytes"}}
	case "Bitset":
		if len(call.Args) < 2 {
			return []sourceOperation{e.unresolved(call, method.Key, field, "bitset size argument is missing", name)}
		}
		length, ok := e.literalInt(call.Args[1], method.File)
		if !ok || length <= 0 {
			return []sourceOperation{e.unresolved(call, method.Key, field, "bitset size is not a static integer", e.nodeString(call.Args[1]))}
		}
		return []sourceOperation{{Kind: "bitset", Field: field, Length: uint64(length)}}
	default:
		return []sourceOperation{e.unresolved(call, method.Key, field, "opaque or unsupported IO helper", name)}
	}
}

func (e *extractor) expandProtocol(call *ast.CallExpr, name string, method *marshalInfo, env map[string]typeRef, base string, depth int, stack map[string]bool) []sourceOperation {
	if prefix, ok := sourceArrayPrefixes[name]; ok {
		if len(call.Args) < 2 {
			return []sourceOperation{e.unresolved(call, method.Key, base, "slice helper has too few arguments", name)}
		}
		field := e.fieldPath(call.Args[1], method, env, base)
		ref := e.resolveExprType(call.Args[1], method, env)
		for ref.Kind == "pointer" && ref.Elem != nil {
			ref = *ref.Elem
		}
		if (ref.Kind != "slice" && ref.Kind != "array") || ref.Elem == nil {
			return []sourceOperation{e.unresolved(call, method.Key, field, "slice field type could not be resolved", e.nodeString(call.Args[1]))}
		}
		element := *ref.Elem
		var children []sourceOperation
		if strings.HasPrefix(name, "Func") || name == "SliceOfFunc" {
			if len(call.Args) < 3 {
				children = []sourceOperation{e.unresolved(call, method.Key, field+"[]", "slice callback is missing", name)}
			} else {
				children = e.callbackOps(call.Args[2], element, method, env, field+"[]", depth, stack)
			}
		} else {
			children = e.expandRef(element, field+"[]", depth+1, stack, call, method)
		}
		return []sourceOperation{{Kind: "array", Field: field, Prefix: prefix, Element: children}}
	}
	if name == "SliceOfLen" || name == "FuncSliceOfLen" || name == "FuncIOSliceOfLen" {
		if len(call.Args) < 3 {
			return []sourceOperation{e.unresolved(call, method.Key, base, "slice-of-length helper has too few arguments", name)}
		}
		field := e.fieldPath(call.Args[2], method, env, base)
		ref := e.resolveExprType(call.Args[2], method, env)
		for ref.Kind == "pointer" && ref.Elem != nil {
			ref = *ref.Elem
		}
		if (ref.Kind != "slice" && ref.Kind != "array") || ref.Elem == nil {
			return []sourceOperation{e.unresolved(call, method.Key, field, "slice field type could not be resolved", e.nodeString(call.Args[2]))}
		}
		length, ok := e.literalInt(call.Args[1], method.File)
		if !ok || length <= 0 {
			return []sourceOperation{e.unresolved(call, method.Key, field, "slice helper uses an external runtime length that cannot be represented", e.nodeString(call.Args[1]))}
		}
		var children []sourceOperation
		if name == "SliceOfLen" {
			children = e.expandRef(*ref.Elem, field+"[]", depth+1, stack, call, method)
		} else if len(call.Args) < 4 {
			children = []sourceOperation{e.unresolved(call, method.Key, field+"[]", "slice callback is missing", name)}
		} else {
			children = e.callbackOps(call.Args[3], *ref.Elem, method, env, field+"[]", depth, stack)
		}
		return []sourceOperation{{Kind: "fixed_array", Field: field, Length: uint64(length), Element: children}}
	}
	if name == "Optional" || name == "OptionalFunc" || name == "DoubleOptionalFunc" || name == "OptionalMarshaler" {
		if len(call.Args) < 2 {
			return []sourceOperation{e.unresolved(call, method.Key, base, "optional helper has too few arguments", name)}
		}
		field := e.fieldPath(call.Args[1], method, env, base)
		ref := e.resolveExprType(call.Args[1], method, env)
		for ref.Kind == "pointer" && ref.Elem != nil {
			ref = *ref.Elem
		}
		if ref.Kind != "optional" || ref.Elem == nil {
			return []sourceOperation{e.unresolved(call, method.Key, field, "optional field type could not be resolved", e.nodeString(call.Args[1]))}
		}
		var children []sourceOperation
		if (name == "OptionalFunc" || name == "DoubleOptionalFunc") && len(call.Args) >= 3 {
			children = e.callbackOps(call.Args[2], *ref.Elem, method, env, field, depth, stack)
		} else {
			children = e.expandRef(*ref.Elem, field, depth+1, stack, call, method)
		}
		inner := sourceOperation{Kind: "optional", Field: field, Presence: "bool", Value: children}
		if name == "DoubleOptionalFunc" {
			return []sourceOperation{{Kind: "optional", Field: field, Presence: "bool", Value: []sourceOperation{inner}}}
		}
		return []sourceOperation{inner}
	}
	if name == "IntegerFunc" {
		if len(call.Args) < 2 {
			return []sourceOperation{e.unresolved(call, method.Key, base, "IntegerFunc has too few arguments", name)}
		}
		field := e.fieldPath(call.Args[0], method, env, base)
		callback, ok := unwrapIndex(call.Args[len(call.Args)-1]).(*ast.SelectorExpr)
		if !ok {
			return []sourceOperation{e.unresolved(call, method.Key, field, "IntegerFunc callback could not be resolved", e.nodeString(call))}
		}
		ioBase, ok := callback.X.(*ast.Ident)
		if !ok || ioBase.Name != method.IO {
			return []sourceOperation{e.unresolved(call, method.Key, field, "IntegerFunc callback is not an IO method", e.nodeString(call))}
		}
		code, ok := sourcePrimitive[callback.Sel.Name]
		if !ok {
			return []sourceOperation{e.unresolved(call, method.Key, field, "IntegerFunc callback is unsupported", callback.Sel.Name)}
		}
		return []sourceOperation{{Kind: "primitive", Field: field, Code: code}}
	}
	if name == "Single" {
		if len(call.Args) < 2 {
			return []sourceOperation{e.unresolved(call, method.Key, base, "Single helper has too few arguments", name)}
		}
		field := e.fieldPath(call.Args[1], method, env, base)
		ref := e.resolveExprType(call.Args[1], method, env)
		for ref.Kind == "pointer" && ref.Elem != nil {
			ref = *ref.Elem
		}
		return e.expandRef(ref, field, depth+1, stack, call, method)
	}
	return []sourceOperation{e.unresolved(call, method.Key, base, "unsupported protocol helper", name)}
}

func (e *extractor) callbackOps(fn ast.Expr, element typeRef, method *marshalInfo, env map[string]typeRef, field string, depth int, stack map[string]bool) []sourceOperation {
	fn = unwrapIndex(fn)
	if selector, ok := fn.(*ast.SelectorExpr); ok {
		if base, yes := selector.X.(*ast.Ident); yes && base.Name == method.IO {
			fake := &ast.CallExpr{Fun: selector, Args: []ast.Expr{ast.NewIdent(field)}}
			if code, yes := sourcePrimitive[selector.Sel.Name]; yes {
				return []sourceOperation{{Kind: "primitive", Field: field, Code: code}}
			}
			return e.expandIOHelper(fake, selector.Sel.Name, field, method)
		}
		if selector.Sel.Name == "Marshal" {
			ref := e.resolveExprType(selector.X, method, env)
			for ref.Kind == "pointer" && ref.Elem != nil {
				ref = *ref.Elem
			}
			if ref.Kind == "named" && e.marshals[ref.Name] != nil {
				return e.callbackDecl(e.marshals[ref.Name], element, field, depth, stack)
			}
		}
	}
	if literal, ok := fn.(*ast.FuncLit); ok {
		nextEnv := map[string]typeRef{}
		for key, value := range env {
			nextEnv[key] = value
		}
		nextMethod := *method
		if literal.Type.Params != nil {
			for index, param := range literal.Type.Params.List {
				for _, name := range param.Names {
					ref := e.parseType(param.Type, method.File)
					if isIORef(ref) {
						nextMethod.IO = name.Name
						nextEnv[name.Name] = ref
					} else if index == len(literal.Type.Params.List)-1 {
						nextEnv[name.Name] = element
					} else {
						nextEnv[name.Name] = ref
					}
				}
			}
		}
		return e.extractBlock(literal.Body.List, &nextMethod, nextEnv, field, depth+1, stack)
	}
	if ident, ok := fn.(*ast.Ident); ok && method.File.Pkg == "protocol" {
		if target := e.functions["protocol."+ident.Name]; target != nil {
			return e.callbackDecl(target, element, field, depth, stack)
		}
	}
	return []sourceOperation{e.unresolved(fn, method.Key, field, "slice/optional callback could not be resolved", e.nodeString(fn))}
}

func (e *extractor) callbackDecl(target *marshalInfo, element typeRef, field string, depth int, stack map[string]bool) []sourceOperation {
	context := *target
	env := map[string]typeRef{}
	if target.Recv != "" {
		env[target.Recv] = typeRef{Kind: "named", Name: target.Key}
	}
	params := target.Decl.Type.Params
	if params != nil {
		for index, param := range params.List {
			for _, name := range param.Names {
				ref := e.parseType(param.Type, target.File)
				if isIORef(ref) {
					context.IO = name.Name
					env[name.Name] = ref
				} else if index == len(params.List)-1 {
					env[name.Name] = element
				} else {
					env[name.Name] = ref
				}
			}
		}
	}
	if context.IO == "" {
		return []sourceOperation{e.unresolved(target.Decl, target.Key, field, "named callback IO parameter could not be resolved", target.Key)}
	}
	return e.extractBlock(target.Decl.Body.List, &context, env, field, depth+1, stack)
}

func (e *extractor) expandRef(ref typeRef, field string, depth int, stack map[string]bool, node ast.Node, method *marshalInfo) []sourceOperation {
	for ref.Kind == "pointer" && ref.Elem != nil {
		ref = *ref.Elem
	}
	switch ref.Kind {
	case "named":
		return e.expandType(ref.Name, field, depth, stack)
	case "primitive":
		if code, ok := sourceBuiltinPrimitive[ref.Name]; ok {
			return []sourceOperation{{Kind: "primitive", Field: field, Code: code}}
		}
		return []sourceOperation{e.unresolved(node, method.Key, field, "element primitive could not be resolved", ref.Name)}
	case "array":
		if ref.Elem == nil || ref.Len <= 0 {
			return []sourceOperation{e.unresolved(node, method.Key, field, "fixed array type is incomplete", ref.String())}
		}
		children := e.expandRef(*ref.Elem, field+"[]", depth+1, stack, node, method)
		return []sourceOperation{{Kind: "fixed_array", Field: field, Length: uint64(ref.Len), Element: children}}
	case "slice":
		return []sourceOperation{e.unresolved(node, method.Key, field, "bare slice has no known length prefix", ref.String())}
	default:
		return []sourceOperation{e.unresolved(node, method.Key, field, "element/value type could not be resolved", ref.String())}
	}
}

func (e *extractor) parseType(expr ast.Expr, sf *sourceFile) typeRef {
	switch current := expr.(type) {
	case *ast.Ident:
		if _, ok := sourceBuiltinPrimitive[current.Name]; ok {
			return typeRef{Kind: "primitive", Name: current.Name}
		}
		return typeRef{Kind: "named", Name: sf.Pkg + "." + current.Name}
	case *ast.SelectorExpr:
		if base, ok := current.X.(*ast.Ident); ok {
			pkg := sf.Imports[base.Name]
			if pkg == "" {
				pkg = base.Name
			}
			return typeRef{Kind: "named", Name: pkg + "." + current.Sel.Name}
		}
	case *ast.StarExpr:
		inner := e.parseType(current.X, sf)
		return typeRef{Kind: "pointer", Elem: &inner}
	case *ast.ArrayType:
		inner := e.parseType(current.Elt, sf)
		if current.Len == nil {
			return typeRef{Kind: "slice", Elem: &inner}
		}
		length, _ := e.literalInt(current.Len, sf)
		return typeRef{Kind: "array", Elem: &inner, Len: length}
	case *ast.IndexExpr:
		if isOptionalType(current.X) {
			inner := e.parseType(current.Index, sf)
			return typeRef{Kind: "optional", Elem: &inner}
		}
	case *ast.IndexListExpr:
		if isOptionalType(current.X) && len(current.Indices) == 1 {
			inner := e.parseType(current.Indices[0], sf)
			return typeRef{Kind: "optional", Elem: &inner}
		}
	}
	return typeRef{Kind: "unknown"}
}

func isOptionalType(expr ast.Expr) bool {
	switch current := expr.(type) {
	case *ast.Ident:
		return current.Name == "Optional"
	case *ast.SelectorExpr:
		return current.Sel.Name == "Optional"
	default:
		return false
	}
}

func isIORef(ref typeRef) bool {
	for ref.Kind == "pointer" && ref.Elem != nil {
		ref = *ref.Elem
	}
	return ref.Kind == "named" && shortType(ref.Name) == "IO"
}

func (e *extractor) resolveExprType(expr ast.Expr, method *marshalInfo, env map[string]typeRef) typeRef {
	switch current := expr.(type) {
	case *ast.UnaryExpr:
		return e.resolveExprType(current.X, method, env)
	case *ast.ParenExpr:
		return e.resolveExprType(current.X, method, env)
	case *ast.StarExpr:
		return e.resolveExprType(current.X, method, env)
	case *ast.Ident:
		if ref, ok := env[current.Name]; ok {
			return ref
		}
		return typeRef{Kind: "unknown"}
	case *ast.IndexExpr:
		ref := e.resolveExprType(current.X, method, env)
		for ref.Kind == "pointer" && ref.Elem != nil {
			ref = *ref.Elem
		}
		if (ref.Kind == "slice" || ref.Kind == "array") && ref.Elem != nil {
			return *ref.Elem
		}
		return typeRef{Kind: "unknown"}
	case *ast.SelectorExpr:
		ref := e.resolveExprType(current.X, method, env)
		for ref.Kind == "pointer" && ref.Elem != nil {
			ref = *ref.Elem
		}
		if ref.Kind != "named" {
			return typeRef{Kind: "unknown"}
		}
		return e.lookupField(ref.Name, current.Sel.Name, map[string]bool{})
	case *ast.CompositeLit:
		return e.parseType(current.Type, method.File)
	default:
		return typeRef{Kind: "unknown"}
	}
}

func (e *extractor) lookupField(key, name string, seen map[string]bool) typeRef {
	if seen[key] {
		return typeRef{Kind: "unknown"}
	}
	seen[key] = true
	info := e.types[key]
	if info == nil {
		return typeRef{Kind: "unknown"}
	}
	if ref, ok := info.Fields[name]; ok {
		return ref
	}
	for _, embedded := range info.Embedded {
		for embedded.Kind == "pointer" && embedded.Elem != nil {
			embedded = *embedded.Elem
		}
		if embedded.Kind == "named" {
			if ref := e.lookupField(embedded.Name, name, seen); ref.Kind != "unknown" {
				return ref
			}
		}
	}
	return typeRef{Kind: "unknown"}
}

func (e *extractor) fieldPath(expr ast.Expr, method *marshalInfo, env map[string]typeRef, fallback string) string {
	var parts []string
	var walk func(ast.Expr) bool
	walk = func(current ast.Expr) bool {
		switch value := current.(type) {
		case *ast.UnaryExpr:
			return walk(value.X)
		case *ast.ParenExpr:
			return walk(value.X)
		case *ast.IndexExpr:
			if !walk(value.X) {
				return false
			}
			parts = append(parts, "[]")
			return true
		case *ast.Ident:
			return value.Name == method.Recv || hasEnv(env, value.Name)
		case *ast.SelectorExpr:
			if !walk(value.X) {
				return false
			}
			parts = append(parts, value.Sel.Name)
			return true
		default:
			return false
		}
	}
	if !walk(expr) || len(parts) == 0 {
		return fallback
	}
	path := fallback
	for _, part := range parts {
		if part == "[]" {
			path += part
		} else if path == "" {
			path = part
		} else {
			path += "." + part
		}
	}
	return path
}

func hasEnv(env map[string]typeRef, name string) bool {
	_, ok := env[name]
	return ok
}

func (e *extractor) literalInt(expr ast.Expr, sf *sourceFile) (int, bool) {
	return e.evalConst(expr, 0, sf.Pkg)
}

func (t typeRef) String() string {
	if t.Kind == "named" || t.Kind == "primitive" {
		return t.Name
	}
	if t.Elem != nil {
		return t.Kind + "[" + t.Elem.String() + "]"
	}
	return t.Kind
}

func unwrapIndex(expr ast.Expr) ast.Expr {
	switch current := expr.(type) {
	case *ast.IndexExpr:
		return current.X
	case *ast.IndexListExpr:
		return current.X
	default:
		return expr
	}
}

func referencesIO(call *ast.CallExpr, ioName string) bool {
	for _, arg := range call.Args {
		if ident, ok := arg.(*ast.Ident); ok && ident.Name == ioName {
			return true
		}
		if selector, ok := arg.(*ast.SelectorExpr); ok {
			if base, ok := selector.X.(*ast.Ident); ok && base.Name == ioName {
				return true
			}
		}
	}
	return false
}

func (e *extractor) unresolved(node ast.Node, typ, field, reason, raw string) sourceOperation {
	file := ""
	line := 0
	site := ""
	if node != nil {
		position := e.fset.Position(node.Pos())
		file = filepath.ToSlash(position.Filename)
		line = position.Line
		site = file + ":" + strconv.Itoa(line)
	}
	e.diagnostics = append(e.diagnostics, diagnostic{Packet: e.packet, Type: typ, Field: field, File: file, Line: line, Reason: reason, Raw: raw})
	return sourceOperation{Kind: "unresolved", Field: field, Reason: reason, Site: site}
}

func (e *extractor) addDiagnostic(node ast.Node, typ, field, reason, raw string) {
	_ = e.unresolved(node, typ, field, reason, raw)
}

func (e *extractor) nodeString(node ast.Node) string {
	if node == nil {
		return ""
	}
	var builder strings.Builder
	_ = printer.Fprint(&builder, e.fset, node)
	result := builder.String()
	if len(result) > 300 {
		return result[:300] + "..."
	}
	return result
}

func shortType(name string) string {
	if index := strings.LastIndexByte(name, '.'); index >= 0 {
		return name[index+1:]
	}
	return name
}

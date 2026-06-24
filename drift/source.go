package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

// WireOp is a single serialization operation extracted from a packet's Marshal
// method, in wire order.
type WireOp struct {
	// Method is the IO method or helper invoked, e.g. "Uint8", "Varuint64", "OptionalFunc".
	Method string
	// Kind is "io" (io.Method(...)), "protocol" (protocol.Helper(io, ...)), or
	// "custom" (any other call that references the io variable).
	Kind string
	// Field is the pk field referenced via &pk.Field, if any.
	Field string
	// Inner is a nested io method passed as a function value, e.g. the "String" in
	// protocol.OptionalFunc(io, &x, io.String).
	Inner string
	// Guarded is true when the op sits inside an if/for/switch (conditional or loop).
	Guarded bool
}

// SourcePacket is a packet type discovered in gophertunnel's source, with the
// ordered wire operations of its Marshal method.
type SourcePacket struct {
	TypeName string
	IDConst  string
	// ID is the numeric packet id, resolved from IDConst via parsePacketIDs.
	// analyzePackets leaves it zero; the orchestrator fills it in.
	ID  int
	Ops []WireOp
}

// analyzePackets parses a gophertunnel packet source file and returns the packet
// types it defines together with the ordered wire ops of each Marshal method.
func analyzePackets(src []byte) ([]SourcePacket, error) {
	f, err := parser.ParseFile(token.NewFileSet(), "packet.go", src, 0)
	if err != nil {
		return nil, fmt.Errorf("parse packet source: %w", err)
	}

	idConsts := map[string]string{} // receiver type -> ID constant name
	hasID := map[string]bool{}      // receiver type -> declares an ID() method
	marshals := map[string]*ast.FuncDecl{}
	order := []string{}

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv == nil || len(fn.Recv.List) != 1 {
			continue
		}
		recv := receiverType(fn.Recv.List[0].Type)
		if recv == "" {
			continue
		}
		switch fn.Name.Name {
		case "ID":
			hasID[recv] = true
			if c := idConstFromReturn(fn); c != "" {
				idConsts[recv] = c
			}
		case "Marshal":
			if _, seen := marshals[recv]; !seen {
				order = append(order, recv)
			}
			marshals[recv] = fn
		}
	}

	var packets []SourcePacket
	for _, recv := range order {
		// A packet has both ID() and Marshal(); a type with only Marshal() is a
		// composite sub-type (e.g. PartyInfo) and must be skipped.
		if !hasID[recv] {
			continue
		}
		fn := marshals[recv]
		ioName := firstParamName(fn)
		var ops []WireOp
		walkStmts(fn.Body.List, ioName, false, &ops)
		packets = append(packets, SourcePacket{
			TypeName: recv,
			IDConst:  idConsts[recv],
			Ops:      ops,
		})
	}
	return packets, nil
}

// receiverType returns the named type of a method receiver, unwrapping a pointer.
func receiverType(expr ast.Expr) string {
	if star, ok := expr.(*ast.StarExpr); ok {
		expr = star.X
	}
	if id, ok := expr.(*ast.Ident); ok {
		return id.Name
	}
	return ""
}

// idConstFromReturn extracts the identifier from `return IDXxx` in an ID method.
func idConstFromReturn(fn *ast.FuncDecl) string {
	if fn.Body == nil {
		return ""
	}
	for _, stmt := range fn.Body.List {
		ret, ok := stmt.(*ast.ReturnStmt)
		if !ok || len(ret.Results) != 1 {
			continue
		}
		if id, ok := ret.Results[0].(*ast.Ident); ok {
			return id.Name
		}
	}
	return ""
}

// firstParamName returns the name of a function's first parameter (the io var).
func firstParamName(fn *ast.FuncDecl) string {
	if fn.Type.Params == nil || len(fn.Type.Params.List) == 0 || len(fn.Type.Params.List[0].Names) == 0 {
		return ""
	}
	return fn.Type.Params.List[0].Names[0].Name
}

// walkStmts walks a statement list in source order, appending wire ops. Ops found
// inside control-flow bodies are marked guarded.
func walkStmts(stmts []ast.Stmt, ioName string, guarded bool, ops *[]WireOp) {
	for _, stmt := range stmts {
		walkStmt(stmt, ioName, guarded, ops)
	}
}

func walkStmt(stmt ast.Stmt, ioName string, guarded bool, ops *[]WireOp) {
	switch s := stmt.(type) {
	case *ast.ExprStmt:
		if call, ok := s.X.(*ast.CallExpr); ok {
			collectCall(call, ioName, guarded, ops)
		}
	case *ast.AssignStmt:
		for _, rhs := range s.Rhs {
			if call, ok := rhs.(*ast.CallExpr); ok {
				collectCall(call, ioName, guarded, ops)
			}
		}
	case *ast.IfStmt:
		walkStmts(s.Body.List, ioName, true, ops)
		if s.Else != nil {
			walkStmt(s.Else, ioName, true, ops)
		}
	case *ast.ForStmt:
		walkStmts(s.Body.List, ioName, true, ops)
	case *ast.RangeStmt:
		walkStmts(s.Body.List, ioName, true, ops)
	case *ast.BlockStmt:
		walkStmts(s.List, ioName, guarded, ops)
	case *ast.SwitchStmt:
		for _, c := range s.Body.List {
			if cc, ok := c.(*ast.CaseClause); ok {
				walkStmts(cc.Body, ioName, true, ops)
			}
		}
	case *ast.TypeSwitchStmt:
		for _, c := range s.Body.List {
			if cc, ok := c.(*ast.CaseClause); ok {
				walkStmts(cc.Body, ioName, true, ops)
			}
		}
	}
}

// collectCall records a wire op if the call is io.Method(...) or references the
// io variable in its arguments. It does not descend into a recorded call's args.
func collectCall(call *ast.CallExpr, ioName string, guarded bool, ops *[]WireOp) {
	op := WireOp{Guarded: guarded}
	fun := unwrapIndex(call.Fun)

	switch f := fun.(type) {
	case *ast.SelectorExpr:
		if base, ok := f.X.(*ast.Ident); ok && base.Name == ioName {
			op.Kind, op.Method = "io", f.Sel.Name
		} else if base, ok := f.X.(*ast.Ident); ok && base.Name == "protocol" {
			if !referencesIO(call, ioName) {
				return
			}
			op.Kind, op.Method = "protocol", f.Sel.Name
		} else {
			if !referencesIO(call, ioName) {
				return
			}
			op.Kind, op.Method = "custom", f.Sel.Name
		}
	case *ast.Ident:
		if !referencesIO(call, ioName) {
			return
		}
		op.Kind, op.Method = "custom", f.Name
	default:
		return
	}

	op.Field = fieldFromArgs(call.Args)
	op.Inner = innerIOMethod(call.Args, ioName)
	*ops = append(*ops, op)
}

// unwrapIndex strips generic instantiation (Foo[T] / Foo[T,U]) to the base expr.
func unwrapIndex(expr ast.Expr) ast.Expr {
	switch e := expr.(type) {
	case *ast.IndexExpr:
		return e.X
	case *ast.IndexListExpr:
		return e.X
	default:
		return expr
	}
}

// referencesIO reports whether the io variable appears as a direct argument.
func referencesIO(call *ast.CallExpr, ioName string) bool {
	for _, arg := range call.Args {
		if id, ok := arg.(*ast.Ident); ok && id.Name == ioName {
			return true
		}
		// io.Method passed as a function value also counts.
		if sel, ok := arg.(*ast.SelectorExpr); ok {
			if base, ok := sel.X.(*ast.Ident); ok && base.Name == ioName {
				return true
			}
		}
	}
	return false
}

// fieldFromArgs returns the pk field name from the first &pk.Field argument.
func fieldFromArgs(args []ast.Expr) string {
	for _, arg := range args {
		unary, ok := arg.(*ast.UnaryExpr)
		if !ok || unary.Op != token.AND {
			continue
		}
		if sel, ok := unary.X.(*ast.SelectorExpr); ok {
			return sel.Sel.Name
		}
	}
	return ""
}

// innerIOMethod returns the method name of an io.Method function value passed as
// an argument, e.g. the "String" in OptionalFunc(io, &x, io.String).
func innerIOMethod(args []ast.Expr, ioName string) string {
	for _, arg := range args {
		sel, ok := arg.(*ast.SelectorExpr)
		if !ok {
			continue
		}
		if base, ok := sel.X.(*ast.Ident); ok && base.Name == ioName {
			return sel.Sel.Name
		}
	}
	return ""
}

// parsePacketIDs evaluates the numeric values of the packet ID constants defined
// in gophertunnel's id.go. The constants are declared in const blocks of the form
// `IDLogin = iota + 1` followed by bare identifiers (and occasional blank `_`
// entries that still consume an iota value).
func parsePacketIDs(src []byte) (map[string]int, error) {
	f, err := parser.ParseFile(token.NewFileSet(), "id.go", src, 0)
	if err != nil {
		return nil, fmt.Errorf("parse id source: %w", err)
	}
	ids := map[string]int{}
	for _, decl := range f.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.CONST {
			continue
		}
		var lastExpr ast.Expr
		for iota, spec := range gen.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			if len(vs.Values) > 0 {
				lastExpr = vs.Values[0]
			}
			val, err := evalIotaExpr(lastExpr, iota)
			if err != nil {
				return nil, err
			}
			for _, name := range vs.Names {
				if name.Name != "_" {
					ids[name.Name] = val
				}
			}
		}
	}
	return ids, nil
}

// evalIotaExpr evaluates a constant expression composed of integer literals, the
// iota identifier, and +/-/* binary operators. It is intentionally minimal: it
// only needs to cover the shapes gophertunnel uses in id.go.
func evalIotaExpr(expr ast.Expr, iota int) (int, error) {
	switch e := expr.(type) {
	case nil:
		return iota, nil
	case *ast.Ident:
		if e.Name == "iota" {
			return iota, nil
		}
		return 0, fmt.Errorf("unsupported identifier %q in const expr", e.Name)
	case *ast.BasicLit:
		if e.Kind != token.INT {
			return 0, fmt.Errorf("unsupported literal kind %s", e.Kind)
		}
		return strconv.Atoi(e.Value)
	case *ast.ParenExpr:
		return evalIotaExpr(e.X, iota)
	case *ast.BinaryExpr:
		x, err := evalIotaExpr(e.X, iota)
		if err != nil {
			return 0, err
		}
		y, err := evalIotaExpr(e.Y, iota)
		if err != nil {
			return 0, err
		}
		switch e.Op {
		case token.ADD:
			return x + y, nil
		case token.SUB:
			return x - y, nil
		case token.MUL:
			return x * y, nil
		default:
			return 0, fmt.Errorf("unsupported operator %s in const expr", e.Op)
		}
	default:
		return 0, fmt.Errorf("unsupported const expr %T", expr)
	}
}

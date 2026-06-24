package codegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

// packetIDsByNumber parses gophertunnel's id.go and returns a map from each
// numeric packet id to its constant name (e.g. 44 -> "IDAnimate").
func packetIDsByNumber(src []byte) (map[int]string, error) {
	f, err := parser.ParseFile(token.NewFileSet(), "id.go", src, 0)
	if err != nil {
		return nil, fmt.Errorf("parse id source: %w", err)
	}
	out := map[int]string{}
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
					out[val] = name.Name
				}
			}
		}
	}
	return out, nil
}

// evalIotaExpr evaluates an integer constant expression of literals, iota, and
// +/-/* operators.
func evalIotaExpr(expr ast.Expr, iota int) (int, error) {
	switch e := expr.(type) {
	case nil:
		return iota, nil
	case *ast.Ident:
		if e.Name == "iota" {
			return iota, nil
		}
		return 0, fmt.Errorf("unsupported identifier %q", e.Name)
	case *ast.BasicLit:
		if e.Kind != token.INT {
			return 0, fmt.Errorf("unsupported literal %s", e.Kind)
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
			return 0, fmt.Errorf("unsupported operator %s", e.Op)
		}
	default:
		return 0, fmt.Errorf("unsupported expr %T", expr)
	}
}

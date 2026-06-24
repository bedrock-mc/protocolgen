package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// packetIDConstsInFile returns the ID constant names of packet types defined in
// a single source file. A packet type has both an ID() method returning an
// identifier and a Marshal method; composite sub-types (Marshal only) are skipped.
func packetIDConstsInFile(src []byte) []string {
	f, err := parser.ParseFile(token.NewFileSet(), "p.go", src, 0)
	if err != nil {
		return nil
	}
	idConst := map[string]string{}
	hasMarshal := map[string]bool{}
	var order []string
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv == nil || len(fn.Recv.List) != 1 {
			continue
		}
		recv := recvType(fn.Recv.List[0].Type)
		if recv == "" {
			continue
		}
		switch fn.Name.Name {
		case "ID":
			if c := returnIdent(fn); c != "" {
				if _, seen := idConst[recv]; !seen {
					order = append(order, recv)
				}
				idConst[recv] = c
			}
		case "Marshal":
			hasMarshal[recv] = true
		}
	}
	var consts []string
	for _, recv := range order {
		if hasMarshal[recv] {
			consts = append(consts, idConst[recv])
		}
	}
	return consts
}

func recvType(expr ast.Expr) string {
	if star, ok := expr.(*ast.StarExpr); ok {
		expr = star.X
	}
	if id, ok := expr.(*ast.Ident); ok {
		return id.Name
	}
	return ""
}

func returnIdent(fn *ast.FuncDecl) string {
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

// existingPacketFiles maps a numeric packet id to the existing source file that
// defines it, but only for files defining exactly one packet (so overwriting
// cannot clobber a second packet in the same file). nameToNum maps ID constant
// names to their numeric values.
func existingPacketFiles(dir string, nameToNum map[string]int) (map[int]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read packets dir: %w", err)
	}
	out := map[int]string{}
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		path := filepath.Join(dir, name)
		b, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		consts := packetIDConstsInFile(b)
		if len(consts) != 1 {
			continue // skip files with zero or multiple packets
		}
		if num, ok := nameToNum[consts[0]]; ok {
			out[num] = path
		}
	}
	return out, nil
}

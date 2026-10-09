package runtimegen

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

// runtimeFixture supplies independent registry, property and supplemental-tag inputs.
func runtimeFixture() (map[string]RegistryEntry, map[string]Properties, map[string][]string, Corrections) {
	return map[string]RegistryEntry{
		"minecraft:apple": {RuntimeID: 1, Version: 2},
		"minecraft:new":   {RuntimeID: 2, ComponentBased: true},
	}, map[string]Properties{
		"minecraft:apple": {MaxStackSize: 64},
		"minecraft:new":   {Tags: []string{"minecraft:harness"}},
	}, map[string][]string{
		"minecraft:allow_offhand": {"minecraft:apple"},
		"minecraft:head":          {"minecraft:apple"},
	}, Corrections{StackSizes: map[string]int{"minecraft:new": 16}, Offhand: []string{"minecraft:new"}, Evidence: map[string]string{"fixture": "synthetic"}}
}

func TestRuntimeCatalogCombinesSourcesDeterministically(t *testing.T) {
	registry, properties, tags, corrections := runtimeFixture()
	first, err := Generate(registry, properties, tags, corrections)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Generate(registry, properties, tags, corrections)
	if err != nil || !bytes.Equal(first, second) {
		t.Fatalf("generation is not deterministic: %v", err)
	}
	if !bytes.Contains(first, []byte(`Name: "minecraft:new", RuntimeID: 2, ComponentBased: true, Version: 0, MaxCount: 16, BodyWearable: true, AllowOffHand: true, ArmourSlot: 0, ArmourWearable: false`)) {
		t.Fatalf("registry, stack and capability inputs were not combined:\n%s", first)
	}
}

func TestRuntimeGeneratorHandlesEverySchemaField(t *testing.T) {
	r, p, tags, c := runtimeFixture()
	output, err := Generate(r, p, tags, c)
	if err != nil {
		t.Fatal(err)
	}
	file, err := parser.ParseFile(token.NewFileSet(), "runtime_generated.go", output, 0)
	if err != nil {
		t.Fatal(err)
	}
	// Read the sibling catalog schema without making either module import the other.
	definition, err := parser.ParseFile(token.NewFileSet(), "../../../generated/data/item/runtime.go", nil, 0)
	if err != nil {
		t.Fatalf("read Runtime schema from the repository checkout: %v", err)
	}
	schema := make(map[string]bool)
	ast.Inspect(definition, func(node ast.Node) bool {
		spec, ok := node.(*ast.TypeSpec)
		if !ok || spec.Name.Name != "Runtime" {
			return true
		}
		structure, ok := spec.Type.(*ast.StructType)
		if !ok {
			t.Fatal("Runtime must be a struct")
		}
		for _, field := range structure.Fields.List {
			if len(field.Names) == 0 {
				t.Fatal("Runtime has an embedded field that the generator cannot handle")
			}
			for _, name := range field.Names {
				schema[name.Name] = true
			}
		}
		return false
	})
	if len(schema) == 0 {
		t.Fatal("Runtime schema was not found or has no fields")
	}
	rows := 0
	ast.Inspect(file, func(node ast.Node) bool {
		literal, ok := node.(*ast.CompositeLit)
		if !ok || literal.Type != nil {
			return true
		}
		rows++
		fields := make(map[string]bool)
		for _, element := range literal.Elts {
			pair, ok := element.(*ast.KeyValueExpr)
			if !ok {
				t.Fatal("Runtime rows must name every field")
			}
			name := pair.Key.(*ast.Ident).Name
			if !schema[name] {
				t.Errorf("generator emits unknown Runtime.%s", name)
			}
			fields[name] = true
		}
		for name := range schema {
			if !fields[name] {
				t.Errorf("generator silently omitted Runtime.%s", name)
			}
		}
		return false
	})
	if rows != len(r) {
		t.Fatalf("checked %d Runtime rows, want %d", rows, len(r))
	}
}

func TestRuntimeCatalogRejectsMissingAndStaleFacts(t *testing.T) {
	for _, scenario := range []struct {
		name string
		edit func(map[string]RegistryEntry, map[string]Properties, map[string][]string, *Corrections)
		want string
	}{
		{"missing size", func(_ map[string]RegistryEntry, _ map[string]Properties, _ map[string][]string, c *Corrections) {
			delete(c.StackSizes, "minecraft:new")
		}, "missing or invalid stack size"},
		{"stale size", func(_ map[string]RegistryEntry, p map[string]Properties, _ map[string][]string, _ *Corrections) {
			p["minecraft:new"] = Properties{MaxStackSize: 16}
		}, "is stale"},
		{"removed item", func(r map[string]RegistryEntry, _ map[string]Properties, _ map[string][]string, _ *Corrections) {
			delete(r, "minecraft:new")
		}, "unknown item"},
		{"stale tag", func(_ map[string]RegistryEntry, _ map[string]Properties, tags map[string][]string, _ *Corrections) {
			tags["minecraft:allow_offhand"] = append(tags["minecraft:allow_offhand"], "minecraft:new")
		}, "stale or duplicated"},
		{"missing tag", func(_ map[string]RegistryEntry, _ map[string]Properties, tags map[string][]string, _ *Corrections) {
			delete(tags, "minecraft:head")
		}, "missing or empty"},
		{"duplicate ID", func(r map[string]RegistryEntry, _ map[string]Properties, _ map[string][]string, _ *Corrections) {
			r["minecraft:apple"] = r["minecraft:new"]
		}, "share runtime ID"},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			r, p, tags, c := runtimeFixture()
			scenario.edit(r, p, tags, &c)
			if _, err := Generate(r, p, tags, c); err == nil || !strings.Contains(err.Error(), scenario.want) {
				t.Fatalf("error = %v, want %q", err, scenario.want)
			}
		})
	}
}

// TestDecodeRegistry preserves zero values and rejects incomplete or duplicate identities.
func TestDecodeRegistry(t *testing.T) {
	const valid = `{"name":"minecraft:air","id":-158,"version":0,"componentBased":false}`
	rows, err := DecodeRegistry([]byte(`[` + valid + `]`))
	if err != nil || len(rows) != 1 || rows["minecraft:air"] != (RegistryEntry{RuntimeID: -158}) {
		t.Fatalf("registry = %v, error = %v", rows, err)
	}
	for _, input := range []string{
		`[` + valid + `,` + valid + `]`,
		`[{"name":"minecraft:air","id":-158,"version":0}]`,
		`[{"name":"minecraft:air","componentBased":false,"version":0}]`,
		`[{"name":"minecraft:air","id":-158,"componentBased":false}]`,
	} {
		if _, err := DecodeRegistry([]byte(input)); err == nil {
			t.Fatalf("accepted invalid registry %s", input)
		}
	}
}

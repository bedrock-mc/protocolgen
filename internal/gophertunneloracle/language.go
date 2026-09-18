package gophertunneloracle

import (
	"sort"
	"strconv"
	"strings"
)

// wireAutomaton is an epsilon-NFA over wire atoms built from a shapeExpr.
// Every expression is finite, so language equality is decided by a product
// walk over epsilon-closed state sets without enumerating paths.
type wireAutomaton struct {
	edges  [][]wireEdge
	eps    [][]int
	start  int
	accept int
}

type wireEdge struct {
	atom atom
	to   int
}

func buildAutomaton(expression shapeExpr) *wireAutomaton {
	a := &wireAutomaton{}
	a.start, a.accept = a.add(expression)
	return a
}

func (a *wireAutomaton) state() int {
	a.edges = append(a.edges, nil)
	a.eps = append(a.eps, nil)
	return len(a.edges) - 1
}

func (a *wireAutomaton) add(expression shapeExpr) (int, int) {
	switch expression.kind {
	case "token":
		start, end := a.state(), a.state()
		a.edges[start] = append(a.edges[start], wireEdge{atom: expression.atom, to: end})
		return start, end
	case "concat":
		start := a.state()
		end := start
		for _, part := range expression.parts {
			partStart, partEnd := a.add(part)
			a.eps[end] = append(a.eps[end], partStart)
			end = partEnd
		}
		return start, end
	case "alt":
		start, end := a.state(), a.state()
		for _, alternative := range expression.alts {
			altStart, altEnd := a.add(alternative)
			a.eps[start] = append(a.eps[start], altStart)
			a.eps[altEnd] = append(a.eps[altEnd], end)
		}
		return start, end
	default:
		start, end := a.state(), a.state()
		a.eps[start] = append(a.eps[start], end)
		return start, end
	}
}

func (a *wireAutomaton) closure(states []int) []int {
	seen := make(map[int]bool, len(states))
	stack := append([]int(nil), states...)
	for len(stack) > 0 {
		state := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if seen[state] {
			continue
		}
		seen[state] = true
		stack = append(stack, a.eps[state]...)
	}
	result := make([]int, 0, len(seen))
	for state := range seen {
		result = append(result, state)
	}
	sort.Ints(result)
	return result
}

// outgoing groups the edges leaving a state set by token, keeping the first
// atom seen for each token as its display.
func (a *wireAutomaton) outgoing(states []int) map[string]outgoingToken {
	result := map[string]outgoingToken{}
	for _, state := range states {
		for _, edge := range a.edges[state] {
			entry, ok := result[edge.atom.Token]
			if !ok {
				entry = outgoingToken{atom: edge.atom}
			}
			entry.targets = append(entry.targets, edge.to)
			result[edge.atom.Token] = entry
		}
	}
	return result
}

type outgoingToken struct {
	atom    atom
	targets []int
}

func (a *wireAutomaton) accepts(states []int) bool {
	for _, state := range states {
		if state == a.accept {
			return true
		}
	}
	return false
}

// completion returns the shortest atom sequence that leads from states to
// acceptance, for rendering a full example wire path.
func (a *wireAutomaton) completion(states []int) []atom {
	type entry struct {
		states []int
		parent int
		atom   atom
	}
	queue := []entry{{states: states, parent: -1}}
	seen := map[string]bool{stateKey(states): true}
	for index := 0; index < len(queue); index++ {
		current := queue[index]
		if a.accepts(current.states) {
			var result []atom
			for cursor := index; queue[cursor].parent >= 0; cursor = queue[cursor].parent {
				result = append(result, queue[cursor].atom)
			}
			for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
				result[i], result[j] = result[j], result[i]
			}
			return result
		}
		outgoing := a.outgoing(current.states)
		for _, token := range sortedTokens(outgoing) {
			next := a.closure(outgoing[token].targets)
			key := stateKey(next)
			if seen[key] {
				continue
			}
			seen[key] = true
			queue = append(queue, entry{states: next, parent: index, atom: outgoing[token].atom})
		}
	}
	return nil
}

func stateKey(states []int) string {
	parts := make([]string, len(states))
	for index, state := range states {
		parts[index] = strconv.Itoa(state)
	}
	return strings.Join(parts, ",")
}

func sortedTokens(outgoing map[string]outgoingToken) []string {
	tokens := make([]string, 0, len(outgoing))
	for token := range outgoing {
		tokens = append(tokens, token)
	}
	sort.Strings(tokens)
	return tokens
}

// languageWitness is the shortest wire path on which the two languages
// disagree: the shared prefix, then each side's next atom (nil when that side
// ends the packet) followed by its shortest completion.
type languageWitness struct {
	prefix       []atom
	manifest     []atom
	gophertunnel []atom
}

const wildcardVariantPrefix = "VARIANT:!"

// variantMatches reports whether a manifest variant token is accepted by a
// gophertunnel token, which may be a default-branch wildcard excluding the
// explicitly handled discriminants.
func variantMatches(manifestToken, sourceToken string) bool {
	if manifestToken == sourceToken {
		return true
	}
	if !strings.HasPrefix(sourceToken, wildcardVariantPrefix) || !strings.HasPrefix(manifestToken, "VARIANT:") {
		return false
	}
	value := strings.TrimPrefix(manifestToken, "VARIANT:")
	for _, excluded := range strings.Split(strings.TrimPrefix(sourceToken, wildcardVariantPrefix), ",") {
		if excluded == value {
			return false
		}
	}
	return true
}

// compareLanguages decides whether two wire languages are equal and returns a
// deterministic witness when they are not.
func compareLanguages(want, got shapeExpr) *languageWitness {
	manifest := buildAutomaton(want)
	source := buildAutomaton(got)
	type pair struct {
		manifest []int
		source   []int
		parent   int
		atom     atom
	}
	queue := []pair{{manifest: manifest.closure([]int{manifest.start}), source: source.closure([]int{source.start}), parent: -1}}
	seen := map[string]bool{stateKey(queue[0].manifest) + "|" + stateKey(queue[0].source): true}
	prefixOf := func(index int) []atom {
		var result []atom
		for cursor := index; queue[cursor].parent >= 0; cursor = queue[cursor].parent {
			result = append(result, queue[cursor].atom)
		}
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
		return result
	}
	for index := 0; index < len(queue); index++ {
		current := queue[index]
		manifestOut := manifest.outgoing(current.manifest)
		sourceOut := source.outgoing(current.source)
		manifestAccepts := manifest.accepts(current.manifest)
		sourceAccepts := source.accepts(current.source)
		if manifestAccepts != sourceAccepts {
			witness := &languageWitness{prefix: prefixOf(index)}
			if !manifestAccepts {
				witness.manifest = manifest.completion(current.manifest)
			}
			if !sourceAccepts {
				witness.gophertunnel = source.completion(current.source)
			}
			return witness
		}
		var unmatchedManifest, unmatchedSource []outgoingToken
		for _, token := range sortedTokens(manifestOut) {
			matched := false
			for _, sourceToken := range sortedTokens(sourceOut) {
				if !variantMatches(token, sourceToken) {
					continue
				}
				matched = true
				nextManifest := manifest.closure(manifestOut[token].targets)
				nextSource := source.closure(sourceOut[sourceToken].targets)
				key := stateKey(nextManifest) + "|" + stateKey(nextSource)
				if seen[key] {
					continue
				}
				seen[key] = true
				queue = append(queue, pair{manifest: nextManifest, source: nextSource, parent: index, atom: manifestOut[token].atom})
			}
			if !matched {
				unmatchedManifest = append(unmatchedManifest, manifestOut[token])
			}
		}
		for _, sourceToken := range sortedTokens(sourceOut) {
			matched := false
			for _, token := range sortedTokens(manifestOut) {
				if variantMatches(token, sourceToken) {
					matched = true
					break
				}
			}
			if !matched {
				unmatchedSource = append(unmatchedSource, sourceOut[sourceToken])
			}
		}
		if len(unmatchedManifest) == 0 && len(unmatchedSource) == 0 {
			continue
		}
		// A side without an unmatched atom is reported as its shortest valid
		// continuation, so the witness shows what it expects instead.
		witness := &languageWitness{prefix: prefixOf(index)}
		if len(unmatchedManifest) > 0 {
			witness.manifest = append([]atom{unmatchedManifest[0].atom}, manifest.completion(manifest.closure(unmatchedManifest[0].targets))...)
		} else {
			witness.manifest = manifest.completion(current.manifest)
		}
		if len(unmatchedSource) > 0 {
			witness.gophertunnel = append([]atom{unmatchedSource[0].atom}, source.completion(source.closure(unmatchedSource[0].targets))...)
		} else {
			witness.gophertunnel = source.completion(current.source)
		}
		return witness
	}
	return nil
}

// isLinear reports whether an expression has no alternatives, in which case
// its token count is a meaningful operation count.
func isLinear(expression shapeExpr) bool {
	switch expression.kind {
	case "alt":
		return false
	case "concat":
		for _, part := range expression.parts {
			if !isLinear(part) {
				return false
			}
		}
	}
	return true
}

func tokenCount(expression shapeExpr) int {
	switch expression.kind {
	case "token":
		return 1
	case "concat":
		total := 0
		for _, part := range expression.parts {
			total += tokenCount(part)
		}
		return total
	default:
		return 0
	}
}

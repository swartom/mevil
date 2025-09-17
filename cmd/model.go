package cmd

import ()

type Module struct {
	Letter     rune
	Parameters [][]string
}

type Rule struct {
	id          string
	predecessor Module
	condition   []string
	sucessor    []Module

	constructions [][]string
}

type System struct {
	// Metadata
	name   string
	inputs []rune

	// Raw Data

	alphabet []rune // raw Alphabet
	space    []rune // raw space

	axiom []Module // Axiom
	rules []Rule   // raw Rules
}

type StaticAnalysis struct {
}

const (
	VERTEX   = iota
	EDGE     = iota
	VARIABLE = iota

)

type ParserRule struct {
	ruleType int
	letter   rune
}

type Parser struct {
}

// Constructions

// Letter
// Params Length
// conditional
// Returnable

type Node struct {
	Letter    rune
	Length    int
	Condition []string

	id int

	Children *[]Node
}

type MetaSyntaxTree []Node

func (g *System) ConstructTree() (m MetaSyntaxTree) {

	return
}

var TestModel = System{
	name:     "scale-free spanning tree",
	inputs:   []rune{'k', 'a', 'b'},
	alphabet: []rune{'A', 'L'},
	space:    []rune{'x', 'y'},
	axiom:    []Module{Module{'A', [][]string{[]string{"1"}, []string{"k"}}}},
	rules: []Rule{
		Rule{id: "p_1",
			predecessor: Module{'A', [][]string{[]string{"x"}, []string{"y"}}},
			condition:   []string{"x", "!", "=", "y"},
			sucessor: []Module{
				Module{'L', [][]string{[]string{"r"}}},
				Module{'A', [][]string{[]string{"x"}, []string{"r", "-", "1"}}},
				Module{'A', [][]string{[]string{"r"}, []string{"y"}}},
			},
			constructions: [][]string{
				[]string{"float", "q", "=", "draw", "S"},
				[]string{"integer", "r", "=", "q", "*", "(", "y", "-", "x", ")", "+", "x", "+", "1"},
			},
		},
	},
}

// Na\"ively

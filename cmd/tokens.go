package cmd

type Kind string

// Overloaded Kind of variable type string
const (
	EOF                  Kind = ""
	RESERVED_WORD_STRING Kind = "string"
)

type Token struct {
	Kind  Kind
	Value string
}

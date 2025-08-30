package cmd

type Kind int

// Overloaded Kind of variable type string
const (
	EOF              Kind = iota
	TYPE             Kind = iota
	NAMESPACE        Kind = iota
	ARITHMETIC_LOGIC Kind = iota
	FLOW_LOGIC       Kind = iota
	SEPARATOR_VALUE  Kind = iota
	SPECIAL_CONTROL  Kind = iota
)

const (
	RESERVED_WORD_INTEGER  string = "integer"
	RESERVED_WORD_STRING   string = "string"
	RESERVED_WORD_RULE     string = "rule"
	RESERVED_WORD_LIST     string = "list"
	RESERVED_WORD_PROPERTY string = "property"
	RESERVED_WORD_ALPHABET string = "alphabet"
)

type Token struct {
	Kind  Kind
	Value string
}

func Classify(s string, t *Token) {

	switch s {
	case RESERVED_WORD_INTEGER:
		fallthrough
	case RESERVED_WORD_STRING:
		fallthrough
	case RESERVED_WORD_RULE:
		t.Kind = TYPE
	case "":
		t.Kind = EOF
	}
	t.Value = s

	return
}

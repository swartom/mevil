package cmd

// This is the language specification for MEVILang

const (
	_PROGRAM_       = iota // <program> ::= <statementlist>;
	_STATEMENTLIST_ = iota // <statementlist> ::= <statement>;<statementlist>|<statement>
	_STATEMENT_     = iota // <statement> ::= <metadata>|<primitive>|<variable>|<return>|<comment><statement>
)

// Under _STATEMENT_
const (
	_METADATA_  = iota
	_PRIMITIVE_ = iota
	_VARIABLE_  = iota
	_RETURN_    = iota
	_COMMENT_   = iota
)

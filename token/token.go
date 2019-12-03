package token

const (
	ILLEGAL = "illegal"
	EOF     = "EOF"

	AND = "&&"
	OR  = "||"

	IDENT  = "IDENT" // identifier
	INT    = "INT"
	STRING = "STRING"
	NULL   = "NULL"
	TRUE   = "TRUE"
	FALSE  = "FALSE"

	LT  = "<"
	LTE = "<="
	GT  = ">"
	GTE = ">="
	IN  = "in"

	EQ     = "=="
	NOT_EQ = "!="

	BANG = "!"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACKET = "["
	RBRACKET = "]"
)

var (
	keywords = map[string]TokenType{
		"true":  TRUE,
		"false": FALSE,
	}
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// check if identifier is keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	} else {
		return IDENT
	}
}

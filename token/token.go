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

	LT     = "<"
	LTE    = "<="
	GT     = ">"
	GTE    = ">="
	IN     = "in"
	EQ     = "=="
	NOT_EQ = "!="

	LPAREN   = "("
	RPAREN   = ")"
	LBRACKET = "["
	RBRACKET = "]"
	COMMA    = ","
)

var (
	keywords = map[string]TokenType{
		"true":  TRUE,
		"false": FALSE,
		"null":  NULL,
		"in":    IN,
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

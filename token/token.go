package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER" // add, foo, x, y
	INT        = "INT"        // 12312

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE   = "{"
	RBRACE   = "}"
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

func GetIdentifierToken(identifier string) Type {
	if t, ok := keywords[identifier]; ok {
		return t
	}

	return IDENTIFIER
}

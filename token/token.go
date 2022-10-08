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

	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	EXCLAMATION = "!"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	FOR      = "FOR"

	GT  = ">"
	LT  = "<"
	EQ  = "=="
	NE  = "!="
	GTE = ">="
	LTE = "<="

	FSLASH   = "/"
	ASTERISK = "*"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"for":    FOR,
}

func GetIdentifierToken(identifier string) Type {
	if t, ok := keywords[identifier]; ok {
		return t
	}

	return IDENTIFIER
}

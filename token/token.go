package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	END     = "END"
	ILLEGAL = "ILLEGAL"
	EOL     = "EOL"

	INTID    = "INTID"
	STRINGID = "STR_IDENT"
	INT      = "INT"
	STRING   = "STR"

	LET   = "LET"
	PRINT = "PRINT"
	GOTO  = "GOTO"

	ASSIGN   = "="
	PLUS     = "+"
	SUBTRACT = "-"
	LESS     = "<"
	GREATER  = ">"

	COMMA     = ","
	SEMICOLON = ";"
)

var keywords = map[string]Type{
	"PRINT": PRINT,
	"LET":   LET,
	"GOTO":  GOTO,
	"END":   END,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	if ident[len(ident)-1] == '$' {
		return STRINGID
	}
	return INTID
}

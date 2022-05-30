package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	END     = "END"
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	INTID    = "INTID"
	STRINGID = "STR_IDENT"
	INT      = "INT"
	STRING   = "STR"

	LET    = "LET"
	PRINT  = "PRINT"
	MOD    = "MOD"
	GOTO   = "GOTO"
	IF     = "IF"
	THEN   = "THEN"
	ELSE   = "ELSE"
	ELSEIF = "ELSEIF"
	TRUE   = "TRUE"
	BOOL   = "BOOL"

	EQ    = "EQUAL"
	NOTEQ = "NOT_EQUAL"
	GTEQ  = "GREATER_THAN_OR_EQUAL"
	LTEQ  = "LESS_THAN_OR_EQUAL"

	EOL      = "EOL"
	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	SUBTRACT = "SUBTRACT"
	LT       = "LESS_THAN"
	GT       = "GREATER_THAN"
	OR       = "OR"
	AND      = "AND"
	NOT      = "NOT"

	COMMA     = "COMMA"
	COLON     = "COLON"
	SEMICOLON = "SEMICOLON"
)

var keywords = func() map[string]Type {
	vs := []Type{END, LET, PRINT, MOD, GOTO, IF, THEN, ELSE, ELSEIF}
	kws := make(map[string]Type, len(vs)+2)
	for _, v := range vs {
		kws[string(v)] = v
	}
	kws["TRUE"] = BOOL
	kws["FALSE"] = BOOL
	return kws
}()

var singleChr = map[string]Type{
	"\n": EOL,
	"=":  ASSIGN,
	"+":  PLUS,
	"-":  SUBTRACT,
	"<":  LT,
	">":  GT,
	"|":  OR,
	"&":  AND,
	"!":  NOT,
	",":  COMMA,
	":":  COLON,
	";":  SEMICOLON,
}

var dualChr = map[string]Type{
	"==": EQ,
	"<>": NOTEQ,
	">=": GTEQ,
	"<=": LTEQ,
}

func LookupChr(chr byte) Type {
	if chr == 0 {
		return END
	}
	if tokenType, exists := singleChr[string(chr)]; exists {
		return tokenType
	}
	return Type("")
}

func LookupDualChr(first byte, sec byte) (Type, string) {
	lit := string([]byte{first, sec})
	if tokenType, exists := dualChr[lit]; exists {
		return tokenType, lit
	}
	return Type(""), ""
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

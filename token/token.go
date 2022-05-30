package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	END     = "END"
	ILLEGAL = "ILLEGAL"

	INTID    = "INTID"
	STRINGID = "STR_IDENT"
	INT      = "INT"
	STRING   = "STR"

	LET    = "LET"
	PRINT  = "PRINT"
	GOTO   = "GOTO"
	IF     = "IF"
	THEN   = "THEN"
	ELSE   = "ELSE"
	ELSEIF = "ELSEIF"

	EOL = "\n"

	ASSIGN   = "="
	PLUS     = "+"
	SUBTRACT = "-"
	LT       = "<"
	GT       = ">"

	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
)

var keywords = func() map[string]Type {
	vs := []Type{PRINT, LET, GOTO, END, IF, THEN, ELSE, ELSEIF}
	kws := make(map[string]Type, len(vs))
	for _, v := range vs {
		kws[string(v)] = v
	}
	return kws
}()

var singleChr = func() map[string]Type {
	vs := []Type{EOL, ASSIGN, PLUS, SUBTRACT, LT, GT, COMMA, COLON, SEMICOLON}
	scs := make(map[string]Type, len(vs))
	for _, v := range vs {
		scs[string(v)] = v
	}
	return scs
}()

func LookupChr(chr byte) Type {
	if chr == 0 {
		return END
	}
	if tokenType, exists := singleChr[string(chr)]; exists {
		return tokenType
	}
	return Type("")
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

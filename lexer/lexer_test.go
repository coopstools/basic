package lexer

import (
	"fmt"
	"github.com/coopstools/basic/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewToken(t *testing.T) {
	lexer := New(`10 LET ME$="JAMES"
15 LET X=0
20 PRINT X; ME$; "\n"
25 LET X = X + 1
30 IF X <> 10: GOTO 20: END IF
?% == < <= > >= TRUE FALSE
`)

	for i, test := range []struct {
		Token token.Type
		Char  string
	}{
		{token.INT, "10"},
		{token.LET, "LET"},
		{token.STRINGID, "ME$"},
		{token.ASSIGN, "="},
		{token.STRING, "JAMES"},
		{token.EOL, "\n"},

		{token.INT, "15"},
		{token.LET, "LET"},
		{token.INTID, "X"},
		{token.ASSIGN, "="},
		{token.INT, "0"},
		{token.EOL, "\n"},

		{token.INT, "20"},
		{token.PRINT, "PRINT"},
		{token.INTID, "X"},
		{token.SEMICOLON, ";"},
		{token.STRINGID, "ME$"},
		{token.SEMICOLON, ";"},
		{token.STRING, "\\n"},
		{token.EOL, "\n"},

		{token.INT, "25"},
		{token.LET, "LET"},
		{token.INTID, "X"},
		{token.ASSIGN, "="},
		{token.INTID, "X"},
		{token.PLUS, "+"},
		{token.INT, "1"},
		{token.EOL, "\n"},

		{token.INT, "30"},
		{token.IF, "IF"},
		{token.INTID, "X"},
		{token.NOTEQ, "<>"},
		{token.INT, "10"},
		{token.COLON, ":"},
		{token.GOTO, "GOTO"},
		{token.INT, "20"},
		{token.COLON, ":"},
		{token.END, "END"},
		{token.IF, "IF"},
		{token.EOL, "\n"},

		{token.ILLEGAL, "?"},
		{token.ILLEGAL, "%"},
		{token.EQ, "=="},
		{token.LT, "<"},
		{token.LTEQ, "<="},
		{token.GT, ">"},
		{token.GTEQ, ">="},
		{token.BOOL, "TRUE"},
		{token.BOOL, "FALSE"},
		{token.EOL, "\n"},
		{token.EOF, ""},
	} {
		tk := lexer.NextToken()
		assert.Equal(t, test.Token, tk.Type)
		assert.Equal(t, test.Char, tk.Literal)
		if t.Failed() {
			fmt.Println("Last character", tk, "at", i)
			return
		}
	}
}

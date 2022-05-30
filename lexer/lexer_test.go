package lexer

import (
	"github.com/coopstools/basic/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewToken(t *testing.T) {
	lexer := New(`10 LET ME$="JAMES"
15 LET X=0
20 PRINT X; ME$; "\n"
25 LET X = X + 1
30 GOTO 20
?%
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
		{token.GOTO, "GOTO"},
		{token.INT, "20"},
		{token.EOL, "\n"},

		{token.ILLEGAL, "?"},
		{token.ILLEGAL, "%"},
		{token.EOL, "\n"},
	} {
		tk := lexer.NextToken()
		assert.Equal(t, test.Token, tk.Type, "Failed on %d", i)
		assert.Equal(t, test.Char, tk.Literal, "Failed on %d", i)
	}
}

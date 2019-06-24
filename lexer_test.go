package lexer

import (
	"testing"
	"github.com/woshahua/chimpanzeelang/chimpanzee/token"
)


func TestNextToken(t *testing.T){
	input := `=+(){},;`  // use `` can ignore the symbol

	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	//TODO: what this mean, like new instance of input
	l := New(input)
	for i, tt := range tests{
		tok := l.NextToken()
		if tok.Type != tt.expectedType{
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}//else{
		//	t.Fatalf("tests[%d] - tokentype right. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		//}
	}
}

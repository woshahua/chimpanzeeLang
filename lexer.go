package lexer

import "github.com/woshahua/chimpanzeelang/chimpanzee/token"

type Lexer struct{
	input string
	position int // get current position of input string

	readPosition int // next reading position ?
	char byte // current char under examination
}

// return pointer of Lexer struct
func New(input string) *Lexer{
	l := &Lexer{input:input}
	l.readChar()
	return l
}

func (l *Lexer) readChar(){
	// if next position is over the input length, set char to 0
	if l.readPosition >= len(l.input){
		l.char = 0 // 0 is ascii code of "NUL", cause all syntax of chimpanzee is handle as string, so we only see ascii code
	} else{
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition = l.readPosition + 1
}

// return token of current char, and move to next char
func (l *Lexer) NextToken() token.Token{
	var tok token.Token

	switch l.char{
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{Type:tokenType, Literal:string(ch)}
}
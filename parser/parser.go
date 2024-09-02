package parser

import (
	"github.com/perryfranks/monkey-interpreter/ast"
	"github.com/perryfranks/monkey-interpreter/lexer"
	"github.com/perryfranks/monkey-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read 2 tokens to set our vars
	p.nextToken()
	p.nextToken()

	return p

}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

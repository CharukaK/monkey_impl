package parser

import (
	"github.com/charukak/monkey_impl/src/monkey/ast"
	"github.com/charukak/monkey_impl/src/monkey/lexer"
	"github.com/charukak/monkey_impl/src/monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	prg := &ast.Program{}
	prg.Statements = make([]ast.Statement, 0)

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			prg.Statements = append(prg.Statements, *stmt)
		}

		p.nextToken()
	}

	return prg
}

func (p *Parser) parseStatement() *ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}

	return stmt
}

func (p *Parser) expectPeekType(tokenType token.TokenType) bool {
	return p.peekToken.Type == tokenType
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

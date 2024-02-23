package ast

import "github.com/charukak/monkey_impl/src/monkey/token"

type ReturnStatement struct {
	Token token.Token // token.RETURN value
	Value Expression
}

func (s *ReturnStatement) statementNode() {}

func (s *ReturnStatement) TokenLiteral() string {
	return s.Token.Literal
}

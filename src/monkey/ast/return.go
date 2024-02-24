package ast

import (
	"bytes"

	"github.com/charukak/monkey_impl/src/monkey/token"
)

type ReturnStatement struct {
	Token token.Token // token.RETURN value
	Value Expression
}

func (s *ReturnStatement) statementNode() {}

func (s *ReturnStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")
	out.WriteString(s.Value.String())
	out.WriteString(";")

	return out.String()
}

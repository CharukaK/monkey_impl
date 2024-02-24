package ast

import (
	"bytes"

	"github.com/charukak/monkey_impl/src/monkey/token"
)

type LetStatement struct {
	Token token.Token // token.LET value
	Name  *Identifier
	Value Expression
}

func (s *LetStatement) statementNode() {

}

func (s *LetStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")
	out.WriteString(s.Name.String() + " ")
	out.WriteString("= ")
	out.WriteString(s.Value.String())
	out.WriteString(";")

	return out.String()
}

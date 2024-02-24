package ast

import "github.com/charukak/monkey_impl/src/monkey/token"

type IntegerLiteral struct {
	Token token.Token // token.INT
	Value int64
}

func (i *IntegerLiteral) expressionNode() {}

func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}

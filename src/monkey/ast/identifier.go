package ast

import "github.com/charukak/monkey_impl/src/monkey/token"

type Identifier struct {
	Token token.Token // token.IDENTIFIER
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}


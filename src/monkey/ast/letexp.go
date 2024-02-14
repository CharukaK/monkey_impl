package ast

import "github.com/charukak/monkey_impl/src/monkey/token"

type LetStatement struct {
	Token token.Token // token.LET value
	Name  *Identifier
	Value Expression
}

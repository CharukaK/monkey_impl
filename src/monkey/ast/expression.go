package ast

import "github.com/charukak/monkey_impl/src/monkey/token"

type ExpressionStatement struct {
	Token      token.Token // first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}


func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es != nil {
		return es.Expression.String()
	}

	return ""
}


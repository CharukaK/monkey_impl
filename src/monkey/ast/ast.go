package ast

type Node interface {
	TokenLiteral() string // returns the literal value of the token
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}


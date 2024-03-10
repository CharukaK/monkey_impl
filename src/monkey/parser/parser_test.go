package parser

import (
	"fmt"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/charukak/monkey_impl/src/monkey/ast"
	"github.com/charukak/monkey_impl/src/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := heredoc.Doc(`
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`)

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf(`Program statement does not contain 3 statements. got %d`, len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tc := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tc.expectedIdentifier) {
			return
		}
	}

}

func TestReturnStatements(t *testing.T) {
	input := heredoc.Doc(`
		return 5;
		return 10;
		return 838383;
	`)

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("prgram.Statements does not contain 3 statements, got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExp(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not an *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("expression is not *ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value is not %s. got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral is not %s, got=%s", "foobar", ident.TokenLiteral())
	}

}

func TestNumericLiteral(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not an *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	integerLit, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("stmt.Expression is not of type *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if integerLit.Value != 5 {
		t.Errorf("integerLit.Value is not %d. got=%d", 5, integerLit.Value)
	}

	if integerLit.TokenLiteral() != "5" {
		t.Errorf("integerLit.TokenLiteral is not %s. got=%s", "5", integerLit.TokenLiteral())
	}
}

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does contain %d statements, got=%d", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not of type ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not of type *ast.PrefixExpression. got %T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tc := range infixTests {
		l := lexer.New(tc.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not of type *ast.ExpressionStatement, got=%T ", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)

		if !ok {
			t.Fatalf("stmt.Expression is not of type *ast.InfixExpression, got=%T", stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, tc.leftValue) {
			return
		}

		if exp.Operator != tc.operator {
			t.Fatalf("exp.Operator is not %s, got=%s", tc.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tc.rightValue) {
			return
		}

	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			"-a * b;",
			"((-a) * b)",
		},
		{
			"!-a;",
			"(!(-a))",
		},
	}

	for _, tt := range testCases {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()

		if actual != tt.expected {
			t.Errorf("expected %q, got=%q", tt.expected, actual)
		}
	}

}

func checkParserErrors(t *testing.T, p *Parser) {
	if len(p.errors) == 0 {
		return
	}

	t.Errorf("parser had %d errors", len(p.errors))

	for _, e := range p.errors {
		t.Errorf("parser error: %q", e)
	}

	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Token literal not `let`. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("not a let statement, got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("identifier name is not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("il is not a *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value is not %d, got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got %s", value, integ.TokenLiteral())
	}

	return true
}

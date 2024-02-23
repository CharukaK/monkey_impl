package parser

import (
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

package lexer

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/charukak/monkey_impl/src/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){}"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Assign, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal type wrong. expected=%q, got%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenLangLike(t *testing.T) {
	input := heredoc.Doc(`
        let five = 5;
        let ten = 10;

        let add = fn(x, y) {
            x+y;
        };

        let result = add(five, ten);

        !-/*5;
        5 < 10 > 5;
    `)

    // some of the statements here wouldn't make sense semantically but this test is based on how we can 
    // identify tokens which will be the job of the lexer

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.Assign, "="},
		{token.NUMERIC, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.Assign, "="},
		{token.NUMERIC, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.Assign, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.Assign, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.NUMERIC, "5"},
		{token.SEMICOLON, ";"},
		{token.NUMERIC, "5"},
		{token.LT, "<"},
		{token.NUMERIC, "10"},
		{token.GT, ">"},
		{token.NUMERIC, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal type wrong. expected=%q, got%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

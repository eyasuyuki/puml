package lexer

import (
	"github.com/eyasuyuki/puml/token"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	input := `@startuml
skinparam linetype ortho
' hide the spot
hide circle

entity address <<Resource>> {
* id: bigint unsigned <<PK>>
}

entity address_history <<Event>> {
* id: bigint unsigned <<PK>>
* address_id: bigint unsigned <<FK>>
* status: enum ('ENABLE', 'DISABLE') = 'ENABLE'
* created_at: timestamp(6) = current_timestamp(6)
}

entity address_entry_type <<Resource>> {
* id: SERIAL <<PK>>
* name: text
* description: text
}

entity address_entry_type_history <<Event>> {
* id: SERIAL <<PK>>
* address_entry_type_id: bigint unsigned <<FK>>
* status: enum('ENABLE', 'DISABLE') = 'ENABLE'
* created_at: timestamp(6) = current_timestamp(6)
}

entity address_entry <<Resource>> {
* id: bigint unsigned <<PK>>
* address_entry_type_id: bigint unsigned <<FK>>
* value: text
}

entity address_detail <<Resource>> {
* id: bigint unsigned <<PK>>
* address_id: bigint unsigned <<FK>>
* address_entry_id: bigint unsigned <<FK>>
}

entity address_detail_history <<Event>> {
* id: bigint unsigned <<PK>>
* address_detail_id: bigint unsigned <<FK>>
* status: enum ('ENABLE', 'DISABLE') = 'ENABLE'
* created_at: timestamp(6) = current_timestamp(6)
}

address ||--|{ address_history

address_entry_type ||--|{ address_entry_type_history
address_entry_type ||--|{ address_entry
address ||--|{ address_detail
address_entry }|--|{ address_detail

address_detail ||--|{ address_detail_history

@enduml
`
	// DEBUG
	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.STARTUML, "@startuml"},
		{token.SKINPARAM, "skinparam"},
		{token.IDENT, "linetype"},
		{token.IDENT, "ortho"},
		{token.QUART, "'"},
		{token.HIDE, "hide"},
		{token.IDENT, "the"},
		{token.IDENT, "spot"},
		{token.HIDE,  "hide"},
		{token.CIRCLE, "circle"},
		{token.ENTITY,"entity"},
		{token.IDENT, "address"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "Resource"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.LBRACE, "{"},
		{token.ASTERISK, "*"},
		{token.IDENT, "id"},
		{token.COLON, ":"},
		{token.IDENT, "bigint"},
		{token.IDENT, "unsigned"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "PK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.RBRACE, "}"},
		{token.ENTITY, "entity"},
		{token.IDENT, "address_history"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "Event"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.LBRACE, "{"},
		{token.ASTERISK, "*"},
		{token.IDENT, "id"},
		{token.COLON, ":"},
		{token.IDENT, "bigint"},
		{token.IDENT, "unsigned"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "PK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.ASTERISK, "*"},
		{token.IDENT, "address_id"},
		{token.COLON, ":"},
		{token.IDENT, "bigint"},
		{token.IDENT, "unsigned"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "FK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.ASTERISK, "*"},
		{token.IDENT, "status"},
		{token.COLON, ":"},
		{token.ENUM, "enum"},
		{token.LPAREN, "("},
		{token.QUART, "'"},
		{token.IDENT, "ENABLE"},
		{token.QUART, "'"},
		{token.COMMA, ","},
		{token.QUART, "'"},
		{token.IDENT, "DISABLE"},
		{token.QUART, "'"},
		{token.RPAREN, ")"},
		{token.ASSIGN, "="},
		{token.QUART, "'"},
		{token.IDENT, "ENABLE"},
		{token.QUART, "'"},
		{token.ASTERISK, "*"},
		{token.IDENT, "created_at"},
		{token.COLON, ":"},
		{token.IDENT, "timestamp"},
		{token.LPAREN, "("},
		{token.INT, "6"},
		{token.RPAREN, ")"},
		{token.ASSIGN, "="},
		{token.IDENT, "current_timestamp"},
		{token.LPAREN, "("},
		{token.INT, "6"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.ENTITY, "entity"},
		{token.IDENT, "address_entry_type"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "Resource"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.LBRACE, "{"},
		{token.ASTERISK, "*"},
		{token.IDENT, "id"},
		{token.COLON, ":"},
		{token.IDENT, "SERIAL"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "PK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.ASTERISK, "*"},
		{token.IDENT, "name"},
		{token.COLON, ":"},
		{token.IDENT, "text"},
		{token.ASTERISK, "*"},
		{token.DESCRIPTION, "description"},
		{token.COLON, ":"},
		{token.IDENT, "text"},
		{token.RBRACE, "}"},
		{token.ENTITY, "entity"},
		{token.IDENT, "address_entry_type_history"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "Event"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.LBRACE, "{"},
		{token.ASTERISK, "*"},
		{token.IDENT, "id"},
		{token.COLON, ":"},
		{token.IDENT, "SERIAL"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "PK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.ASTERISK, "*"},
		{token.IDENT, "address_entry_type_id"},
		{token.COLON, ":"},
		{token.IDENT, "bigint"},
		{token.IDENT, "unsigned"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "FK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.ASTERISK, "*"},
		{token.IDENT, "status"},
		{token.COLON, ":"},
		{token.ENUM, "enum"},
		{token.LPAREN, "("},
		{token.QUART, "'"},
		{token.IDENT, "ENABLE"},
		{token.QUART, "'"},
		{token.COMMA, ","},
		{token.QUART, "'"},
		{token.IDENT, "DISABLE"},
		{token.QUART, "'"},
		{token.RPAREN, ")"},
		{token.ASSIGN, "="},
		{token.QUART, "'"},
		{token.IDENT, "ENABLE"},
		{token.QUART, "'"},
		{token.ASTERISK, "*"},
		{token.IDENT, "created_at"},
		{token.COLON, ":"},
		{token.IDENT, "timestamp"},
		{token.LPAREN, "("},
		{token.INT, "6"},
		{token.RPAREN, ")"},
		{token.ASSIGN, "="},
		{token.IDENT, "current_timestamp"},
		{token.LPAREN, "("},
		{token.INT, "6"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.ENTITY, "entity"},
		{token.IDENT, "address_entry"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "Resource"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.LBRACE, "{"},
		{token.ASTERISK, "*"},
		{token.IDENT, "id"},
		{token.COLON, ":"},
		{token.IDENT, "bigint"},
		{token.IDENT, "unsigned"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "PK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.ASTERISK, "*"},
		{token.IDENT, "address_entry_type_id"},
		{token.COLON, ":"},
		{token.IDENT, "bigint"},
		{token.IDENT, "unsigned"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "FK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.ASTERISK, "*"},
		{token.IDENT, "value"},
		{token.COLON, ":"},
		{token.IDENT, "text"},
		{token.RBRACE, "}"},
		{token.ENTITY, "entity"},
		{token.IDENT, "address_detail"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "Resource"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.LBRACE, "{"},
		{token.ASTERISK, "*"},
		{token.IDENT, "id"},
		{token.COLON, ":"},
		{token.IDENT, "bigint"},
		{token.IDENT, "unsigned"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "PK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.ASTERISK, "*"},
		{token.IDENT, "address_id"},
		{token.COLON, ":"},
		{token.IDENT, "bigint"},
		{token.IDENT, "unsigned"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "FK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.ASTERISK, "*"},
		{token.IDENT, "address_entry_id"},
		{token.COLON, ":"},
		{token.IDENT, "bigint"},
		{token.IDENT, "unsigned"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "FK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.RBRACE, "}"},
		{token.ENTITY, "entity"},
		{token.IDENT, "address_detail_history"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "Event"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.LBRACE, "{"},
		{token.ASTERISK, "*"},
		{token.IDENT, "id"},
		{token.COLON, ":"},
		{token.IDENT, "bigint"},
		{token.IDENT, "unsigned"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "PK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.ASTERISK, "*"},
		{token.IDENT, "address_detail_id"},
		{token.COLON, ":"},
		{token.IDENT, "bigint"},
		{token.IDENT, "unsigned"},
		{token.LT, "<"},
		{token.LT, "<"},
		{token.IDENT, "FK"},
		{token.GT, ">"},
		{token.GT, ">"},
		{token.ASTERISK, "*"},
		{token.IDENT, "status"},
		{token.COLON, ":"},
		{token.ENUM, "enum"},
		{token.LPAREN, "("},
		{token.QUART, "'"},
		{token.IDENT, "ENABLE"},
		{token.QUART, "'"},
		{token.COMMA, ","},
		{token.QUART, "'"},
		{token.IDENT, "DISABLE"},
		{token.QUART, "'"},
		{token.RPAREN, ")"},
		{token.ASSIGN, "="},
		{token.QUART, "'"},
		{token.IDENT, "ENABLE"},
		{token.QUART, "'"},
		{token.ASTERISK, "*"},
		{token.IDENT, "created_at"},
		{token.COLON, ":"},
		{token.IDENT, "timestamp"},
		{token.LPAREN, "("},
		{token.INT, "6"},
		{token.RPAREN, ")"},
		{token.ASSIGN, "="},
		{token.IDENT, "current_timestamp"},
		{token.LPAREN, "("},
		{token.INT, "6"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.IDENT, "address"},
		{token.PIPE, "|"},
		{token.PIPE, "|"},
		{token.MINUS, "-"},
		{token.MINUS, "-"},
		{token.PIPE, "|"},
		{token.LBRACE, "{"},
		{token.IDENT, "address_history"},
		{token.IDENT, "address_entry_type"},
		{token.PIPE, "|"},
		{token.PIPE, "|"},
		{token.MINUS, "-"},
		{token.MINUS, "-"},
		{token.PIPE, "|"},
		{token.LBRACE, "{"},
		{token.IDENT, "address_entry_type_history"},
		{token.IDENT, "address_entry_type"},
		{token.PIPE, "|"},
		{token.PIPE, "|"},
		{token.MINUS, "-"},
		{token.MINUS, "-"},
		{token.PIPE, "|"},
		{token.LBRACE, "{"},
		{token.IDENT, "address_entry"},
		{token.IDENT, "address"},
		{token.PIPE, "|"},
		{token.PIPE, "|"},
		{token.MINUS, "-"},
		{token.MINUS, "-"},
		{token.PIPE, "|"},
		{token.LBRACE, "{"},
		{token.IDENT, "address_detail"},
		{token.IDENT, "address_entry"},
		{token.RBRACE, "}"},
		{token.PIPE, "|"},
		{token.MINUS, "-"},
		{token.MINUS, "-"},
		{token.PIPE, "|"},
		{token.LBRACE, "{"},
		{token.IDENT, "address_detail"},
		{token.IDENT, "address_detail"},
		{token.PIPE, "|"},
		{token.PIPE, "|"},
		{token.MINUS, "-"},
		{token.MINUS, "-"},
		{token.PIPE, "|"},
		{token.LBRACE, "{"},
		{token.IDENT, "address_detail_history"},
		{token.ENDUML, "@enduml"},
		{token.EOF, ""},
	}


	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		//fmt.Printf("%v: %v\n", tt, tok)
		if tok.Type == token.EOF {
			break
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong: expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong: expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}


	}
}
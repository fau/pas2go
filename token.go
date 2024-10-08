// Turbo Pascal lexer tokens

package main

import (
	"strings"
)

// Token is the type of a single token.
type Token int

const (
	ILLEGAL Token = iota
	EOF

	// Symbols
	ASSIGN
	AT
	COLON
	COMMA
	DOT
	DOT_DOT
	EQUALS
	GREATER
	GTE
	LBRACKET
	LESS
	LPAREN
	LTE
	MINUS
	NOT_EQUALS
	PLUS
	POINTER
	RBRACKET
	RPAREN
	SEMICOLON
	SLASH
	STAR
	SLASH_SLASH
	LPAREN_STAR
	RPAREN_STAR

	// Keywords
	AND
	ARRAY
	BEGIN
	CASE
	CONST
	DIV
	DO
	DOWNTO
	ELSE
	END
	FALSE
	FILE
	FINALIZATION
	FOR
	FUNCTION
	GOTO
	IF
	IMPLEMENTATION
	IN
	INITIALIZATION
	INLINE
	INTERFACE
	INTERRUPT
	LABEL
	MOD
	NIL
	NOT
	OF
	OR
	PROCEDURE
	PROGRAM
	RECORD
	REPEAT
	SHL
	SHR
	THEN
	TO
	TRUE
	TYPE
	UNIT
	UNTIL
	USES
	VAR
	WHILE
	WITH
	XOR

	// Literals and names
	IDENT
	NUM
	HEX
	STR
)

var keywordTokens = map[string]Token{
	"AND":            AND,
	"ARRAY":          ARRAY,
	"BEGIN":          BEGIN,
	"CASE":           CASE,
	"CONST":          CONST,
	"DIV":            DIV,
	"DO":             DO,
	"DOWNTO":         DOWNTO,
	"ELSE":           ELSE,
	"END":            END,
	"FALSE":          FALSE,
	"FILE":           FILE,
	"FINALIZATION":   FINALIZATION,
	"FOR":            FOR,
	"FUNCTION":       FUNCTION,
	"GOTO":           GOTO,
	"IF":             IF,
	"IMPLEMENTATION": IMPLEMENTATION,
	"IN":             IN,
	"INITIALIZATION": INITIALIZATION,
	"INLINE":         INLINE,
	"INTERFACE":      INTERFACE,
	"INTERRUPT":      INTERRUPT,
	"LABEL":          LABEL,
	"MOD":            MOD,
	"NIL":            NIL,
	"NOT":            NOT,
	"OF":             OF,
	"OR":             OR,
	"PROCEDURE":      PROCEDURE,
	"PROGRAM":        PROGRAM,
	"RECORD":         RECORD,
	"REPEAT":         REPEAT,
	"SHL":            SHL,
	"SHR":            SHR,
	"THEN":           THEN,
	"TO":             TO,
	"TRUE":           TRUE,
	"TYPE":           TYPE,
	"UNIT":           UNIT,
	"UNTIL":          UNTIL,
	"USES":           USES,
	"VAR":            VAR,
	"WHILE":          WHILE,
	"WITH":           WITH,
	"XOR":            XOR,
}

// KeywordToken returns the token associated with the given keyword
// string, or ILLEGAL if given name is not a keyword.
func KeywordToken(name string) Token {
	return keywordTokens[strings.ToUpper(name)]
}

var tokenNames = map[Token]string{
	ILLEGAL: "<illegal>",
	EOF:     "EOF",

	ASSIGN:      ":=",
	AT:          "@",
	COLON:       ":",
	COMMA:       ",",
	DOT:         ".",
	DOT_DOT:     "..",
	EQUALS:      "=",
	GREATER:     ">",
	GTE:         ">=",
	LBRACKET:    "[",
	LESS:        "<",
	LPAREN:      "(",
	LTE:         "<=",
	MINUS:       "-",
	NOT_EQUALS:  "<>",
	PLUS:        "+",
	POINTER:     "^",
	RBRACKET:    "]",
	RPAREN:      ")",
	SEMICOLON:   ";",
	SLASH:       "/",
	STAR:        "*",
	SLASH_SLASH: "//",
	LPAREN_STAR: "(*",
	RPAREN_STAR: "*)",

	AND:            "AND",
	ARRAY:          "ARRAY",
	BEGIN:          "BEGIN",
	CASE:           "CASE",
	CONST:          "CONST",
	DIV:            "DIV",
	DO:             "DO",
	DOWNTO:         "DOWNTO",
	ELSE:           "ELSE",
	END:            "END",
	FALSE:          "FALSE",
	FILE:           "FILE",
	FINALIZATION:   "FINALIZATION",
	FOR:            "FOR",
	FUNCTION:       "FUNCTION",
	GOTO:           "GOTO",
	IF:             "IF",
	IMPLEMENTATION: "IMPLEMENTATION",
	IN:             "IN",
	INITIALIZATION: "INITIALIZATION",
	INLINE:         "INLINE",
	INTERFACE:      "INTERFACE",
	INTERRUPT:      "INTERRUPT",
	LABEL:          "LABEL",
	MOD:            "MOD",
	NIL:            "NIL",
	NOT:            "NOT",
	OF:             "OF",
	OR:             "OR",
	PROCEDURE:      "PROCEDURE",
	PROGRAM:        "PROGRAM",
	RECORD:         "RECORD",
	REPEAT:         "REPEAT",
	SHL:            "SHL",
	SHR:            "SHR",
	THEN:           "THEN",
	TO:             "TO",
	TRUE:           "TRUE",
	TYPE:           "TYPE",
	UNIT:           "UNIT",
	UNTIL:          "UNTIL",
	USES:           "USES",
	VAR:            "VAR",
	WHILE:          "WHILE",
	WITH:           "WITH",
	XOR:            "XOR",

	IDENT: "IDENT",
	NUM:   "NUM",
	HEX:   "HEX",
	STR:   "STR",
}

// String returns the string name of this token.
func (t Token) String() string {
	return tokenNames[t]
}

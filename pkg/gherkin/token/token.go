//go:generate stringer -type=TokenType

package token

type TokenType uint8

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = iota
	EOF
	FEATURE
	BACKGROUND
	SCENARIO
	GIVEN
	WHEN
	THEN
	AND
)

// keywords is a map of gherkin keywords in different languages to the corresponding tokens.
var keywords = map[string]TokenType{
	"feature":    FEATURE,
	"background": BACKGROUND,
	"scenario":   SCENARIO,
	"given":      GIVEN,
	"when":       WHEN,
	"then":       THEN,
	"and":        AND,
}

// LookupKeyword returns corresponding to the literal token.
func LookupKeyword(literal string) TokenType {
	if tok, ok := keywords[literal]; ok {
		return tok
	}
	return ILLEGAL
}

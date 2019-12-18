package lexer

import (
	"bytes"
	"github.com/proofhouse/proofhouse/pkg/gherkin/token"
	"testing"
)

var feature = `
Feature: agagaga
	test tst

	Scenario: gagagaga
		
		dsadasdsada
`

func TestAga(t *testing.T) {
	lexer := New(bytes.NewReader([]byte(feature)))

	i := 0
	for tok := lexer.NextToken(); i < 100 && tok.Type != token.EOF; {
		t.Logf("%v %v", tok.Type.String(), tok.Literal)
		i++
	}
}

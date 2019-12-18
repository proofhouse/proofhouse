package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/proofhouse/proofhouse/pkg/gherkin/token"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

// Lexer structure.
type Lexer struct {
	reader       *bufio.Reader
	position     int
	readPosition int
	rune         rune
	line         int
	linePos      int
}

// New creates new Lexer structure.
func New(reader io.Reader) *Lexer {
	lexer := &Lexer{
		reader: bufio.NewReader(reader),
	}

	var sb strings.Builder

	for {
		if rn, _, err := lexer.reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			sb.WriteRune(rn)
		}
	}

	aga := sb.String()
	_ = aga

	lexer.readRune()

	return lexer
}

// NextToken reads next sequence of runes and returns matching token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	var b bytes.Buffer
	_, _ = b.ReadFrom(l.reader)
	aga := b.String()
	_ = aga

	switch l.rune {
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if unicode.IsLetter(l.rune) {
			tok.Literal = l.readLiteral()
			tok.Type = token.LookupKeyword(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.rune)
		}
	}

	l.readRune()
	return tok
}

// readLiteral continuously reads runes until it reaches non-letter rune.
func (l *Lexer) readLiteral() string {
	var sb strings.Builder
	for unicode.IsLetter(l.rune) {
		sb.WriteRune(l.rune)
		l.readRune()
	}
	return sb.String()
}

// readRune reads next character as a rune and sets it to the current state. Also updates current line and cursor
// positions.
func (l *Lexer) readRune() {
	rn, _, err := l.reader.ReadRune()
	if err != nil {
		l.rune = 0
		fmt.Printf("ReadRune() error: %v", err)
		os.Exit(1)
	} else {
		fmt.Print(123)
		if l.rune == '\n' {
			l.line++
			l.linePos = 0
		}
		l.linePos++
		l.rune = rn
	}
}

// newToken creates new token.
func newToken(tokenType token.TokenType, rn rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(rn)}
}

package analysis

import (
	"bytes"
)

type Filter interface {
	Filter([]Token) []Token
}

type LowercaseFilter struct{}

func (filter *LowercaseFilter) Filter(tokens []Token) []Token {
	for i := 0; i < len(tokens); i++ {
		tokens[i].Text = bytes.ToLower(tokens[i].Text)
	}
	return tokens
}

type StripFilter struct{}

func (filter *StripFilter) Filter(tokens []Token) []Token {
	for i := 0; i < len(tokens); i++ {
		tokens[i].Text = bytes.TrimSpace(tokens[i].Text)
	}
	return tokens
}

package analysis

import (
	"regexp"
)

type Tokenizer interface {
	Token(value []byte) []Token
}

type IDTokenizer struct{}

func (tokenizer *IDTokenizer) Token(value []byte) []Token {
	t := Token{
		Text:  value,
		Boost: 1.0,
	}
	return []Token{t}
}

type RegexTokenizer struct {
	pattern *regexp.Regexp
}

func NewRegexTokenizer(exp string) (*RegexTokenizer, error) {
	pattern, err := regexp.Compile(exp)
	if err != nil {
		return nil, nil
	}

	return &RegexTokenizer{
		pattern: pattern,
	}, nil
}

func (tokenizer *RegexTokenizer) Token(value []byte) []Token {
	values := tokenizer.pattern.FindAll(value, -1)
	if values == nil {
		return nil
	}
	results := make([]Token, 0, len(values))
	for _, v := range values {
		results = append(results, Token{
			Text:  v,
			Boost: 1.0,
		})
	}
	return results
}

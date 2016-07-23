package analysis

import (
	"bytes"
	"testing"
)

func TestLowerFilter(t *testing.T) {
	tokens := []Token{
		Token{
			Text: []byte("Abc"),
		},
		Token{
			Text: []byte("aBc"),
		},
	}

	filter := &LowercaseFilter{}
	tokens = filter.Filter(tokens)
	for _, token := range tokens {
		if !bytes.Equal(token.Text, []byte("abc")) {
			t.Errorf("wrong in lowercase filter %v %v", string(token.Text), "abc")
		}
	}
}

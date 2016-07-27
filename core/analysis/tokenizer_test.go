package analysis

import "testing"

func verifyTokenizer(t *testing.T, tokenizer Tokenizer, value []byte, tokenResults []string) {
	tokens := tokenizer.Token(value)
	if tokens == nil {
		t.Errorf("error in token content")
	}
	if len(tokenResults) != len(tokens) {
		t.Errorf("error in return token length %v %v", tokens, tokenResults)
	}

	for index, token := range tokens {
		if string(token.Text) != tokenResults[index] {
			t.Errorf("error in compare tokens %v : %v ", token.Text, tokenResults[index])
		}
	}
}

func TestRegrexTokenizer(t *testing.T) {

	value := []byte("i am Chinese 4.56")
	tokenResults := []string{"i", "am", "Chinese", "4.56"}
	var tokenizer Tokenizer
	tokenizer, err := NewRegexTokenizer(`\w+(\.?\w*)*`)
	if err != nil {
		t.Errorf("error in building tokenizer")
	}

	verifyTokenizer(t, tokenizer, value, tokenResults)

	tokenizer = NewSpaceSeparatedTokenizer()

	verifyTokenizer(t, tokenizer, value, tokenResults)

}

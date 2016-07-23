package analysis

type Term struct {
	Value []byte
	Pos   int32
	Boost float32
}

type Analyzer interface {
	Parse(string) []Term
}

type CompositeAnalyzer struct {
	tokenizer Tokenizer
	filters   []Filter
}

func NewCompositeAnalyzer(tokenizer Tokenizer, filters ...Filter) *CompositeAnalyzer {
	return &CompositeAnalyzer{
		tokenizer: tokenizer,
		filters:   filters,
	}
}

func (analyzer *CompositeAnalyzer) Parse(doc []byte) []Token {
	tokens := analyzer.tokenizer.Token(doc)
	if analyzer == nil {
		return tokens
	}

	for _, filter := range analyzer.filters {
		tokens = filter.Filter(tokens)
	}
	return tokens
}

var IDAnalyzer *CompositeAnalyzer = NewCompositeAnalyzer(&IDTokenizer{}, &LowercaseFilter{})

package analysis

import "github.com/liuzz1983/scalesearch/utils/types"

type Term struct {
	Value []byte
	Pos   int32
	Boost float32
}

// Analyzer parse method need to convert type into it needs
type Analyzer interface {
	Parse(interface{}) ([]Token, error)
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

func (analyzer *CompositeAnalyzer) Parse(input interface{}) ([]Token, error) {
	value, err := types.ToBytes(input)
	if err != nil {
		return nil, err
	}
	tokens := analyzer.tokenizer.Token(value)
	if analyzer.filters == nil {
		return tokens, nil
	}

	for _, filter := range analyzer.filters {
		tokens = filter.Filter(tokens)
	}
	return tokens, nil
}

var IDAnalyzer *CompositeAnalyzer = NewCompositeAnalyzer(&IDTokenizer{}, &LowercaseFilter{})

var RegrexAnalyzer *CompositeAnalyzer = NewCompositeAnalyzer(NewSpaceSeparatedTokenizer())

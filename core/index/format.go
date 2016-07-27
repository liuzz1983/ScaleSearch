package index

import (
	"github.com/liuzz1983/scalesearch/core/analysis"
	"github.com/liuzz1983/scalesearch/core/compact"
)

// IndexValue for
// Token is the token
// frequency is the number of times "tokentext" appeared in the value, weight is the
// weight (a float usually equal to frequency in the absence of per-term boosts) and
// valuestring is encoded field-specific posting value for thetoken.
type IndexValue struct {
	Token     []byte
	Frequency uint32
	Weight    float32
	Value     []byte
}

// Format interface for representing storage format for a field or vector
type Format interface {
	FixValueSize() int
	WordValues(value string, analyzer analysis.Analyzer) []IndexValue
}
type BaseFormat struct {
	FieldBoost float32
}

type Positions struct {
	BaseFormat
}

const DefaultPosLen = 8

func (pos *Positions) WordValues(value string, analyzer analysis.Analyzer) ([]IndexValue, error) {
	fb := pos.FieldBoost
	weights := make(map[string]float32)
	poses := make(map[string][]int32)
	tokens, err := analyzer.Parse(value)
	if err != nil {
		return nil, err
	}
	for _, term := range tokens {
		name := string(term.Text)
		datas, ok := poses[name]
		if !ok {
			datas = make([]int32, 0, DefaultPosLen)
		}
		datas = append(datas, term.Pos)
		poses[name] = datas

		boost, ok := weights[name]
		if !ok {
			boost = 0.0
		}
		weights[name] = boost + term.Boost
	}

	result := make([]IndexValue, len(poses))
	for key, posList := range poses {
		value, err := pos.Encode(posList)
		if err != nil {
			return nil, err
		}
		result = append(result, IndexValue{
			Token:     []byte(key),
			Frequency: uint32(len(posList)),
			Weight:    weights[string(key)] * fb,
			Value:     value,
		})
	}
	return result, nil
}

func (pos *Positions) Encode(posList []int32) ([]byte, error) {
	encoder := compact.NewInt64Encoder()
	for _, pos := range posList {
		encoder.Write(int64(pos))
	}
	return encoder.Bytes()
}

func (pos *Positions) Decode(values []byte) ([]int32, error) {
	decoder := compact.NewInt64Decoder(values)
	result := make([]int32, 0, 256)
	for decoder.Next() {
		v := decoder.Read()
		result = append(result, int32(v))
	}
	err := decoder.Error()
	if err != nil {
		return nil, err
	}
	return result, nil
}

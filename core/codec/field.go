package codec

import (
	"github.com/liuzz1983/scalesearch/core/index"
	"github.com/liuzz1983/scalesearch/core/schema"
)

type FieldWriter interface {
	StartField(fieldName string, fieldObj interface{})
	FinishField()

	StartTerm(term string)
	FinishTerm()

	Add(docNum int64, weight float32, value []byte, length int) error
	AddSpellWord(fieldName string, text []byte) error

	AddPostings(writer FieldWriter, schema schema.Schema, lengths []int, items []interface{}) error

	Close()
}

type FieldCursor interface {
	First()
	Find(string)
	Next()
	Term()
}

type TermsReader interface {
	Contains(term *index.FieldTerm) bool
	Cursor(fieldName string, fieldObj interface{}) FieldCursor
	Terms() []index.FieldTerm
	TermsFrom(fieldName string, prefix []byte) []index.FieldTerm
	TermInfo(fieldName string, term []byte) *index.TermInfo
	Frequency(fieldName string, term []byte)
	DocFrequency(fieldName string, term []byte)
	Matcher(fieldName string, term []byte, formate interface{}, scorer interface{})
	IndexedFieldNames() []string
	Close() error
}

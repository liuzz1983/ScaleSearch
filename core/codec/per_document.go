package codec

import "github.com/liuzz1983/scalesearch/core/index"

type PerDocumentWriter interface {
	StartDoc(docNum int64) error
	FinishDoc() error

	AddField(fieldName string, fieldObj interface{}, value []byte, length int) error
	AddColumnValue(fieldName string, columnObj interface{}, value []byte) error
	AddVectorItems(fieldName string, fieldObe interface{}) ([]index.VectorItem, error)

	Close() error
}

type PerDocumentReader interface {
	Close() error

	// doc statistics
	DocCount() int
	DocCountAll() int
	HasDeletions() bool
	IsDeleted(docNum int64) bool
	DeletedDocs() []int64
	AllDocIds() []int64

	// columns
	SupportColumns() bool
	HasColumn(columnName string) bool
	ListCoumns() []string

	FieldDocs(fieldName string) []int64

	// Lengths
	DocFieldLength(docNum int64, fieldName string, defaultValue int)
	FieldLength(fieldName string)
	MinFieldLength(fieldName string)
	MaxFieldLength(fieldName string)

	// Vectors
	HasVector(docNum int64, fieldName string) bool
	Vector(docNum int64, fieldName string, format interface{}) index.VectorItem

	StoredFields(docNum int64) []string
	AllStoredFields() []string
}

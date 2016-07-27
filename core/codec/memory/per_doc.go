package memory

import (
	"github.com/liuzz1983/scalesearch/core/filedb"
	"github.com/liuzz1983/scalesearch/core/index"
)

type MemPerDocWriter struct {
	storage  filedb.Storage
	segment  *MemSegment
	isClosed bool
	columns  map[string]interface{}

	docCount int
	docNum   int64
	stored   map[string][]byte
	lengths  map[string]int
	vectors  map[string][]index.VectorItem
}

func NewMemPerDocWriter(storage filedb.Storage, segment *MemSegment) *MemPerDocWriter {
	return &MemPerDocWriter{
		storage:  storage,
		segment:  segment,
		isClosed: false,
		docCount: 0,
	}
}

func (writer *MemPerDocWriter) hasColumn(fieldName string) bool {
	_, ok := writer.columns[fieldName]
	return ok
}

func (writer *MemPerDocWriter) StartDoc(docNum int64) error {
	writer.docCount++
	writer.docNum = docNum
	writer.stored = make(map[string][]byte)
	writer.lengths = make(map[string]int)
	writer.vectors = make(map[string][]index.VectorItem)
	return nil
}

func (writer *MemPerDocWriter) FinishDoc() error {
	writer.segment.lock.Lock()
	defer writer.segment.lock.Unlock()
	docNum := writer.docNum
	writer.segment.stored[docNum] = writer.stored
	writer.segment.lengths[docNum] = writer.lengths
	writer.segment.vectors[docNum] = writer.vectors
	return nil
}

func (writer *MemPerDocWriter) AddField(fieldName string, fieldObj interface{}, value []byte, length int) error {
	if value != nil {
		writer.stored[fieldName] = value
	}
	if length >= 0 {
		writer.lengths[fieldName] = length
	}
	return nil
}
func (writer *MemPerDocWriter) AddColumnValue(fieldName string, columnObj interface{}, value []byte) error {
	writer.columns[fieldName] = value
	return nil
}
func (writer *MemPerDocWriter) AddVectorItems(fieldName string, fieldObj interface{}, items []index.VectorItem) error {
	writer.vectors[fieldName] = items
	return nil
}

func (writer *MemPerDocWriter) Close() error {
	writer.isClosed = true
	return nil
}

type MemPerDocReader struct {
	storage filedb.Storage
	segment *MemSegment
}

func NewMemPerDocReader(storage filedb.Storage, segment *MemSegment) *MemPerDocReader {
	return &MemPerDocReader{
		storage: storage,
		segment: segment,
	}
}

func (reader *MemPerDocReader) Close() error {
	return nil
}

// doc statistics
func (reader *MemPerDocReader) DocCount() int {
	return reader.segment.DocCount()
}
func (reader *MemPerDocReader) DocCountAll() int {
	return reader.segment.DocCountAll()
}
func (reader *MemPerDocReader) HasDeletions() int {
	return reader.segment.HasDeletions()
}
func (reader *MemPerDocReader) IsDeleted(docNum int64) bool {
	return reader.segment.IsDeleted(docNum)
}
func (reader *MemPerDocReader) DeletedDocs() []int64 {
	return reader.segment.DeletedDocs()
}

// columns
func (reader *MemPerDocReader) SupportColumns() bool {
	return false
}
func (reader *MemPerDocReader) HasColumn(columnName string) bool {
	return false
}
func (reader *MemPerDocReader) ListCoumns() []string {
	panic("not implemnt")
}

func (reader *MemPerDocReader) FieldDocs(fieldName string) []int64 {
	panic("not support")
}

// Lengths
func (reader *MemPerDocReader) DocFieldLength(docNum int64, fieldName string, defaultValue int) {
	panic("not implement")
}
func (reader *MemPerDocReader) FieldLength(fieldName string) int {
	panic("not implement")
}
func (reader *MemPerDocReader) MinFieldLength(fieldName string) int {
	panic("not implement")
}
func (reader *MemPerDocReader) MaxFieldLength(fieldName string) int {
	panic("not implement")
}

// Vectors
func (reader *MemPerDocReader) HasVector(docNum int64, fieldName string) bool {
	panic("not implement")
}
func (reader *MemPerDocReader) Vector(docNum int64, fieldName string, format interface{}) index.VectorItem {
	panic("not implement")
}

func (reader *MemPerDocReader) StoredFields(docNum int64) []string {
	panic("not implement")
}
func (reader *MemPerDocReader) AllStoredFields() []string {
	panic("not implement")
}

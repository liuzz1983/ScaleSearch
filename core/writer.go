package core

import (
	"github.com/liuzz1983/scalesearch/core/codec"
	"github.com/liuzz1983/scalesearch/core/filedb"
)

// schema = Schema(date=DATETIME, size=NUMERIC(float), content=TEXT)
// myindex = index.create_in("indexdir", schema)
// w = myindex.writer()
// w.add_document(date=datetime.now(), size=5.5, content=u"Hello")
// w.commit()
type IndexWriter interface {
	StartGroup()
	EndGroup()

	AddField(fieldName string, fieldType FieldType)
	RemoveField(fieldName string)

	//Reader()
	//Searcher()
	//DeleteByTerm
	//DeleteByQuery
	//DeleteDocument
	//AddReader
	//docBoost
	//fieldBoost
	//uniqueFields
	//UpdateDocument
	//Commit
	//Cancel

	// add document fields into index/store
	AddDocument(doc *Document)
}
type SegmentWriter struct {
	WriteLock interface{}
	codec     codec.Codec
	storage   filedb.Storage
	indexName string

	// toc information
	generation int
	schema     Schema
	segments   []codec.Segment
	docNum     int64
	docBase    int64
	docOffsets []int64

	//
	tmpStorage filedb.Storage
}

func (writer *SegmentWriter) DocSegment(docNum int64) int {
	if len(writer.docOffsets) == 0 {
		return 0
	}

}
func (writer *SegmentWriter) SegmentAndDocNum(docNum int64) {

}

package core

import (
	"fmt"

	"github.com/liuzz1983/scalesearch/core/codec"
	"github.com/liuzz1983/scalesearch/core/document"
	"github.com/liuzz1983/scalesearch/core/filedb"
	"github.com/liuzz1983/scalesearch/core/index"
	"github.com/liuzz1983/scalesearch/core/schema"
)

// schema = Schema(date=DATETIME, size=NUMERIC(float), content=TEXT)
// myindex = index.create_in("indexdir", schema)
// w = myindex.writer()
// w.add_document(date=datetime.now(), size=5.5, content=u"Hello")
// w.commit()
type IndexWriter interface {
	StartGroup()
	EndGroup()

	AddField(fieldName string, fieldType index.FieldType)
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
	AddDocument(doc *document.Document)
}
type SegmentWriter struct {
	WriteLock interface{}
	codec     codec.Codec
	storage   filedb.Storage
	indexName string

	// toc information
	generation int
	schema     schema.Schema
	segments   []codec.Segment
	docNum     int64
	docBase    int64
	docOffsets []int64

	//
	tmpStorage   filedb.Storage
	perDocWriter codec.PerDocumentWriter
}

// TODO
func (writer *SegmentWriter) DocSegment(docNum int64) int {
	if len(writer.docOffsets) == 0 {
		return 0
	}
	return -1
}
func (writer *SegmentWriter) SegmentAndDocNum(docNum int64) {
	panic("not impl")
}

func (writer *SegmentWriter) AddDocument(doc *document.Document) error {

	docNum := writer.docNum
	writer.perDocWriter.StartDoc(docNum)

	for fieldName, fieldValue := range doc.Fields {
		analyzer, err := writer.schema.Analyzer(fieldName)
		terms, err := analyzer.Parse(fieldValue)
		if err != nil {
			return nil
		}
		fmt.Println("%v", terms)
	}
	return nil
}

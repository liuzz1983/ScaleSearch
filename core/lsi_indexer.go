package core

import (
	"github.com/balzaczyy/golucene/core/document"
	"github.com/balzaczyy/golucene/core/index"
)

type LsiIndexer struct {
	lsi    *LsiIndex
	writer *index.IndexWriter
}

func NewLsiIndexer(index *LsiIndex) (*LsiIndexer, error) {
	writer, err := index.openWriter()
	if err != nil {
		return nil, err
	}
	indexer := &LsiIndexer{
		index,
		writer,
	}

	return indexer, nil
}
func (indexer *LsiIndexer) add(docId string, doc *Document) {
	d := document.NewDocument()
	d.Add(document.NewFieldFromString("documentId", docId, document.STRING_FIELD_TYPE_NOT_STORED))
	for _, field := range doc.Fields {
		indexField := document.NewTextFieldFromString(field.name, field.value, document.STORE_YES)
		d.Add(indexField)
	}
	indexer.writer.AddDocument(d.Fields())
	indexer.writer.Commit()
}

func (index *LsiIndexer) del(docId string) {

}

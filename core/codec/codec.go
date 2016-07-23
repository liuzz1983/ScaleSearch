package codec

import (
	"github.com/liuzz1983/scalesearch/core/filedb"
)

// document will be seperated into
// field, term -> postlist(docNum)
// docNum -> vector_items,columns

// PerDocumentWriter, store field value,vector_items for document
//  StartDoc
//  AddField
//  AddColumnValue
//  AddVectorItems

// PerDocumentReader

// Codec factory index read and write interface
type Codec interface {
	PerDocumentWriter(storage filedb.Storage, segment string) (PerDocumentWriter, error)
	PerDocumentReader(storage filedb.Storage, segment string) (PerDocumentReader, error)

	FieldWriter(storage filedb.Storage, segment string) (FieldWriter, error)
	PostingsWriter(dbfile string) (PostingWriter, error)
	//PostingsReader(dbfile string, info *TermInfo, format_ string, term string, scorer float32) (PostingReader, error)
	TermsReader(storage filedb.Storage, segment string) (TermsReader, error)
}

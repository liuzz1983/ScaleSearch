package codec

import (
	"github.com/liuzz1983/ScaleSearch/core"
	"github.com/liuzz1983/ScaleSearch/core/filedb"
)

type Codec interface {
	PerDocumentWriter(storage filedb.Storage, segment string) (PerDocumentWriter, error)
	PerDocumentReader(storage filedb.Storage, segment string) (PerDocumentReader, error)

	FieldWriter(storage filedb.Storage, segment string) (interface{}, error)
	PostingsWriter(dbfile string) (PostingWriter, error)
	//PostingsReader(dbfile string, info *TermInfo, format_ string, term string, scorer float32) (PostingReader, error)
	TermsReader(storage filedb.Storage, segment string) (TermsReader, error)
}

type PostingWriter interface {
	StartPostings(format core.Format, terminfo TermInfo) error
	AddPosting(id int64, weight float32, vbytes []byte) error
	FinishPostings() error
	IsWritten() bool
}

type PerDocumentWriter interface {
}
type PerDocumentReader interface {
}

type TermsReader interface {
}

type FieldWriter interface {
}

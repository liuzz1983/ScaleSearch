package codec

import (
	"github.com/liuzz1983/scalesearch/core"
	"github.com/liuzz1983/scalesearch/core/filedb"
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

/*
def start_doc(self, docnum):
        raise NotImplementedError

    @abstractmethod
    def add_field(self, fieldname, fieldobj, value, length):
        raise NotImplementedError

    @abstractmethod
    def add_column_value(self, fieldname, columnobj, value):
        raise NotImplementedError("Codec does not implement writing columns")

    @abstractmethod
    def add_vector_items(self, fieldname, fieldobj, items):
        raise NotImplementedError

    def add_vector_matcher(self, fieldname, fieldobj, vmatcher):
        def readitems():
            while vmatcher.is_active():
                text = vmatcher.id()
                weight = vmatcher.weight()
                valuestring = vmatcher.value()
                yield (text, weight, valuestring)
                vmatcher.next()
        self.add_vector_items(fieldname, fieldobj, readitems())

    def finish_doc(self):
        pass

    def close(self):
*/
type PerDocumentWriter interface {
}
type PerDocumentReader interface {
}

type TermsReader interface {
}

type FieldWriter interface {
}

package codec

import (
	"fmt"
	"strings"
	"github.com/liuzz1983/ScaleSearch/core/filedb"
	"github.com/liuzz1983/ScaleSearch/core"
	"github.com/liuzz1983/ScaleSearch/utils"
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

type Segment interface {
	randomId(size int32) []byte
	Codec() Codec

	IndexName() string 
	SegmentId() string 
	IsCompound() bool 
	MakeFileName(ext string) string 
	ListFiles( storage filedb.Storage ) []string

	CreateFile(storage filedb.Storage, ext string, args map[string]interface{} ) (filedb.File, error)
	OpenFile( storage filedb.Storage, ext string, args map[string]interface{}) (filedb.File, error)

	CreateCompoundFile(storage filedb.Storage) error
	DocCountAll() int64
	DocCount() int64
	SetDocCount(int64)
	HasDeletions() bool
	DeletedCount() int64
	DeletedDocs() int64
	DeleteDocument( doc *core.Document, delete bool ) error 
	IsDeleted( docnum int64) bool 
	ShouldAssemble() bool 

}

const (
	COMPOUND_EXT = ".seg"
	DEFAULT_RANDOM_SIZE = 12
	)

type DefaultSegment struct {
	indexName string
	segId string
	compound bool 
}


func NewDefaultSegment(indexName string) Segment {
	segment := &DefaultSegment{
		indexName: indexName,
	}
	segment.segId = segment.randomId(DEFAULT_RANDOM_SIZE)
	segment.compound = true
	return segment
}

func (seg *DefaultSegment) randomId(size int32 ) string {
	return utils.RandomName( size )
}

func( seg *DefaultSegment) IndexName() string {
	return seg.indexName
}

func(seg *DefaultSegment) Codec() Codec {
	panic(core.ErrNotImplement)
	return nil
}

func(seg *DefaultSegment) SegmentId() string {
	return seg.segId
}

func(seg *DefaultSegment) IsCompound() bool {
	return seg.compound
}

func(seg *DefaultSegment) MakeFileName( ext string) string {
	return fmt.Sprintf("%s%s", seg.SegmentId(), ext)
}

func(seg *DefaultSegment) ListFiles( storage filedb.Storage ) []string {
	prefix := fmt.Sprintf("%s.", seg.SegmentId() ) 
	fullNames, _ := storage.List()
	fileNames := make([]string, 10 )
	for _, name := range( fullNames) {
		if strings.HasPrefix(name, prefix ) {
			fileNames = append( fileNames, name )
		}
	}
	return fileNames
}

func(seg *DefaultSegment) CreateFile( storage filedb.Storage, ext string, args map[string]string)(filedb.File, error){
	fname := seg.MakeFileName(ext)
	return storage.CreateFile(fname,args)
}

func(seg *DefaultSegment) OpenFile( storage filedb.Storage, ext string, args map[string]string)(filedb.File, error){
	fname := seg.MakeFileName(ext)
	return storage.OpenFile(fname, args )
}

func(seg *DefaultSegment) CreateCompoundFile( storage filedb.Storage) {
/*
        segfiles = self.list_files(storage)
        assert not any(name.endswith(self.COMPOUND_EXT) for name in segfiles)
        cfile = self.create_file(storage, self.COMPOUND_EXT)
        CompoundStorage.assemble(cfile, storage, segfiles)
        for name in segfiles:
            storage.delete_file(name)
        self.compound = True
        */
    panic(core.ErrNotImplement)
}

func(seg* DefaultSegment) OpenCompoundFile( storage filedb.Storage) {
 /*       name = self.make_filename(self.COMPOUND_EXT)
    dbfile = storage.open_file(name)
    return CompoundStorage(dbfile, use_mmap=storage.supports_mmap)*/
    panic(core.ErrNotImplement)
}


func(seg* DefaultSegment) DocCountAll() int64 {
	panic(core.ErrNotImplement)
	return 0
}
/*

func(seg* DefaultSegment) 

    def doc_count_all(self):
        """
        Returns the total number of documents, DELETED OR UNDELETED, in this
        segment.
        """

        raise NotImplementedError

    def doc_count(self):
        """
        Returns the number of (undeleted) documents in this segment.
        """

        return self.doc_count_all() - self.deleted_count()

    def set_doc_count(self, doccount):
        raise NotImplementedError

    def has_deletions(self):
        """
        Returns True if any documents in this segment are deleted.
        """

        return self.deleted_count() > 0

    @abstractmethod
    def deleted_count(self):
        """
        Returns the total number of deleted documents in this segment.
        """

        raise NotImplementedError

    @abstractmethod
    def deleted_docs(self):
        raise NotImplementedError

    @abstractmethod
    def delete_document(self, docnum, delete=True):
        """Deletes the given document number. The document is not actually
        removed from the index until it is optimized.

        :param docnum: The document number to delete.
        :param delete: If False, this undeletes a deleted document.
        """

        raise NotImplementedError

    @abstractmethod
    def is_deleted(self, docnum):
        """
        Returns True if the given document number is deleted.
        """

        raise NotImplementedError

    def should_assemble(self):
        return True*/




type PerDocumentWriter interface {
}
type PerDocumentReader interface {
}

type TermsReader interface {
}

type FieldWriter interface {
}

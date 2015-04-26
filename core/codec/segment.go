package codec

import (
	"fmt"
	"github.com/liuzz1983/ScaleSearch/core"
	"github.com/liuzz1983/ScaleSearch/core/filedb"
	"github.com/liuzz1983/ScaleSearch/utils"
	"strings"
)

type Segment interface {
	randomId(size int32) string
	Codec() Codec

	IndexName() string
	SegmentId() string
	IsCompound() bool
	MakeFileName(ext string) string
	ListFiles(storage filedb.Storage) []string

	CreateFile(storage filedb.Storage, ext string, args map[string]string) (filedb.File, error)
	OpenFile(storage filedb.Storage, ext string, args map[string]string) (filedb.File, error)

	CreateCompoundFile(storage filedb.Storage) error
	DocCountAll() int64
	DocCount() int64
	SetDocCount(int64)
	HasDeletions() bool
	DeletedCount() int64
	DeletedDocs() int64
	DeleteDocument(doc *core.Document, delete bool) error
	IsDeleted(docnum int64) bool
	ShouldAssemble() bool
}

const (
	COMPOUND_EXT        = ".seg"
	DEFAULT_RANDOM_SIZE = 12
)

type DefaultSegment struct {
	indexName string
	segId     string
	compound  bool
}

func NewDefaultSegment(indexName string) Segment {
	segment := &DefaultSegment{
		indexName: indexName,
	}
	segment.segId = segment.randomId(DEFAULT_RANDOM_SIZE)
	segment.compound = true
	return segment
}

func (seg *DefaultSegment) randomId(size int32) string {
	return utils.RandomName(size)
}

func (seg *DefaultSegment) IndexName() string {
	return seg.indexName
}

func (seg *DefaultSegment) Codec() Codec {
	panic(core.ErrNotImplement)
	return nil
}

func (seg *DefaultSegment) SegmentId() string {
	return seg.segId
}

func (seg *DefaultSegment) IsCompound() bool {
	return seg.compound
}

func (seg *DefaultSegment) MakeFileName(ext string) string {
	return fmt.Sprintf("%s%s", seg.SegmentId(), ext)
}

func (seg *DefaultSegment) ListFiles(storage filedb.Storage) []string {
	prefix := fmt.Sprintf("%s.", seg.SegmentId())
	fullNames, _ := storage.List()
	fileNames := make([]string, 10)
	for _, name := range fullNames {
		if strings.HasPrefix(name, prefix) {
			fileNames = append(fileNames, name)
		}
	}
	return fileNames
}

func (seg *DefaultSegment) CreateFile(storage filedb.Storage, ext string, args map[string]string) (filedb.File, error) {
	fname := seg.MakeFileName(ext)
	return storage.CreateFile(fname, args)
}

func (seg *DefaultSegment) OpenFile(storage filedb.Storage, ext string, args map[string]string) (filedb.File, error) {
	fname := seg.MakeFileName(ext)
	return storage.OpenFile(fname, args)
}

func (seg *DefaultSegment) CreateCompoundFile(storage filedb.Storage) error {
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

func (seg *DefaultSegment) OpenCompoundFile(storage filedb.Storage) error {
	/*       name = self.make_filename(self.COMPOUND_EXT)
	dbfile = storage.open_file(name)
	return CompoundStorage(dbfile, use_mmap=storage.supports_mmap)*/
	panic(core.ErrNotImplement)
}

func (seg *DefaultSegment) DocCountAll() int64 {
	panic(core.ErrNotImplement)
	return 0
}

func (seg *DefaultSegment) DocCount() int64 {
	panic(core.ErrNotImplement)
}
func (seg *DefaultSegment) SetDocCount(count int64) {
	panic(core.ErrNotImplement)
}
func (seg *DefaultSegment) HasDeletions() bool {
	panic(core.ErrNotImplement)
}
func (seg *DefaultSegment) DeletedCount() int64 {
	panic(core.ErrNotImplement)
}
func (seg *DefaultSegment) DeletedDocs() int64 {
	panic(core.ErrNotImplement)
}

func (seg *DefaultSegment) DeleteDocument(doc *core.Document, del bool) error {
	panic(core.ErrNotImplement)
}
func (seg *DefaultSegment) IsDeleted(docnum int64) bool {
	panic(core.ErrNotImplement)
}
func (seg *DefaultSegment) ShouldAssemble() bool {
	return true
}

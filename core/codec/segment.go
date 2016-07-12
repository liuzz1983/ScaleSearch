package codec

import (
	"fmt"
	"github.com/liuzz1983/scalesearch/core"
	"github.com/liuzz1983/scalesearch/core/errors"
	"github.com/liuzz1983/scalesearch/core/filedb"
	"github.com/liuzz1983/scalesearch/utils"
	"github.com/liuzz1983/scalesearch/utils/fs"
	"strings"
	_ "sync"
)

type Segment interface {
	randomId(size int32) string
	Codec() Codec

	IndexName() string
	SegmentId() string
	IsCompound() bool
	MakeFileName(ext string) string
	ListFiles(storage filedb.Storage) []string

	CreateFile(storage filedb.Storage, ext string, args map[string]string) (fs.File, error)
	OpenFile(storage filedb.Storage, ext string, args map[string]string) (fs.File, error)

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

type BaseSegment struct {
	indexName string
	segId     string
	compound  bool
}

func NewBaseSegment(indexName string) Segment {
	segment := &BaseSegment{
		indexName: indexName,
	}
	segment.segId = segment.randomId(DEFAULT_RANDOM_SIZE)
	segment.compound = true
	return segment
}

func (seg *BaseSegment) randomId(size int32) string {
	return utils.RandomName(size)
}

func (seg *BaseSegment) IndexName() string {
	return seg.indexName
}

func (seg *BaseSegment) Codec() Codec {
	panic(errors.ErrNotImplement)
	return nil
}

func (seg *BaseSegment) SegmentId() string {
	return seg.segId
}

func (seg *BaseSegment) IsCompound() bool {
	return seg.compound
}

func (seg *BaseSegment) MakeFileName(ext string) string {
	return fmt.Sprintf("%s%s", seg.SegmentId(), ext)
}

func (seg *BaseSegment) ListFiles(storage filedb.Storage) []string {
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

func (seg *BaseSegment) CreateFile(storage filedb.Storage, ext string,
	args map[string]string) (fs.File, error) {

	fname := seg.MakeFileName(ext)
	return storage.CreateFile(fname, args)
}

func (seg *BaseSegment) OpenFile(storage filedb.Storage, ext string,
	args map[string]string) (fs.File, error) {

	fname := seg.MakeFileName(ext)
	return storage.OpenFile(fname, args)
}

func (seg *BaseSegment) CreateCompoundFile(storage filedb.Storage) error {
	/*
	   segfiles = self.list_files(storage)
	   assert not any(name.endswith(self.COMPOUND_EXT) for name in segfiles)
	   cfile = self.create_file(storage, self.COMPOUND_EXT)
	   CompoundStorage.assemble(cfile, storage, segfiles)
	   for name in segfiles:
	       storage.delete_file(name)
	   self.compound = True
	*/
	panic(errors.ErrNotImplement)
}

func (seg *BaseSegment) OpenCompoundFile(storage filedb.Storage) error {
	/*       name = self.make_filename(self.COMPOUND_EXT)
	dbfile = storage.open_file(name)
	return CompoundStorage(dbfile, use_mmap=storage.supports_mmap)*/
	panic(errors.ErrNotImplement)
}

func (seg *BaseSegment) DocCountAll() int64 {
	panic(errors.ErrNotImplement)
	return 0
}

func (seg *BaseSegment) DocCount() int64 {
	panic(errors.ErrNotImplement)
}

func (seg *BaseSegment) SetDocCount(count int64) {
	panic(errors.ErrNotImplement)
}

func (seg *BaseSegment) HasDeletions() bool {
	panic(errors.ErrNotImplement)
}

func (seg *BaseSegment) DeletedCount() int64 {
	panic(errors.ErrNotImplement)
}

func (seg *BaseSegment) DeletedDocs() int64 {
	panic(errors.ErrNotImplement)
}

func (seg *BaseSegment) DeleteDocument(doc *core.Document, del bool) error {
	panic(errors.ErrNotImplement)
}
func (seg *BaseSegment) IsDeleted(docnum int64) bool {
	panic(errors.ErrNotImplement)
}
func (seg *BaseSegment) ShouldAssemble() bool {
	return true
}

type InvIndex map[string]interface{}

type FieldTerm struct {
	Field string
	Term  string
}

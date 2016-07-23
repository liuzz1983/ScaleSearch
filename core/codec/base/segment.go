package base

import (
	"fmt"
	"strings"

	"github.com/liuzz1983/scalesearch/core/codec"
	"github.com/liuzz1983/scalesearch/core/errors"
	"github.com/liuzz1983/scalesearch/core/filedb"
	"github.com/liuzz1983/scalesearch/utils"
	"github.com/liuzz1983/scalesearch/utils/fs"
)

type BaseSegment struct {
	InxName  string
	SegId    string
	Compound bool
}

func NewBaseSegment(indexName string) codec.Segment {
	segment := &BaseSegment{
		InxName: indexName,
	}
	segment.SegId = segment.RandomId(codec.DEFAULT_RANDOM_SIZE)
	segment.Compound = true
	return segment
}

func (seg *BaseSegment) RandomId(size int32) string {
	return utils.RandomName(size)
}

func (seg *BaseSegment) IndexName() string {
	return seg.InxName
}

func (seg *BaseSegment) Codec() codec.Codec {
	panic(errors.ErrNotImplement)
	return nil
}

func (seg *BaseSegment) SegmentId() string {
	return fmt.Sprintf("%s,%v", seg.InxName, seg.SegId)
}

func (seg *BaseSegment) IsCompound() bool {
	return seg.Compound
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
	return seg.DocCountAll() - seg.DeletedCount()
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

func (seg *BaseSegment) DeleteDocument(docNum int64, del bool) error {
	panic(errors.ErrNotImplement)
}
func (seg *BaseSegment) IsDeleted(docNum int64) bool {
	panic(errors.ErrNotImplement)
}

func (seg *BaseSegment) ShouldAssemble() bool {
	return true
}

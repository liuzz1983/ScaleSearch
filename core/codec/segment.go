package codec

import (
	"github.com/liuzz1983/scalesearch/core/filedb"
	"github.com/liuzz1983/scalesearch/utils/fs"
)

type Segment interface {
	RandomId(size int32) string
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
	DeleteDocument(docNum int64, delete bool) error
	IsDeleted(docnum int64) bool
	ShouldAssemble() bool
}

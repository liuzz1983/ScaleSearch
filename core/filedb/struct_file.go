package filedb

import (
	"github.com/liuzz1983/scalesearch/utils/fs"
)

type StructFile struct {
	file     fs.File
	name     string
	isClosed bool
}

type OnClose func(*StructFile)

// provide addition method for file operation
func NewStructFile(fileObj fs.File, name string, onClose OnClose) *StructFile {
	return &StructFile{
		file:     fileObj,
		name:     name,
		isClosed: false,
	}
}

func (fs *StructFile) RawFile() fs.File {
	return fs.file
}

func (fs *StructFile) Read(b []byte) (n int, err error) {
	return fs.file.Read(b)
}

func (fs *StructFile) Write(b []byte) (n int, err error) {
	return fs.file.Write(b)
}

func (fs *StructFile) WriteAt(b []byte, offset int64) (n int, err error) {
	return fs.file.WriteAt(b, offset)
}

func (fs *StructFile) ReadAt(b []byte, offset int64) (n int, err error) {
	return fs.file.ReadAt(b, offset)
}

func (fs *StructFile) Tell() (n int64, err error) {
	stat, err := fs.file.Stat()
	if err == nil {
		return 0, err
	}
	return stat.Size(), nil
}

func (fs *StructFile) Flush() error {
	return fs.file.Sync()
}

func (fs *StructFile) Close() error {
	return fs.file.Close()
}

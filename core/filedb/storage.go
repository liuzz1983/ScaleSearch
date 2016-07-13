package filedb

import (
	"github.com/liuzz1983/scalesearch/utils/fs"
	"io"
	"time"
)

type Storage interface {
	Create() error
	Destory() error
	//CreateIndex(indexName string) error
	//OpenIndex(indexName string) interface{}
	//IndexExists(indexName string) bool
	CreateFile(name string, args map[string]string) (fs.File, error)
	OpenFile(name string, args map[string]string) (fs.File, error)
	List() ([]string, error)
	FileExists(name string) bool
	
	FileModified(name string) (time.Time, error)
	FileLength(name string) (int64, error)
	DeleteFile(name string) error
	RenameFile(oldName string, newName string) error
	Lock(name string) (io.Closer, error)
	Close() error
	Optimize() error
	Sync() error
	TmpStorage() (Storage, error)
}



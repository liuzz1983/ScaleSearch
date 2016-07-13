package filedb

import (
	"sync"
	"github.com/liuzz1983/scalesearch/core/errors"
	"time"
)

// storage in memory, for testing
type RamStorage struct {
	files  map[string]string
	locks  map[string]sync.Mutex
	folder string
}

func NewMemStorage() *RamStorage {
	return &RamStorage{
		files:  make(map[string]string),
		locks:  make(map[string]sync.Mutex),
		folder: "",
	}
}

func (rs *RamStorage) Destory() error {
	rs.files = nil
	rs.locks = nil
	return nil
}

func (rs *RamStorage) List() []string {
	files := make([]string, len(rs.files))
	for key, _ := range rs.files {
		files = append(files, key)
	}
	return files
}

func (rs *RamStorage) Clean() error {
	rs.files = make(map[string]string)
	return nil
}

func (rs *RamStorage) TotalSize() (int64, error) {
	var total int64 = 0
	for key, _ := range rs.files {
		size, _ := rs.FileLength(key)
		total += size
	}
	return total, nil
}

func (rs *RamStorage) FileExists(name string) bool {
	_, ok := rs.files[name]
	return ok
}
func (rs *RamStorage) FileLength(name string) (int64, error) {
	if !rs.FileExists(name) {
		return -1, errors.ErrFileNotExists
	}
	memLen := len(rs.files[name])
	return int64(memLen), nil
}

func (rs *RamStorage) FileModified(name string) (*time.Time, error) {
	return nil, errors.New("not implement this method")
}

func (rs *RamStorage) DeleteFile(name string) error {
	if !rs.FileExists(name) {
		return errors.ErrFileNotExists
	}
	delete(rs.files, name)
	return nil
}

func (rs *RamStorage) RenameFile(oldName string, newName string) error {
	if !rs.FileExists(oldName) {
		return errors.ErrFileNotExists
	}
	if rs.FileExists(newName) {
		return errors.ErrFileExists
	}
	content := rs.files[oldName]
	delete(rs.files, oldName)
	rs.files[newName] = content
	return nil
}

func (rs *RamStorage) CreateFile(name string) error {
	return nil
}

func (rs *RamStorage) OpenFile(name string) {

}

func (rs *RamStorage) Lock(name string) {

}

func (rs *RamStorage) TmpStorage(name string) {

}
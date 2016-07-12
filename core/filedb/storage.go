package filedb

import (
	"github.com/liuzz1983/scalesearch/core/errors"
	"github.com/liuzz1983/scalesearch/utils/fs"
	"io"
	"os"
	"path/filepath"
	"sync"
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
	//FileModified(name string) (*time.Time, error)
	FileLength(name string) (int64, error)
	DeleteFile(name string) error
	RenameFile(oldName string, newName string) error
	Lock(name string) (io.Closer, error)
	Close() error
	Optimize() error
	Sync() error
	TmpStorage() (Storage, error)
}

type FileStorage struct {
	readOnly    bool
	supportMMap bool
	folder      string
	system      fs.FileSystem
}

func NewFileStorage(name string, fileSystem fs.FileSystem) (Storage, error) {
	file := FileStorage{
		readOnly:    false,
		supportMMap: false,
		folder:      name,
		system:      fileSystem,
	}
	return &file, nil
}

func (store *FileStorage) Create() error {
	dirname, err := filepath.Abs(store.folder)
	err = store.system.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		return err
	}
	dir, err := store.system.Stat(dirname)
	if err != nil {
		return err
	}
	if !dir.IsDir() {
		return err
	}
	return nil
}

func (store *FileStorage) filePath(name string) string {
	return filepath.Join(store.folder, name)
}

func (store *FileStorage) fileStat(name string) (os.FileInfo, error) {
	path := store.filePath(name)
	info, err := store.system.Stat(path)
	return info, err
}

//
func (store *FileStorage) Destory() error {
	if store.readOnly {
		return errors.ErrIsReadOnly
	}
	names, err := store.List()
	if err != nil {
		return err
	}
	for _, name := range names {
		fileName := filepath.Join(store.folder, name)
		err := store.system.Remove(fileName)
		if err != nil {
			return err
		}
	}

	err = store.system.Remove(store.folder)
	return err
}

// need consider
func (store *FileStorage) List() ([]string, error) {
	return store.system.List(store.folder)
}

func (store *FileStorage) FileExists(name string) bool {
	fileName := store.filePath(name)
	_, err := store.system.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (rs *FileStorage) CreateFile(name string, args map[string]string) (fs.File, error) {
	panic(errors.ErrNotImplement)
}

func (rs *FileStorage) OpenFile(name string, args map[string]string) (fs.File, error) {
	panic(errors.ErrNotImplement)
}

/*func (store *FileStorage) FileModified(name string) (*time.Time, error) {
	if !store.FileExists(name) {
		return nil, ErrFileNotExists
	}
	return nil, nil
}*/

func (store *FileStorage) FileLength(name string) (int64, error) {
	stat, err := store.fileStat(name)
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}

func (store *FileStorage) DeleteFile(name string) error {
	fileName := store.filePath(name)
	err := store.system.Remove(fileName)
	return err
}

func (store *FileStorage) RenameFile(oldName string, newName string) error {
	oldPath := store.filePath(oldName)
	newPath := store.filePath(newName)
	return store.system.Rename(oldPath, newPath)
}

func (store *FileStorage) Lock(name string) (fl io.Closer, err error) {
	path := store.filePath(name)
	return store.system.Lock(path)
}

func (store *FileStorage) Close() error {
	return nil
}

func (store *FileStorage) Optimize() error {
	return nil
}

func (store *FileStorage) TmpStorage() (Storage, error) {
	tmpDir := os.TempDir()
	return NewFileStorage(tmpDir, store.system)
}

func (store *FileStorage) Sync() error {
	return fs.SyncDir(store.folder)
}

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

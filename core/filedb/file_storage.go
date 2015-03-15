package filedb

import (
	"io"
	"os"
	"path/filepath"
	_ "time"
	"github.com/liuzz1983/ScaleSearch/core"
)

type Storage interface {
	Create() error
	Destory() error
	//CreateIndex(indexName string) error
	//OpenIndex(indexName string) interface{}
	//IndexExists(indexName string) bool
	CreateFile(name string,args map[string]string) (File, error)
	OpenFile(name string,args map[string]string) (File, error)
	List() ([]string,error)
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
	system FileSystem
}

func NewFileStorage(name string,fileSystem FileSystem ) (Storage, error) {
	file := FileStorage{
		readOnly:    false,
		supportMMap: false,
		folder:      name,
		system: fileSystem,
	}
	return &file, nil
}

func (fs *FileStorage) Create() error {
	dirname, err := filepath.Abs(fs.folder)
	err = fs.system.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		return err
	}
	dir, err := fs.system.Stat(dirname)
	if err != nil {
		return err
	}
	if !dir.IsDir() {
		return err
	}
	return nil
}

func (fs *FileStorage) filePath(name string) string {
	return filepath.Join(fs.folder, name)
}

func (fs *FileStorage) fileStat(name string) (os.FileInfo, error) {
	path := fs.filePath(name)
	info, err := fs.system.Stat(path)
	return info, err
}

//
func (fs *FileStorage) Destory() error {
	if fs.readOnly {
		return core.ErrIsReadOnly
	}
	names, err := fs.List()
	if err != nil {
		return err
	}
	for _, name := range names {
		fileName := filepath.Join(fs.folder, name)
		err := fs.system.Remove(fileName)
		if err != nil {
			return err
		}
	}

	err = fs.system.Remove(fs.folder)
	return err
}

// need consider
func (fs *FileStorage) List() ([]string, error) {
	return fs.system.List(fs.folder)
}

func (fs *FileStorage) FileExists(name string) bool {
	fileName := fs.filePath(name)
	_, err := fs.system.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (rs *FileStorage) CreateFile(name string)(File,error) {
	return nil,core.ErrNotImplement
}

func (rs *FileStorage) OpenFile(name string)(File,error) {
	return nil, core.ErrNotImplement
}

/*func (fs *FileStorage) FileModified(name string) (*time.Time, error) {
	if !fs.FileExists(name) {
		return nil, ErrFileNotExists
	}
	return nil, nil
}*/

func (fs *FileStorage) FileLength(name string) (int64, error) {
	stat, err := fs.fileStat(name)
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}

func (fs *FileStorage) DeleteFile(name string) error {
	fileName := fs.filePath(name)
	err := fs.system.Remove(fileName)
	return err
}

func (fs *FileStorage) RenameFile(oldName string, newName string) error {
	oldPath := fs.filePath(oldName)
	newPath := fs.filePath(newName)
	return fs.system.Rename(oldPath, newPath)
}

func (fs *FileStorage) Lock(name string) (fl io.Closer, err error) {
	path := fs.filePath(name)
	return fs.system.Lock(path)
}

func (fs *FileStorage) Close() error {
	return nil
}

func (fs *FileStorage) Optimize() error {
	return nil
}

func (fs *FileStorage) TmpStorage() (Storage, error) {
	tmpDir := os.TempDir()
	return NewFileStorage(tmpDir,fs.system)
}

func (fs *FileStorage) Sync() error {
	return syncDir(fs.folder)
}

/*
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
		return -1, ErrFileNotExists
	}
	memLen := len(rs.files[name])
	return int64(memLen), nil
}

func (rs *RamStorage) FileModified(name string) (*time.Time, error) {
	return nil, errors.New("not implement this method")
}

func (rs *RamStorage) DeleteFile(name string) error {
	if !rs.FileExists(name) {
		return ErrFileNotExists
	}
	delete(rs.files, name)
	return nil
}

func (rs *RamStorage) RenameFile(oldName string, newName string) error {
	if !rs.FileExists(oldName) {
		return ErrFileNotExists
	}
	if rs.FileExists(newName) {
		return ErrFileExists
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
*/
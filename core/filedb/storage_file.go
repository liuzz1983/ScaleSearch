package filedb

import (
	"github.com/liuzz1983/scalesearch/core/errors"
	"github.com/liuzz1983/scalesearch/utils/fs"
	"io"
	"os"
	"path/filepath"
	"time"
)

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

// destroy the whole directory of storage,
// but if the hidden file exists, this will error, we need use removeall ?
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

func (store *FileStorage) CreateFile(name string, args map[string]string) (fs.File, error) {
	return store.system.Create(name)
}

func (store *FileStorage) OpenFile(name string, args map[string]string) (fs.File, error) {
	return store.system.Open(name)
}

func (store *FileStorage) FileModified(name string) (time.Time, error) {
	if !store.FileExists(name) {
		return time.Time{}, errors.ErrFileNotExists
	}

	stats, error := store.fileStat(name)
	if error != nil {
		return time.Time{}, error
	}

	return stats.ModTime(), nil
}

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

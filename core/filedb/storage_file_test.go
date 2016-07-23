package filedb

import (
	"github.com/liuzz1983/scalesearch/utils/fs"
	"os"
	"path/filepath"
	"testing"
)

func testFileStorage(t *testing.T, fileSystem fs.FileSystem) {

	maindir := filepath.Join(os.TempDir(), "tmp-test")
	storage, error := NewFileStorage(maindir, fileSystem)
	if error != nil {
		t.Errorf("error in create storage %v", maindir)
	}

	error = storage.Create()
	if error != nil {
		t.Errorf("error create main storage %v", maindir)
	}

	exists := fileSystem.IsExist(maindir)
	if !exists {
		stat, error := fileSystem.Stat(maindir)
		t.Errorf("error main storage is not exists %v, %v, %v %v", maindir, error, os.IsExist(error), stat)
	}

	error = storage.Destory()
	if error != nil {
		t.Errorf(" error in destory storage %v", error)
	}

	if fileSystem.IsExist(maindir) {
		t.Errorf("error main storage should be destoried %v", maindir)
	}
}

func TestCeateMemFileStorage(t *testing.T) {

	testFileStorage(t, fs.DefaultFileSystem)
	testFileStorage(t, fs.NewMemFileSystem())

}

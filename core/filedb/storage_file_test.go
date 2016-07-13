package filedb

import  (
	"os"
	"testing"
	"github.com/liuzz1983/scalesearch/utils/fs"
)


func TestCreateFileStorage(t *testing.T) { 
	maindir := os.TempDir()
	storage, error := NewFileStorage(maindir, fs.DefaultFileSystem)
	if error != nil {
		t.Errorf("error in create storage %v", maindir)
	}

	error = storage.Create()
	if error != nil {
		t.Errorf("error create main storage %v", maindir)
	}

	exists := fs.DefaultFileSystem.IsExist(maindir)
	if ! exists {
		stat, error := fs.DefaultFileSystem.Stat(maindir)
		t.Errorf("error main storage is not exists %v, %v, %v %v", maindir, error, os.IsExist(error),stat)
	}

	error = storage.Destory()
	if error != nil {
		t.Errorf(" error in destory storage %v", error)
	}

	if fs.DefaultFileSystem.IsExist(maindir) {
		t.Errorf("error main storage should be destoried %v", maindir)
	}
}
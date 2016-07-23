package fs

import (
	"bytes"
	"testing"
)

func TestFileSystem(t *testing.T) {

	fileSystem := NewMemFileSystem()

	fileName := "test_file"
	file, err := fileSystem.Create(fileName)
	if err != nil {
		t.Errorf("error create file  %v", fileName)
	}

	if file == nil {
		t.Errorf("error in return nil file %v", fileName)
	}

	content := []byte("zhongguo")
	l, err := file.Write(content)

	if err != nil {
		t.Errorf("error in write file %v", err)
	}

	if l != len(content) {
		t.Errorf("size is wrong in write content %v %v ", l, len(content))
	}

	err = file.Close()
	if err != nil {
		t.Errorf("wrong in close file")
	}

	// reopen file
	file, err = fileSystem.Open(fileName)
	if err != nil {
		t.Errorf("error open file  %v", fileName)
	}

	if file == nil {
		t.Errorf("error in return nil file %v", fileName)
	}

	value := make([]byte, len(content))

	l, err = file.Read(value)
	if err != nil {
		t.Errorf("error in read file %v", err)
	}

	if bytes.Compare(content, value) != 0 {
		t.Errorf("error in compare readcontent")
	}

}

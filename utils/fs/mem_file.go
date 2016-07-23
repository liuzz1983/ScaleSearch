// Copyright 2012 The LevelDB-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package memfs provides a memory-backed db.MemFileSystem implementation.
//
// It can be useful for tests, and also for LevelDB instances that should not
// ever touch persistent storage, such as a web browser's private browsing mode.
package fs

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

const sep = string(os.PathSeparator)

type nopCloser struct{}

func (nopCloser) Close() error {
	return nil
}

// node holds a file's data or a directory's children, and implements os.FileInfo.
type node struct {
	name     string
	data     []byte
	modTime  time.Time
	children map[string]*node
	isDir    bool
}

func (f *node) IsDir() bool {
	return f.isDir
}

func (f *node) ModTime() time.Time {
	return f.modTime
}

func (f *node) Mode() os.FileMode {
	if f.isDir {
		return os.ModeDir | 0755
	}
	return 0755
}

func (f *node) Name() string {
	return f.name
}

func (f *node) Size() int64 {
	return int64(len(f.data))
}

func (f *node) Sys() interface{} {
	return nil
}

func (f *node) dump(w *bytes.Buffer, level int) {
	if f.isDir {
		w.WriteString("          ")
	} else {
		fmt.Fprintf(w, "%8d  ", len(f.data))
	}
	for i := 0; i < level; i++ {
		w.WriteString("  ")
	}
	w.WriteString(f.name)
	if !f.isDir {
		w.WriteByte('\n')
		return
	}
	w.WriteByte(os.PathSeparator)
	w.WriteByte('\n')
	names := make([]string, 0, len(f.children))
	for name := range f.children {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		f.children[name].dump(w, level+1)
	}
}

// file is a reader or writer of a node's data, and implements db.File.
type MemFile struct {
	n           *node
	rpos        int
	read, write bool
}

func (f *MemFile) Close() error {
	return nil
}

func (f *MemFile) Read(p []byte) (int, error) {
	if !f.read {
		return 0, errors.New("leveldb/memfs: file was not opened for reading")
	}
	if f.n.isDir {
		return 0, errors.New("leveldb/memfs: cannot read a directory")
	}
	if f.rpos >= len(f.n.data) {
		return 0, io.EOF
	}
	n := copy(p, f.n.data[f.rpos:])
	f.rpos += n
	return n, nil
}

func (f *MemFile) ReadAt(p []byte, off int64) (int, error) {
	if !f.read {
		return 0, errors.New("leveldb/memfs: file was not opened for reading")
	}
	if f.n.isDir {
		return 0, errors.New("leveldb/memfs: cannot read a directory")
	}
	if off >= int64(len(f.n.data)) {
		return 0, io.EOF
	}
	return copy(p, f.n.data[off:]), nil
}

func (f *MemFile) Write(p []byte) (int, error) {
	if !f.write {
		return 0, errors.New("leveldb/memfs: file was not created for writing")
	}
	if f.n.isDir {
		return 0, errors.New("leveldb/memfs: cannot write a directory")
	}
	f.n.modTime = time.Now()
	f.n.data = append(f.n.data, p...)
	return len(p), nil
}

func (f *MemFile) WriteAt(p []byte, off int64) (int, error) {
	if !f.write {
		return 0, errors.New("leveldb/memfs: file was not created for writing")
	}
	if f.n.isDir {
		return 0, errors.New("leveldb/memfs: cannot write a directory")
	}
	if off >= int64(len(f.n.data)) {
		return 0, io.EOF
	}
	return copy(f.n.data[off:], p), nil
}

func (f *MemFile) Stat() (os.FileInfo, error) {
	return f.n, nil
}

func (f *MemFile) Sync() error {
	return nil
}

func (f *MemFile) Seek(offset int64, whence int) (int64, error) {
	return 0, errors.New("not implement this method")
}

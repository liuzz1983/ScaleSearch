// Copyright 2012 The LevelDB-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"io"
	"os"
)

// File is a readable, writable sequence of bytes.
//
// Typically, it will be an *os.File, but test code may choose to substitute
// memory-backed implementations.
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Writer
	io.WriterAt
	io.Seeker
	Stat() (os.FileInfo, error)
	Sync() error
}

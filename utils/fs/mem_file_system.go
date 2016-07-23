package fs

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"sync"
)

// New returns a new memory-backed 112.MemFileSystem implementation.
func NewMemFileSystem() *MemFileSystem {
	return &MemFileSystem{
		root: &node{
			children: make(map[string]*node),
			isDir:    true,
		},
	}
}

// MemFileSystem implements db.MemFileSystem.
type MemFileSystem struct {
	mu   sync.Mutex
	root *node
}

func (y *MemFileSystem) String() string {
	y.mu.Lock()
	defer y.mu.Unlock()

	s := new(bytes.Buffer)
	y.root.dump(s, 0)
	return s.String()
}

// walk walks the directory tree for the fullname, calling f at each step. If
// f returns an error, the walk will be aborted and return that same error.
//
// Each walk is atomic: y's mutex is held for the entire operation, including
// all calls to f.
//
// dir is the directory at that step, frag is the name fragment, and final is
// whether it is the final step. For example, walking "/foo/bar/x" will result
// in 3 calls to f:
//   - "/", "foo", false
//   - "/foo/", "bar", false
//   - "/foo/bar/", "x", true
// Similarly, walking "/y/z/", with a trailing slash, will result in 3 calls to f:
//   - "/", "y", false
//   - "/y/", "z", false
//   - "/y/z/", "", true
func (y *MemFileSystem) walk(fullname string, f func(dir *node, frag string, final bool) error) error {
	y.mu.Lock()
	defer y.mu.Unlock()

	// For memfs, the current working directory is the same as the root directory,
	// so we strip off any leading "/"s to make fullname a relative path, and
	// the walk starts at y.root.
	for len(fullname) > 0 && fullname[0] == os.PathSeparator {
		fullname = fullname[1:]
	}
	dir := y.root

	for {
		frag, remaining := fullname, ""
		i := strings.IndexRune(fullname, os.PathSeparator)
		final := i < 0
		if !final {
			frag, remaining = fullname[:i], fullname[i+1:]
			for len(remaining) > 0 && remaining[0] == os.PathSeparator {
				remaining = remaining[1:]
			}
		}
		if err := f(dir, frag, final); err != nil {
			return err
		}
		if final {
			break
		}
		child := dir.children[frag]
		if child == nil {
			return errors.New("leveldb/memfs: no such directory")
		}
		if !child.isDir {
			return errors.New("leveldb/memfs: not a directory")
		}
		dir, fullname = child, remaining
	}
	return nil
}

func (y *MemFileSystem) Create(fullname string) (File, error) {
	var ret *MemFile
	err := y.walk(fullname, func(dir *node, frag string, final bool) error {
		if final {
			if frag == "" {
				return errors.New("leveldb/memfs: empty file name")
			}
			n := &node{name: frag}
			dir.children[frag] = n
			ret = &MemFile{
				n:     n,
				write: true,
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (y *MemFileSystem) Open(fullname string) (File, error) {
	var ret *MemFile
	err := y.walk(fullname, func(dir *node, frag string, final bool) error {
		if final {
			if frag == "" {
				return errors.New("leveldb/memfs: empty file name")
			}
			if n := dir.children[frag]; n != nil {
				ret = &MemFile{
					n:    n,
					read: true,
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: fullname,
			Err:  os.ErrNotExist,
		}
	}
	return ret, nil
}

func (y *MemFileSystem) Remove(fullname string) error {
	return y.walk(fullname, func(dir *node, frag string, final bool) error {
		if final {
			if frag == "" {
				return errors.New("leveldb/memfs: empty file name")
			}
			_, ok := dir.children[frag]
			if !ok {
				return errors.New("leveldb/memfs: no such file or directory")
			}
			delete(dir.children, frag)
		}
		return nil
	})
}

func (y *MemFileSystem) Rename(oldname, newname string) error {
	var n *node
	err := y.walk(oldname, func(dir *node, frag string, final bool) error {
		if final {
			if frag == "" {
				return errors.New("leveldb/memfs: empty file name")
			}
			n = dir.children[frag]
			delete(dir.children, frag)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if n == nil {
		return errors.New("leveldb/memfs: no such file or directory")
	}
	return y.walk(newname, func(dir *node, frag string, final bool) error {
		if final {
			if frag == "" {
				return errors.New("leveldb/memfs: empty file name")
			}
			dir.children[frag] = n
		}
		return nil
	})
}

func (y *MemFileSystem) MkdirAll(dirname string, perm os.FileMode) error {
	return y.walk(dirname, func(dir *node, frag string, final bool) error {
		if frag == "" {
			if final {
				return nil
			}
			return errors.New("leveldb/memfs: empty file name")
		}
		child := dir.children[frag]
		if child == nil {
			dir.children[frag] = &node{
				name:     frag,
				children: make(map[string]*node),
				isDir:    true,
			}
			return nil
		}
		if !child.isDir {
			return errors.New("leveldb/memfs: not a directory")
		}
		return nil
	})
}

func (y *MemFileSystem) Lock(fullname string) (io.Closer, error) {
	// MemFileSystem.Lock excludes other processes, but other processes cannot
	// see this process' memory, so Lock is a no-op.
	return nopCloser{}, nil
}

func (y *MemFileSystem) List(dirname string) ([]string, error) {
	if !strings.HasSuffix(dirname, sep) {
		dirname += sep
	}
	var ret []string
	err := y.walk(dirname, func(dir *node, frag string, final bool) error {
		if final {
			if frag != "" {
				panic("unreachable")
			}
			ret = make([]string, 0, len(dir.children))
			for s := range dir.children {
				ret = append(ret, s)
			}
		}
		return nil
	})
	return ret, err
}

func (y *MemFileSystem) Stat(name string) (os.FileInfo, error) {
	f, err := y.Open(name)
	if err != nil {
		if pe, ok := err.(*os.PathError); ok {
			pe.Op = "stat"
		}
		return nil, err
	}

	defer f.Close()

	return f.Stat()
}

func (y *MemFileSystem) IsExist(name string) bool {
	_, error := y.Stat(name)
	if error == nil {
		return true
	}
	return os.IsExist(error)
}

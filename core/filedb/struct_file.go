package filedb


type StructFile struct {
	file     File
	name     string
	isClosed bool
}

func NewStructFile(fileObj File, name string) *StructFile {
	return &StructFile{
		file:     fileObj,
		name:     name,
		isClosed: false,
	}
}

func (fs *StructFile) RawFile() File {
	return fs.file
}

func (fs *StructFile) Read(b []byte) (n int, err error) {
	return fs.file.Read(b)
}

func (fs *StructFile) Write(b []byte) (n int, err error) {
	return fs.file.Write(b)
}

func (fs *StructFile) WriteAt(b []byte, offset int64 ) (n int, err error) {
	return fs.file.WriteAt(b,offset)
}

func (fs *StructFile) ReadAt( b []byte,offset int64) (n int, err error ) {
	return fs.file.ReadAt( b,offset)
}

func (fs *StructFile) Tell() (n int64, err error) {
	stat, err := fs.file.Stat()
	if err == nil {
		return 0, err
	}
	return stat.Size(), nil
}

func (fs *StructFile) Flush() error {
	return fs.file.Sync()
}

func (fs *StructFile) Close() error {
	return fs.file.Close()
}

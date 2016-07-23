package whoosh

import (
	"os"

	"github.com/liuzz1983/scalesearch/core"
	"github.com/liuzz1983/scalesearch/core/errors"
	"github.com/liuzz1983/scalesearch/utils/fs"
)

type W3PostingsWriter struct {
	PostFile    fs.File
	BlockLimit  int
	Compression int
	InlineLimit int

	BlockCount int
	Format     core.Format
	TermInfo   *core.TermInfo

	StartOffset int64

	Ids     []int32
	Weights []float32
	Values  [][]byte

	MinLength int32
	MaxLength int32
	MaxWeight float32
}

func NewW3PostingsWriter(postfile fs.File, blocklimit int,
	compression int, inlinelimit int) *W3PostingsWriter {

	return &W3PostingsWriter{
		PostFile:    postfile,
		BlockLimit:  blocklimit,
		Compression: compression,
		InlineLimit: inlinelimit,
	}
}

func (writer *W3PostingsWriter) StartPostings(format core.Format, termInfo *core.TermInfo) error {
	if writer.TermInfo != nil {
		return errors.New("can not start in a term")
	}
	writer.Format = format
	writer.BlockCount = 0
	writer.NewBlock()
	writer.TermInfo = termInfo
	writer.StartOffset, _ = writer.PostFile.Seek(0, os.SEEK_CUR)
	return nil
}

func (writer *W3PostingsWriter) NewBlock() {
	writer.Ids = make([]int32, 0)
	writer.Weights = make([]float32, 0)
	writer.Values = make([][]byte, 0)
	writer.MinLength = -1
	writer.MaxLength = 0
	writer.MaxWeight = 0.0
}

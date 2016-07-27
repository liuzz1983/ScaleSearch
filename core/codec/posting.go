package codec

import "github.com/liuzz1983/scalesearch/core/index"

type PostingWriter interface {
	StartPostings(format index.Format, terminfo index.TermInfo) error
	AddPosting(id int64, weight float32, vbytes []byte) error
	FinishPostings() error
	IsWritten() bool
}

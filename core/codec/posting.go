package codec

import "github.com/liuzz1983/scalesearch/core"

type PostingWriter interface {
	StartPostings(format core.Format, terminfo core.TermInfo) error
	AddPosting(id int64, weight float32, vbytes []byte) error
	FinishPostings() error
	IsWritten() bool
}

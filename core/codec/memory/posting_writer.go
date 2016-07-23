package memory

import "github.com/liuzz1983/scalesearch/core"

type MemPostringWriter struct {
}

func (writer *MemPostringWriter) StartPostings(format core.Format, terminfo core.TermInfo) error {
	panic("not imp")
}
func (writer *MemPostringWriter) AddPosting(id int64, weight float32, vbytes []byte) error {
	panic("not imp")
}
func (writer *MemPostringWriter) FinishPostings() error {
	panic("not imp")
}
func (writer *MemPostringWriter) IsWritten() bool {
	panic("not imp")
}

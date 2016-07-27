package memory

import (
	"sync"

	"github.com/liuzz1983/scalesearch/core/codec"
	"github.com/liuzz1983/scalesearch/core/codec/base"
	"github.com/liuzz1983/scalesearch/core/index"
)

// use memory structure to store document and inverted index
type MemSegment struct {
	base.BaseSegment

	codec codec.Codec
	// document count
	docCount int
	// docNum , fieldName ,value
	stored map[int64]map[string][]byte
	// docNum, fieldName, lengths
	lengths map[int64]map[string]int
	// document vector term,
	vectors map[int64]map[string][]index.VectorItem

	// inverted index for store
	invIndex map[string]InvIndex
	// term information
	termInfos map[string]*index.TermInfo

	// lock for the structure
	lock sync.Mutex
}

func (seg *MemSegment) Codec() codec.Codec {
	return seg.codec
}

// lock support for memory segment
func (seg *MemSegment) Lock() {
	seg.lock.Lock()
}
func (seg *MemSegment) Unlock() {
	seg.lock.Unlock()
}

func (seg *MemSegment) SetDocCount(docCount int) {
	seg.docCount = docCount
}

func (seg *MemSegment) DocCount() int {
	return len(seg.stored)
}

func (seg *MemSegment) DocCountAll() int {
	return seg.docCount
}

func (seg *MemSegment) DeleteDocument(docNum int64, del bool) {
	if !del {
		panic("memory can not undelete")
	}
	// lock the segment
	seg.Lock()
	defer seg.Unlock()

	// delete fields in memory segment
	delete(seg.stored, docNum)
	delete(seg.lengths, docNum)
	delete(seg.vectors, docNum)

}

func (seg *MemSegment) HasDeletions() int {

	// lock the segment
	seg.Lock()
	defer seg.Unlock()
	return seg.docCount - len(seg.stored)
}

func (seg *MemSegment) IsDeleted(docNum int64) bool {
	if _, ok := seg.stored[docNum]; ok {
		return false
	}
	return true
}

// TODO need reconsider this code
func (seg *MemSegment) DeletedDocs() []int64 {
	result := make([]int64, 10)
	for i := 0; i < seg.DocCountAll(); i++ {
		num := int64(i)
		if _, ok := seg.stored[num]; !ok {
			result = append(result, num)
		}
	}
	return result
}

// whether the term in the segment
func (seg *MemSegment) ContainTerm(term *index.FieldTerm) bool {
	_, ok := seg.termInfos[term.Id()]
	return ok
}

// for assemble
func (seg *MemSegment) ShouldAssemble() bool {
	return false
}

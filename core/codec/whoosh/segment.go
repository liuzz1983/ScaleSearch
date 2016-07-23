package whoosh

import (
	"github.com/dropbox/godropbox/container/set"
	"github.com/liuzz1983/scalesearch/core/codec"
	"github.com/liuzz1983/scalesearch/core/codec/base"
)

type W3Segment struct {
	base.BaseSegment
	codec    codec.Codec
	docCount int64
	deleted  set.Set
}

func NewW3Segment(c codec.Codec, indexName string, docCount int64, segId string, deleted set.Set) {
	segment := &W3Segment{}
	segment.InxName = indexName
	if segId == "" {
		segment.SegId = segment.RandomId(codec.DEFAULT_RANDOM_SIZE)
	} else {
		segment.SegId = segId
	}
	segment.Compound = false
	segment.docCount = docCount
	if deleted != nil {
		segment.deleted = deleted
	} else {
		segment.deleted = set.NewSet()
	}

	segment.codec = c
}

func (seg *W3Segment) SetDocCount(count int64) {
	seg.docCount = count
}

func (seg *W3Segment) HasDeletions() bool {
	return seg.deleted.Len() > 0
}

func (seg *W3Segment) DeletedCount() int {
	return seg.deleted.Len()
}

func (seg *W3Segment) DeletedDocs() []int64 {
	docs := make([]int64, 0, seg.deleted.Len())
	for docNum := range seg.deleted.Iter() {
		docs = append(docs, docNum.(int64))
	}
	return docs
}

func (seg *W3Segment) DeleteDocument(docNum int64, del bool) error {
	seg.deleted.Add(docNum)
	return nil
}
func (seg *W3Segment) IsDeleted(docNum int64) bool {
	return seg.deleted.Contains(docNum)
}

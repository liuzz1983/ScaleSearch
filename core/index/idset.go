package index

import (
	"github.com/liuzz1983/scalesearch/core/errors"
)

// DocIdSetIter
type DocIdSetIter interface {
	Next() (DocId, error)
}

//DocIdSet interface for basic docidset operation
type DocIdSet interface {
	Copy() (DocIdSet, error)
	Add(DocId) error
	Discard(DocId) error
	Update(DocIdSet) error
	IntersectionUpdate(DocIdSet) error
	DifferenceUpdate(DocIdSet) error
	InvertUpdate(DocIdSet) error
	Intersection(DocIdSet) (DocIdSet, error)
	Union(DocIdSet) (DocIdSet, error)

	First() (DocId, error)
	Last() (DocId, error)
}

// BaseDocIdSet default struct for DocIdSet interface, define basic method for  doc idset implement
type BaseDocIdSet struct{}

func (sets *BaseDocIdSet) Copy() (DocIdSet, error) {
	return nil, errors.ErrNotImplement
}
func (sets *BaseDocIdSet) Add(DocId) error {
	return errors.ErrNotImplement
}
func (sets *BaseDocIdSet) Discard(DocId) error {
	return errors.ErrNotImplement
}
func (sets *BaseDocIdSet) Update(DocIdSet) error {
	return errors.ErrNotImplement
}
func (sets *BaseDocIdSet) IntersectionUpdate(DocIdSet) error {
	return errors.ErrNotImplement
}
func (sets *BaseDocIdSet) DifferenceUpdate(DocIdSet) error {
	return errors.ErrNotImplement
}
func (sets *BaseDocIdSet) InvertUpdate(DocIdSet) error {
	return errors.ErrNotImplement
}
func (sets *BaseDocIdSet) Intersection(DocIdSet) (DocIdSet, error) {
	return nil, errors.ErrNotImplement
}
func (sets *BaseDocIdSet) Union(DocIdSet) (DocIdSet, error) {
	return nil, errors.ErrNotImplement
}

func (sets *BaseDocIdSet) After(DocId) (DocId, error) {
	return DefaultDocIdNum, errors.ErrNotImplement
}

func (sets *BaseDocIdSet) Before(DocId) (DocId, error) {
	return DefaultDocIdNum, errors.ErrNotImplement
}

func (sets *BaseDocIdSet) First() (DocId, error) {
	return DefaultDocIdNum, errors.ErrNotImplement
}
func (sets *BaseDocIdSet) Last() (DocId, error) {
	return DefaultDocIdNum, errors.ErrNotImplement
}

type SortedIntSet struct {
	BaseDocIdSet
	ids *DocIdArray
}

func NewSortedIntSet() *SortedIntSet {
	return &SortedIntSet{
		ids: NewDocIdArray(),
	}
}

func (sets *SortedIntSet) Add(docId DocId) error {
	data := sets.ids
	if data == nil {
		data = NewDocIdArray()
	}
	return data.Add(docId)

}

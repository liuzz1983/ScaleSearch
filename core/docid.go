package core

import "github.com/liuzz1983/scalesearch/core/errors"

// DocId - for document id
type DocId uint32

var DefaultDocIdNum DocId = 0

const (
	DEFAUL_ID_NUM = 1024
)

type DocIdArray struct {
	data []DocId
}

func NewDocIdArray() *DocIdArray {
	return &DocIdArray{
		data: make([]DocId, 0, DEFAUL_ID_NUM),
	}
}

func (ids *DocIdArray) Length() int {
	return len(ids.data)
}

func (ids *DocIdArray) First() DocId {
	return ids.data[0]
}
func (ids *DocIdArray) Last() DocId {
	return ids.data[len(ids.data)-1]
}

func (ids *DocIdArray) Append(docIds ...DocId) {
	ids.data = append(ids.data, docIds...)
}

func (ids *DocIdArray) Add(docId DocId) error {

	length := ids.Length()
	if length == 0 || ids.Last() < docId {
		ids.Append(docId)
	} else {
		min := ids.First()
		max := ids.Last()
		if docId == min || docId == max {
			return nil
		} else if docId > max {
			ids.Append(docId)
		} else if docId < min {
			err := ids.Insert(0, docId)
			if err != nil {
				return err
			}
		} else {
			pos := ids.FindInsertPos(docId)
			err := ids.Insert(pos, docId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (ids *DocIdArray) Insert(index int, docId DocId) error {
	if index < 0 || index > len(ids.data) {
		return errors.ErrOutOfRange
	}
	result := make([]DocId, len(ids.data)+1)
	result = append(result, ids.data[:index]...)
	result = append(result, docId)
	result = append(result, ids.data[index:]...)
	ids.data = result
	return nil
}

func (ids *DocIdArray) Remove(id DocId) {
	pos := ids.Find(id)
	if pos == -1 {
		return
	}
	ids.data = append(ids.data[:pos], ids.data[pos+1:]...)
}

func (ids *DocIdArray) FindInsertPos(id DocId) int {
	for pos := 0; pos < len(ids.data); pos++ {
		if ids.data[pos] > id {
			return pos
		}
	}
	return len(ids.data)
}

func (ids *DocIdArray) Find(id DocId) int {
	right := len(ids.data) - 1
	left := 0
	for left < right {
		middle := (right-left)/2 + left
		if ids.data[middle] == id {
			return middle
		} else if ids.data[middle] > id {
			right = middle
		} else {
			left = middle
		}
	}
	return -1
}

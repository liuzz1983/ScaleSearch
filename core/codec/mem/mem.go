package mem

import (
	"strings"

	"github.com/liuzz1983/scalesearch/core/document"
)

type Indexer interface {
	Add(docId string, doc document.Document)
	Del(docId string)
}

type InvertIndex struct {
	// for max doc num
	maxDocCount int
	// doc count
	docCount int
	// store all dockids
	docids []string

	// docid map into internal id
	docidsIndexes map[string]int
	// inverted index for store
	invertedIndex map[string]interface{}
	deletes       map[string]string
}

func NewInvertIndex() *InvertIndex {
	return &InvertIndex{
		docids:        make([]string, 0),
		docidsIndexes: make(map[string]int),
		invertedIndex: make(map[string]interface{}),
		deletes:       make(map[string]string),
	}
}

func (indexer *InvertIndex) Add(docId string, doc document.Document) {

}

type Compare interface {
}

type ScoreMatch struct {
	Score float64
	DocId string
}

func (match *ScoreMatch) GetScore() float64 {
	return match.Score
}
func (match *ScoreMatch) GetDocId() string {
	return match.DocId
}
func (match *ScoreMatch) Compare(other ScoreMatch) int {
	if match.Score > other.Score {
		return 1
	} else if match.Score < other.Score {
		return -1
	} else {
		return strings.Compare(match.DocId, other.DocId)
	}
}

type Iterator interface {
	Next() interface{}
	HasNext() bool
}

type SkippableIterator interface {
	Iterator
	SkipTo(i int)
}

type TopMatches interface {
	getLimit()
	getTotalMatches()
	//String, Multiset<String>
	getFacetingResults() map[string]interface{}
}

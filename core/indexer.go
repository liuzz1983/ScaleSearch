package core

import "github.com/liuzz1983/scalesearch/core/document"

type Indexer interface {
	Add(docId string, doc document.Document)
	Del(docId string)
}

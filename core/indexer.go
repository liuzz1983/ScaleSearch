package core

type Indexer interface {
	add(docId string, doc Document)
	del(docId string)
}

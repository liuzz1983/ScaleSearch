package core

import (
	"fmt"
	"github.com/balzaczyy/golucene/core/index"
	"github.com/balzaczyy/golucene/core/search"
	"os"
	"testing"
)

func Test_Create_Index(t *testing.T) {
	dirName := "test_index"
	_, err := os.Getwd()

	lsiIndex, err := NewLsiIndex(dirName)
	if err != nil {
		t.Errorf("fail to build index %s", err)
		return
	}

	indexer, err := NewLsiIndexer(lsiIndex)
	if err != nil {
		t.Error("can not build indexer")
		return
	}
	if indexer == nil {
		t.Error("fail to build indexer")
		return
	}
	document := &Document{
		ID:     "1",
		Fields: []Field{Field{name: "liu", value: "zhong"}},
	}

	document2 := &Document{
		ID:     "2",
		Fields: []Field{Field{name: "liu", value: "hua"}},
	}
	indexer.add("1", document)
	indexer.add("2", document2)

	searcher, err := lsiIndex.openSearcher()
	reader, err := lsiIndex.openReader()

	q := search.NewTermQuery(index.NewTerm("liu", "zhong"))
	res, _ := searcher.Search(q, nil, 1000)
	fmt.Printf("Found %v hit(s).\n", res.TotalHits)
	for _, hit := range res.ScoreDocs {
		fmt.Printf("Doc %v score: %v\n", hit.Doc, hit.Score)
		doc, _ := reader.Document(hit.Doc)
		fmt.Printf("foo -> %v\n", doc.Get("liu"))
	}

}

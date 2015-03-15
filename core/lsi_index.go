package core

import (
	std "github.com/balzaczyy/golucene/analysis/standard"
	_ "github.com/balzaczyy/golucene/core/codec/lucene410"
	"github.com/balzaczyy/golucene/core/index"
	"github.com/balzaczyy/golucene/core/search"
	"github.com/balzaczyy/golucene/core/store"
	"github.com/balzaczyy/golucene/core/util"
	"os"
)

type LsiIndex struct {
	dirName   string
	directory store.Directory
	writer    *index.IndexWriter
}

func NewLsiIndex(dirName string) (lsi *LsiIndex, err error) {

	util.SetDefaultInfoStream(util.NewPrintStreamInfoStream(os.Stdout))
	index.DefaultSimilarity = func() index.Similarity {
		return search.NewDefaultSimilarity()
	}

	directory, err := store.OpenFSDirectory(dirName)
	if err != nil {
		return nil, err
	}

	lsiIndex := &LsiIndex{
		dirName,
		directory,
		nil,
	}
	writer, err := lsiIndex.openWriter()
	if err != nil {
		return nil, err
	}

	lsiIndex.writer = writer
	return lsiIndex, nil
}

func (lsi *LsiIndex) openReader() (index.DirectoryReader, error) {
	reader, err := index.OpenDirectoryReader(lsi.directory)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func (lsi *LsiIndex) openSearcher() (*search.IndexSearcher, error) {

	reader, err := index.OpenDirectoryReader(lsi.directory)
	if err != nil {
		return nil, err
	}
	searcher := search.NewIndexSearcher(reader)
	return searcher, nil
}

func (lsi *LsiIndex) openWriter() (writer *index.IndexWriter, err error) {

	analyzer := std.NewStandardAnalyzer()
	conf := index.NewIndexWriterConfig(util.VERSION_LATEST, analyzer)

	writer, err = index.NewIndexWriter(lsi.directory, conf)
	if err != nil {
		return nil, err
	}
	return writer, nil
}

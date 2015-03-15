package core

import (
	_ "errors"
	std "github.com/balzaczyy/golucene/analysis/standard"
	_ "github.com/balzaczyy/golucene/core/codec/lucene410"
	_ "github.com/balzaczyy/golucene/core/document"
	"github.com/balzaczyy/golucene/core/index"
	_ "github.com/balzaczyy/golucene/core/search"
	"github.com/balzaczyy/golucene/core/store"
	"github.com/balzaczyy/golucene/core/util"
)

type IndexEngine struct {
	baseDir   string
	indexer   Indexer
	directory store.Directory
	analyzer  std.StandardAnalyzer
	conf      *index.IndexWriterConfig
	writer    *index.IndexWriter
}

func NewIndexEngine(baseDir string, ops int) (*IndexEngine, error) {
	directory, _ := store.OpenFSDirectory(baseDir)
	analyzer := std.NewStandardAnalyzer()
	conf := index.NewIndexWriterConfig(util.VERSION_LATEST, analyzer)
	writer, _ := index.NewIndexWriter(directory, conf)
	engine := &IndexEngine{
		baseDir:   baseDir,
		directory: directory,
		conf:      conf,
		writer:    writer,
	}

	return engine, nil
}

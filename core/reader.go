package core

type IndexReader interface {
	Codec()
	Segment()
	Storage()
	IsAtomic()
	textToBytes(fieldname string, text string) (string, error)

	Close()
	Generation()

	IndexedFieldNames() []string
	AllTerms()
	TermFrom(fieldname string, prefix string) (string, string)

	AllDocIds() []string
	IterDocs() string

	IsDeleted(docnum int )
}

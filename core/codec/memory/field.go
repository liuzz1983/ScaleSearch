package memory

import (
	"errors"

	"github.com/liuzz1983/scalesearch/core"
	"github.com/liuzz1983/scalesearch/core/filedb"
)

type MemFieldWriter struct {
	storage filedb.Storage
	segment *MemSegment

	fieldName string
	fieldObj  interface{}
	term      []byte
	isClosed  bool

	postings []int64
	termInfo *core.TermInfo
}

func (writer *MemFieldWriter) StartField(fieldName string, fieldObj interface{}) error {
	if fieldName == "" {
		return errors.New("field name should not be empty")
	}

	writer.segment.Lock()
	defer writer.segment.Unlock()

	_, ok := writer.segment.invIndex[fieldName]
	if !ok {
		writer.segment.invIndex[fieldName] = NewInvIndex()
	}
	writer.fieldName = fieldName
	writer.fieldObj = fieldObj
	return nil
}
func (writer *MemFieldWriter) StartTerm(term []byte) {

	fieldName := writer.fieldName
	fieldDict := writer.segment.invIndex[fieldName]
	termInfos := writer.segment.termInfos

	writer.segment.Lock()
	defer writer.segment.Unlock()

	if !fieldDict.Contains(term) {
		writer.postings = make([]int64, 0)
	} else {
		writer.postings = fieldDict.Postings(term)
	}
	fieldTerm := core.FieldTerm{
		Field: []byte(fieldName),
		Term:  term,
	}

	fieldTermId := fieldTerm.Id()

	_, ok := termInfos[fieldTermId]
	if !ok {
		termInfos[fieldTermId] = &core.TermInfo{}
	}
	writer.termInfo = termInfos[fieldTermId]
	writer.term = term

}

type MemTermsReader struct {
	storage filedb.Storage
	segment *MemSegment
}

func NewMemTermReader(storage filedb.Storage, segment *MemSegment) *MemTermsReader {
	return &MemTermsReader{
		storage: storage,
		segment: segment,
	}
}

func (reader *MemTermsReader) Contains(term *core.FieldTerm) bool {
	return reader.segment.ContainTerm(term)
}

func (reader *MemTermsReader) Terms() []core.FieldTerm {
	terms := make([]core.FieldTerm, 0)
	for fieldName, invIndex := range reader.segment.invIndex {
		for term, _ := range invIndex {
			terms = append(terms, core.FieldTerm{[]byte(fieldName), []byte(term)})
		}
	}
	return terms
}

func (reader *MemTermsReader) TermsFrom(fieldName string, prefix []byte) []core.FieldTerm {
	/*
	       if fieldname not in self._invindex:
	       raise TermNotFound("Unknown field %r" % (fieldname,))
	   terms = sorted(self._invindex[fieldname])
	   if not terms:
	       return
	   start = bisect_left(terms, prefix)
	   for i in xrange(start, len(terms)):
	       yield (fieldname, terms[i])
	*/
	panic("not impl")
}

func (reader *MemTermsReader) TermInfo(fieldName string, term []byte) *core.TermInfo {
	v := &core.FieldTerm{[]byte(fieldName), term}
	return reader.segment.termInfos[v.Id()]
}

func (reader *MemTermsReader) Matcher(fieldName string, term []byte, format interface{}, scorer interface{}) {
	/*items = self._invindex[fieldname][btext]
	  ids, weights, values = zip(*items)
	  return ListMatcher(ids, weights, values, format_, scorer=scorer)*/
	panic("mot impl")

}

func (reader *MemTermsReader) IndexedFieldNames() []string {
	names := make([]string, len(reader.segment.invIndex))
	for key, _ := range reader.segment.invIndex {
		names = append(names, key)
	}
	return names
}

func (reader *MemTermsReader) Close() error {
	return nil
}

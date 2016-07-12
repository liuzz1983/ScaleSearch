package codec

import (
	"github.com/liuzz1983/scalesearch/core/filedb"
    "sync"
    "io"
)

type MemWriter struct {
}

func (writer MemWriter) commit() {
	//writer.FinalizeSegment()
}


type MemTermsReader struct {
	storage  filedb.Storage
	segment  MemSegment
	invIndex map[string]InvIndex
}

func NewMemTermReader(storage filedb.Storage, segment MemSegment) *MemTermsReader {
	return &MemTermsReader{
		storage:  storage,
		segment:  segment,
		invIndex: segment.invindex,
	}
}

func (reader *MemTermsReader) Contains(term FieldTerm) bool {
	_, ok := reader.segment.terminfos[term]
	return ok
}

type Term struct {
	FieldName string
	Text      string
}


type MemSegment struct {
    BaseSegment

    codec Codec

    docCount int
    stored   map[int64]interface{}
    lengths  map[int64]interface{}
    vectors  map[int64]interface{}

    invindex  map[string]InvIndex
    terminfos map[FieldTerm]TermInfo
    lock      sync.Mutex
}

func (seg *MemSegment) Codec() Codec {
    return seg.codec
}

func (seg *MemSegment) SetDocCount(docCount int) {
    seg.docCount = docCount
}

func (seg *MemSegment) DocCount() int {
    return len(seg.stored)
}

func (seg *MemSegment) DocCountAll() int {
    return seg.docCount
}

func (seg *MemSegment) DeleteDocument(docNum int64, del bool) {
    if !del {
        panic("memory can not undelete")
    }

    // lock the segment
    seg.lock.Lock()
    defer seg.lock.Unlock()

    delete(seg.stored, docNum)
    delete(seg.lengths, docNum)
    delete(seg.vectors, docNum)

}

func (seg *MemSegment) HasDeletions() int {

    // lock the segment
    seg.lock.Lock()
    defer seg.lock.Unlock()

    return seg.docCount - len(seg.stored)
}

func (seg *MemSegment) IsDeleted(docNum int64) bool {
    if _, ok := seg.stored[docNum]; ok {
        return false
    }
    return true
}

// TODO need reconsider this code
func (seg *MemSegment) DeletedDocs() []int64 {
    result := make([]int64, 10)
    for i := 0; i < seg.DocCountAll(); i++ {
        num := int64(i)
        if _, ok := seg.stored[num]; !ok {
            result = append(result, num)
        }
    }
    return result
}

func (seg *MemSegment) ShouldAssemble() bool {
    return false
}


func (reader *MemTermsReader) Terms(ch chan<- *Term) {
	for fieldName, v := range reader.invIndex {
		for text, _ := range v {
			ch <- &Term{fieldName, text}
		}
	}
}

/*func (reader *MemTermsReader) TermsFrom(fieldName string, prefix string) error {
	if _, ok := reader.invIndex ; !ok {
        return  errors.TermNotFound;
    }
    terms := sorted(reader.invIndex[fieldName])
    if len(tems) == 0:
        return nil

}*/

func (reader *MemTermsReader) Close() {

}

type MemPerDocWriter struct {
    storage filedb.Storage
    segment MemSegment
    isClosed bool
    colWriters map[string]io.Writer
    
    docCount int64
}

func NewMemPerDocWriter( storage filedb.Storage, segment MemSegment) *MemPerDocWriter{
    return &MemPerDocWriter {
        storage: storage,
        segment: segment,
        isClosed: false,
        colWriters: make( map[string]io.Writer, 10),
        docCount: 0,
    }
}

func ( writer *MemPerDocWriter) hasColumn( fieldName string) bool {
    _, ok := writer.colWriters[fieldName]
    return ok
}

/**
func (writer *MemPerDocWriter) createColumn( fieldName, )


    def _has_column(self, fieldname):
        return fieldname in self._colwriters

    def _create_column(self, fieldname, column):
        colfile = self._storage.create_file("%s.c" % fieldname)
        self._colwriters[fieldname] = (colfile, column.writer(colfile))

    def _get_column(self, fieldname):
         return self._colwriters[fieldname][1]
*/
         /*
    def start_doc(self, docnum):
        self._doccount += 1
        self._docnum = docnum
        self._stored = {}
        self._lengths = {}
        self._vectors = {}

    def add_field(self, fieldname, fieldobj, value, length):
        if value is not None:
            self._stored[fieldname] = value
        if length is not None:
            self._lengths[fieldname] = length

    def add_vector_items(self, fieldname, fieldobj, items):
        self._vectors[fieldname] = tuple(items)

    def finish_doc(self):
        with self._segment._lock:
            docnum = self._docnum
            self._segment._stored[docnum] = self._stored
            self._segment._lengths[docnum] = self._lengths
            self._segment._vectors[docnum] = self._vectors

    def close(self):
        colwriters = self._colwriters
        for fieldname in colwriters:
            colfile, colwriter = colwriters[fieldname]
            colwriter.finish(self._doccount)
            colfile.close()
        self.is_closed = True

/*
   def terms_from(self, fieldname, prefix):
       if fieldname not in self._invindex:
           raise TermNotFound("Unknown field %r" % (fieldname,))
       terms = sorted(self._invindex[fieldname])
       if not terms:
           return
       start = bisect_left(terms, prefix)
       for i in xrange(start, len(terms)):
           yield (fieldname, terms[i])

   def term_info(self, fieldname, text):
       return self._segment._terminfos[fieldname, text]

   def matcher(self, fieldname, btext, format_, scorer=None):
       items = self._invindex[fieldname][btext]
       ids, weights, values = zip(*items)
       return ListMatcher(ids, weights, values, format_, scorer=scorer)

   def indexed_field_names(self):
       return self._invindex.keys()

   def close(self):
       pass
*/

/*
class MemPerDocWriter(base.PerDocWriterWithColumns):
    def __init__(self, storage, segment):
        self._storage = storage
        self._segment = segment
        self.is_closed = False
        self._colwriters = {}
        self._doccount = 0

    def _has_column(self, fieldname):
        return fieldname in self._colwriters

    def _create_column(self, fieldname, column):
        colfile = self._storage.create_file("%s.c" % fieldname)
        self._colwriters[fieldname] = (colfile, column.writer(colfile))

    def _get_column(self, fieldname):
        return self._colwriters[fieldname][1]

    def start_doc(self, docnum):
        self._doccount += 1
        self._docnum = docnum
        self._stored = {}
        self._lengths = {}
        self._vectors = {}

    def add_field(self, fieldname, fieldobj, value, length):
        if value is not None:
            self._stored[fieldname] = value
        if length is not None:
            self._lengths[fieldname] = length

    def add_vector_items(self, fieldname, fieldobj, items):
        self._vectors[fieldname] = tuple(items)

    def finish_doc(self):
        with self._segment._lock:
            docnum = self._docnum
            self._segment._stored[docnum] = self._stored
            self._segment._lengths[docnum] = self._lengths
            self._segment._vectors[docnum] = self._vectors

    def close(self):
        colwriters = self._colwriters
        for fieldname in colwriters:
            colfile, colwriter = colwriters[fieldname]
            colwriter.finish(self._doccount)
            colfile.close()
        self.is_closed = True


class MemPerDocReader(base.PerDocumentReader):
    def __init__(self, storage, segment):
        self._storage = storage
        self._segment = segment

    def doc_count(self):
        return self._segment.doc_count()

    def doc_count_all(self):
        return self._segment.doc_count_all()

    def has_deletions(self):
        return self._segment.has_deletions()

    def is_deleted(self, docnum):
        return self._segment.is_deleted(docnum)

    def deleted_docs(self):
        return self._segment.deleted_docs()

    def supports_columns(self):
        return True

    def has_column(self, fieldname):
        filename = "%s.c" % fieldname
        return self._storage.file_exists(filename)

    def column_reader(self, fieldname, column):
        filename = "%s.c" % fieldname
        colfile = self._storage.open_file(filename)
        length = self._storage.file_length(filename)
        return column.reader(colfile, 0, length, self._segment.doc_count_all())

    def doc_field_length(self, docnum, fieldname, default=0):
        return self._segment._lengths[docnum].get(fieldname, default)

    def field_length(self, fieldname):
        return sum(lens.get(fieldname, 0) for lens
                   in self._segment._lengths.values())

    def min_field_length(self, fieldname):
        return min(lens[fieldname] for lens in self._segment._lengths.values()
                   if fieldname in lens)

    def max_field_length(self, fieldname):
        return max(lens[fieldname] for lens in self._segment._lengths.values()
                   if fieldname in lens)

    def has_vector(self, docnum, fieldname):
        return (docnum in self._segment._vectors
                and fieldname in self._segment._vectors[docnum])

    def vector(self, docnum, fieldname, format_):
        items = self._segment._vectors[docnum][fieldname]
        ids, weights, values = zip(*items)
        return ListMatcher(ids, weights, values, format_)

    def stored_fields(self, docnum):
        return self._segment._stored[docnum]

    def close(self):
        pass


class MemFieldWriter(base.FieldWriter):
    def __init__(self, storage, segment):
        self._storage = storage
        self._segment = segment
        self._fieldname = None
        self._btext = None
        self.is_closed = False

    def start_field(self, fieldname, fieldobj):
        if self._fieldname is not None:
            raise Exception("Called start_field in a field")

        with self._segment._lock:
            invindex = self._segment._invindex
            if fieldname not in invindex:
                invindex[fieldname] = {}

        self._fieldname = fieldname
        self._fieldobj = fieldobj

    def start_term(self, btext):
        if self._btext is not None:
            raise Exception("Called start_term in a term")
        fieldname = self._fieldname

        fielddict = self._segment._invindex[fieldname]
        terminfos = self._segment._terminfos
        with self._segment._lock:
            if btext not in fielddict:
                fielddict[btext] = []

            if (fieldname, btext) not in terminfos:
                terminfos[fieldname, btext] = TermInfo()

        self._postings = fielddict[btext]
        self._terminfo = terminfos[fieldname, btext]
        self._btext = btext

    def add(self, docnum, weight, vbytes, length):
        self._postings.append((docnum, weight, vbytes))
        self._terminfo.add_posting(docnum, weight, length)

    def finish_term(self):
        if self._btext is None:
            raise Exception("Called finish_term outside a term")

        self._postings = None
        self._btext = None
        self._terminfo = None

    def finish_field(self):
        if self._fieldname is None:
            raise Exception("Called finish_field outside a field")
        self._fieldname = None
        self._fieldobj = None

    def close(self):
        self.is_closed = True


class MemTermsReader(base.TermsReader):
    def __init__(self, storage, segment):
        self._storage = storage
        self._segment = segment
        self._invindex = segment._invindex

    def __contains__(self, term):
        return term in self._segment._terminfos

    def terms(self):
        for fieldname in self._invindex:
            for btext in self._invindex[fieldname]:
                yield (fieldname, btext)

    def terms_from(self, fieldname, prefix):
        if fieldname not in self._invindex:
            raise TermNotFound("Unknown field %r" % (fieldname,))
        terms = sorted(self._invindex[fieldname])
        if not terms:
            return
        start = bisect_left(terms, prefix)
        for i in xrange(start, len(terms)):
            yield (fieldname, terms[i])

    def term_info(self, fieldname, text):
        return self._segment._terminfos[fieldname, text]

    def matcher(self, fieldname, btext, format_, scorer=None):
        items = self._invindex[fieldname][btext]
        ids, weights, values = zip(*items)
        return ListMatcher(ids, weights, values, format_, scorer=scorer)

    def indexed_field_names(self):
        return self._invindex.keys()

    def close(self):
        pass


class MemSegment(base.Segment):
    def __init__(self, codec, indexname):
        base.Segment.__init__(self, indexname)
        self._codec = codec
        self._doccount = 0
        self._stored = {}
        self._lengths = {}
        self._vectors = {}
        self._invindex = {}
        self._terminfos = {}
        self._lock = Lock()

    def codec(self):
        return self._codec

    def set_doc_count(self, doccount):
        self._doccount = doccount

    def doc_count(self):
        return len(self._stored)

    def doc_count_all(self):
        return self._doccount

    def delete_document(self, docnum, delete=True):
        if not delete:
            raise Exception("MemoryCodec can't undelete")
        with self._lock:
            del self._stored[docnum]
            del self._lengths[docnum]
            del self._vectors[docnum]

    def has_deletions(self):
        with self._lock:
            return self._doccount - len(self._stored)

    def is_deleted(self, docnum):
        return docnum not in self._stored

    def deleted_docs(self):
        stored = self._stored
        for docnum in xrange(self.doc_count_all()):
            if docnum not in stored:
                yield docnum

    def should_assemble(self):
        return False

*/

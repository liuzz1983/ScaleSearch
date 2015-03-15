package filedb

import (
	_ "bytes"
	_ "encoding/binary"
	"encoding/json"
	_ "io"
)

type hashFunc func(key []byte) uint32

var cdbHash hashFunc = func(key []byte) uint32 {
	h := uint32(5381)
	u := uint32(0xffffffff)
	for c := range key {
		h = (h+h<<5)&u ^ uint32(c)
	}
	return h
}

var (
	HASH_FUNCTIOONS = [...]hashFunc{cdbHash}
)

type bucketEntry struct {
	key uint32
	pos int64
}

type bucketEntries struct {
	items []bucketEntry
}

func (e *bucketEntries) Add(key uint32, pos int64) {
	if e.items == nil {
		e.items = make([]bucketEntry, 0, 10)
	}
	e.items = append(e.items, bucketEntry{key, pos})
}

type dictionaryEntry struct {
	pos int64
	num int
}

type HashWriter struct {
	file     *StructFile
	magic    []byte
	hashType byte
	keyFunc  hashFunc

	startOffset int64
	buckets     []bucketEntries
	dictionary  []dictionaryEntry

	extras map[string]interface{}
}

func NewHashWriter(file *StructFile, magic []byte) (*HashWriter, error) {
	if magic == nil || len(magic) == 0 {
		magic = DEFAULT_MAGIC
	}

	writer := &HashWriter{
		file:     file,
		magic:    magic,
		hashType: 0,
	}

	writer.keyFunc = cdbHash
	writer.startOffset, _ = file.Tell()
	writer.buckets = make([]bucketEntries, 0, 256)
	writer.dictionary = make([]dictionaryEntry, 0, 256)

	//buf := new(bytes.Buffer)

	//binary.Write(buf, binary.LittleEndian, &p)

	err := writeBinary(writer.file, magic)
	if err != nil {
		return writer, err
	}
	err = writeBinary(writer.file, writer.hashType)
	if err != nil {
		return writer, err
	}

	//Unused future expansion bits
	err = writeBinary(writer.file, int32(0))
	if err != nil {
		return writer, err
	}
	err = writeBinary(writer.file, int32(0))
	if err != nil {
		return writer, err
	}

	return writer, nil
}

func (writer *HashWriter) Add(key []byte, value []byte) error {

	dbfile := writer.file
	pos, err := dbfile.Tell()

	//write key,value length
	err = writeBinary(dbfile, uint32(len(key)))
	if err != nil {
		return err
	}
	err = writeBinary(dbfile, uint32(len(value)))
	if err != nil {
		return err
	}

	//write key and value
	err = writeBinary(dbfile, key)
	if err != nil {
		return err
	}
	err = writeBinary(dbfile, value)
	if err != nil {
		return err
	}

	//get Hash for the key
	h := writer.keyFunc(key)
	writer.buckets[h&255].Add(h, pos)

	return nil
}

func (writer *HashWriter) writePointer(entry *bucketEntry) error {
	err := writeBinary(writer.file, entry.key)
	if err != nil {
		return err
	}
	err = writeBinary(writer.file, entry.pos)
	if err != nil {
		return err
	}
	return nil
}

func (writer *HashWriter) writeDictionary(entry *dictionaryEntry) error {
	err := writeBinary(writer.file, entry.pos)
	if err != nil {
		return err
	}
	err = writeBinary(writer.file, entry.num)
	if err != nil {
		return err
	}
	return nil
}

func (writer *HashWriter) writeHash() error {
	dbfile := writer.file
	//null := []int64{0, 0}

	nullEntry := bucketEntry{0, 0}

	for _, entry := range writer.buckets {
		pos, _ := dbfile.Tell()
		numslots := 2 * len(entry.items)
		hashtable := make([]bucketEntry, numslots)

		writer.dictionary = append(writer.dictionary, dictionaryEntry{pos, numslots})
		for _, entry := range entry.items {
			slot := (entry.key >> uint32(8)) % uint32(numslots)
			for hashtable[slot] != nullEntry {
				slot = (slot + 1) % uint32(numslots)
			}
			hashtable[slot] = entry
		}

		for _, entry := range hashtable {
			err := writer.writePointer(&entry)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (writer *HashWriter) writeDictionaries() error {
	//dbfile := writer.file
	for _, entry := range writer.dictionary {
		err := writer.writeDictionary(&entry)
		if err != nil {
			return err
		}
	}
	return nil
}

func (writer *HashWriter) writerExtra() error {

	//serialize the extra values
	values, err := json.Marshal(writer.extras)
	if err != nil {
		return err
	}

	// write extra value
	err = writeBinary(writer.file, values)
	if err != nil {
		return err
	}
	return nil
}

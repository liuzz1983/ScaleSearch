package index

type FieldTerm struct {
	Field []byte
	Term  []byte
}

func (term *FieldTerm) Id() string {
	key := make([]byte, 0, len(term.Field)+len(term.Term)+1)
	key = append(key, term.Field...)
	key = append(key, ':')
	key = append(key, term.Term...)
	return string(key)
}

type VectorItem struct {
	Text   []byte
	Weight float32
	Value  []byte
}

type FieldTerms map[string][]string

type TermFieldVector struct {
	Field          string
	ArrayPositions []uint64
	Pos            uint64
	Start          uint64
	End            uint64
}

type TermFieldDoc struct {
	Term    string
	ID      string
	Freq    uint64
	Norm    float64
	Vectors []*TermFieldVector
}

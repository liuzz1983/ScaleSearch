package memory

type InvIndex map[string][]int64

func NewInvIndex() InvIndex {
	return make(InvIndex)
}

func (inv InvIndex) Contains(term []byte) bool {
	_, ok := inv[string(term)]
	return ok
}

func (inv InvIndex) Add(term []byte, docNum int64) {
}

func (inv InvIndex) Postings(term []byte) []int64 {
	postings, ok := inv[string(term)]
	if ok {
		return postings
	}
	return make([]int64, 0)
}

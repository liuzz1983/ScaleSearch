package matching

// used for read posting list
type Matcher interface {
	IsActive() bool
	Reset()
	// return (fieldname, termtext)
	Term() ([]byte, []byte)
	TermMatchers()
	MatchingTerms(id int64)

	IsLeaf() bool
	Children() []Matcher
}

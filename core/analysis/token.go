package analysis

type Token struct {
	Positions   bool
	Chars       bool
	Stopped     bool
	Removestops bool

	Boost float32
	Text  []byte
	Pos   int32
}

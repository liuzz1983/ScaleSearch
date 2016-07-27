package analysis

type Token struct {
	Boost float32
	Text  []byte
	Pos   int32
}

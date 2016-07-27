package index

type FieldType interface {
	Index(document string)
}

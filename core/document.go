package core

type Field struct {
	name  string `json:name`
	value string `json:value`
}

type Document struct {
	ID     string  `json:"id"`
	Fields []Field `json:"fields"`
}

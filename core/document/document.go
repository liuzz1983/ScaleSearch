package document

import (
	"errors"

	"github.com/liuzz1983/scalesearch/utils/types"
)

type Document struct {
	Fields map[string]interface{}
}

func (doc *Document) GetField(name string) (interface{}, error) {
	field, ok := doc.Fields[name]
	if !ok {
		return nil, errors.New("cant not find fields")
	}
	return field, nil
}

func (doc *Document) GetFieldNames() []string {
	names := make([]string, 0, len(doc.Fields))
	for name, _ := range doc.Fields {
		names = append(names, name)
	}
	return names
}

func (doc *Document) GetString(name string) ([]byte, error) {
	field, err := doc.GetField(name)
	if err != nil {
		return nil, err
	}
	value, err := types.ToBytes(field)
	return value, err
}
